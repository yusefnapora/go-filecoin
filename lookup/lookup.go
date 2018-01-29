package lookup

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	logging "gx/ipfs/QmSpJByNKFX1sCsHBEp3R73FL4NF6FnQTEGyNAXHm2GS52/go-log"
	floodsub "gx/ipfs/Qma2TkMxcFLVGkYECTo4hrQohBYPx7uhpYL9EejEi8y3Nm/go-libp2p-floodsub"
	peer "gx/ipfs/Qma7H6RW8wRrfZpNSXwxYGcd1E149s42FpWNpDNieSVrnU/go-libp2p-peer"

	pubsub "github.com/briantigerchow/pubsub"
	types "github.com/filecoin-project/go-filecoin/types"
)

var FilLookupTopic = "/fil/lookup/1.0.0"

var log = logging.Logger("lookup")

type LookupEngine struct {
	lk    sync.Mutex
	cache map[types.Address]peer.ID

	ourAddresses map[types.Address]struct{}
	ourPeerID    peer.ID

	reqPubsub *pubsub.PubSub

	ps *floodsub.PubSub
}

func NewLookupEngine(ps *floodsub.PubSub, self peer.ID) (*LookupEngine, error) {
	sub, err := ps.Subscribe(FilLookupTopic)
	if err != nil {
		return nil, err
	}

	le := &LookupEngine{
		ps:           ps,
		cache:        make(map[types.Address]peer.ID),
		ourPeerID:    self,
		ourAddresses: make(map[types.Address]struct{}),
		reqPubsub:    pubsub.New(128),
	}

	go le.HandleMessages(sub)
	return le, nil
}

type message struct {
	Address types.Address
	Peer    string
	Request bool
}

func (le *LookupEngine) HandleMessages(s *floodsub.Subscription) {
	defer s.Cancel()
	ctx := context.TODO()
	for {
		msg, err := s.Next(ctx)
		if err != nil {
			log.Error("from subscription.Next(): ", err)
			return
		}
		if msg.GetFrom() == le.ourPeerID {
			continue
		}

		var m message
		if err := json.Unmarshal(msg.GetData(), &m); err != nil {
			log.Error("malformed message: ", err)
			continue
		}

		le.lk.Lock()
		if m.Request {
			if _, ok := le.ourAddresses[m.Address]; ok {
				go le.SendMessage(&message{
					Address: m.Address,
					Peer:    le.ourPeerID.Pretty(),
				})
			}
		} else {
			pid, err := peer.IDB58Decode(m.Peer)
			if err != nil {
				log.Error("bad peer ID: ", err)
				continue
			}
			le.cache[m.Address] = pid
			le.reqPubsub.Pub(pid, string(m.Address))
		}
		le.lk.Unlock()
	}
}

func (le *LookupEngine) SendMessage(m *message) {
	d, err := json.Marshal(m)
	if err != nil {
		log.Error("failed to marshal message: ", err)
		return
	}

	if err := le.ps.Publish(FilLookupTopic, d); err != nil {
		log.Error("publish failed: ", err)
	}
}

func (le *LookupEngine) Lookup(a types.Address) (peer.ID, error) {
	le.lk.Lock()
	v, ok := le.cache[a]
	le.lk.Unlock()
	if ok {
		return v, nil
	}

	ch := le.reqPubsub.SubOnce(string(a))

	le.SendMessage(&message{
		Address: a,
		Request: true,
	})

	select {
	case out := <-ch:
		return out.(peer.ID), nil
	case <-time.After(time.Second * 10):
		return "", fmt.Errorf("timed out waiting for response")
	}
}

func (le *LookupEngine) AddAddress(a types.Address) {
	le.lk.Lock()
	le.ourAddresses[a] = struct{}{}
	le.lk.Unlock()
}
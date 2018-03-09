package commands

import (
	"io"
	"math/big"

	cmds "gx/ipfs/QmRv6ddf7gkiEgBs1LADv3vC1mkVGPZEfByoiiVybjE9Mc/go-ipfs-cmds"
	"gx/ipfs/QmcZfnkapfECQGcLZaf9B79NRg7cRa9EnZh4LSbkCzwNvY/go-cid"
	cmdkit "gx/ipfs/QmceUdzxkimdYsgtX733uNgzf1DLHyBKN6ehGSp85ayppM/go-ipfs-cmdkit"

	"github.com/filecoin-project/go-filecoin/types"
)

var sendMsgCmd = &cmds.Command{
	Helptext: cmdkit.HelpText{
		Tagline: "Send a message", // This feels too generic...
	},
	Arguments: []cmdkit.Argument{
		cmdkit.StringArg("target", true, false, "address to send message to"),
	},
	Options: []cmdkit.Option{
		cmdkit.IntOption("value", "value to send with message"),
		cmdkit.StringOption("from", "address to send message from"),
	},
	Run: func(req *cmds.Request, re cmds.ResponseEmitter, env cmds.Environment) {
		n := GetNode(env)

		target, err := types.ParseAddress(req.Arguments[0])
		if err != nil {
			re.SetError(err, cmdkit.ErrNormal)
			return
		}

		val, ok := req.Options["value"].(int)
		if !ok {
			val = 0
		}

		fromAddr, err := addressWithDefault(req.Options["from"], n)
		if err != nil {
			re.SetError(err, cmdkit.ErrNormal)
			return
		}

		msg := types.NewMessage(fromAddr, target, big.NewInt(int64(val)), "", nil)

		if err := n.AddNewMessage(req.Context, msg); err != nil {
			re.SetError(err, cmdkit.ErrNormal)
			return
		}

		c, err := msg.Cid()
		if err != nil {
			re.SetError(err, cmdkit.ErrNormal)
			return
		}

		re.Emit(c) // nolint: errcheck
	},
	Type: cid.Cid{},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request, w io.Writer, c *cid.Cid) error {
			return PrintString(w, c)
		}),
	},
}
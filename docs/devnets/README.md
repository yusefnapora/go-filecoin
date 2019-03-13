<!---
title: "Devnets"
weight: 60
pre: "<i class='fas fa-fw fa-globe'></i> "
--->

There are currently 3 Filecoin developer networks (aka devnets) available for development and testing.

Both are subject to _**frequent downtimes**_ for a variety of reasons, including bugs, breaking changes, or the polar vortex ⛄. There is no [SLA](https://en.wikipedia.org/wiki/Service-level_agreement); until we have a status page, please ping in [#fil-chat](https://github.com/filecoin-project/community#chat) if you would like to confirm status.

### User

At the moment, **this is the network to use by default, for testing and experimenting**. Nodes are deployed from master manually.

- NetworkStats: https://stats.kittyhawk.wtf
- Faucet: http://user.kittyhawk.wtf:9797/
- Dashboard: http://user.kittyhawk.wtf:8010/
- Genesis File: http://user.kittyhawk.wtf:8020/genesis.car
- Block explorer: http://user.kittyhawk.wtf:8000/
- Prometheus Endpoint: http://user.kittyhawk.wtf:9082/metrics
- Connected Nodes PeerID's: http://user.kittyhawk.wtf:9082/nodes
- Sector size: **256MiB**

The [Getting Started](/getting-started) page contains
instructions for connecting your nodes to this network.

### Nightly

Deployed from master by CI every day at 0600 UTC. **This network
is for people contributing to `go-filecoin`. You should probably avoid it unless that describes you.**

- Faucet: https://faucet.nightly.kittyhawk.wtf/
- Dashboard: http://nightly.kittyhawk.wtf:8010/
- Genesis File: https://genesis.nightly.kittyhawk.wtf/genesis.car
- Block explorer: https://explorer.nightly.kittyhawk.wtf/
- Prometheus Endpoint: http://nightly.kittyhawk.wtf:9082/metrics
- Connected Nodes PeerID's: http://nightly.kittyhawk.wtf:9082/nodes
- Sector size: **1KiB**

### Test (for Infra)

Deployed via CI by tagging a commit with `redeploy_test_devnet`. **This network
is for people working on infra. You should probably avoid it unless that describes you.**

- Faucet: https://faucet.test.kittyhawk.wtf/
- Dashboard: http://test.kittyhawk.wtf:8010/
- Genesis File: https://genesis.test.kittyhawk.wtf/genesis.car
- Block explorer: https://explorer.test.kittyhawk.wtf/
- Prometheus Endpoint: http://test.kittyhawk.wtf:9082/metrics
- Connected Nodes PeerID's: http://test.kittyhawk.wtf:9082/nodes
- Sector size: **not stable**

## Configuration
Both the Test and Nightly networks run with `FIL_USE_SMALL_SECTORS=true`, this means that the proofs system only ever seals the first 1016 bytes of a sector. Because of this, users connecting to the these networks must:

- Set `FIL_USE_SMALL_SECTORS=true` before starting the daemon, else risk experiencing undefined behavior.
- Only store files that are <= 1016 bytes.

## Network Health
To get an idea of the state of each network, you can visit the network dashboard (port 8010).
Once the page is loaded, wait about 5 to 10 seconds for heartbeats to arrive. A healthy
network should show at least 5 nodes with the nickname "boot".

Each network is set up to mine a new block about every 30 seconds. If the "Last Block" column
on the far right remains blank, or the number increments far beyond the 30 second mining time,
this indicates that the network is no longer mining, or is not mining at the desired rate.

Another indication that a network is not in a healthy state is the background
color of the tipset value. If all colors are the same, the nodes are in consensus. Two or more
colors that persist may indicate that the nodes have fallen out of consensus.

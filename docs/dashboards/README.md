---
title: "Dashboards"
weight: 50
pre: "<i class='fas fa-tachometer-alt'></i>&nbsp;"
---

The Filecoin dashboards provides a visual way to see nodes on the Filecoin network.

View stats about the whole devnet network:
- [User devnet network stats](https://stats.kittyhawk.wtf/)
- [User devnet explorer](http://user.kittyhawk.wtf:8000/)

See [Devnets](../devnets) for links to dashboards on the nightly and infra/test devnets.

In order to see your node's activity, configure your node to [stream its activity](../getting-started#start-streaming-activity-from-your-node):

```Shell
# User devnet
go-filecoin config heartbeat.beatTarget "/dns4/stats-infra.kittyhawk.wtf/tcp/8080/ipfs/QmUWmZnpZb6xFryNDeNU7KcJ1Af5oHy7fB9npU67sseEjR"
# Nightly devnet
go-filecoin config heartbeat.beatTarget "/dns4/nightly.kittyhawk.wtf/tcp/9081/ipfs/QmVR3UFv588pSu8AxSw9C6DrMHiUFkWwdty8ajgPvtWaGU"
```

Look for your node's miner address, which you can obtain by running `go-filecoin id`.

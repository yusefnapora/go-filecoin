<!---
title: "More HowTos"
weight: 90
pre: "<i class='fas fa-fw fa-ellipsis-h'></i> "
--->

# HOW TO: Connect Two Nodes

This guide is primarily for getting two nodes mining and connected to a swarm on a single machine, for local development. Some of the instructions can still apply to connecting to a devnet.

Optionally, you could look at and modify your own copy of [this script](https://github.com/filecoin-project/go-filecoin/blob/master/functional-tests/retrieval), which sets up two local nodes and does a file store/retrieval.

## Initial setup

Let's first set the location of the filecoin source (substitute with where ever your source tree is cloned):

`export GO_FILECOIN_PATH=$HOME/go/src/github.com/filecoin-project/go-filecoin`

### Node 1

Initialize go-filecoin in the default directory, and use the genesis file from go-filecoin source. These files will now be in $HOME/.filecoin

```
go-filecoin init --genesisfile ./fixtures/genesis.car
```

Run the daemon:
```
go-filecoin daemon
```

Get the address of Node 1 with `go-filecoin id`:
```
$ go-filecoin id
    {
	"Addresses": [
		"/ip4/127.0.0.1/tcp/6000/ipfs/QmVk7A2vEBFr9GyKyQ3wvDmTWj8M4H3jubHUDc3CktdoXL",
		"/ip4/172.16.200.201/tcp/6000/ipfs/QmVk7A2vEBFr9GyKyQ3wvDmTWj8M4H3jubHUDc3CktdoXL"
	],
	"ID": "QmVk7A2vEBFr9GyKyQ3wvDmTWj8M4H3jubHUDc3CktdoXL",
	"AgentVersion": "",
	"ProtocolVersion": "",
	"PublicKey": ""
    }
```

Then for convenience, we'll export the node address to a bash variable.
```
export NODE1_ADDR=/ip4/127.0.0.1/tcp/6000/ipfs/QmVk7A2vEBFr9GyKyQ3wvDmTWj8M4H3jubHUDc3CktdoXL    
```

### Node 2
This assumes you wish to run two nodes on the same machine for development/testing purposes. If you are trying to connect two separate machines then you will not need to use `--repodir` everywhere unless you do not want to use the default filecoin directory in $HOME/.filecoin.

In another terminal, choose where you want the other FileCoin RepoDir to be, and supply this to the intialization script.

You have two choices when running commands for a node that is not located in the default directory - either `export FIL_PATH=<repodir>` or else you must always specify `--repodir`. In these instructions, we have chosen the latter, to make it easy to distinguish which node the commands are being run against.

Then specify non-default values for the api.address and swarm.address. NOTE ANY VALUE HAS TO BE SINGLE AND DOUBLE-QUOTED.

Then you can run the daemon.

```
export FCRD=$HOME/.filecoin2
go-filecoin init --genesisfile $GO_FILECOIN_PATH/fixtures/genesis.car --repodir=$FCRD
```

Edit config.json file and change the api address and the swarm address:
```
       "api": {                                                                                                                                                            
                "address": "/ip4/127.0.0.1/tcp/3453",                                                                                                                       
                "accessControlAllowOrigin": [ ^^^^ change this to a different port/value                                                                                                                              
                        "http://localhost:8080",                                                                                                                            
                        "https://localhost:8080",                                                                                                                           
                        "http://127.0.0.1:8080",                                                                                                                            
                        "https://127.0.0.1:8080"                                                                                                                            
                ],                                                                                                                                                          
                "accessControlAllowCredentials": false,                                                                                                                     
                "accessControlAllowMethods": [                                                                                                                              
                        "GET",                                                                                                                                              
                        "POST",                                                                                                                                             
                        "PUT"                                                                                                                                               
                ]                                                                                                                                                           
        },
...                                                                                 

        "swarm": {                                                                                                                                                          
                "address": "/ip4/0.0.0.0/tcp/6000"                                                                                                                          
        },                                   ^^^^ change this to a different port/value                                                   
...
```

Once you're done updating the config from the default values, launch the daemon (running in the background is optional; otherwise open a new terminal):

```
go-filecoin daemon --repodir=$FCRD &
```

Get the address of Node 2:

```
# you can also run commands without specifying --repodir if you set FIL_PATH
$ go-filecoin id --repodir=$FCRD
    {
   	"Addresses": [
   		"/ip4/127.0.0.1/tcp/6001/ipfs/QmXcUJ7YoFQEY7w8bpxuFvQtY9VHUkYfx6AZW6Bi2MDFbs",
   		"/ip4/172.16.200.201/tcp/6001/ipfs/QmXcUJ7YoFQEY7w8bpxuFvQtY9VHUkYfx6AZW6Bi2MDFbs"
   	],
   	"ID": "QmXcUJ7YoFQEY7w8bpxuFvQtY9VHUkYfx6AZW6Bi2MDFbs",
   	"AgentVersion": "",
   	"ProtocolVersion": "",
   	"PublicKey": ""
   }
$ export NODE2_ADDR=/ip4/127.0.0.1/tcp/6001/ipfs/QmXcUJ7YoFQEY7w8bpxuFvQtY9VHUkYfx6AZW6Bi2MDFbs    
```

Now connect Node 2 to Node 1 using the address retrieved for Node 1:

```
go-filecoin swarm connect $NODE1_ADDR --repodir=$FCRD
```

You should be able to see who connected peers are:

```
# Peers of node 1
go-filecoin swarm peers
/ip4/127.0.0.1/tcp/6001/ipfs/QmXcUJ7YoFQEY7w8bpxuFvQtY9VHUkYfx6AZW6Bi2MDFbs

# Peers of node 2
# you can also run commands without specifying --repodir if you set FIL_PATH
go-filecoin swarm peers --repodir=$FCRD
/ip4/127.0.0.1/tcp/6000/ipfs/QmVk7A2vEBFr9GyKyQ3wvDmTWj8M4H3jubHUDc3CktdoXL
```

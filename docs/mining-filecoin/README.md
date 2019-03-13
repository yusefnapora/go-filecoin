<!---
title: "Mining Filecoin"
weight: 20
pre: "<i class='fas fa-fw fa-gem'></i> "
--->


This guide provides an overview of how mining works, and a step-by-step of how to mine on the Filecoin network.

## Table of contents

* [What is mining?](#what-is-mining)
* [Start mining](#start-mining)
* [Explore mined blocks](#explore-mined-blocks)
* [Create an ask](#create-an-ask)
* [Accept a deal and get paid](#accept-a-deal-and-get-paid)
* [Stop mining](#stop-mining)

## <div id="what-is-mining" />What is mining?

In most blockchain protocols, "miners" are the participants on the network that do the work necessary to keep the blockchain valid and secure. For providing these services, miners are compensated in the native cryptocurrency. The term "miner" emerged to compare the work of securing blockchains with that of gold miners who expended resources to expand the supply of gold.
<!--
One key difference between Proof-of-Work systems (such as Bitcoin) and Filecoin is that Filecoin is designed to generate a native token as. More specifically, in the case of Filecoin, miners secure the network by computing proofs of storage and the overall purpose of the network is for miners to provide storage to other users. Contrast this to Bitcoin, in which miners secure the network by computing wasteful proofs of work, while the overall purposes of the network are transactions and store-of-value.-->

The Filecoin network will have multiple types of miners:
* Storage miners
* Retrieval miners
* Repair miners

In the current implementation, we focus mostly on storage miners. A storage miner sells storage capacity in exchange for filecoin.

## Start mining

After you `init` the Filecoin node and run the go-filecoin daemon, you can create a miner. By default, Filecoin nodes are not set up to mine. (See [Getting Started](/getting-started) for how to initialize a Filecoin node).

1. Create a miner that pledges 10 sectors (currently 256 MiB each) of storage and 100 FIL as collateral, with a message gas price of 0 FIL/unit and limit of 1000 gas units. When successful, it returns the miner address of the newly created miner.

    *Note: This step may take about a minute to process, but if it hangs for longer, double-check that `price` is less than `$YOUR_WALLET_BALANCE / LIMIT`.*

    ```shell
    go-filecoin miner create 10 100 --price=0 --limit=1000 --peerid `go-filecoin id | jq -r '.ID'`   # this may take a minute
    ```
1. Once the miner is created, we can run the following to start mining:
    ```shell
    go-filecoin mining start
    ```
:star2: Congrats, you are now mining filecoin! For now, you are mining blocks that contain activity on the network. Let's take a detour to explore these blocks. (Or, to begin mining your unused storage, skip directly to [Create an ask](#create-an-ask).)

## Explore mined blocks

You can explore the Filecoin blockchain using the [Filecoin Block Explorer](http://nightly.kittyhawk.wtf:8000/), or via the command line.  For example, let's get the `blockID` of the very first block of the blockchain. This is known as the _head_.

1. Show the chain head and copy the a block ID (there may be more than one):

    ```Shell    
    go-filecoin chain head # returns JSON including the <blockID> of the chain head
    ```

1. Then, view the contents of that block with `show block`:

    ```Shell    
    go-filecoin show block <blockID>
    ```

Many commands also support a `--enc=json` option for machine-readable output.

## Create an ask

In the Filecoin storage market, miners submit *ask orders* that provide some detail on their available storage space. Clients propose *deals* to miners for the files they want to store. Creating an _ask_ requires the following values:
1. your miner address (created earlier in this tutorial via [Start Mining](#start-mining))
1. your miner owner address (created automatically when you started the daemon)
1. the price at which you are willing to sell that much storage (in FIL/byte/block)
1. the number of blocks for which this asking price is valid
1. the price to pay for each gas unit consumed mining this message (in FIL)
1. the maximum number of gas units to be consumed by this message

Let's submit an ask order!

1. Get the miner address, and export it into a variable:
    ```Shell
    export MINER_ADDR=`go-filecoin config mining.minerAddress | tr -d \"`
    ```

1. Get the miner owner address, and export it into a variable:
    ```shell
    export MINER_OWNER_ADDR=`go-filecoin miner owner $MINER_ADDR`
    ```

1. Add an ask at a price of 0.000000001 FIL/byte/block, valid for 2880 blocks, with a message gas price of 0 FIL/unit and limit of 1000 gas units:
    ```shell
    go-filecoin miner set-price --from=$MINER_OWNER_ADDR --miner=$MINER_ADDR --price=0 --limit=1000 0.000000001 2880 # output: CID of the ask
    ```
1. Once you have placed the order, wait for the block containing your ask to be mined (around 30 seconds), and then check `client list-asks` to confirm that your ask has been added (look for your $MINER_ADDR):
    ```Shell
    go-filecoin client list-asks --enc=json | jq
    ```

## Accept a deal and get paid

Clients propose storage deals to miners who have enough storage and at a price that is lower than their willingness to pay.

A few notes on the current implementation as of Dec 2018:
- Currently, miners accept all deals that are proposed to them by clients with sufficient funds. Payment validation is done automatically so you don't have to take any action to accept a deal.
- Deal payments and payment channels are implemented. Thus, miners are periodically credited funds in a payment channel throughout the lifetime of the deal.

## Stop mining

If at any point you want to stop mining, you can always stop:
```Shell
go-filecoin mining stop
```
You can also remove all data associated with your Filecoin node instance:
```Shell
rm -rf ~/.filecoin
```

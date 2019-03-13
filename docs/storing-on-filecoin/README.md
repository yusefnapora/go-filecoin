---
title: "Storing on Filecoin"
weight: 30
pre: "<i class='fas fa-hdd'></i>&nbsp;"
---

This guide will show you how to store data on Filecoin by making storage deals with storage miners.

## Table of contents

* [How to store on Filecoin](#how-to-store-on-filecoin)
* [Test Filecoin with sample data](#test-filecoin-with-sample-data)
* [Import data](#import-data)
* [Propose a storage deal](#propose-a-storage-deal)
* [Send data and pay miner](#send-data-and-pay-miner)
* [Retrieve your data](#retrieve-your-data)

## How to store on Filecoin

Because Filecoin is a decentralized storage marketplace, storage clients (our term for anyone who wants to store data on the Filecoin network) propose storage deals with storage miners that they feel can meet their storage needs. Once a storage client has proposed a deal, the miner can choose to accept or reject that deal. The client then sends over the data, and the storage deal officially begins!

In order to run the remainder of these commands, make sure you have Filecoin installed and running. See [Getting Started](/getting-started) for more information.

## Test Filecoin with sample data

You can use your own data to test Filecoin. However, we recommend using small test files while you're still getting used to the different commands.

We put together a folder of [sample data](https://github.com/filecoin-project/sample-data) (images, GIFs, and text files) that you can get started with today. Download this data, unzip it, and save it wherever you like. In the examples below, we will assume that you have saved the repo to your Desktop (`~/Desktop`).

## Import data

Before storing data on the Filecoin network, you'll first need to import the data into your local Filecoin instance:

```Shell
# Start by creating a simple .txt file
echo "Hi my name is $USER"> hello.txt

# Import the created file into go-filecoin.
# RETURN VAL: Content ID (CID) of this file.
# You'll need this CID for future commands to
# view, store, or retrieve this data.
export CID=`go-filecoin client import ./hello.txt`

# View the imported file
# RETURN VAL: "Hi my name is user1"
go-filecoin client cat $CID
```
You can also import and view files from your computer:
```Shell
# Import an image from the sample data. Returns
# its CID which you'll need for the next step.
export CID=`go-filecoin client import ~/Desktop/sample-data-master/camel.jpg`

# View imported image
go-filecoin client cat $CID > image.png && open image.png
```

## Propose a storage deal

Once the data is imported, you can look through available _ask orders_ from miners. For now,  you'll manually choose one of them and propose a storage deal.

<!-- Clients can find storage miners that they would like to store with. When miners place ask orders on the network, they specify the amount of storage space they are pledging (in bytes) and the price at which they're willing to sell that amount of space (in FIL/MB/month). Clients find miners by looking through the asks of miners on the network and identifying those miners who are selling at a low enough price. Once a client finds a miner, she keeps track of the miner's ID and can propose a storage deal to the miner. -->

<!--# Check size of the imported file (in bytes)
go-filecoin client cat <CID> | wc -c
-->

To see all available asks, run this command:
```Shell
go-filecoin client list-asks --enc=json | jq
```
You should see a list of asks, nicely formatted as JSON objects:
```Shell
{
  "Miner": "fcqxvnl37zdv8clc26j6r43zn8md7tc2mrfx77vru",
  "Price": "2.5",
  "Expiry": 588,
  "ID": 0,
  "Error": null
}
```

Choose an ask from that list. For now, the biggest difference between miners will be price. In the future, miners will have other characteristics.

Now, let's propose a storage deal. You'll need to include these 4 parameters:
* `<miner>` address of the miner from `list-asks`
* `<data>` CID of the imported data that you want to store
* `<ask>` ID of the ask, also from `list-asks` (usually 0)
* `<duration>` how long you want to store (in # of ~30sec blocks). For example, storing for 1 day (2 blocks/min * 60 min/hr * 24 hr/day) = 2880 blocks.
```Shell
go-filecoin client propose-storage-deal <miner> <data> <ask> <duration>
```
If the deal is proposed correctly, it will automatically be accepted
and return `Status: accepted`. Currently, deals will fail if the data exceeds what that miner has available. Deals will also fail if your node's wallet balance is insufficient to cover the asking cost over the proposed data size and duration.

## Send data and pay miner

As part of negotiating the deal, your `propose-storage-deal` command will automatically move the funds necessary to cover the cost of the deal into a payment channel with the storage miner.  This payment channel will periodically release payments to the storage miner across the lifetime of the deal.

Your data is automatically transferred to a staging area via the underlying [bitswap](https://github.com/ipfs/specs/tree/master/bitswap) mechanism (implementation from IPFS) of Filecoin. Miners will now be required to store your files on their machines and begin the sealing/proving process. Under the hood, filecoin miners automatically fire off a sealing process at some configured interval (default 120 seconds) to take all data in the miner's staging area and seal it into the miner's pledged storage sectors.  Your deal status will move to "staged" at this point.  After sealing is complete and the miner posts its commitment to the chain your deal status will move to "posted".

## Retrieve your data

You have now successfully stored your data! Now, let's try retrieving it back.

<!-- We currently have two versions of retrieval that are implemented: -->

<!--
* Version 1 uses `go-filecoin client cat` to retrieve your data from the blockstore and is almost instantaneous.
* Version 2 is more representative of how Filecoin will work in the longer term. In its fullest implementation, data on Filecoin will have to be retrieved from sealed storage sectors. This is at least one mechanism by which we can ensure that the miner was correctly storing the data you sent. However, before data can be retrieved from a sealed sector, the sector needs to be unsealed. We have currently sped up the sealing and unsealing process so it will take around 5-10 minutes for you to retrieve a piece of data after you proposed a storage deal. Note that when Filecoin is working "for real," this number may be longer or shorter.

We show how to use both versions of retrieval here. Remember that you must keep track of the `<CID>` for each file you store on Filecoin.
-->

<!--
#### Version 1
```Shell
# Retrieve a text file
go-filecoin client cat <CID>

# Retrieve an image file
go-filecoin client cat <CID> > image.png && open image.png

# Retrieve a video file
go-filecoin client cat <CID> > movie.mp4 && open movie.mp4
```
#### Version 2
-->

Before retrieving, you must wait for the sector in which your data is to be stored is sealed (cryptographically proving storage). A deal starts out in state `accepted`, then changes to `posted` when a proof has been published. Check your deal state with:
```Shell
go-filecoin client query-storage-deal <dealID>
```

After your deal is `posted`, you can retrieve your data:
```Shell
# Retrieve your data, using the address of the
# miner you made a deal with and the <CID> of the data.
go-filecoin retrieval-client retrieve-piece <minerAddress> <CID> # can take a minute
```

Congratulations! You have now walked through several of the client commands available on Filecoin. Let us know if you have any questions by hitting us up on Matrix (#fil-chat) or the [community forum](https://discuss.filecoin.io/). You can also take a look at our [Troubleshooting & FAQs page](../troubleshooting-and-faq).

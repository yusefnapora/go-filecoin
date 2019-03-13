<!---
title: "Troubleshooting & FAQ"
weight: 80
pre: "<i class='fas fa-fw fa-question-circle'></i> "
--->

_Having trouble with `go-filecoin`? Here are some common errors (and their fixes), as well as answers to frequently asked questions._

_Note: This wiki focuses on `go-filecoin`. For questions about the Filecoin Project at large, see the [Filecoin Project FAQ](https://filecoin.io/faqs/)._

### Installing from binary

* **Error: “go-filecoin” is damaged and can’t be opened. You should move it to the Trash." (MacOS) <br />**
This is due to Mac's default protection settings. Go to `System Preferences > Security & Privacy > General`. Select `Allow apps downloaded from: Anywhere`. (Allowing the individual application doesn't work.)

### Downloading and building from source

* **Error when cloning: authenticity of host can’t be established<br />**
Make sure your Github account has [SSH keys added](https://help.github.com/articles/adding-a-new-ssh-key-to-your-github-account/).

* **Error: invalid active developer path (/Library/Developer/CommandLineTools), missing xcrun<br />**
Make sure you have [XCode](https://developer.apple.com/xcode/) installed.

* **Error when pulling: SharedFrameworks not found or unable to locate xcodebuild<br />**
This is due to an out-of-sync Xcode/command line tools installation (see [golang/go/issues/26073](https://github.com/golang/go/issues/26073)). You can try updating Xcode (if you’re running an outdated version) and then launch it to install the latest command-line tools.

* **$GOPATH or $PATH errors such as "panic: exec: gx-go executable file is not found in $PATH"<br />**
Make sure that the deps step ran successfully. If for example you are on a Mac and you installed go and rust with brew, GOPATH isn't explicitly set for you, the binaries will be installed in `/usr/local/opt/go/libexec/bin`, and also your default GOPATH is `~/go/`. Check your path and make sure all the go bin paths are there, and if not, add them:
    ```
    export PATH=${PATH}:/usr/local/opt/go/libexec/bin:~/go/bin
    ```
    Then retry `go run ./build/*.go deps` before retrying the build.

    If you are still stuck, or you see an error message with `$GOPATH` or `$PATH`, it could be an issue with your Go workspace setup. This [tutorial](https://www.ardanlabs.com/blog/2016/05/installing-go-and-your-workspace.html) may help.

* **Error: package 'go-filecoin' requires at least go version 1.11.1.**
...however, your gx-go binary was compiled with go1.10.3. Please update gx-go (or recompile with your current go compiler)
If you installed `ipfs` from homebrew, you may have the wrong version of `gx` and `gx-go` on your path. As with the previous point, check your `$PATH` and make sure all the go bin paths are there, and if not, add them, before your homebrew path.

* **Rust error while running build or deps<br />**
If you encounter a rust compiler error, i.e. ` cargo build --release --all --manifest-path proofs/rust-proofs/Cargo.toml' failed:` try updating rust to the latest version with `rustup update`.

* **Error while running install<br />**
If you encounter an error while running install, i.e. `/System/Library/Frameworks//Security.framework/Security are out of sync. Falling back to library file for linking`, this may be due to outdated installs. Reinstall Go via the [installer](https://golang.org/doc/install). Then remove filecoin (`rm -rf ./go-filecoin`) and [reinstall it](../getting-started).

* **Can't build OS X Mojave `fatal error: 'stdio.h' file not found`**
You may see this error if you are building everything from source and not installing anything with homebrew, for example, when building go-secp256k1:
    ```
    go get -u github.com/ipsn/go-secp256k1
    #github.com/ipsen/go-secp256k1
    In file included from ../../ipsn/go-secp256k1/secp256.go:17
    In file included from ././libsecp256k1/src/secp256k1.c:9:
    ../../ipsn/go-secp256k1/libsecp256k1/src/util.h:14:10: fatal error: 'stdlib.h' file not found
    #include <stdlib.h>
             ^~~~~~~~~~
    ```

    OS X Mojave moved the location of `stdlib.h` out of `/usr/include`.  This issue exists for other packages and there are several possible solutions suggested in this thread for [neovim issue #9050](https://github.com/neovim/neovim/issues/9050). Thanks to Filecoin community member *A_jinbin_filecoin.cn* for the link.

### Mining and deals

* **Miner create fails**
    * Make sure that your node has a valid wallet and the wallet has a nonzero balance.
    * Make sure that your node is connected to the swarm/network. If the daemon has been restarted, run `swarm connect`.

* **Transaction hangs (message never mined)**
    * If running two nodes on the same machine and you see something like `ERROR consensus.: Nonce too high: 5 0 <UnknownActor (0xc0137edda0); balance: 1000000; nonce: 0>`, you probably tried to create a miner at least once before. This updated the local nonce, even though the creation may have failed. Your nonce is too high for valid mining. *There is no known way around this except re-initing the node.*
    * If only running a single node, this problem is unobservable to you at present (the log message only appears in other nodes). If you suspect this, re-init your node.

* **Proposing deal fails: error sending proposal: context deadline exceeded**
    * Deals are proposed directly to the miner in question (off-chain), so your node needs a direct connection to the miner, e.g. to create a payment channel. The miner is offline or otherwise inaccessible to you. Try another miner.
    * The maximum piece (i.e. file) size must be less than the miner sector size, currently 256Mib.

* **Proposing deal fails: error creating payment: context deadline exceeded**
    * Our best guess is that a prior message failed to be mined, but increased your actor's nonce. Tracking in [#1936](https://github.com/filecoin-project/go-filecoin/issues/1936). You probably need to re-init your node.

### Daemon won't start

* If you see `Error: failed to load config file: failed to read config file at "~/.filecoin/config.json": invalid checksum` when trying to start a daemon, check that the defaultAddress and miner.address are correct in config.json.

### Network
* **How do I connect to the network?<br />**
See `go-filecoin swarm --help`.  In order to connect to a particular network, you must have initialized your filecoin repodir with the right genesis.car file for that network.

### Upgrading

* **How do I upgrade my version of go-filecoin?<br />**
To upgrade `go-filecoin`, you will need to re-run the full download and build process in [README.md](https://github.com/filecoin-project/go-filecoin/blob/master/README.md). In the future, we plan to add automatic updating ([#8](https://github.com/filecoin-project/go-filecoin/issues/8)).

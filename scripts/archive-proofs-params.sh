#!/usr/bin/env bash

#large sector params
GZIP=-1 tar -czvf "filecoin-proof-params-large-sectors-$CIRCLE_TAG.tar.gz" -C $HOME/filecoin-proof-parameters v9-zigzag-proof-of-replication-52431242c129794fe51d373ae29953f2ff52abd94c78756e318ce45f3e4946d8 v9-zigzag-proof-of-replication-52431242c129794fe51d373ae29953f2ff52abd94c78756e318ce45f3e4946d8.vk

#small sector params
GZIP=-1 tar -czvf "filecoin-proof-params-small-sectors-$CIRCLE_TAG.tar.gz" -C $HOME/filecoin-proof-parameters v9-zigzag-proof-of-replication-f8b6b5b4f1015da3984944b4aef229b63ce950f65c7f41055a995718a452204d v9-zigzag-proof-of-replication-f8b6b5b4f1015da3984944b4aef229b63ce950f65c7f41055a995718a452204d.vk

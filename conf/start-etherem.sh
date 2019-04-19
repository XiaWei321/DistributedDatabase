#!/bin/bash
./geth --datadir chain  --rpc --rpcapi "db,web3,eth,net,personal" --rpcaddr "0.0.0.0" --rpccorsdomain "*" --nodiscover --port 16333 --rpcport 8545  console
#!/bin/bash

set -e
set -x

make install

certik unsafe-reset-all
rm -rf ~/.certik

certik init node0 --chain-id certikchain
# certikcli config chain-id certikchain
# certikcli config keyring-backend test
certik keys add jack
jack=$(certik keys show jack -a)
certik keys add mary --ledger
mary=$(certik keys show mary -a)
certik add-genesis-account $jack 1000000000uctk
certik add-genesis-account $mary 1000000000uctk
certik add-genesis-certifier $jack
certik add-genesis-certifier $mary
certik gentx jack 2000000uctk --chain-id certikchain
certik collect-gentxs
sed -i 's/enable = false/enable = true/g' ~/.certik/config/app.toml # enable rest server
certik start
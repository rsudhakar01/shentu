#!/bin/bash

make install

certikd unsafe-reset-all
rm -rf ~/.certikd
rm -rf ~/.certikcli

certikd init node0 --chain-id certikchain
certikcli config chain-id certikchain
certikcli config keyring-backend test

certikcli keys add jack
jack=$(certikcli keys show jack -a)
certikd add-genesis-account $jack 1000000000uctk --vesting-amount=200000000uctk --period=10 --num-periods=5

certikd gentx --name jack --amount 2000000uctk --keyring-backend test
certikd collect-gentxs

certikd start

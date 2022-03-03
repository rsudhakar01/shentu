#!/bin/bash

set -x
set -e

make install

binary=~/go/bin/certik # default

$binary unsafe-reset-all
rm -rf ~/.certik

$binary init test-node --chain-id test-chain

$binary keys add jack
jack=$($binary keys show jack -a)

$binary add-genesis-account $jack 100000000uctk

$binary gentx jack 20000000uctk --chain-id test-chain
$binary collect-gentxs

$binary start

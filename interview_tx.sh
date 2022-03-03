#!/bin/bash

set -x
set -e

binary=~/go/bin/certik # default

$binary tx interview lock-user 1 --from jack --chain-id test-chain -y
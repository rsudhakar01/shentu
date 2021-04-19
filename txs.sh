#!/bin/bash

set -e
set -x

jack=$(certik keys show jack -a)
echo
mary=$(certik keys show mary -a)
echo

jack_pubkey=$(certik tendermint show-validator)

certik tx cert certify-platform $jack_pubkey test-platform --from $jack --chain-id certikchain -y
#!/bin/bash

set -e
set -x

jack=$(certik keys show jack -a)
echo
mary=$(certik keys show mary -a)
echo

certik tx cert issue-certificate COMPILATION CCC --compiler A --bytecode-hash B --from $jack --chain-id certikchain -y
sleep 5
echo

certik tx cert issue-certificate AUDITING ccccc --from $mary --chain-id certikchain -y
sleep 5
echo

<<COMMENT1
certik tx cert issue-certificate PROOF ccccc --from $jack --chain-id certikchain -y
sleep 5
echo

certik tx cert issue-certificate ORACLEOPERATOR ccccc --from $jack --chain-id certikchain -y
sleep 5
echo

certik tx cert issue-certificate SHIELDPOOLCREATOR ccccc --from $jack --chain-id certikchain -y
sleep 5
echo

certik tx cert issue-certificate IDENTITY ccccc --from $jack --chain-id certikchain -y
sleep 5
echo

certik tx cert issue-certificate GENERAL ccccc --from $jack --chain-id certikchain -y
sleep 5
echo
COMMENT1

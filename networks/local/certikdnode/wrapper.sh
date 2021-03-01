#!/usr/bin/env sh

##
## Input parameters
##
BINARY=/certik/${BINARY:-certik}
ID=${ID:-0}
LOG=${LOG:-certik.log}

##
## Assert linux binary
##
if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'certik' E.g.: -e BINARY=certikd_my_test_version"
	exit 1
fi
BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"
if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

##
## Run binary with all parameters
##
export CERTIKDHOME="/certik/node${ID}/certik"

if [ -d "$(dirname "${CERTIKDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${CERTIKDHOME}" "$@" | tee "${CERTIKDHOME}/${LOG}"
else
  "${BINARY}" --home "${CERTIKDHOME}" "$@"
fi

#!/usr/bin/env sh

##
## Input parameters
##
ID=${ID:-0}
LOG=${LOG:-brczerod.log}

##
## Run binary with all parameters
##
export brczeroDHOME="/brczerod/node${ID}/brczerod"

if [ -d "$(dirname "${brczeroDHOME}"/"${LOG}")" ]; then
  brczerod --chain-id brczero-1 --home "${brczeroDHOME}" "$@" | tee "${brczeroDHOME}/${LOG}"
else
  brczerod --chain-id brczero-1 --home "${brczeroDHOME}" "$@"
fi


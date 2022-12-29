#!/bin/bash

ARGS=$(./bin/getargs $@)

VERBOSE=$(echo $ARGS | jq '.v')

echo "Verbose: $VERBOSE"

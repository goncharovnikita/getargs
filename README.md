# getargs

Command line arguments parser for scripts

## Example

In your script bypass arguments to `getargs`

script.sh:

```bash
ARGS=$(getargs $@)

VERBOSE=$(echo $ARGS | jq '.v')

echo "Verbose: $VERBOSE"
```

Then `ARGS` variable will contain json formatted string with all flags in key-map format:

```
$ ./script.sh -v
Verbose: true

#ARGS='{"v": true}'
#VERBOSE=true
```

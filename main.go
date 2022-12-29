package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	l := log.New(os.Stdout, "", 0)

	result := parseArgs(os.Args[1:])

	if err := json.NewEncoder(os.Stdout).Encode(&result); err != nil {
		l.Fatal(err)
	}
}

type parserState int

const (
	parserStateFlagName parserState = iota
	parserStateFlagValue
)

func parseArgs(args []string) map[string]interface{} {
	result := make(map[string]interface{})
	state := parserStateFlagName

	var flagName string

	for _, arg := range args {
		switch state {
		case parserStateFlagName:
			{
				if isFlagName(arg) {
					flagName = stripFlagPrefix(arg)
					state = parserStateFlagValue
				}
			}

		case parserStateFlagValue:
			{
				if isFlagName(arg) {
					result[flagName] = true
					flagName = stripFlagPrefix(arg)
				} else {
					result[flagName] = arg
					state = parserStateFlagName
				}
			}
		}
	}

	return result
}

func isFlagName(s string) bool {
	return strings.HasPrefix(s, "--") || strings.HasPrefix(s, "-")
}

func stripFlagPrefix(s string) string {
	return strings.TrimLeft(strings.TrimLeft(s, "--"), "-")
}

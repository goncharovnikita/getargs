package main

import (
	"reflect"
	"testing"
)

func Test_parseArgs(t *testing.T) {
	cases := []struct {
		name     string
		args     []string
		expected map[string]interface{}
	}{
		{
			name: "parse args",
			args: []string{
				"-v",
				"--config",
				"./some-file.yaml",
			},
			expected: map[string]interface{}{
				"v":      true,
				"config": "./some-file.yaml",
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			actual := parseArgs(c.args)
			if !reflect.DeepEqual(c.expected, actual) {
				t.Logf("expected: %+v", c.expected)
				t.Log("p-----p")
				t.Logf("actual: %+v", actual)
				t.Fatalf("expected value not equal to actual")
			}
		})
	}
}

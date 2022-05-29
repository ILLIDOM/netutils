package multicast

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteMACCommand(t *testing.T) {
	tc := []struct {
		args        []string
		expectedOut string
	}{
		{
			args:        []string{"01:00:5e:00:00:01"},
			expectedOut: "224.0.0.1\n",
		},
		{
			args:        []string{"01:00:00:AA:BB:CC"},
			expectedOut: "No valid Multicast MAC provided: remember Multicast MAC addresses start with 01:00:5e\n",
		},
	}

	for _, tc := range tc {
		cmd := newMacCommand()
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.SetArgs(tc.args)
		cmd.Execute()
		out, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, tc.expectedOut, string(out))
	}

}

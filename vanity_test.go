package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

var nilString []string

func TestVanityLabels(t *testing.T) {
	v := vanity{
		ShortPackage: "pkg/test",
		Package:      "github.com/pkg/test",
		Name:         "test",
	}

	buf := bytes.NewBuffer([]byte{})
	err := vanityLabels(buf, v)
	require.NoError(t, err)
	golden, err := ioutil.ReadFile("hack/vanity.gold")
	require.Equal(t, golden, buf.Bytes())
}

func TestWriteTravis(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	err := writeTravis(buf)
	require.NoError(t, err)
	golden, err := ioutil.ReadFile("hack/travis.gold")
	require.NoError(t, err)
	require.Equal(t, golden, buf.Bytes())
}

func TestParse(t *testing.T) {
	testcases := []struct {
		Test   []string
		Result []string
		Err    bool
	}{
		{
			Test:   []string{"testcase1", "github.com/pkg/test"},
			Result: []string{"github.com", "pkg", "test"},
			Err:    false,
		},
		{
			Test:   []string{"testcase2"},
			Result: nilString,
			Err:    true,
		},
		{
			Test:   []string{"testcase3", "arg1", "arg2-cause-failure"},
			Result: nilString,
			Err:    true,
		},
		{
			Test:   []string{"testcase4", "github.com"},
			Result: []string{"github.com"},
			Err:    true,
		},
	}

	for i, test := range testcases {
		r, err := parse(test.Test)
		require.Equal(t, test.Result, r, "Failed test case %d", i+1)
		if test.Err {
			require.Error(t, err, "Failed test case %d", i+1)
		} else {
			require.NoError(t, err, "Failed test case %d", i+1)
		}
	}
}

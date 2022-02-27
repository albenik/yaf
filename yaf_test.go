package yaf_test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCMD(t *testing.T) {
	t.Run("Include", func(t *testing.T) {
		outStr := new(strings.Builder)
		errStr := new(strings.Builder)

		cmd := exec.Command("go", "run", "github.com/albenik/yaf/cmd/kyaf", "-i", "operator.victoriametrics.com/v1beta1:VMAgent")
		cmd.Stdin = strings.NewReader(testResourceInput)
		cmd.Stdout = outStr
		cmd.Stderr = errStr
		err := cmd.Run()

		if assert.NoError(t, err) {
			assert.Equal(t, testResourceOutput, outStr.String())
			assert.Empty(t, errStr.String())
		}
	})

	t.Run("Exclude", func(t *testing.T) {
		outStr := new(strings.Builder)
		errStr := new(strings.Builder)

		cmd := exec.Command("go", "run", "github.com/albenik/yaf/cmd/kyaf", "-x", "operator.victoriametrics.com/v1beta1:VMServiceScrape")
		cmd.Stdin = strings.NewReader(testResourceInput)
		cmd.Stdout = outStr
		cmd.Stderr = errStr
		err := cmd.Run()

		if assert.NoError(t, err) {
			assert.Equal(t, testResourceOutput, outStr.String())
			assert.Empty(t, errStr.String())
		}
	})

	t.Run("Error", func(t *testing.T) {
		outStr := new(strings.Builder)
		errStr := new(strings.Builder)

		cmd := exec.Command("go", "run", "github.com/albenik/yaf/cmd/kyaf", "-x", "invalid")
		cmd.Stdin = strings.NewReader(testResourceInput)
		cmd.Stdout = outStr
		cmd.Stderr = errStr
		err := cmd.Run()

		assert.Error(t, err)
		assert.Empty(t, outStr.String())
		assert.Equal(t, "ERROR: invalid filter string\nexit status 1\n", errStr.String())
	})
}

package yaf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/albenik/yaf"
)

func TestParseArgument(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		filters, err := yaf.ParseArgument([]string{
			"example.net/v1:Kind1",
			"example.net/v1:Kind2,Kind3",
			"example.net/v2beta1:Kind1,Kind2,Kind3,",
		})
		if assert.NoError(t, err) {
			assert.Equal(t, map[string]map[string]struct{}{
				"example.net/v1":      {"Kind1": {}, "Kind2": {}, "Kind3": {}},
				"example.net/v2beta1": {"Kind1": {}, "Kind2": {}, "Kind3": {}},
			}, filters)
		}
	})

	t.Run("Incorrect", func(t *testing.T) {
		_, err := yaf.ParseArgument([]string{
			"",
			":",
			"example1.net/v1:",
			":Kind1",
			"Kind1",
			"example.net/v1beta1:Kind1,,Kind2",
		})
		assert.ErrorIs(t, err, yaf.ErrInvalidFilterString)
	})

	t.Run("Mixed", func(t *testing.T) {
		_, err := yaf.ParseArgument([]string{
			"example1.net/v1:Kind1",
			"invalid",
		})
		assert.ErrorIs(t, err, yaf.ErrInvalidFilterString)
	})
}

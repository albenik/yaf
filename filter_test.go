package yaf_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/albenik/yaf"
)

const testResourceInput = `apiVersion: operator.victoriametrics.com/v1beta1
kind: VMServiceScrape
metadata:
  labels:
    control-plane: controller-manager
  name: vm-operator-controller-manager-metrics-monitor
  namespace: operators
spec:
  endpoints:
  - path: /metrics
    port: https
  selector:
    matchLabels:
      control-plane: controller-manager
---
apiVersion: operator.victoriametrics.com/v1beta1
kind: VMAgent
metadata:
  labels:
    control-plane: controller-manager
  name: vm-operator-controller-manager-metrics-monitor
  namespace: operators
spec:
  endpoints:
  - path: /metrics
    port: https
  selector:
    matchLabels:
      control-plane: controller-manager
`

const testResourceOutput = `apiVersion: operator.victoriametrics.com/v1beta1
kind: VMAgent
metadata:
  labels:
    control-plane: controller-manager
  name: vm-operator-controller-manager-metrics-monitor
  namespace: operators
spec:
  endpoints:
  - path: /metrics
    port: https
  selector:
    matchLabels:
      control-plane: controller-manager
`

func TestFilter(t *testing.T) {
	t.Run("Exclude", func(t *testing.T) {
		s := new(strings.Builder)
		err := yaf.Filter(s, strings.NewReader(testResourceInput),
			map[string]map[string]struct{}{
				"operator.victoriametrics.com/v1beta1": {
					"VMServiceScrape": struct{}{},
				},
			},
			true,
		)
		require.NoError(t, err)
		require.Equal(t, testResourceOutput, s.String())
	})

	t.Run("Include", func(t *testing.T) {
		s := new(strings.Builder)
		err := yaf.Filter(s, strings.NewReader(testResourceInput),
			map[string]map[string]struct{}{
				"operator.victoriametrics.com/v1beta1": {
					"VMAgent": struct{}{},
				},
			},
			false,
		)
		require.NoError(t, err)
		require.Equal(t, testResourceOutput, s.String())
	})
}

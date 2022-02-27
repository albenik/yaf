package yaf

import (
	"errors"
	"strings"
)

var ErrInvalidFilterString = errors.New("invalid filter string")

func ParseArgument(strs []string) (map[string]map[string]struct{}, error) {
	m := make(map[string]map[string]struct{})

	for _, s := range strs {
		parts := strings.SplitN(s, ":", 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return nil, ErrInvalidFilterString
		}

		var (
			kindsMap map[string]struct{}
			ok       bool
		)
		kinds := strings.Split(parts[1], ",")

		if kindsMap, ok = m[parts[0]]; !ok {
			kindsMap = make(map[string]struct{}, len(kinds))
			m[parts[0]] = kindsMap
		}
		for i, k := range kinds {
			if k == "" {
				if i != len(kinds)-1 {
					return nil, ErrInvalidFilterString
				}
				break
			}
			kindsMap[k] = struct{}{}
		}
		m[parts[0]] = kindsMap
	}

	return m, nil
}

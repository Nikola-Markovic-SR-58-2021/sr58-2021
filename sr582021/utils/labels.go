package utils

import (
	"errors"
	"strings"
)

func ParseLabels(raw string) (map[string]string, error) {
	labels := make(map[string]string)
	if raw == "" {
		return labels, nil
	}
	for _, pair := range strings.Split(raw, ";") {
		parts := strings.SplitN(strings.TrimSpace(pair), ":", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid label format: " + pair)
		}
		labels[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}
	return labels, nil
}

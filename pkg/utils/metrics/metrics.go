package metrics

import (
	"sort"
	"strings"
)

func BuildMetricName(base string, labels map[string]string) string {
	if labels == nil || len(labels) == 0 {
		return base
	}

	pairs := make([]string, 0, len(labels))
	for k, v := range labels {
		pairs = append(pairs, k+"="+v)
	}
	sort.Strings(pairs)

	return base + "[" + strings.Join(pairs, ",") + "]"
}
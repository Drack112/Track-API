package observability

import (
	"net/url"
	"strings"
)

func ParseHeaders(headersStr string) map[string]string {
	result := make(map[string]string)
	if headersStr == "" {
		return result
	}

	pairs := strings.Split(headersStr, ",")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			result[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	return result
}

func NormalizeEndpoint(endpoint string) (string, error) {
	if strings.TrimSpace(endpoint) == "" {
		return "", nil
	}

	e := strings.TrimSpace(endpoint)
	if !strings.HasPrefix(e, "http://") && !strings.HasPrefix(e, "https://") {
		e = "http://" + e
	}
	u, err := url.Parse(e)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

package headers

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type Headers map[string]string

const crlf = "\r\n"

var allowed = regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.\^_` + "`" + `|~]+$`)

func NewHeaders() Headers {
	return map[string]string{}
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		return 0, false, nil
	}
	if idx == 0 {
		return 2, true, nil
	}

	parts := bytes.SplitN(data[:idx], []byte(":"), 2)
	namePart := string(parts[0])

	if namePart != strings.TrimRight(namePart, " ") {
		return 0, false, fmt.Errorf("invalid name for header: %s", namePart)
	}

	valuePart := bytes.TrimSpace(parts[1])
	namePart = strings.TrimSpace(namePart)
	if ok := Allowed(namePart); !ok {
		return 0, false, fmt.Errorf("name contains unallowed characters: %s", namePart)
	}
	namePart = strings.ToLower(namePart)
	if value, exists := h[namePart]; !exists {
		h.Set(namePart, string(valuePart))
	} else {
		newValuePart := value + ", " + string(valuePart)
		h.Set(namePart, newValuePart)
	}

	return idx + 2, false, nil
}

func (h Headers) Set(key, value string) {
	h[key] = value
}

func Allowed(str string) bool {
	return allowed.MatchString(str)
}

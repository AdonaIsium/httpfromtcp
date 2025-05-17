package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine
	// Headers     map[string]string
	// Body        []byte
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (Request, error) {
	req, err := io.ReadAll(reader)
	if err != nil {
		fmt.Printf("Error reading request: %v", err)
		return Request{}, err
	}

	reqLine, err := parseRequestLine(req)
	if err != nil {
		return Request{}, err
	}

	var request Request
	request.RequestLine = reqLine

	return request, nil

}

func parseRequestLine(req []byte) (RequestLine, error) {
	strReq := string(req)

	splitReq := strings.Split(strReq, "\r\n")

	smaller := strings.Split(splitReq[0], " ")

	if len(smaller) != 3 {
		return RequestLine{}, errors.New("request line does not contain expected number of parts")
	}
	if ok := isAlpha(smaller[0]); !ok {
		return RequestLine{}, errors.New("method must contain only alphabetic characters")
	}
	var requestLine RequestLine
	requestLine.RequestTarget = smaller[1]
	requestLine.Method = smaller[0]
	httpVersion := strings.Split(smaller[2], "/")
	if httpVersion[1] != "1.1" {
		return RequestLine{}, errors.New("only http 1.1 is supported at this time")
	}
	requestLine.HttpVersion = httpVersion[1]

	return requestLine, nil
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return len(s) > 0
}

package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

const CRLF = "\r\n"

func parseRequestLine(line []byte) (*RequestLine, error) {
	parts := bytes.Split(line, []byte(" "))
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request-line: expected 3 parts, got %d", len(parts))
	}

	method, target, version := parts[0], parts[1], parts[2]

	if !bytes.Equal(bytes.ToUpper(method), method) {
		return nil, fmt.Errorf("invalid method: %q", method)
	}

	if !bytes.HasPrefix(version, []byte("HTTP/")) {
		return nil, fmt.Errorf("invalid HTTP version (missing HTTP/ prefix): %q", version)
	}

	v := bytes.TrimPrefix(version, []byte("HTTP/"))
	if !bytes.Equal(v, []byte("1.1")) {
		return nil, fmt.Errorf("unsupported HTTP version: %q", v)
	}

	return &RequestLine{HttpVersion: string(v), RequestTarget: string(target), Method: string(method)}, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("read request: %w", err)
	}

	i := bytes.Index(data, []byte(CRLF))
	if i < 0 {
		return nil, errors.New("malformed request: missing CRLF after request-line")
	}

	rl, err := parseRequestLine(data[:i])
	if err != nil {
		return nil, err
	}

	return &Request{RequestLine: *rl}, nil
}

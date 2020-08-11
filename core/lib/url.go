package lib

import (
	"strings"
)

// URLPath represent url path of request
// it is useful for checking handler paths
type URLPath struct {
	parts      []urlPart
	raw        string
	paramCount int
}

// Raw return raw string of url path
func (p URLPath) Raw() string {
	return p.raw
}

// Check checks URLPath satisfy url or not
func (p URLPath) Check(url string) bool {
	url = cleanURLTrials(url)
	parts := splitPath(url)
	if len(parts) != len(p.parts) {
		return false
	}
	for i, v := range parts {
		if p.parts[i].partType == parameterURLPart {
			continue
		}
		if p.parts[i].content != v {
			return false
		}
	}
	return true
}

// GetURLParams parse url parameter from urlPart
func (p URLPath) GetURLParams(url string) map[string]string {
	m := make(map[string]string)
	parts := splitPath(url)
	for i, v := range parts {
		if p.parts[i].partType == parameterURLPart {
			m[p.parts[i].content] = v
		}
	}
	return m
}

// NewURLPath returns new url path
func NewURLPath(s string) *URLPath {
	s = cleanURLTrials(s)
	sarr := strings.Split(s, "/")
	parts := make([]urlPart, len(sarr))
	pcount := 0
	for i, v := range sarr {
		if v == "" {
			continue
		}
		part := newURLPart(v)
		if part.partType == parameterURLPart {
			pcount++
		}
		parts[i] = part
	}
	return &URLPath{
		parts:      parts,
		raw:        s,
		paramCount: pcount,
	}
}

const (
	parameterURLPart = iota
	constantURLPart
)

type urlPart struct {
	partType int
	content  string
}

func newURLPart(part string) urlPart {
	pt := constantURLPart
	if strings.HasPrefix(part, "_") {
		pt = parameterURLPart
		part = part[1:]
	}
	return urlPart{
		partType: pt,
		content:  part,
	}
}

func splitPath(path string) []string {
	sarr := strings.Split(path, "/")
	if len(sarr) > 1 && sarr[0] == "" {
		sarr = sarr[1:]
	}
	li := len(sarr) - 1
	if len(sarr) > 0 && sarr[li] == "" {
		sarr = sarr[:li]
	}
	return sarr
}

func cleanURLTrials(url string) string {
	if len(url) > 0 && url[0] == '/' {
		url = url[1:]
	}
	if len(url) > 0 && url[len(url)-1] == '/' {
		url = url[:(len(url) - 1)]
	}
	return url
}

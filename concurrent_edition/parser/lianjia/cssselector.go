package lianjia

import (
	"bytes"
	"io"
	"regexp"
	"strings"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

func cssSelect(contents []byte, sel *css.Selector) (string, error) {
	node, err := html.Parse(strings.NewReader(string(contents)))
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	w := io.Writer(&buf)
	for _, ele := range sel.Select(node) {
		html.Render(w, ele)
	}
	return buf.String(), nil
}

func extractString(contents []byte, sel *css.Selector, re *regexp.Regexp) ([][]string, error) {
	sBasic, err := cssSelect(contents, sel)
	var s [][]string
	if err != nil {
		return s, err
	}

	return re.FindAllStringSubmatch(sBasic, -1), nil
}

func getInfoString(contents []byte, sel *css.Selector, re *regexp.Regexp) (string, error) {
	s, err := extractString(contents, sel, re)
	if err != nil {
		return "", err
	}
	if len(s) == 1 {
		return s[0][1], nil
	} else {
		Q := ""
		for _, v := range s {
			Q = Q + " " + v[1]
		}
		return Q, nil
	}
}

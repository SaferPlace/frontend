package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	invalidConfig = `
invalid:
  - invalid
  invalid
`
	validConfig = `
home:
  blurb: valid
  `
)

func TestLoadLanguage(t *testing.T) {
	testCases := []struct {
		lang     string
		config   string
		expected string
		errWord  string
	}{
		{
			"",
			"",
			"",
			"load",
		}, {
			"invalid",
			invalidConfig,
			"",
			"parse",
		}, {
			"valid",
			validConfig,
			"valid",
			"",
		},
	}

	for _, tc := range testCases {
		createTestLang(tc.lang, tc.config)
		defer removeTestLang(tc.lang)

		l, err := loadLanguage(tc.lang)
		if err != nil {
			if !strings.Contains(err.Error(), tc.errWord) {
				t.Errorf("%s: expected error with word %s, have error %v",
					tc.lang, tc.errWord, err)
			}
		}
		if l.Home.Blurb != tc.expected {
			t.Errorf("%s: expected .Home.Blurb to be %s, have %s",
				tc.lang, tc.expected, l.Home.Blurb)
		}
	}
}

func createTestLang(lang, data string) {
	// special case when lang == "" for tests
	if lang == "" {
		return
	}
	ioutil.WriteFile(fmt.Sprintf("./languages/%s.yaml", lang), []byte(data), 0777)
}

func removeTestLang(lang string) {
	os.Remove(fmt.Sprintf("./languages/%s.yaml", lang))
}

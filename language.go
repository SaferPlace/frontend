package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type language struct {
	Header struct {
		Lang              string // This field is filled at runtime
		Input             string // This field is filled at runtime
		Title             string `yaml:"title"`
		AddressFieldLabel string `yaml:"address_field_label"`
		AboutLink         string `yaml:"about_link"`
	}
	Home struct {
		Blurb string `yaml:"blurb"`
	}
	About struct {
		AboutTitle   string `yaml:"about_title"`
		WhatQuestion string `yaml:"what_question"`
		WhatAnswer   string `yaml:"what_answer"`
		HowQuestion  string `yaml:"how_question"`
		HowAnswer    string `yaml:"how_answer"`
		QandA        []struct {
			Question string `yaml:"question"`
			Answer   string `yaml:"answer"`
		} `yaml:"q_and_a"`
	}
}

func loadLanguage(lang string) (language, error) {
	l := language{}
	data, err := ioutil.ReadFile(fmt.Sprintf("./languages/%s.yaml", lang))
	if err != nil {
		return l, fmt.Errorf("Unable to load language %s: %v", lang, err)
	}

	if err = yaml.Unmarshal(data, &l); err != nil {
		return l, fmt.Errorf("Unable to parse YAML for language %s: %v", lang, err)
	}

	return l, nil
}

func languages() map[string]language {
	langPaths, err := filepath.Glob("languages/*.yaml")
	if err != nil {
		panic(fmt.Sprintf("Unable to Glob languages: %v", err))
	}

	langs := make(map[string]language)
	for _, lang := range langPaths {
		lang = strings.TrimPrefix(lang, "languages/")
		lang = strings.TrimSuffix(lang, ".yaml")

		langs[lang], err = loadLanguage(lang)
		if err != nil {
			panic(err)
		}
		log.Printf("Loaded \"%s\" language", lang)
	}
	return langs
}

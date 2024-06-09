package i18n

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"sync"
)

type Options map[string]interface{}

var cache map[string]map[string]interface{}
var mu sync.RWMutex

const defaultLang = "en"

func flatten(data map[string]interface{}) map[string]interface{} {
	flatted := make(map[string]interface{})

	var aux func(obj map[string]interface{}, base string)
	aux = func(obj map[string]interface{}, base string) {
		for key := range obj {
			newBase := key
			if base != "" {
				newBase = fmt.Sprintf("%s.%s", base, key)
			}

			switch value := obj[key].(type) {
			case string:
				flatted[newBase] = value
			case map[string]interface{}:
				var keys []string
				for key := range value {
					keys = append(keys, key)
				}

				slices.Sort(keys)

				if strings.Join(keys, " ") == "many one" {
					flatted[newBase] = value
				} else {
					aux(value, newBase)
				}
			}
		}
	}
	aux(data, "")

	return flatted
}

func Init(localesPath string) (err error) {
	mu.Lock()
	defer mu.Unlock()

	cache = make(map[string]map[string]interface{})

	var languages []fs.DirEntry

	if languages, err = os.ReadDir(localesPath); err != nil {
		log.Println("error reading locales path:", err)
		return
	}

	for _, langFile := range languages {
		lang := strings.TrimSuffix(filepath.Base(langFile.Name()), filepath.Ext(langFile.Name()))
		cache[lang] = make(map[string]interface{})

		var langFullPath string
		if langFullPath, err = filepath.Abs(path.Join(localesPath, langFile.Name())); err != nil {
			log.Println("error converting lang path to absolute", err)
			return
		}

		var data []byte
		if data, err = os.ReadFile(langFullPath); err != nil {
			log.Println("error reading lang file", err)
			return
		}

		var content map[string]interface{}
		if err = json.Unmarshal(data, &content); err != nil {
			log.Println("error parsing lang file", err)
			return
		}

		cache[lang] = flatten(content)

	}

	return
}

func T(key string, lang string) (text string, err error) {
	return T2(key, lang, Options{})
}

func T2(key string, lang string, options Options) (text string, err error) {
	mu.RLock()
	defer mu.RUnlock()

	if _, ok := cache[lang]; !ok {
		fmt.Printf("requested language (%s) is not supported\n", lang)
		lang = defaultLang
	}

	switch value := cache[lang][key].(type) {
	case string:
		text = value
	case map[string]interface{}:
		var ok bool
		var opts interface{}
		if opts, ok = options["count"]; !ok {
			err = fmt.Errorf("expected options[\"count\"] in key %s", key)
			return
		}

		var count int
		if count, ok = opts.(int); !ok {
			err = fmt.Errorf("expected int type options[\"count\"] in key %s", key)
			return
		}

		if count == 1 {
			var one string
			if one, ok = value["one"].(string); !ok {
				err = fmt.Errorf("expected a string in key %s.one", key)
				return
			}

			text = one
		} else {
			var many string
			if many, ok = value["many"].(string); !ok {
				err = fmt.Errorf("expected a string in key %s.many", key)
				return
			}

			text = many
		}
	}

	for key, value := range options {
		var re *regexp.Regexp
		if re, err = regexp.Compile(fmt.Sprintf("({{%s}})", key)); err != nil {
			fmt.Println("regex could not be compiled", err)
			return
		}

		text = re.ReplaceAllString(text, fmt.Sprintf("%v", value))
	}

	return
}

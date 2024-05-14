package i18n

import (
	"alice/env"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/telebot.v3"
)

type Locale struct {
	Ok            string `json:"ok"`
	StartMessagee string `json:"start_message"`
}

type LocaleKey string

const (
	Ok           LocaleKey = "ok"
	StartMessage LocaleKey = "start_message"
)

var locales = make(map[string]Locale)
var bundle = make(map[string]map[LocaleKey]string)

func loadLocale(path string) (l Locale, err error) {
	var content []byte

	if content, err = os.ReadFile(path); err != nil {
		return
	}

	var lt = &Locale{}
	if err = json.Unmarshal(content, lt); err != nil {
		return
	}

	return *lt, nil
}

func Load(p string) (err error) {

	var files []os.DirEntry
	if files, err = os.ReadDir(p); err != nil {
		return
	}

	for _, f := range files {
		fileName := f.Name()
		filePath := fmt.Sprintf("%s/%s", p, fileName)

		var l Locale
		if l, err = loadLocale(filePath); err != nil {
			return
		}

		parts := strings.Split(fileName, ".")
		if len(parts) != 2 || parts[0] == "" {
			err = fmt.Errorf("invalid locale file name: %s", fileName)
		}

		locales[parts[0]] = l
	}

	return
}

func GetLocale(lang string) Locale {
	if locale := locales[lang]; locale.Ok != "" {
		return locale
	}

	return locales["en"]
}

func GetText(c telebot.Context, locale LocaleKey) (text string) {
	sender := c.Sender()

	replacer := strings.NewReplacer(
		"{me}", env.ClientName,
		"{username}", sender.Username,
	)

	text = replacer.Replace(bundle[sender.LanguageCode][locale])

	return
}

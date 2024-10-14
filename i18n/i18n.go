package i18n

import (
	"embed"
	"fmt"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type LanguageCode = string
type Options = map[string]any

const defaultLocale = "en"

var Locales = make(map[LanguageCode]*Locale)

//go:embed locales/*.toml
var content embed.FS

func Init() error {
	entries, err := content.ReadDir("locales")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		name := entry.Name()

		if entry.IsDir() {
			continue
		}

		file, err := content.Open("locales/" + name)
		if err != nil {
			return err
		}

		defer file.Close()

		locale := new(Locale)

		decoder := toml.NewDecoder(file)
		if err := decoder.Decode(locale); err != nil {
			return err
		}

		var languageCode string

		dotIndex := strings.LastIndex(name, ".")

		if dotIndex == -1 {
			languageCode = name
		} else {
			languageCode = name[:dotIndex]
		}

		Locales[languageCode] = locale
	}

	return nil
}

func GetLocale(code LanguageCode) (*Locale, error) {
	locale, ok := Locales[code]
	if ok {
		return locale, nil
	}

	locale, ok = Locales[defaultLocale]
	if ok {
		return locale, nil
	}

	return nil, ErrNotFound
}

func WithOptions(text string, opts Options) string {
	for key, value := range opts {
		old := fmt.Sprintf("{{%s}}", key)
		new := fmt.Sprintf("%s", value)

		text = strings.ReplaceAll(text, old, new)
	}

	return text
}

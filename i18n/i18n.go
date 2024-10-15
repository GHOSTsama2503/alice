package i18n

import (
	"embed"
	"fmt"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type (
	LocalesMap   = map[LanguageCode]*Locale
	LanguageCode = string
	Options      = map[string]any
)

var (
	//go:embed locales/*.toml
	content  embed.FS
	locales  = make(LocalesMap)
	fallback string
)

func Locales() LocalesMap {
	return locales
}

func Fallback() string {
	return fallback
}

func Init(defaultLocale string) error {
	fallback = defaultLocale

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

		locales[languageCode] = locale
	}

	return check()
}

func check() error {
	locale := GetLocale(fallback)
	if locale.LocaleName == "" {
		return ErrNameEmpty
	}

	return nil
}

func GetLocale(code LanguageCode) *Locale {
	locale, ok := locales[code]
	if ok {
		return locale
	}

	locale, ok = locales[fallback]
	if ok {
		return locale
	}

	return new(Locale)
}

func WithOptions(text string, opts Options) string {
	for key, value := range opts {
		old := fmt.Sprintf("{{%s}}", key)
		new := fmt.Sprintf("%s", value)

		text = strings.ReplaceAll(text, old, new)
	}

	return text
}

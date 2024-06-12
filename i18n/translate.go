package i18n

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"slices"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/apicontext"
	"github.com/techx/portal/config"
	"golang.org/x/text/language"
)

const fileSuffix = ".json"

type Translator struct {
	bundle          *i18n.Bundle
	defaultLanguage language.Tag
}

var translator *Translator

var basepath = func() string {
	// See https://stackoverflow.com/questions/31873396/is-it-possible-to-get-the-current-root-of-package-structure-as-a-string-in-golan
	_, f, _, _ := runtime.Caller(0)
	// Return the project root directory path.
	return filepath.Dir(filepath.Dir(f))
}()

func Initialize(cfg config.Translation) {
	defaultLanguageTag := language.Make(cfg.DefaultLanguage)
	translator = &Translator{
		defaultLanguage: defaultLanguageTag,
		bundle:          i18n.NewBundle(defaultLanguageTag),
	}

	translator.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	if !filepath.IsAbs(cfg.FilePath) {
		cfg.FilePath = filepath.Join(basepath, cfg.FilePath)
	}

	files, err := filepath.Glob(path.Join(cfg.FilePath, "*"+fileSuffix))
	if err != nil {
		log.Panic().Msgf("error: %v in loading translation file from path: %s", err, cfg.FilePath)
	}

	for _, file := range files {
		_, err := translator.bundle.LoadMessageFile(file)
		if err != nil {
			log.Panic().Msgf("error: %v in loading translation file: %s", err, file)
		}
	}

	languageTags := translator.bundle.LanguageTags()
	if !slices.Contains(languageTags, defaultLanguageTag) {
		log.Panic().Msgf("tranlation is missing for default language: %s. check translation file: %s is present",
			defaultLanguageTag.String(), filepath.Join(cfg.FilePath, defaultLanguageTag.String()+fileSuffix))
	}
}

func Title(ctx context.Context, key string, args ...interface{}) string {
	locale := apicontext.RequestContextFromContext(ctx).GetLocale()
	return Translate(locale, getTitleKey(key), args...)
}

func Message(ctx context.Context, key string, args ...interface{}) string {
	locale := apicontext.RequestContextFromContext(ctx).GetLocale()
	return Translate(locale, getMessageKey(key), args...)
}

func HasTitle(ctx context.Context, key string) bool {
	return Title(ctx, key) != getTitleKey(key)
}

func Translate(language, key string, args ...interface{}) string {
	localizerConfig := &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    key,
			Other: key,
		},
		TemplateData: args,
	}
	localizer := i18n.NewLocalizer(translator.bundle, language)
	translation, err := localizer.Localize(localizerConfig)
	if err != nil {
		log.Error().Err(err).Msgf("error translating key: %s", key)
		return key
	}
	return translation
}

func getTitleKey(key string) string {
	return fmt.Sprintf("%s_title", strings.ToLower(key))
}

func getMessageKey(key string) string {
	return fmt.Sprintf("%s_message", strings.ToLower(key))
}

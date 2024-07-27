package i18n

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/apicontext"
	"github.com/techx/portal/config"
	"github.com/techx/portal/utils"
	"golang.org/x/text/language"
)

const (
	fileSuffixJSON = ".json"
	fileSuffixHTML = ".html"
)

type Translator struct {
	defaultLanguage language.Tag
	jsonDirectory   string
	htmlDirectory   string
	bundle          *i18n.Bundle
}

var translator *Translator

func Initialize(cfg config.Translation) {
	defaultLanguageTag := language.Make(cfg.DefaultLanguage)
	translator = &Translator{
		defaultLanguage: defaultLanguageTag,
		jsonDirectory:   cfg.JSONDirectory,
		htmlDirectory:   cfg.HTMLDirectory,
		bundle:          i18n.NewBundle(defaultLanguageTag),
	}

	translator.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	if !filepath.IsAbs(cfg.JSONDirectory) {
		cfg.JSONDirectory = filepath.Join(utils.GetProjectDirectoryPath(), cfg.JSONDirectory)
	}

	files, err := filepath.Glob(path.Join(cfg.JSONDirectory, "*"+fileSuffixJSON))
	if err != nil {
		log.Panic().Msgf("error: %v in loading translation file from path: %s", err, cfg.JSONDirectory)
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
			defaultLanguageTag.String(), filepath.Join(cfg.JSONDirectory, defaultLanguageTag.String()+fileSuffixJSON))
	}
}

func HasTitle(ctx context.Context, key string) bool {
	return Title(ctx, key) != getTitleKey(key)
}

func Title(ctx context.Context, key string, args ...map[string]interface{}) string {
	return Translate(ctx, getTitleKey(key), args...)
}

func Message(ctx context.Context, key string, args ...map[string]interface{}) string {
	return Translate(ctx, getMessageKey(key), args...)
}

func Translate(ctx context.Context, key string, args ...map[string]interface{}) string {
	locale := apicontext.RequestContextFromContext(ctx).GetLocale()
	localizerConfig := &i18n.LocalizeConfig{
		MessageID: key,
		DefaultMessage: &i18n.Message{
			ID:    key,
			Other: key,
		},
	}

	if len(args) > 0 && args[0] != nil {
		localizerConfig.TemplateData = args[0]
	}

	localizer := i18n.NewLocalizer(translator.bundle, locale)
	translation, err := localizer.Localize(localizerConfig)
	if err != nil {
		log.Error().Err(err).Msgf("error translating key: %s", key)
		return key
	}
	return translation
}

func HTML(_ context.Context, fileName string, args ...map[string]interface{}) string {
	content, err := os.ReadFile(translator.htmlDirectory + "/" + fileName + fileSuffixHTML)
	if err != nil {
		log.Error().Err(err).Msgf("error reading HTML file: %s", fileName)
		return ""
	}

	htmlContent := string(content)
	if len(args) > 0 && args[0] != nil {
		for key, value := range args[0] {
			placeholder := fmt.Sprintf("{{.%s}}", key)
			htmlContent = strings.ReplaceAll(htmlContent, placeholder, fmt.Sprintf("%v", value))
		}
	}

	return htmlContent
}

func getTitleKey(key string) string {
	return fmt.Sprintf("%s_title", strings.ToLower(key))
}

func getMessageKey(key string) string {
	return fmt.Sprintf("%s_message", strings.ToLower(key))
}

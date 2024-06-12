package i18n_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/techx/portal/apicontext"
	"github.com/techx/portal/config"
	"github.com/techx/portal/i18n"
)

func TestI18nTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(I18nTestSuite))
}

type I18nTestSuite struct {
	suite.Suite
}

func (s *I18nTestSuite) SetupSuite() {
	i18n.Initialize(config.Translation{
		FilePath:        "./i18n/testdata",
		DefaultLanguage: "en",
	})
}

func (s *I18nTestSuite) TestTranslate() {
	testCases := []struct {
		lang        string
		key         string
		expectedMsg string
	}{
		{"en-US", "msg_needs_translation", "This message needs translation"},
		{"en-US", "congrats", "Congratulations!"},
		{"en-IN", "congrats", "Congratulations!"},
		{"en_IN", "congrats", "Congratulations!"},
		{"", "congrats", "Congratulations!"},
	}

	for _, tc := range testCases {
		s.Equal(tc.expectedMsg, i18n.Translate(tc.lang, tc.key))
	}
}

func (s *I18nTestSuite) TestTitle() {
	testCases := []struct {
		lang        string
		key         string
		expectedMsg string
	}{
		{"en-IN", "invalid_user", "Invalid user title"},
		{"en_IN", "INVALID_USER", "Invalid user title"},
		{"en-IN", "not_exist", "not_exist_title"},
	}

	for _, tc := range testCases {
		reqCtx := apicontext.RequestContext{Language: tc.lang}
		ctx := apicontext.NewContextWithRequestContext(context.Background(), reqCtx)
		s.Equal(tc.expectedMsg, i18n.Title(ctx, tc.key))
	}
}

func (s *I18nTestSuite) TestMessage() {
	testCases := []struct {
		lang        string
		key         string
		expectedMsg string
	}{
		{"en-IN", "invalid_user", "Couldn't fetch user profile"},
		{"en_IN", "INVALID_USER", "Couldn't fetch user profile"},
		{"en-IN", "not_exist", "not_exist_message"},
	}

	for _, tc := range testCases {
		reqCtx := apicontext.RequestContext{Language: tc.lang}
		ctx := apicontext.NewContextWithRequestContext(context.Background(), reqCtx)
		s.Equal(tc.expectedMsg, i18n.Message(ctx, tc.key))
	}
}

func (s *I18nTestSuite) TestHasTitle() {
	testCases := []struct {
		name          string
		language      string
		key           string
		expectedValue bool
	}{
		{name: "when key is present", language: "en-IN", key: "invalid_user", expectedValue: true},
		{name: "when key not present", language: "en-IN", key: "unavailable_translation_key", expectedValue: false},
	}

	for _, tc := range testCases {
		reqCtx := apicontext.RequestContext{Language: tc.language}
		ctx := apicontext.NewContextWithRequestContext(context.Background(), reqCtx)
		s.Equal(tc.expectedValue, i18n.HasTitle(ctx, tc.key))
	}
}

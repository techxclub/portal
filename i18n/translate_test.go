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
		JSONDirectory:   "./i18n/testdata",
		DefaultLanguage: "en",
	})
}

func (s *I18nTestSuite) TestHasTitle() {
	testCases := []struct {
		name          string
		language      string
		key           string
		expectedValue bool
	}{
		{name: "when key is present", language: "en-IN", key: "test_user", expectedValue: true},
		{name: "when key is present with template", language: "en-IN", key: "test_user_with_template", expectedValue: true},
		{name: "when key not present", language: "en-IN", key: "unavailable_translation_key", expectedValue: false},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			reqCtx := apicontext.RequestContext{Language: tc.language}
			ctx := apicontext.NewContextWithRequestContext(context.Background(), reqCtx)
			s.Equal(tc.expectedValue, i18n.HasTitle(ctx, tc.key))
		})
	}
}

func (s *I18nTestSuite) TestTitle() {
	testCases := []struct {
		name        string
		lang        string
		key         string
		values      map[string]interface{}
		expectedMsg string
	}{
		{
			name:        "when key not present",
			lang:        "en-IN",
			key:         "test_user_not_present",
			values:      map[string]interface{}{},
			expectedMsg: "test_user_not_present_title",
		},
		{
			name:        "when key is present without template",
			lang:        "en-IN",
			key:         "test_user",
			values:      map[string]interface{}{},
			expectedMsg: "User title test passed!",
		},
		{
			name: "when key is present without template but template values are sent",
			lang: "en-IN",
			key:  "test_user",
			values: map[string]interface{}{
				"User": "John",
			},
			expectedMsg: "User title test passed!",
		},
		{
			name: "when key is present with template but invalid template values",
			lang: "en-IN",
			key:  "test_user_with_template",
			values: map[string]interface{}{
				"NotPresent": "John",
			},
			expectedMsg: "<no value> title test passed!",
		},
		{
			name: "when key is present with template and valid template values",
			lang: "en-IN",
			key:  "test_user_with_template",
			values: map[string]interface{}{
				"User": "John",
			},
			expectedMsg: "John title test passed!",
		},
		{
			name:        "when locale is not present",
			key:         "test_user",
			values:      map[string]interface{}{},
			expectedMsg: "User title test passed!",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			reqCtx := apicontext.RequestContext{Language: tc.lang}
			ctx := apicontext.NewContextWithRequestContext(context.Background(), reqCtx)
			s.Equal(tc.expectedMsg, i18n.Title(ctx, tc.key, tc.values))
		})
	}
}

func (s *I18nTestSuite) TestMessage() {
	testCases := []struct {
		name        string
		lang        string
		key         string
		values      map[string]interface{}
		expectedMsg string
	}{
		{
			name:        "when key not present",
			lang:        "en-IN",
			key:         "test_user_not_present",
			values:      map[string]interface{}{},
			expectedMsg: "test_user_not_present_message",
		},
		{
			name:        "when message suffix not present in key",
			lang:        "en-IN",
			key:         "test_user",
			values:      map[string]interface{}{},
			expectedMsg: "User message test passed!",
		},
		{
			name:        "when key is present without template",
			lang:        "en-IN",
			key:         "test_user",
			values:      map[string]interface{}{},
			expectedMsg: "User message test passed!",
		},
		{
			name: "when key is present without template but template values are sent",
			lang: "en-IN",
			key:  "test_user",
			values: map[string]interface{}{
				"User": "John",
			},
			expectedMsg: "User message test passed!",
		},
		{
			name: "when key is present with template but invalid template values",
			lang: "en-IN",
			key:  "test_user_with_template",
			values: map[string]interface{}{
				"NotPresent": "John",
			},
			expectedMsg: "<no value> message test passed!",
		},
		{
			name: "when key is present with template and valid template values",
			lang: "en-IN",
			key:  "test_user_with_template",
			values: map[string]interface{}{
				"User": "John",
			},
			expectedMsg: "John message test passed!",
		},
		{
			name:        "when locale is not present",
			key:         "test_user",
			values:      map[string]interface{}{},
			expectedMsg: "User message test passed!",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			reqCtx := apicontext.RequestContext{Language: tc.lang}
			ctx := apicontext.NewContextWithRequestContext(context.Background(), reqCtx)
			s.Equal(tc.expectedMsg, i18n.Message(ctx, tc.key, tc.values))
		})
	}
}

func (s *I18nTestSuite) TestTranslate() {
	testCases := []struct {
		name        string
		lang        string
		key         string
		values      map[string]interface{}
		expectedMsg string
	}{
		{
			name:        "when key not present",
			lang:        "en-IN",
			key:         "test_user_not_present",
			values:      map[string]interface{}{},
			expectedMsg: "test_user_not_present",
		},
		{
			name:        "when key is present without template",
			lang:        "en-IN",
			key:         "test_user",
			values:      map[string]interface{}{},
			expectedMsg: "User test passed!",
		},
		{
			name: "when key is present without template but template values are sent",
			lang: "en-IN",
			key:  "test_user",
			values: map[string]interface{}{
				"User": "John",
			},
			expectedMsg: "User test passed!",
		},
		{
			name: "when key is present with template but invalid template values",
			lang: "en-IN",
			key:  "test_user_with_template",
			values: map[string]interface{}{
				"NotPresent": "John",
			},
			expectedMsg: "<no value> test passed!",
		},
		{
			name: "when key is present with template and valid template values",
			lang: "en-IN",
			key:  "test_user_with_template",
			values: map[string]interface{}{
				"User": "John",
			},
			expectedMsg: "John test passed!",
		},
		{
			name:        "when locale is not present",
			key:         "test_user",
			values:      map[string]interface{}{},
			expectedMsg: "User test passed!",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			reqCtx := apicontext.RequestContext{Language: tc.lang}
			ctx := apicontext.NewContextWithRequestContext(context.Background(), reqCtx)
			s.Equal(tc.expectedMsg, i18n.Translate(ctx, tc.key, tc.values))
		})
	}
}

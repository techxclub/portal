package utils_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/techx/portal/utils"
)

type UtilsTestSuite struct {
	suite.Suite
}

func TestUtilsTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UtilsTestSuite))
}

func (s *UtilsTestSuite) Test_ParseInt64WithDefault() {
	cases := []struct {
		name       string
		input      string
		defaultVal int64
		expected   int64
	}{
		{
			name:       "When invalid input and zero default",
			input:      "abc",
			defaultVal: 0,
			expected:   0,
		},
		{
			name:       "When invalid input and non-zero default",
			input:      "abc",
			defaultVal: 101,
			expected:   101,
		},
		{
			name:       "When valid input",
			input:      "123",
			defaultVal: 101,
			expected:   123,
		},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			s.Equal(tc.expected, utils.ParseInt64WithDefault(tc.input, tc.defaultVal))
		})
	}
}

func (s *UtilsTestSuite) TestGenerateRandomNumber() {
	cases := []struct {
		name        string
		length      int
		expectedLen int
		expectedErr error
	}{
		{
			name:        "When length is 0",
			length:      0,
			expectedLen: 0,
			expectedErr: errors.New("invalid length"),
		},
		{
			name:        "When length is 1",
			length:      1,
			expectedLen: 1,
		},
		{
			name:        "When length is n",
			length:      10,
			expectedLen: 10,
		},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			randomNumber, _ := utils.GenerateRandomNumber(tc.length)
			actualLen := len(randomNumber)
			s.Equal(tc.expectedLen, actualLen)
		})
	}
}

package model_llm_test

import (
	"errors"
	"github.com/Paulo-Lopes-Estevao/make-words-chunk-with-limit-token-model-LLM/model_llm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func countTokens(text, model string) (int, error) {
	NewToken := model_llm.NewTokenCounter()

	totalTokenText, err := NewToken.CountTokens(model, text)
	if err != nil {
		return 0, err
	}
	return totalTokenText, nil
}

type MockTokenizer struct {
	CountTokensFunc func(model, text string) (int, error)
}

func (m *MockTokenizer) CountTokens(model, text string) (int, error) {
	return m.CountTokensFunc(model, text)
}

func TestTokenizer_CountTokens(t *testing.T) {

	mockTokenizer := &MockTokenizer{
		CountTokensFunc: func(model, text string) (int, error) {
			return len(text), nil
		},
	}

	tests := []struct {
		text          string
		model         string
		maxToken      int
		expectedError error
	}{
		{
			text:          "this is a long text that needs to be chunked",
			model:         "gpt-3.5-turbo",
			maxToken:      11,
			expectedError: nil,
		},
		{
			text:          "text with tokenizer error",
			model:         "gpt-3.5-turbo",
			maxToken:      30,
			expectedError: errors.New("tokenizer error"),
		},
		{
			text:          "model not exist",
			model:         "",
			maxToken:      8,
			expectedError: errors.New("no encoding for model"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			if tt.expectedError != nil {
				mockTokenizer.CountTokensFunc = func(model, text string) (int, error) {
					return 0, tt.expectedError
				}
			} else {
				mockTokenizer.CountTokensFunc = func(model, text string) (int, error) {
					return countTokens(tt.text, tt.model)
				}
			}

			totalTokenChunk, err := mockTokenizer.CountTokens(tt.model, tt.text)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.maxToken, totalTokenChunk)
			}

		})
	}
}

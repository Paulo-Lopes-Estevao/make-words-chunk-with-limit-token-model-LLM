package chunk_test

import (
	"errors"
	"github.com/Paulo-Lopes-Estevao/make-words-chunk-with-limit-token-model-LLM/chunk"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockTokenizer struct {
	CountTokensFunc func(model, text string) (int, error)
}

func (m *MockTokenizer) CountTokens(model, text string) (int, error) {
	return m.CountTokensFunc(model, text)
}

func TestChunkTextToken(t *testing.T) {
	mockTokenizer := &MockTokenizer{
		CountTokensFunc: func(model, text string) (int, error) {
			return len(text), nil
		},
	}

	tests := []struct {
		name           string
		text           string
		model          string
		maxToken       int
		expectedChunks []chunk.WordsChunk
		expectedError  error
	}{
		{
			name:           "Text within token limit",
			text:           "short text",
			model:          "dummy-model",
			maxToken:       10,
			expectedChunks: []chunk.WordsChunk{{Words: "short text", TotalTokenChunk: 10}},
			expectedError:  nil,
		},
		{
			name:           "Text exceeds token limit",
			text:           "this is a long text that needs to be chunked",
			model:          "dummy-model",
			maxToken:       10,
			expectedChunks: []chunk.WordsChunk{{Words: "this is a ", TotalTokenChunk: 10}, {Words: "long text ", TotalTokenChunk: 10}, {Words: "that needs", TotalTokenChunk: 10}, {Words: " to be chu", TotalTokenChunk: 10}, {Words: "nked", TotalTokenChunk: 4}},
			expectedError:  nil,
		},
		{
			name:           "Tokenizer error",
			text:           "text with tokenizer error",
			model:          "dummy-model",
			maxToken:       10,
			expectedChunks: nil,
			expectedError:  errors.New("tokenizer error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedError != nil {
				mockTokenizer.CountTokensFunc = func(model, text string) (int, error) {
					return 0, tt.expectedError
				}
			} else {
				mockTokenizer.CountTokensFunc = func(model, text string) (int, error) {
					return len(text), nil
				}
			}

			chunks, err := chunk.ChunkTextToken(tt.text, tt.model, tt.maxToken, mockTokenizer)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedChunks, chunks)
			}
		})
	}
}

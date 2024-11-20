package model_llm

import (
	"fmt"
	"github.com/pkoukk/tiktoken-go"
)

type Tokenizer interface {
	CountTokens(model string, text string) (int, error)
}

type tokenizer struct{}

func (*tokenizer) CountTokens(model string, text string) (int, error) {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return 0, err
	}

	return len(tkm.Encode(text, nil, nil)), nil
}

func NewTokenCounter() Tokenizer {
	return &tokenizer{}
}

package main

import (
	"fmt"
	"github.com/Paulo-Lopes-Estevao/make-words-chunk-with-limit-token-model-LLM/chunk"
	"github.com/Paulo-Lopes-Estevao/make-words-chunk-with-limit-token-model-LLM/model_llm"
)

type LLMConfig struct {
	ModelName string
	MaxTokens int
	tokenizer model_llm.Tokenizer
}

func main() {
	NewToken := model_llm.NewTokenCounter()

	llmConfig := LLMConfig{
		ModelName: "gpt-4o",
		MaxTokens: 80,
		tokenizer: NewToken,
	}

	content := `Large-scale language models (LLMs) represent a significant advance at the intersection of 
artificial intelligence and natural language processing. They are algorithms that use
deep neural networks trained on vast amounts of textual data to understand,generate 
and translate human language in a variety of contexts, from commercial applications
such as virtual assistants and chatbots to academic fields such as historical 
text analysis and technical language modeling.`

	chunks, err := chunk.ChunkTextToken(content, llmConfig.ModelName, llmConfig.MaxTokens, llmConfig.tokenizer)
	if err != nil {
		panic(err)
	}

	for i, wordsChunk := range chunks {
		fmt.Printf("%d - Total Token: %d Words %s \n", i, wordsChunk.TotalTokenChunk, wordsChunk.Words)
	}

}

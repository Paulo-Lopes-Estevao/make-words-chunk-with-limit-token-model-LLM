package chunk

import "github.com/Paulo-Lopes-Estevao/make-words-chunk-with-limit-token-model-LLM/model_llm"

type WordsChunk struct {
	Words           string
	TotalTokenChunk int
}

func ChunkTextToken(text, model string, maxToken int, adapterToken model_llm.Tokenizer) ([]WordsChunk, error) {
	var words []WordsChunk

	totalTokenText, err := countTokens(text, model, adapterToken)
	if err != nil {
		return nil, err
	}

	if totalTokenText > maxToken {
		for i := 0; i < len(text); i += maxToken {
			end := i + maxToken
			if end > len(text) {
				end = len(text)
			}

			chunk := text[i:end]
			totalTokenChunk, err := countTokens(chunk, model, adapterToken)
			if err != nil {
				return nil, err
			}
			words = append(words, WordsChunk{
				Words:           chunk,
				TotalTokenChunk: totalTokenChunk,
			})

			if totalTokenChunk > maxToken {
				for totalTokenChunk > maxToken {
					for j := 0; j < len(chunk); j += maxToken {
						end := j + maxToken
						if end > len(chunk) {
							end = len(chunk)
						}

						chunk := chunk[j:end]
						totalTokenChunk, err := countTokens(chunk, model, adapterToken)
						if err != nil {
							return nil, err
						}
						words = append(words, WordsChunk{
							Words:           chunk,
							TotalTokenChunk: totalTokenChunk,
						})
					}
				}
			}
		}
	} else {
		words = append(words, WordsChunk{
			Words:           text,
			TotalTokenChunk: totalTokenText,
		})
	}

	return words, nil
}

func countTokens(text, model string, adapterToken model_llm.Tokenizer) (int, error) {
	totalTokenText, err := adapterToken.CountTokens(model, text)
	if err != nil {
		return 0, err
	}
	return totalTokenText, nil
}

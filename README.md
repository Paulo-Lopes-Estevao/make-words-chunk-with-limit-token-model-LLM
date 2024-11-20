# Make Words Chunk with Limit Token Model LLM

## Project Description

This project provides a simple and efficient functionality to split texts into "chunks" based on a token limit, using a Language Model (LLM). It is particularly useful for applications that need to handle token limits, such as interacting with APIs of language models with token processing restrictions.

### Objective

The main objectives of this project are:

- Split large texts into smaller chunks while respecting the maximum token limit.
- Ensure that each chunk stays within the allowed limit without compromising text integrity.
- Offer a modular and extensible structure for token counting and text splitting.

## How It Works

The project uses three main components:

1. **`model_llm`**: Implements token counting using the `tiktoken-go` library, which is compatible with models like GPT-4.
2. **`chunk`**: Provides the logic for splitting texts into smaller chunks based on the token limit.
3. **`main`**: The entry point, where the model, text, and token limit configuration are defined.

The basic workflow of the project is:

1. Configure the model and token limit in the main code.
2. Pass the text to the `ChunkTextToken` function to split it.
3. Process and display the generated chunks, each containing the total number of tokens and the corresponding text.

## Usage Example

### Input Text

```plaintext
Large-scale language models (LLMs) represent a significant advance at the intersection of artificial intelligence and natural language processing...
```

### Model Configuration

```go
llmConfig := LLMConfig{
	ModelName: "gpt-4o",
	MaxTokens: 80,
	tokenizer: NewToken,
}
```

### Output

```plaintext
0 - Total Token: 16 Words Large-scale language models (LLMs) represent a significant advance at the inters 
1 - Total Token: 14 Words ection of artificial intelligence and natural language processing. They are alg
...
```

## Project Structure

```plaintext
make-words-chunk-with-limit-token-model-LLM/
│
├── chunk/
│   └── chunk_text_token.go             # Logic to split text into chunks
│
├── model_llm/
│   └── tokenizer.go         # Token counting implementation
│
├── main.go                  # Main entry point of the program
│
└── README.md                # Project documentation
```

## How to Run

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Paulo-Lopes-Estevao/make-words-chunk-with-limit-token-model-LLM.git
   cd make-words-chunk-with-limit-token-model-LLM
   ```

2. **Install dependencies**:
   Ensure you have the `tiktoken-go` library installed. If not, install it:
   ```bash
   go get github.com/pkoukk/tiktoken-go
   ```

3. **Compile and run**:
   ```bash
   go run main.go
   ```

4. **Result**:
   The program will split the text into chunks based on the configured token limit and display the results in the terminal.

## Dependencies

- **Go**: Programming language used in the project.
- **tiktoken-go**: Library for token counting compatible with OpenAI models.

## Contributing

Contributions are welcome! Feel free to open issues or pull requests for improvements, bug fixes, or new features.

## License

This project is licensed under the MIT license. Feel free to use, modify, and distribute it as needed.

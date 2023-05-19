# PDF Text Summarizer with ChatGPT (WIP)

This project demonstrates how to read the text from a PDF file and generate a summary using ChatGPT, a language model developed by OpenAI.

## Requirements

- Go programming language
- OpenAI API key

## Installation

1. Clone the repository:

```shell
git clone https://github.com/your-username/pdf-text-summarizer.git
```

2. Navigate to the project directory:

```shell
cd pdf-text-summarizer
```

3. Install project dependencies:

```shell
go mod download
```

4. Set your api key

Your api key needs to be set in your bashrc or zshrc

## Usage

1. Place your PDF file in the project directory.

2. Execute the following command to run the PDF text summarizer:

```shell
go run main.go
```

3. The program will prompt you to enter the name of the PDF file. Provide the name of the PDF file (including the extension) and press Enter.

4. The program will read the text from the PDF file, send it to ChatGPT, and generate a summary.

5. The generated summary will be displayed on the console.

## How it Works

1. PDF Text Extraction:

   - The project uses the Go standard library to extract text from a PDF file. It utilizes the `github.com/unidoc/unipdf/v3` library to open the PDF file and extract the text from each page.

2. Text Summarization with ChatGPT:

   - The extracted text is passed to the ChatGPT language model for summarization. The project uses the OpenAI GPT-3.5 model for generating summaries.

   - The program sends the extracted text as a message to ChatGPT and receives a response with a summary of the text.

3. Displaying the Summary:

   - The generated summary is displayed on the console for easy reading and reference.

## Limitations

- The accuracy of the summary heavily relies on the quality and structure of the text extracted from the PDF file. Complex formatting, tables, or images may affect the summarization results.

- The summarization is performed by the ChatGPT model, which might generate summaries that are subjective or might not capture the entire context of the text.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

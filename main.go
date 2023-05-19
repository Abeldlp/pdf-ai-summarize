package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Abeldlp/pdf-reader/ai"
	"github.com/Abeldlp/pdf-reader/utils"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	chatGPT := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	text, err := utils.ReadPDFText("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	prompt := fmt.Sprintf("Sumarize the following text: %v", text)

	ai.Ask(prompt, chatGPT)
}

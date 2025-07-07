package analyzer

import (
	"context"
	"encoding/json"
	"etsy-agent/model"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

func AnalyzeProducts(products []model.Product) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	data, _ := json.Marshal(products)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are an Etsy product analyst. Get trend analysis from the product list below.",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Product List:\n```json\n%s\n```", string(data)),
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

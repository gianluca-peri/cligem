package main

// Useful link: https://ai.google.dev/gemini-api/docs/get-started/tutorial?lang=go

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	// Check if the environment variable is set
	if os.Getenv("GEMINI_API_KEY") == "" {
		fmt.Println("GEMINI_API_KEY environment variable is not set, please set it to your API key.")
		return
	}

	// Get the API key from the environment variable
	var gemini_api_key = os.Getenv("GEMINI_API_KEY")

	// Choose the model to use
	var model_name string = "gemini-2.0-flash"

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(gemini_api_key))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Create a new GenerativeModel instance and start a chat
	model := client.GenerativeModel(model_name)
	cs := model.StartChat()

	// Set a reader to read user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Using model:", model_name)
	fmt.Println("Please enter a string (or type 'quit' to quit):")
	for {
		// Read user input
		user_input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Failed to read input:", err)
		}

		// Remove the newline character from the input
		user_input = user_input[:len(user_input)-1]

		// Check if the user wants to quit
		if user_input == "quit" {
			break
		}

		// Send the user input to the model and get the response
		resp, err := cs.SendMessage(ctx, genai.Text(user_input))
		if err != nil {
			log.Fatal("Failed to generate content (text):", err)
		}

		// Print the generated content
		for _, candidate := range resp.Candidates {
			for _, part := range candidate.Content.Parts {
				if text, ok := part.(genai.Text); ok {
					fmt.Println(text)
				}
			}
		}
	}
}

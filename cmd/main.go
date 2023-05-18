package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type OpenAIRequest struct {
	Prompt string `json:"prompt"`
}

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	endpoint := "https://api.openai.com/v1/engines/davinci-codex/completions"

	prompt := "Hello, World!"

	reqBody, err := json.Marshal(OpenAIRequest{Prompt: prompt})
	if err != nil {
		log.Fatal(err)
	}

	// make request to OpenAI API
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var openAIResponse OpenAIResponse
	err = json.Unmarshal(respBody, &openAIResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(openAIResponse.Choices[0].Text)
}

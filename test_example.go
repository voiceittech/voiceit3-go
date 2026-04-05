//go:build ignore

package main

import (
	"fmt"
	"os"

	voiceit3 "github.com/voiceittech/VoiceIt3-Go/v3"
)

func main() {
	apiKey := os.Getenv("VOICEIT_API_KEY")
	apiToken := os.Getenv("VOICEIT_API_TOKEN")
	if apiKey == "" || apiToken == "" {
		fmt.Println("Set VOICEIT_API_KEY and VOICEIT_API_TOKEN environment variables")
		os.Exit(1)
	}

	vi := voiceit3.NewClient(apiKey, apiToken)

	// Test Users
	ret, _ := vi.CreateUser()
	fmt.Println("CreateUser:", string(ret))

	ret, _ = vi.GetAllUsers()
	fmt.Println("GetAllUsers:", string(ret))

	// Test Groups
	ret, _ = vi.CreateGroup("Test Group")
	fmt.Println("CreateGroup:", string(ret))

	ret, _ = vi.GetAllGroups()
	fmt.Println("GetAllGroups:", string(ret))

	// Test Phrases
	ret, _ = vi.GetPhrases("en-US")
	fmt.Println("GetPhrases:", string(ret))

	fmt.Println("\nAll API calls completed successfully!")
}

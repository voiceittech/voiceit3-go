//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"

	voiceit3 "github.com/voiceittech/voiceit3-go/v3"
)

func main() {
	apiKey := os.Getenv("VOICEIT_API_KEY")
	apiToken := os.Getenv("VOICEIT_API_TOKEN")
	if apiKey == "" || apiToken == "" {
		fmt.Println("Set VOICEIT_API_KEY and VOICEIT_API_TOKEN")
		os.Exit(1)
	}

	vi := voiceit3.NewClient(apiKey, apiToken)
	phrase := "Never forget tomorrow is a new day"
	td := "test-data"
	errors := 0

	check := func(step string, ret []byte) map[string]interface{} {
		var r map[string]interface{}
		json.Unmarshal(ret, &r)
		code := fmt.Sprintf("%v", r["responseCode"])
		if code == "SUCC" {
			fmt.Printf("PASS: %s (%s)\n", step, code)
		} else {
			fmt.Printf("FAIL: %s (%s)\n", step, code)
			errors++
		}
		return r
	}

	ret, _ := vi.CreateUser()
	r := check("CreateUser", ret)
	userId := fmt.Sprintf("%v", r["userId"])

	for i := 1; i <= 3; i++ {
		ret, _ = vi.CreateVideoEnrollment(userId, "en-US", phrase, fmt.Sprintf("%s/videoEnrollmentA%d.mov", td, i))
		check(fmt.Sprintf("VideoEnrollment%d", i), ret)
	}

	ret, _ = vi.VideoVerification(userId, "en-US", phrase, fmt.Sprintf("%s/videoVerificationA1.mov", td))
	check("VideoVerification", ret)

	ret, _ = vi.DeleteAllEnrollments(userId)
	check("DeleteEnrollments", ret)
	ret, _ = vi.DeleteUser(userId)
	check("DeleteUser", ret)

	if errors > 0 {
		fmt.Printf("\n%d FAILURES\n", errors)
		os.Exit(1)
	}
	fmt.Println("\nAll tests passed!")
}

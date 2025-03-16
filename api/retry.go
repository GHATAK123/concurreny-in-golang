package api

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

// When making API requests, failures can occur due to network issues, rate limits, or temporary server errors.
// Instead of immediately failing, a retry mechanism with backoff helps improve reliability.
// A retry mechanism automatically reattempts a failed operation.
// If a request fails, it waits for a specific time and tries again
// Backoff is a strategy where each retry waits longer than the previous one before reattempting the request.

// Types of Backoff Strategies
// Constant Backoff → Fixed delay between retries (e.g., 2s, 2s, 2s).
// Linear Backoff → Increases delay linearly (e.g., 1s, 2s, 3s).
// Exponential Backoff → Doubles the delay each time (e.g., 1s, 2s, 4s, 8s).
// Exponential Backoff with Jitter → Adds randomness to avoid synchronization issues (e.g., 1s, 2.3s, 3.8s, 6.5s).

// 🚀 Exponential Backoff Strategy
// 1️⃣ First Retry → Wait 1 second
// 2️⃣ Second Retry → Wait 2 seconds
// 3️⃣ Third Retry → Wait 4 seconds
// 4️⃣ Fourth Retry → Wait 8 seconds
// 5️⃣ Fifth Retry → Wait 16 seconds

// 💡 Formula: wait_time = baseDelay * (2^retry_attempt)

func retryWithExponentialBackoff(url string, maxRetries int, baseDelay time.Duration) (*http.Response, error) {
	var resp *http.Response
	var err error

	for attempt := 0; attempt < maxRetries; attempt++ {
		// Attempt the HTTP request
		resp, err = http.Get(url)

		// If request is successful, return response
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}

		// If error is a client error (4xx), do not retry
		if resp != nil && resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return resp, fmt.Errorf("client error: %d, not retrying", resp.StatusCode)
		}

		// Calculate exponential backoff time: baseDelay * (2^attempt)
		backoffTime := baseDelay * time.Duration(math.Pow(2, float64(attempt)))

		// Print retry message
		fmt.Printf("Retry %d failed. Retrying in %v...\n", attempt+1, backoffTime)

		// Wait before retrying
		time.Sleep(backoffTime)
	}

	// Return last error if max retries are reached
	return nil, fmt.Errorf("max retries reached, last error: %v", err)

}

func Retry() {
	url := "https://httpbin.org/status/500" // Simulated failure URL
	maxRetries := 5
	baseDelay := 1 * time.Second // Start with 1 second delay

	resp, err := retryWithExponentialBackoff(url, maxRetries, baseDelay)
	if err != nil {
		fmt.Println("Request failed:", err)
	} else {
		fmt.Println("Request succeeded with status:", resp.StatusCode)
	}

}

// 📌 How This Works
// 1️⃣ Sends an HTTP request to the API.
// 2️⃣ If successful (2xx) → Returns response.
// 3️⃣ If client error (4xx) → Does not retry.
// 4️⃣ If server error (5xx) → Waits using exponential backoff and retries.
// 5️⃣ Stops after 5 failed attempts or successful request.

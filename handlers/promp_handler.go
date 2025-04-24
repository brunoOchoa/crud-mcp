package handlers

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

// func HandlePrompt(request PromptRequest) string {
// 	body, _ := io.ReReadAll(r.Body)
// 	var req PromptRequest
// 	json.Unmarshal(body, &req)

// 	result := prompt.Pro

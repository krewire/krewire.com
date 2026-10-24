---
title: "AI & Agents (ai)"
description: "Reference guide for Krewire's unified Go AI SDK: provider-agnostic abstractions, streaming SSE, structured outputs, and reflection-based tool calling."
date: "2026-10-18"
---

# AI & Agents (`pkg/ai`)

Import paths:
- `github.com/krewire/krewire/pkg/ai`
- `github.com/krewire/krewire/pkg/ai/providers`

The `ai` package provides a unified, provider-agnostic Go SDK for integrating Large Language Models into Krewire web apps, background workers, and autonomous agent workflows with zero external runtime dependencies.

---

## 1. Supported Providers

The SDK supports leading commercial and local open-source LLM providers out of the box:

- **Google Gemini** (`providers.NewGemini(apiKey)`)
- **Anthropic Claude** (`providers.NewAnthropic(apiKey)`)
- **OpenAI & Compatible** (`providers.NewOpenAI(apiKey, opts...)` — works with OpenAI, Groq, DeepSeek, vLLM)
- **Ollama** (`providers.NewOllama(opts...)` — local Llama 3, Qwen, Mistral, DeepSeek)
- **Mock Provider** (`providers.NewMock(content)` — 100% offline unit test fixture)

---

## 2. Basic Chat & Prompt Completion

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/krewire/krewire/pkg/ai"
	"github.com/krewire/krewire/pkg/ai/providers"
)

func main() {
	provider := providers.NewOpenAI(os.Getenv("OPENAI_API_KEY"))
	client := ai.NewClient(provider, "gpt-4o").
		WithSystem("You are a helpful software architecture assistant.")

	// Simple completion
	answer, err := client.Complete(context.Background(), "Explain Krewire downward dependencies.")
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}
```

---

## 3. Type-Safe Structured Outputs (`GenerateObject[T]`)

Extract strictly typed Go structs directly from LLM completions using reflection and JSON schema generation:

```go
type UserReport struct {
	Summary    string   `json:"summary" description:"Executive summary of the user issue"`
	Severity   string   `json:"severity" description:"low, medium, high, critical"`
	Categories []string `json:"categories"`
	Actionable bool     `json:"actionable"`
}

report, err := ai.GenerateObject[UserReport](
	ctx,
	provider,
	"claude-3-5-sonnet-20241022",
	"Analyze error logs: database timeout on pool exhaustion",
)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Severity: %s (Actionable: %v)\n", report.Severity, report.Actionable)
```

---

## 4. Reflection-Based Tool & Function Calling

Register standard Go functions as LLM tools without manually writing JSON Schema definitions:

```go
type SearchQuery struct {
	Keywords []string `json:"keywords" description:"Search keywords"`
	Limit    int      `json:"limit" description:"Maximum results to return"`
}

func searchDocs(ctx context.Context, args SearchQuery) ([]string, error) {
	// Query local documentation index...
	return []string{"Result 1", "Result 2"}, nil
}

// Convert Go func into a validated AI Tool
searchTool, err := ai.NewTool("search_docs", "Search documentation index", searchDocs)

// Pass tool to client or completion request
client.WithTools(searchTool)
```

---

## 5. Streaming & Server-Sent Events (SSE)

Stream tokens in real-time to web components and HTTP route handlers:

```go
func handleStream(w http.ResponseWriter, r *http.Request) {
	stream, err := client.Stream(r.Context(), []ai.Message{
		ai.UserMessage("Stream an explanation of Go concurrency"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Flushes SSE events: event: delta, event: tool_call, event: done
	_ = ai.ServeSSE(w, stream)
}
```

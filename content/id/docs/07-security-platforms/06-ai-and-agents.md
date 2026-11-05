---
title: "AI & Agen Otonom (ai)"
description: "Panduan referensi untuk SDK AI Go terpadu Krewire: abstraksi multi-provider, streaming SSE, output terstruktur, dan tool calling berbasis reflection."
date: "2026-10-18"
---

# AI & Agen Otonom (`pkg/ai`)

Jalur impor (*import paths*):
- `github.com/krewire/krewire/pkg/ai`
- `github.com/krewire/krewire/pkg/ai/providers`

Paket `ai` menyediakan SDK Go terpadu dan agnostik terhadap penyedia (*provider-agnostic*) untuk mengintegrasikan model bahasa besar (LLM) ke dalam aplikasi web Krewire, worker latar belakang, dan alur agen otonom tanpa dependensi runtime eksternal.

---

## 1. Provider yang Didukung

SDK ini mendukung penyedia komersial terkemuka maupun model open-source lokal secara bawaan:

- **Google Gemini** (`providers.NewGemini(apiKey)`)
- **Anthropic Claude** (`providers.NewAnthropic(apiKey)`)
- **OpenAI & Kompatibel** (`providers.NewOpenAI(apiKey, opts...)` — dapat digunakan bersama OpenAI, Groq, DeepSeek, vLLM)
- **Ollama** (`providers.NewOllama(opts...)` — model lokal seperti Llama 3, Qwen, Mistral, DeepSeek)
- **Mock Provider** (`providers.NewMock(content)` — fixture pengujian unit 100% offline)

---

## 2. Percakapan Dasar & Penyelesaian Prompt

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
		WithSystem("Anda adalah asisten arsitektur perangkat lunak yang andal.")

	// Penyelesaian prompt sederhana
	answer, err := client.Complete(context.Background(), "Jelaskan aturan dependensi menurun Krewire.")
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}
```

---

## 3. Output Terstruktur dengan Type-Safety (`GenerateObject[T]`)

Mengekstrak struct Go bertipe ketat secara langsung dari hasil komplesi LLM menggunakan reflection dan pembuatan skema JSON:

```go
type UserReport struct {
	Summary    string   `json:"summary" description:"Ringkasan eksekutif laporan pengguna"`
	Severity   string   `json:"severity" description:"low, medium, high, critical"`
	Categories []string `json:"categories"`
	Actionable bool     `json:"actionable"`
}

report, err := ai.GenerateObject[UserReport](
	ctx,
	provider,
	"claude-3-5-sonnet-20241022",
	"Analisis log error berikut: database timeout on pool exhaustion",
)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Tingkat Keparahan: %s (Dapat Ditindaklanjuti: %v)\n", report.Severity, report.Actionable)
```

---

## 4. Pemanggilan Fungsi & Alat (*Tool Calling*) Berbasis Reflection

Mendaftarkan fungsi standar Go sebagai alat (tools) bagi LLM tanpa perlu menulis definisi JSON Schema secara manual:

```go
type SearchQuery struct {
	Keywords []string `json:"keywords" description:"Kata kunci pencarian"`
	Limit    int      `json:"limit" description:"Jumlah maksimal hasil yang dikembalikan"`
}

func searchDocs(ctx context.Context, args SearchQuery) ([]string, error) {
	// Menjalankan kueri ke indeks dokumentasi lokal...
	return []string{"Hasil 1", "Hasil 2"}, nil
}

// Mengonversi fungsi Go menjadi Alat AI yang tervalidasi
searchTool, err := ai.NewTool("search_docs", "Cari indeks dokumentasi", searchDocs)

// Mendaftarkan alat ke client
client.WithTools(searchTool)
```

---

## 5. Streaming & Server-Sent Events (SSE)

Mengalirkan token secara real-time ke komponen web dan handler rute HTTP:

```go
func handleStream(w http.ResponseWriter, r *http.Request) {
	stream, err := client.Stream(r.Context(), []ai.Message{
		ai.UserMessage("Alirkan penjelasan mengenai konkurensi di Go"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengirim event SSE: event: delta, event: tool_call, event: done
	_ = ai.ServeSSE(w, stream)
}
```

---

## 6. Mesin AI Agent Otonom (`github.com/krewire/krew`)

Sementara `pkg/ai` menyediakan abstraksi tingkat rendah untuk LLM, **`krew`** adalah platform AI agent otonom Krewire:

- **CLI Terdedikasi (`cmd/krew`):** Binary terminal mandiri (`krew run`, `krew session`, `krew exec`, `krew doctor`).
- **Daemon HTTP/SSE Transport:** `transport.NewHTTPServer` yang diamankan token bearer dari `krewire/auth` dan isolasi workspace tenant (sandboxing).
- **Delegasi Sub-Agen:** Dekomposisi tugas otomatis dengan ringkasan eksekusi terstruktur.
- **RBAC & Audit Berbasis Tool:** Integrasi dengan `krewire/auth` yang mengevaluasi kebijakan `ToolPermission` dan menghasilkan audit log forensik.
- **Integrasi MCP:** Dukungan penuh untuk Model Context Protocol.

---

## 7. Workspace Virtual Berbasis AI (`github.com/krewire/dash`)

**`dash`** adalah platform aplikasi web yang menyediakan virtual IDE yang siap di-deploy ke cloud (`cmd/krewire-dash-server`):

- **Editor Kode Monaco:** Penyorotan sintaks dan buffer editing real-time.
- **Terminal PTY Interaktif:** Terminal xterm.js yang terhubung melalui WebSocket aman.
- **AI Pair Programmer:** Terhubung dengan Krew melalui streaming SSE, melakukan pencarian semantik AST codebase, dan mengajukan proposal perubahan kode (diff) dengan tombol Accept/Reject interaktif.
- **Presensi Kolaboratif:** Roster pengguna aktif dan sinkronisasi kursor multi-pengguna secara real-time.

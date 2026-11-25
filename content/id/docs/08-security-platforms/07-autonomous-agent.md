---
title: "Mesin Agen AI Otonom"
description: "Arsitektur, execution loop ReAct, sandboxed tools, memori sliding-window, persona, multi-protocol transport (HTTP/SSE & MCP), dan embedding API di ai dan toolchain krew."
date: "2026-11-24"
---

# Mesin Agen AI Otonom

Krewire menyediakan mesin agen AI otonom tingkat produksi yang dibangun secara native dalam bahasa Go murni pada repositori fondasi di `krewire/ai`. Mesin ini digerakkan secara langsung dalam alur kerja pengembang harian melalui toolchain CLI terpadu `krew` (`cmd/krew`).

Mesin ini menyediakan runtime agen lengkap dengan reasoning ReAct multi-langkah, operasi workspace terisolasi (sandbox), pemadatan konteks memori geser (sliding-window), dukungan native Model Context Protocol (MCP), serta embedding API yang sangat ergonomis.

---

## Subsistem & Arsitektur

### 1. `ai/agent` — ReAct Reasoning Loop
- **Multi-Turn Reasoning:** Loop otonom observe-orient-decide-act dengan batas langkah eksekusi yang terkendali (`MaxSteps`).
- **Concurrent Tool Calls:** Pemanggilan tool LLM yang independen dieksekusi secara konkuren dengan pemulihan error agregat.
- **Sandboxed Workspace Tools:** Tool manipulasi berkas dan shell yang terisolasi (`read_file`, `write_file`, `edit_file`, `list_dir`, `run_command`, `fetch_url`).
- **Keamanan & Otorisasi:** Interface `Authorizer` dan `Auditor` untuk kebijakan izin tool granular serta logging audit kepatuhan.

### 2. `ai/persona` — Spesialisasi Karakter Agen
Persona terkonfigurasi yang menyesuaikan system prompt, temperature, dan batasan perilaku:

- `CodingAssistant`: Pair programming, refaktorisasi struktural, dan implementasi berbasis TDD.
- `DevOpsRunner`: Manajemen infrastruktur, diagnostik pipeline CI/CD, dan telemetri kontainer.
- `AutonomousWorker`: Penyelesaian tugas mandiri tanpa pengawasan, otomatisasi skrip batch, dan migrasi massal.
- `GeneralAssistant`: Asisten percakapan untuk dokumentasi, riset, dan perencanaan proyek.

### 3. `ai/memory` — Buffer Konteks Berbasis Token
- **Sliding-Window Buffer:** `TokenBuffer` melacak riwayat pesan dan konsumsi token secara dinamis.
- **Kompaksi Konteks:** Ringkasan otomatis saat ukuran pesan mendekati batas jendela konteks model.
- **Penyimpanan Sesi Persisten:** Format sesi serial JSON dengan kemampuan melanjutkan sesi (`krew chat --resume=<id>`).
- **Ekspor Transkrip:** Rendering transkrip percakapan ke format Markdown bersih untuk keperluan audit dan review kode.

### 4. `ai/transport` — Antarmuka Multi-Protokol
- **HTTP / REST Server:** Endpoint eksekusi streaming dengan Server-Sent Events (`POST /api/agent/run?stream=true`).
- **Model Context Protocol (MCP):** Server stdio MCP untuk integrasi instan dengan IDE modern (Cursor, Claude Desktop, Antigravity).
- **Autentikasi & RBAC:** Verifikasi token sesi dengan isolasi tenant dan middleware bearer token.

### 5. `ai/embed` — SDK Go In-Process
Menyematkan agen otonom ke dalam layanan Go atau CLI Krewire hanya dalam tiga baris kode:

```go
package main

import (
	"context"
	"fmt"

	"github.com/krewire/krewire/ai/embed"
	"github.com/krewire/krewire/ai/persona"
)

func main() {
	agent, err := embed.New(embed.Options{
		Workspace: ".",
		Persona:   &persona.CodingAssistant,
	})
	if err != nil {
		panic(err)
	}

	result, err := agent.Run(context.Background(), "Audit codebase untuk literal hardcoded")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Content)
}
```

---

## Integrasi Toolchain Developer (`krew`)

CLI developer `krew` mengekspos mesin agen AI otonom secara langsung melalui perintah alur kerja inti:

- `krew chat`: Memulai sesi coding percakapan interaktif di terminal.
- `krew agent <prompt>`: Menjalankan tugas otonom atau prompt secara headless hingga tuntas.
- `krew mcp`: Menjalankan server stdio Model Context Protocol untuk alur kerja copilot di IDE.
- `krew sessions`: Menampilkan daftar sesi percakapan terdahulu beserta jumlah pesan dan token.
- `krew export <id>`: Mengekspor transkrip percakapan sesi ke dalam berkas Markdown bersih.
- `krew daemon`: Menjalankan daemon server HTTP/REST dan SSE untuk eksekusi antrean headless jarak jauh.

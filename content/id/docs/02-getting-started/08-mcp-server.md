---
title: "Server Model Context Protocol (MCP)"
description: "Hubungkan asisten coding AI langsung ke dokumentasi, aturan arsitektur, dan spesifikasi komponen Krewire melalui protokol MCP."
order: 8
tags:
  - mcp
  - ai
  - agents
  - ide
---

# Server Model Context Protocol (MCP)

Krewire menyediakan dukungan bawaan untuk **Model Context Protocol (MCP)** dari Anthropic melalui dua tingkatan yang saling melengkapi:

1. **Server MCP Pengembang Lokal (`krew mcp`)**: Terintegrasi langsung dalam toolchain pengembang `krew`. Menginspeksi rute proyek lokal, menjalankan pengujian unit, memantau alokasi memori RAM/goroutine, dan memanggil generator scaffolding melalui `stdio` atau HTTP/SSE lokal.
2. **Server MCP Publik Jarak Jauh (`https://krewire.com/mcp/sse`)**: Dihosting secara global di portal resmi Krewire. Memberikan asisten AI (Cursor, Claude Desktop, Antigravity, Windsurf) kemampuan pencarian semantik atas lebih dari 70 bab dokumentasi, spesifikasi komponen Pop-Brutalist `.kiw`, dan aturan Arsitektur Piramida 6 Lapisan.

---

## 1. Panduan Koneksi Cepat

### A. Konfigurasi Cursor (`.cursor/mcp.json`)

Untuk menghubungkan Cursor ke server dokumentasi jarak jauh Krewire, tambahkan konfigurasi berikut pada `.cursor/mcp.json` di direktori utama proyek Anda:

```json
{
  "mcpServers": {
    "krewire": {
      "url": "https://krewire.com/mcp/sse"
    }
  }
}
```

Atau untuk inspeksi proyek lokal secara luring menggunakan CLI `krew`:

```json
{
  "mcpServers": {
    "krewire-local": {
      "command": "krew",
      "args": ["mcp", "--stdio"]
    }
  }
}
```

### B. Konfigurasi Claude Desktop (`claude_desktop_config.json`)

Tambahkan Krewire ke berkas konfigurasi Claude Desktop:

- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Linux**: `~/.config/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "krewire": {
      "url": "https://krewire.com/mcp/sse"
    }
  }
}
```

---

## 2. Perkakas (Tools) MCP yang Tersedia

Saat asisten AI terhubung ke `https://krewire.com/mcp/sse`, agen memperoleh akses ke perkakas berkeyakinan tinggi berikut:

- `search_docs(query, locale, limit)`: Pencarian teks lengkap terindeks di seluruh panduan dokumentasi Krewire.
- `get_doc(path, locale)`: Mengambil konten markdown lengkap dari bab dokumentasi mana pun.
- `list_docs(locale)`: Menampilkan daftar hierarkis daftar isi dokumentasi.
- `list_components(category)`: Menampilkan katalog komponen UI Pop-Brutalist `.kiw` beserta varian gayanya.
- `get_component(name)`: Menampilkan spesifikasi properti lengkap dan contoh kode `.kiw` siap pakai.
- `get_ecosystem()`: Menampilkan matriks 11 repositori dan pembagian layer (Layer 0 hingga Layer 5).
- `get_architecture_rules()`: Menyediakan aturan Arsitektur Piramida 6 Lapisan, Tiga Pilar (3S), dan invarian pure-Go.

---

## 3. Sumber Daya (Resources) & Perintah (Prompts)

### Resources
- `krewire://docs/toc`: Daftar isi lengkap naskah dokumentasi.
- `krewire://ecosystem`: Pemetaan repositori resmi dan batas kepemilikan arsitektur.
- `krewire://components`: Spesifikasi komprehensif komponen `.kiw`.
- `krewire://rules`: Filter arsitektur ketat dan batasan invarian.

### Prompts
- `krewire-architect`: Memandu agen dalam merancang aplikasi di dalam Piramida 6 Lapisan Krewire.
- `kiw-component-builder`: Mengarahkan agen dalam menulis komponen brutalist `.kiw` yang aksesibel.
- `krewire-migrator`: Menjelaskan cara menulis migrasi basis data Go bergaya Laravel menggunakan `schema.Builder`.

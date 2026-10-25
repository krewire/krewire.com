---
title: "Worker Latar Belakang (worker)"
description: "Pemrosesan tugas asinkron, jadwal cron berulang, kebijakan retry dengan backoff, dan graceful shutdown di Go dengan pkg/cloud/worker."
date: "2026-10-08"
---

# Worker Latar Belakang (`worker`)

Workload **`worker`** dirancang untuk menjalankan daemon latar belakang berumur panjang (long-lived), consumer antrean tugas asinkron, dan alur kerja terjadwal berbasis cron.

Ditenagai oleh `pkg/cloud/worker` dan `pkg/resilience`, worker beroperasi secara andal di bawah beban tinggi dengan kebijakan retry otomatis, backoff eksponensial, dead-letter queue (DLQ), dan penanganan shutdown yang anggun (graceful shutdown).

---

## Karakteristik Arsitektur

1. **Backend Antrean Pluggable:** Antarmuka antrean standar dengan driver untuk channel in-memory (untuk pengujian/pengembangan lokal), Redis, NATS, dan AWS SQS.
2. **Ketahanan & Kebijakan Retry:** Batas percobaan maksimal yang dapat dikonfigurasi, backoff eksponensial dengan jitter, dan pengiriman ke DLQ untuk pesan bermasalah (poison pill).
3. **Penjadwal Cron:** Pemicu tugas terjadwal presisi menggunakan sintaks cron 5-bagian standar.
4. **Penanganan Sinyal OS yang Bersih:** Menyelesaikan tugas yang sedang berjalan secara anggun saat menerima sinyal `SIGINT` atau `SIGTERM`, mencegah terjadinya korupsi data.

---

## Scaffolding Proyek `worker`

```bash
kiw new email-worker --worker
cd email-worker
```

Struktur direktori:

```text
email-worker/
├── krewire.yaml          # Konfigurasi proyek
├── go.mod
├── cmd/
│   └── worker/
│       └── main.go       # Titik masuk worker
└── internal/
    ├── jobs/             # Definisi payload tugas & handler
    └── services/         # Klien API email pihak ketiga
```

---

## Implementasi Worker

```go
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/krewire/krewire/pkg/cloud/worker"
	"github.com/krewire/krewire/pkg/resilience"
)

type SendEmailPayload struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Inisialisasi dispatcher worker
	dispatcher := worker.NewDispatcher(worker.Config{
		Concurrency: 10,
		Logger:      logger,
	})

	// Daftarkan handler tugas dengan kebijakan retry
	dispatcher.RegisterHandler("send_email", func(ctx context.Context, job worker.Job) error {
		var payload SendEmailPayload
		if err := job.Decode(&payload); err != nil {
			return err
		}

		logger.Info("Processing email job", "recipient", payload.Recipient)

		// Simulasi logika bisnis
		if payload.Recipient == "" {
			return fmt.Errorf("invalid recipient")
		}

		return nil
	}, worker.WithRetryPolicy(resilience.RetryPolicy{
		MaxAttempts: 3,
		InitialBackoff: 500 * time.Millisecond,
		MaxBackoff:     5 * time.Second,
	}))

	// Daftarkan tugas cron berulang (setiap 5 menit)
	dispatcher.RegisterCron("*/5 * * * *", func(ctx context.Context) error {
		logger.Info("Running 5-minute health check")
		return nil
	})

	// Mulai memproses tugas hingga interupsi OS diterima
	if err := dispatcher.Run(context.Background()); err != nil {
		logger.Error("Worker terminated with error", "err", err)
		os.Exit(1)
	}
}
```

---

## Konfigurasi (`krewire.yaml`)

```yaml
project:
  name: "email-worker"
  kind: worker
  version: "0.1.0"

worker:
  concurrency: 8
  queues:
    - default
    - email
```

`concurrency` dan `queues` adalah satu-satunya kunci di bawah `worker:`. Backend antrean—baik in-memory, berbasis KV, atau implementasi kustom `pkg/cloud/worker.Queue` milik Anda—dipilih langsung di dalam kode Go, bukan di `krewire.yaml`.

---

## Menjalankan Worker

Jalankan daemon worker secara lokal:

```bash
kiw worker
```

Jalankan dengan konkurensi atau antrean kustom:

```bash
kiw worker --concurrency 16

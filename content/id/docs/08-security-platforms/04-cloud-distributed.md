---
title: "Cloud & Sistem Terdistribusi (cloud, worker, service, infra)"
description: "Manual referensi untuk pustaka Krewire Cloud: antrean worker latar belakang asinkron, registry microservices dan RPC, serta Infrastructure as Code deklaratif."
date: "2026-10-10"
---

# Cloud & Sistem Terdistribusi (`cloud`)

Jalur impor (*import path*): `github.com/krewire/krewire/pkg/cloud`

Rangkaian paket **`cloud`** membekali aplikasi Krewire dengan primitif sistem terdistribusi kelas produksi, penjadwalan tugas latar belakang, perutean microservices, dan deklarasi infrastruktur cloud sebagai kode (*Infrastructure as Code*).

---

## 1. Worker Latar Belakang & Antrean Tugas (`cloud/worker`)

Jalur impor (*import path*): `github.com/krewire/krewire/pkg/cloud/worker`

Paket `worker` menyediakan eksekusi tugas latar belakang asinkron, penjadwalan cron periodik, kebijakan retry bertingkat, dan dead-letter queue (DLQ).

### Contoh Pemrosesan Tugas

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/krewire/krewire/pkg/cloud/worker"
)

type EmailPayload struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
}

func main() {
	// Inisialisasi runner dengan worker pool konkuren dan kebijakan retry default
	runner := worker.NewRunner(worker.Config{
		Concurrency: 10,
		MaxRetries:  3,
		Backoff:     500 * time.Millisecond,
	})

	// Mendaftarkan handler pekerjaan
	runner.Handle("send_email", func(ctx context.Context, job worker.Job) error {
		var payload EmailPayload
		if err := job.Decode(&payload); err != nil {
			return err
		}
		fmt.Printf("Mengirim email ke %s: %s\n", payload.Recipient, payload.Subject)
		return nil
	})

	// Mengirimkan pekerjaan ke antrean
	jobID, err := runner.Dispatch(context.Background(), "send_email", EmailPayload{
		Recipient: "user@example.com",
		Subject:   "Selamat Datang di Krewire!",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pekerjaan berhasil masuk antrean: %s\n", jobID)
}
```

---

## 2. Mesh Microservice & Gateway (`cloud/service`)

Jalur impor (*import path*): `github.com/krewire/krewire/pkg/cloud/service`

Paket `service` menyediakan perutean microservice cloud-native, penemuan layanan dinamis (*dynamic service discovery*), sinkronisasi konfigurasi, dan tracing terdistribusi.

### Sub-paket Utama:
- **`service/registry`**: Pendaftaran layanan dan heartbeat status kesehatan dengan dukungan untuk Consul, etcd, NATS, dan DNS.
- **`service/messaging`**: Abstraksi event broker dengan semantik publisher/subscriber.
- **`service/tracing`**: Propagasi konteks terdistribusi OpenTelemetry melintasi batas HTTP dan gRPC.
- **`service/resilience`**: Concurrency bulkhead, rate limiter, dan circuit breaker.

### Contoh Penemuan Layanan (Service Discovery)

```go
package main

import (
	"context"

	"github.com/krewire/krewire/pkg/cloud/service/registry"
)

func main() {
	reg := registry.NewMemoryRegistry()

	// Mendaftarkan instance layanan
	instance := registry.Instance{
		ID:      "auth-srv-01",
		Name:    "auth-service",
		Address: "10.0.1.15:8080",
		Tags:    []string{"v1", "auth"},
	}
	_ = reg.Register(context.Background(), instance)

	// Menemukan node layanan aktif
	nodes, _ := reg.Lookup(context.Background(), "auth-service")
	for _, node := range nodes {
		println("Node tersedia:", node.Address)
	}
}
```

---

## 3. Infrastruktur sebagai Kode (`cloud/infra`)

Jalur impor (*import path*): `github.com/krewire/krewire/pkg/cloud/infra`

Paket `infra` memungkinkan pengembang mendefinisikan infrastruktur cloud secara deklaratif menggunakan bahasa Go. Paket ini menegakkan resolusi dependensi deterministik dan menghasilkan rencana eksekusi (*execution plan*) bebas efek samping sebelum menerapkan perubahan nyata.

### Skema & Status Deklaratif

```go
package main

import (
	"github.com/krewire/krewire/pkg/cloud/infra/schema"
	"github.com/krewire/krewire/pkg/cloud/infra/state"
)

func main() {
	// Mendeklarasikan sumber daya cloud
	vpc := schema.Resource{
		Type: "aws_vpc",
		Name: "production-vpc",
		Properties: map[string]any{
			"cidr_block": "10.0.0.0/16",
			"tags": map[string]string{
				"Environment": "production",
			},
		},
	}

	// Memverifikasi rencana sebelum menerapkan perubahan
	plan := schema.NewPlan([]schema.Resource{vpc})
	if plan.HasChanges() {
		println("Rencana membutuhkan perubahan infrastruktur:")
		println(plan.Summary())
	}
}
```

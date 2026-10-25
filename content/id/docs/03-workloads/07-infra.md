---
title: "Infrastruktur Cloud sebagai Kode (infra)"
description: "Deklarasikan, rencanakan, dan sediakan sumber daya cloud langsung di Go dengan pkg/cloud/infra, rencana eksekusi deterministik, dan bebas bloat DSL eksternal."
date: "2026-10-08"
---

# Infrastruktur Cloud sebagai Kode (`infra`)

Workload **`infra`** memungkinkan tim rekayasa mendefinisikan topologi cloud, orkestrasi kontainer, jaringan, dan penyimpanan langsung menggunakan **struktur Go bertipe data kuat (typed structures)** alih-alih berkas HCL terpisah atau template YAML yang rumit.

Ditenagai oleh `pkg/cloud/infra`, engine mengompilasi deklarasi Go menjadi rencana eksekusi deterministik dan menyediakan target cloud secara aman.

---

## Karakteristik Arsitektur

1. **Go sebagai Bahasa Terpadu:** Lakukan refaktor dan uji infrastruktur Anda menggunakan jaminan compiler Go standar, pemeriksaan tipe (type checking), dan pengujian unit (`go test`).
2. **Rencana Eksekusi Deterministik:** Jalankan `kiw deploy --target infra --plan` untuk melihat pratinjau sumber daya cloud yang ditambahkan, diubah, atau dihapus sebelum mutasi apa pun dilakukan.
3. **Topologi Bebas Penyedia (Provider-Agnostic):** Deklarasi abstrak untuk komputasi (kontainer, serverless), jaringan (VPC, load balancer), dan persistensi (PostgreSQL, object store).
4. **Integritas Status (State):** Menyimpan snapshot state terenkripsi di backend jarak jauh (AWS S3, Google Cloud Storage, atau Kubernetes Secrets).

---

## Scaffolding Proyek `infra`

```bash
kiw new cloud-env --infra
cd cloud-env
```

Tata letak direktori:

```text
cloud-env/
├── krewire.yaml          # Konfigurasi proyek
├── go.mod
└── main.go               # Definisi topologi infrastruktur
```

---

## Deklarasi Topologi (`main.go`)

```go
package main

import (
	"context"
	"os"

	"github.com/krewire/krewire/pkg/cloud/infra"
	"github.com/krewire/krewire/pkg/cloud/infra/providers/aws"
)

func main() {
	stack := infra.NewStack("production-ecommerce")

	// Deklarasi VPC
	network := aws.NewVPC(stack, "ecommerce-vpc", aws.VPCConfig{
		CIDRBlock: "10.0.0.0/16",
		Subnets: []string{
			"10.0.1.0/24", // publik
			"10.0.2.0/24", // privat
		},
	})

	// Deklarasi Basis Data Terkelola PostgreSQL
	db := aws.NewRDSInstance(stack, "ecommerce-db", aws.RDSConfig{
		VPC:           network,
		Engine:        "postgres",
		Version:       "16.2",
		InstanceClass: "db.t4g.medium",
		AllocatedGB:   50,
	})

	// Deklarasi Layanan Kontainer
	aws.NewECSCluster(stack, "ecommerce-compute", aws.ECSConfig{
		VPC:         network,
		Database:    db,
		DesiredTasks: 4,
		CPU:         "512",
		Memory:      "1024",
	})

	// Eksekusi Plan atau Apply berdasarkan flag CLI
	if err := infra.Execute(context.Background(), stack); err != nil {
		os.Exit(1)
	}
}
```

---

## Konfigurasi (`krewire.yaml`)

```yaml
project:
  name: "cloud-env"
  kind: infra
  version: "0.1.0"

infra:
  provider: "aws" # aws, gcp, kubernetes
  region: "ap-southeast-1"
  state:
    backend: "s3"
    bucket: "krewire-infra-state"
    lock_table: "krewire-locks"
```

---

## Perintah CLI untuk Infrastruktur

### Melihat Pratinjau Perubahan (Dry Run Plan)

```bash
kiw deploy --target infra --plan
```

*Output:*
```text
Plan: 3 to add, 0 to change, 0 to destroy.
+ aws_vpc.ecommerce-vpc (10.0.0.0/16)
+ aws_rds_instance.ecommerce-db (postgres 16.2)
+ aws_ecs_cluster.ecommerce-compute (4 tasks)
```

### Menerapkan Perubahan (Apply)

```bash
kiw deploy --target infra

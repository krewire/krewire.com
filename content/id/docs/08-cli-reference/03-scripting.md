---
title: "Otomasi Tugas & Skrip"
description: "Deklarasikan tugas alur kerja kustom, pipeline build, dan skrip otomasi langsung di krewire.yaml dengan penerusan argumen serta eksekusi lintas platform."
date: "2026-10-08"
---

# Otomasi Tugas & Skrip

Krewire menyertakan mesin otomasi tugas bawaan langsung di dalam `kiw`, menghilangkan kebutuhan akan `Makefile`, `npm run`, atau task runner eksternal lainnya.

Tugas-tugas dideklarasikan di dalam blok `scripts:` pada `krewire.yaml` dan dieksekusi melalui `kiw run <tugas>`.

---

## Mendefinisikan Skrip di `krewire.yaml`

```yaml
scripts:
  # Tugas baris tunggal sederhana
  lint: "golangci-lint run ./..."

  # Pipeline multi-perintah berurutan
  ci:
    - "kiw fmt"
    - "kiw vet"
    - "kiw test"

  # Tugas dengan variabel lingkungan
  seed:
    cmd: "go run ./cmd/seed/main.go"
    env:
      DB_MODE: "test"

  # Runner migrasi basis data
  migrate: "go run ./cmd/migrate/main.go up"
```

---

## Mengeksekusi Skrip

Jalankan skrip bernama apa pun menggunakan `kiw run`:

```bash
# Menjalankan skrip lint
kiw run lint

# Menjalankan seluruh rantai validasi CI
kiw run ci
```

### Penerusan Argumen (Argument Forwarding)

Argumen yang diberikan setelah tanda `--` secara otomatis diteruskan ke perintah target:

```bash
kiw run migrate -- --steps 2
```

Perintah di balik layar yang dieksekusi:
```bash
go run ./cmd/migrate/main.go up --steps 2
```

---

## Tugas Ekosistem Bawaan

Jika tidak ada skrip dengan nama yang cocok di `krewire.yaml`, `kiw run` akan mencoba mengeksekusi berkas atau titik masuk (entry point) Go yang sesuai:

```bash
kiw run ./cmd/server/main.go
```

Hal ini menjamin alur kerja yang mulus dan bebas hambatan, baik untuk tugas pengembangan Go terprogram maupun tugas ad-hoc.

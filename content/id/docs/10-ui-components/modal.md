---
title: "Modal"
description: "Dialog modal aksesibel yang dibangun di atas elemen native dialog HTML."
date: "2026-11-09"
---

# Modal (`Modal.kiw`)

Kategori: **Lapisan Mengambang & Dialog**  
Berkas Komponen: `components/Modal.kiw`  
Sumber: `github.com/krewire/ui/components/Modal.kiw`

Komponen `<Modal />` merender kotak dialog modal aksesibel berbasis elemen native `<dialog>` browser. Komponen ini memiliki animasi pop-in halus, latar belakang buram, penutupan otomatis via tombol `Escape`, dan isolasi fokus keyboard (*focus trap*) tanpa perlu dependensi JavaScript eksternal.

---

## 1. Instalasi (Model Copy & Own)

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Modal.kiw ./components/
```

---

## 2. Contoh Penggunaan

```html
<!-- Tombol Pembuka -->
<Button variant="primary" onClick="document.getElementById('confirm-modal').showModal()">
  Buka Modal Konfirmasi
</Button>

<!-- Dialog Modal Native -->
<Modal id="confirm-modal" title="Konfirmasi Deployment Workload">
  <div slot="body">
    <p>Apakah Anda yakin ingin menerapkan rilis <strong>v1.2.0</strong> ke kluster produksi?</p>
  </div>
  <div slot="actions">
    <Button variant="ghost" onClick="document.getElementById('confirm-modal').close()">
      Batal
    </Button>
    <Button variant="primary" onClick="executeDeployment()">
      Terapkan Sekarang
    </Button>
  </div>
</Modal>
```

---

## 3. Aksesibilitas & Kontrol Keyboard

- **Tombol Escape**: Otomatis menutup dialog menggunakan event native browser.
- **Klik Backdrop**: Menutup dialog saat pengguna mengklik area luar kotak dialog.
- **Isolasi Fokus Keyboard**: Fokus keyboard otomatis terkurung di dalam dialog selama aktif.

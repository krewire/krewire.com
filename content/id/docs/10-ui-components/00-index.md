---
title: "Katalog Komponen UI"
description: "Katalog referensi komprehensif untuk seluruh 66 Single-File Component (.kiw) pop-brutalist Krewire UI."
date: "2026-11-09"
---

# Katalog Komponen Krewire UI

Jalur Import: `github.com/krewire/ui`  
Stylesheet Design System: `/styles/krewire-ui.css`

Selamat datang di **Katalog Komponen Krewire UI (`kui`)**.

Krewire UI menyediakan **66 Single-File Component (`.kiw`) bergaya pop-brutalist tanpa dependensi eksternal** yang dibangun secara native untuk ekosistem Go Krewire.

Seluruh komponen mengikuti filosofi **shadcn/ui "Copy & Own"**: Anda menyalin berkas `.kiw` langsung ke direktori `./components/` proyek Anda, memberi Anda kepemilikan penuh tanpa dependensi runtime tersembunyi atau risiko breaking changes dari upstream.

---

## Kategori Komponen

### Navigasi & Struktur

- [**`<Navbar />`**](/id/docs/ui-components/navbar) — Header aplikasi utama dengan logo brand, tautan navigasi, dan pemicu drawer.
- [**`<NavMenu />`**](/id/docs/ui-components/nav-menu) — Menu navigasi horizontal dengan menu dropdown interaktif.
- [**`<Sidebar />`**](/id/docs/ui-components/sidebar) — Bilah navigasi samping collapsible untuk dokumentasi dan dasbor aplikasi.
- [**`<SideMenu />`**](/id/docs/ui-components/side-menu) — Daftar navigasi vertikal untuk bilah samping administratif.
- [**`<Breadcrumb />`**](/id/docs/ui-components/breadcrumb) — Navigasi jejak hierarkis yang menunjukkan lokasi halaman saat ini.
- [**`<BreadcrumbItem />`**](/id/docs/ui-components/breadcrumb-item) — Segmen individual di dalam jejak navigasi breadcrumb.
- [**`<Tabs />`**](/id/docs/ui-components/tabs) — Pengalih konten bertab yang aksesibel dengan dukungan status aktif persisten.
- [**`<Pagination />`**](/id/docs/ui-components/pagination) — Bilah navigasi halaman dengan tombol sebelum, sesudah, dan tautan angka.
- [**`<Drawer />`**](/id/docs/ui-components/drawer) — Panel geser mulus (*drawer*) untuk navigasi mobile dan panel samping.
- [**`<Hamburger />`**](/id/docs/ui-components/hamburger) — Tombol menu hamburger 3 baris teranimasi untuk membuka drawer mobile.
- [**`<Brand />`**](/id/docs/ui-components/brand) — Logo brand ekosistem lengkap dengan lambang ikon dan badge versi.
- [**`<Head />`**](/id/docs/ui-components/head) — Penyuntik tag head standar untuk metadata, OpenGraph, dan favicon.
- [**`<Header />`**](/id/docs/ui-components/header) — Banner header bagian atas halaman dengan judul dan subjudul.
- [**`<Footer />`**](/id/docs/ui-components/footer) — Footer situs multi-kolom komprehensif dengan brand, tautan navigasi, dan hak cipta.

### Aksi & Tombol

- [**`<Button />`**](/id/docs/ui-components/button) — Tombol interaktif inti dengan varian primary, secondary, accent, outline, dan ghost.
- [**`<ButtonGroup />`**](/id/docs/ui-components/button-group) — Kelompok tombol bersambung dengan border terpadu dan status aktif.
- [**`<Dropdown />`**](/id/docs/ui-components/dropdown) — Menu dropdown tanpa dependensi yang dibangun di atas elemen native details dan summary.
- [**`<ActionMenu />`**](/id/docs/ui-components/action-menu) — Menu dropdown aksi kontekstual ringkas untuk baris tabel dan kartu.
- [**`<Link />`**](/id/docs/ui-components/link) — Tautan hiperteks aksesibel dengan transisi garis bawah dan keamanan target.
- [**`<Toggle />`**](/id/docs/ui-components/toggle) — Sakelar toggle biner untuk pengubahan pengaturan secara langsung.
- [**`<ThemeSwitch />`**](/id/docs/ui-components/theme-switch) — Tombol pengalih tema gelap dan terang dengan penyimpanan otomatis di localStorage.
- [**`<LangSwitch />`**](/id/docs/ui-components/lang-switch) — Tombol pengalih bahasa untuk aplikasi multi-bahasa internasional.

### Umpan Balik & Status

- [**`<Alert />`**](/id/docs/ui-components/alert) — Banner pemberitahuan kontekstual dengan status info, success, warning, dan error.
- [**`<Banner />`**](/id/docs/ui-components/banner) — Bilah pengumuman tingkat atas yang mencolok dengan tombol penutup.
- [**`<Toast />`**](/id/docs/ui-components/toast) — Notifikasi mengambang ringan (*toast*) dengan pewaktu tutup otomatis.
- [**`<Badge />`**](/id/docs/ui-components/badge) — Lencana mikro pil dengan varian warna untuk tag, jumlah, dan status indikator.
- [**`<Tag />`**](/id/docs/ui-components/tag) — Label kata kunci yang dapat dihapus atau diklik dengan varian warna.
- [**`<Progress />`**](/id/docs/ui-components/progress) — Indikator bilah progres kemajuan berstatus pasti maupun indeterminat.
- [**`<Spinner />`**](/id/docs/ui-components/spinner) — Pemutar loader SVG teranimasi yang aksesibel untuk operasi yang sedang berjalan.
- [**`<Skeleton />`**](/id/docs/ui-components/skeleton) — Placeholder kilau teranimasi saat memuat konten asinkron.
- [**`<Empty />`**](/id/docs/ui-components/empty) — Placeholder status kosong dengan ikon ilustrasi, pesan, dan tombol aksi.
- [**`<Callout />`**](/id/docs/ui-components/callout) — Kotak catatan berbingkai untuk tips, peringatan, dan catatan dokumentasi.

### Tampilan Data & Konten

- [**`<Card />`**](/id/docs/ui-components/card) — Permukaan kartu terangkat dengan slot header, media, isi, dan footer.
- [**`<Stat />`**](/id/docs/ui-components/stat) — Kartu metrik indikator kinerja utama (KPI) dengan angka dan indikator tren.
- [**`<Table />`**](/id/docs/ui-components/table) — Tabel data responsif dengan gaya bergaris, bingkai, dan sorot baris.
- [**`<Accordion />`**](/id/docs/ui-components/accordion) — Panel konten lipat buka-tutup dengan aksesibilitas keyboard lengkap.
- [**`<Collapse />`**](/id/docs/ui-components/collapse) — Kontainer lipat tunggal untuk menampilkan atau menyembunyikan konten.
- [**`<Timeline />`**](/id/docs/ui-components/timeline) — Garis waktu kronologis vertikal untuk jejak aktivitas dan riwayat peristiwa.
- [**`<CodeWindow />`**](/id/docs/ui-components/code-window) — Jendela pratinjau kode bergaya macOS dengan tombol traffic light dan tombol salin.
- [**`<Terminal />`**](/id/docs/ui-components/terminal) — Pratinjau terminal interaktif baris tunggal dengan tombol salin perintah.
- [**`<Code />`**](/id/docs/ui-components/code) — Cuplikan kode monospace sebaris dengan style kontras tinggi.
- [**`<Avatar />`**](/id/docs/ui-components/avatar) — Avatar profil pengguna yang menampilkan gambar dengan inisial cadangan otomatis.
- [**`<Kbd />`**](/id/docs/ui-components/kbd) — Indikator visual tombol keyboard untuk pintasan dan tombol hotkey.
- [**`<Divider />`**](/id/docs/ui-components/divider) — Garis pemisah horizontal atau vertikal dengan label teks opsional.
- [**`<Heading />`**](/id/docs/ui-components/heading) — Pembungkus tipografi heading semantik dari h1 hingga h6.
- [**`<Text />`**](/id/docs/ui-components/text) — Pembungkus tipografi untuk teks paragraf isi, teks pengantar, dan teks diredam.
- [**`<List />`**](/id/docs/ui-components/list) — Kontainer daftar berurutan dan tak berurutan dengan penataan rapi.
- [**`<Ecosystem />`**](/id/docs/ui-components/ecosystem) — Komponen grid interaktif untuk menampilkan arsitektur 10 repositori Krewire.

### Formulir & Masukan

- [**`<Form />`**](/id/docs/ui-components/form) — Pembungkus formulir aksesibel dengan method spoofing dan status pengiriman.
- [**`<Input />`**](/id/docs/ui-components/input) — Kolom input teks standar dengan label mengambang, status error, dan teks bantuan.
- [**`<SearchInput />`**](/id/docs/ui-components/search-input) — Kolom input pencarian khusus dengan ikon cari dan tombol hapus instan.
- [**`<Select />`**](/id/docs/ui-components/select) — Kolom pilihan dropdown select bergaya native.
- [**`<Textarea />`**](/id/docs/ui-components/textarea) — Kolom input teks multi-baris dengan dukungan penyesuaian ukuran otomatis.
- [**`<Checkbox />`**](/id/docs/ui-components/checkbox) — Kotak centang aksesibel berkontras tinggi dengan gaya pop-border kaku.
- [**`<Radio />`**](/id/docs/ui-components/radio) — Tombol radio pilihan tunggal dengan bulatan bergaya pop-brutalist.
- [**`<Slider />`**](/id/docs/ui-components/slider) — Input bilah geser rentang angka dengan gelembung nilai dinamis.

### Lapisan Mengambang & Dialog

- [**`<Modal />`**](/id/docs/ui-components/modal) — Dialog modal aksesibel yang dibangun di atas elemen native dialog HTML.
- [**`<Popover />`**](/id/docs/ui-components/popover) — Kontainer popover kontekstual mengambang yang terhubung ke elemen target.
- [**`<Tooltip />`**](/id/docs/ui-components/tooltip) — Petunjuk tooltip informasi yang muncul saat disorot atau difokuskan.

### Tata Letak Primitif

- [**`<Container />`**](/id/docs/ui-components/container) — Kontainer halaman terpusat responsif dengan batas lebar maksimal standar.
- [**`<Box />`**](/id/docs/ui-components/box) — Kotak tata letak serbaguna dengan dukungan border, padding, dan bayangan pop-brutalist.
- [**`<Flex />`**](/id/docs/ui-components/flex) — Kontainer tata letak fleksibel dengan konfigurasi arah, align, dan justify.
- [**`<Grid />`**](/id/docs/ui-components/grid) — Kontainer CSS Grid dengan pengaturan jumlah kolom dan token jarak.
- [**`<HStack />`**](/id/docs/ui-components/h-stack) — Kontainer flex horizontal dengan jarak seragam antar elemen anak.
- [**`<VStack />`**](/id/docs/ui-components/v-stack) — Kontainer flex vertikal dengan jarak seragam antar elemen anak.
- [**`<Stack />`**](/id/docs/ui-components/stack) — Kontainer tata letak terarah dengan pengaturan jarak fleksibel.
- [**`<Section />`**](/id/docs/ui-components/section) — Pembungkus bagian halaman semantik selebar layar dengan token padding vertikal.

---

## Cara Memasang Komponen

Untuk menambahkan komponen apa pun ke proyek Anda, gunakan perintah salin standar:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/<NamaKomponen>.kiw ./components/
```

Dan hubungkan stylesheet kanonikal pada tag `<head>` tata letak halaman Anda:

```html
<link rel="stylesheet" href="/styles/krewire-ui.css">
```

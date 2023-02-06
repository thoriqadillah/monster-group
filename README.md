# Techinical Assessment PT. Monster Group
## Langkah untuk melakukan percobaan
### Requirement
- MySQL

### Langkah-langkah untuk melakukan percobaan
- Clone project ini dengan run 
```
git clone https://github.com/thoriqadillah/monster-group.git
```
- Di dalam project, run pada terminal command di bawah ini untuk install dan update dependency (jika ada)
```
go mod tidy
```
- Buat database bernama `monster_group` di MySQL client kesayangan anda
- Buat file `.env` pada root folder project ini, atau hapus `.example` pada `.env.example` dan sesuaikan dengan setting database anda
- import database dari file .sql yang ada pada project
- Sebelum melangkah lebih jauh, perhatikan terlebih dahulu data pada tabel product. Pastikan bahwa kolom `price` adalah 0 untuk semua data
- Jalankan command di bawah ini untuk run project
```
go run .
```
- Lihat kembali data yang telah terupdate pada database

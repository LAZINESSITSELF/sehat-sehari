# Sehat Sehari

Aplikasi web untuk membantu pengguna mencapai gaya hidup sehat melalui rekomendasi **menu makan**, **aktivitas fisik**, dan **perhitungan kalori** berdasarkan data personal.

---

## 🧠 Pengenalan Singkat

**Sehat Sehari** menggabungkan:
- ✅ Frontend modern (Vue 3 + Tailwind CSS)
- ✅ Backend API (Golang + Gin + MongoDB)
- ✅ Layanan Machine Learning (FastAPI + Python)

Pengguna dapat:
- Mendaftar, login, dan mengisi data profil
- Mendapat rekomendasi menu sehat & olahraga berdasarkan profil
- Melihat tips hidup sehat dan kalori harian

---

## ⚙️ Teknologi yang Digunakan

| Layer     | Teknologi                            | Keterangan |
|-----------|--------------------------------------|------------|
| Frontend  | Vue 3, Vue Router, Tailwind CSS      | UI interaktif dan responsif |
| Backend   | Golang Gin, MongoDB Atlas            | API server dan database |
| ML        | FastAPI, Uvicorn, NumPy, Pandas, scikit-learn | Prediksi kalori, rekomendasi menu & aktivitas |
| Komunikasi | REST API (JSON)                     | Pertukaran data antar komponen |

---

## 📁 Struktur Proyek

```
sehat-sehari/
│
├── frontend/          # Vue + Tailwind CSS app
├── backend/           # Gin server with MongoDB
└── ml-learning/       # Python ML service (FastAPI)
```

---

## 🚀 Langkah Setup Proyek

### 1. 🔧 Menjalankan Backend (Golang Gin + MongoDB)

#### a. Masuk ke folder backend:
```bash
cd backend
```

#### b. Isi `.env` file:
```env
PORT=8080
MONGODB_URI=mongodb+srv://bayu:YdrsP55MkEtydEoI@rental.kkgrt.mongodb.net/?retryWrites=true&w=majority&appName=rental
DB_NAME=sehatsehari
JWT_SECRET=your_jwt_secret_key
ML_SERVICE_URL=http://localhost:5000/api/recommend
```

#### c. Jalankan server:
```bash
go run cmd/main.go
```

---

### 2. 🧠 Menjalankan ML Service (FastAPI)

#### a. Masuk ke folder:
```bash
cd ml-learning
```

#### b. Buat virtual environment & install dependencies:
```bash
python -m venv venv
source venv/bin/activate        # Linux/macOS
venv\Scripts\activate           # Windows

pip install -r requirements.txt
```

#### c. Jalankan ML service:
```bash
uvicorn main:app --host 0.0.0.0 --port 5000 --reload
```

Server berjalan di `http://localhost:5000`

---

### 3. 🌐 Menjalankan Frontend (Vue.js)

#### a. Masuk ke folder:
```bash
cd frontend
```

#### b. Install dependencies:
```bash
npm install
```

#### c. Jalankan development server:
```bash
npm run dev
```

Frontend dapat diakses di `http://localhost:5173`

---

## 🔄 Alur Komunikasi

```mermaid
graph LR
    A[User] --> B[Vue Frontend]
    B --> C[Golang Backend API]
    C --> D[MongoDB Atlas]
    C --> E[ML API - FastAPI]
```

---

## ✨ Fitur Utama

- ✅ Autentikasi pengguna: register, login, logout
- ✅ Manajemen profil: usia, tinggi, berat, gender, aktivitas, tujuan diet
- ✅ Rekomendasi menu sehat (via ML)
- ✅ Rekomendasi olahraga (via ML)
- ✅ Tips hidup sehat

---

## 📌 Catatan Penting

- Pastikan semua service (frontend, backend, ML) berjalan **secara paralel**.
- Jika ada CORS error di frontend, pastikan backend sudah mengatur header CORS dengan benar.

---

## 📂 Contoh API Endpoint

| Method | Endpoint               | Keterangan            |
|--------|------------------------|------------------------|
| POST   | `/api/register`        | Daftar user baru       |
| POST   | `/api/login`           | Login user             |
| GET    | `/api/user`            | Ambil data profil      |
| PUT    | `/api/user/update`     | Update data profil     |
| POST   | `/api/recommendation`  | Ambil rekomendasi dari ML |

---

## 📌 Kontak

Untuk pertanyaan, silakan hubungi:  
📧 bayu@example.com  
📱 IG: @sehatsehari.app

---

Terima kasih telah menggunakan **Sehat Sehari** 💚

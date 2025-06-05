<template>
    <div class="max-w-xl mx-auto mt-10 bg-white p-6 rounded-lg shadow-md">
      <h2 class="text-2xl font-semibold mb-4 text-center">Isi Profil Anda</h2>
      <form @submit.prevent="saveProfile">
        <div class="grid grid-cols-2 gap-4">
          <div class="mb-4">
            <label class="block text-gray-700">Usia</label>
            <input
              v-model.number="profile.age"
              type="number"
              class="w-full border border-gray-300 p-2 rounded mt-1"
              required
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700">Gender</label>
            <select
              v-model="profile.gender"
              class="w-full border border-gray-300 p-2 rounded mt-1"
              required
            >
              <option disabled value="">Pilih Gender</option>
              <option value="male">Laki-laki</option>
              <option value="female">Perempuan</option>
            </select>
          </div>
          <div class="mb-4">
            <label class="block text-gray-700">Tinggi (cm)</label>
            <input
              v-model.number="profile.height_cm"
              type="number"
              class="w-full border border-gray-300 p-2 rounded mt-1"
              required
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700">Berat (kg)</label>
            <input
              v-model.number="profile.weight_kg"
              type="number"
              class="w-full border border-gray-300 p-2 rounded mt-1"
              required
            />
          </div>
          <div class="mb-4 col-span-2">
            <label class="block text-gray-700">Tingkat Aktivitas</label>
            <select
              v-model="profile.activity_level"
              class="w-full border border-gray-300 p-2 rounded mt-1"
              required
            >
              <option disabled value="">Pilih Aktivitas</option>
              <option value="sedentary">Sedentary (Jarang bergerak)</option>
              <option value="light">Light (Ringan)</option>
              <option value="moderate">Moderate (Sedang)</option>
              <option value="active">Active (Aktif)</option>
              <option value="very_active">Very Active (Sangat Aktif)</option>
            </select>
          </div>
          <div class="mb-4 col-span-2">
            <label class="block text-gray-700">Tujuan Diet</label>
            <select
              v-model="profile.goal"
              class="w-full border border-gray-300 p-2 rounded mt-1"
              required
            >
              <option disabled value="">Pilih Tujuan</option>
              <option value="lose_weight">Menurunkan Berat Badan</option>
              <option value="maintain_weight">Menjaga Berat Badan</option>
              <option value="gain_weight">Menambah Berat Badan</option>
            </select>
          </div>
          <div class="mb-4 col-span-2">
            <label class="block text-gray-700">Durasi Olahraga (menit/hari)</label>
            <input
              v-model.number="profile.exercise_duration_min"
              type="number"
              class="w-full border border-gray-300 p-2 rounded mt-1"
              required
            />
          </div>
        </div>
        <button
          type="submit"
          class="w-full bg-indigo-500 hover:bg-indigo-600 text-white p-2 rounded"
        >
          Simpan Profil
        </button>
      </form>
    </div>
  </template>
  
  <script>
  import api from "../utils/api";
  
  export default {
    data() {
      return {
        profile: {
          age: null,
          gender: "",
          height_cm: null,
          weight_kg: null,
          activity_level: "",
          goal: "",
          exercise_duration_min: null,
        },
        error: null,
      };
    },
    methods: {
      async saveProfile() {
        try {
          await api.post("/profile", this.profile);
          this.$router.push("/dashboard");
        } catch (err) {
          this.error = err.response?.data?.error || "Gagal menyimpan profil";
        }
      },
    },
  };
  </script>
  
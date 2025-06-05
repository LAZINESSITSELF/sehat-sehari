<template>
    <div>
      <h2 class="text-xl font-semibold mb-4">Sehat Hidup</h2>
      <button
        @click="fetchRecommendation"
        class="mb-4 bg-purple-500 hover:bg-purple-600 text-white px-4 py-2 rounded"
      >
        Ambil Rekomendasi Hari Ini
      </button>
      <div v-if="recommendation">
        <div class="bg-white p-6 rounded-lg shadow-md">
          <p class="mb-2">Tanggal: {{ formatDate(recommendation.date) }}</p>
          <p class="mb-2">Kalori Harian: <b>{{ recommendation.calorie_need }} kcal</b></p>
          <p class="mb-2">
            Makronutrien – Protein: {{ recommendation.macros.protein_g }}g, 
            Lemak: {{ recommendation.macros.fat_g }}g, 
            Karbohidrat: {{ recommendation.macros.carbs_g }}g
          </p>
          <p class="mt-4 font-semibold">Menu Hari Ini:</p>
          <ul class="list-disc pl-6">
            <li v-for="(item, idx) in recommendation.menu" :key="idx">
              {{ item.name }} – {{ item.calories }} kcal
            </li>
          </ul>
          <p class="mt-4 font-semibold">Aktivitas Hari Ini:</p>
          <ul class="list-disc pl-6">
            <li v-for="(act, idx) in recommendation.activities" :key="idx">
              {{ act.name }} – {{ act.duration_min }} menit ({{ act.calories_burned }} kcal terbakar)
            </li>
          </ul>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import api from "../utils/api";
  
  export default {
    data() {
      return {
        recommendation: null,
        error: null,
      };
    },
    methods: {
      async fetchRecommendation() {
        try {
          const res = await api.post("/recommend");
          this.recommendation = res.data;
        } catch (err) {
          this.error = err.response?.data?.error || "Gagal mengambil rekomendasi";
        }
      },
      formatDate(dateStr) {
        const date = new Date(dateStr);
        return date.toLocaleDateString("id-ID", {
          weekday: "long",
          year: "numeric",
          month: "long",
          day: "numeric",
        });
      },
    },
  };
  </script>
  
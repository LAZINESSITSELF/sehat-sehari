<template>
    <div>
      <h2 class="text-xl font-semibold mb-4">Rekomendasi Sehat Makan</h2>
      <button
        @click="fetchRecommendation"
        class="mb-4 bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded"
      >
        Ambil Rekomendasi Hari Ini
      </button>
      <div v-if="recommendation">
        <p class="mb-2">Kebutuhan Kalori: <b>{{ recommendation.calorie_need }} kcal</b></p>
        <p class="mb-2">
          Makronutrien – Protein: {{ recommendation.macros.protein_g }}g, 
          Lemak: {{ recommendation.macros.fat_g }}g, 
          Karbohidrat: {{ recommendation.macros.carbs_g }}g
        </p>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mt-4">
          <CardMenu
            v-for="(item, idx) in recommendation.menu"
            :key="idx"
            :item="item"
          />
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import api from "../utils/api";
  import CardMenu from "../components/CardMenu.vue";
  
  export default {
    components: { CardMenu },
    data() {
      return {
        recommendation: null,
        error: null,
      };
    },
    methods: {
      async fetchRecommendation() {
        try {
          const res = await api.post("/recommend"); // endpoint untuk mem-generate rekomendasi
          this.recommendation = res.data;
        } catch (err) {
          this.error = err.response?.data?.error || "Gagal mengambil rekomendasi";
        }
      },
    },
  };
  </script>
  
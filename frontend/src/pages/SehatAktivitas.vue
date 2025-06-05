<template>
    <div>
      <h2 class="text-xl font-semibold mb-4">Rekomendasi Sehat Aktivitas</h2>
      <button
        @click="fetchRecommendation"
        class="mb-4 bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded"
      >
        Ambil Rekomendasi Hari Ini
      </button>
      <div v-if="recommendation">
        <p class="mb-2">Durasi Olahraga: {{ totalDuration }} menit</p>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mt-4">
          <CardActivity
            v-for="(act, idx) in recommendation.activities"
            :key="idx"
            :activity="act"
          />
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import api from "../utils/api";
  import CardActivity from "../components/CardActivity.vue";
  
  export default {
    components: { CardActivity },
    data() {
      return {
        recommendation: null,
        error: null,
      };
    },
    computed: {
      totalDuration() {
        if (!this.recommendation) return 0;
        return this.recommendation.activities.reduce(
          (sum, a) => sum + a.duration_min,
          0
        );
      },
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
    },
  };
  </script>
  
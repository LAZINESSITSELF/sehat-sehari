<template>
    <div>
      <h1 class="text-3xl font-bold mb-4">Selamat Datang di Sehat Sehari</h1>
      <p class="mb-4">
        Dapatkan rekomendasi menu dan olahraga yang sesuai dengan kebutuhan Anda.
      </p>
      
      <form @submit.prevent="getRecommendation" class="space-y-4">
        <div>
          <label class="block">Usia:</label>
          <input type="number" v-model="form.age" class="border rounded p-2" required/>
        </div>
        <div>
          <label class="block">Tinggi (cm):</label>
          <input type="number" v-model="form.height" class="border rounded p-2" required/>
        </div>
        <div>
          <label class="block">Berat (kg):</label>
          <input type="number" v-model="form.weight" class="border rounded p-2" required/>
        </div>
        <div>
          <label class="block">Jenis Kelamin:</label>
          <select v-model="form.gender" class="border rounded p-2" required>
            <option value="male">Pria</option>
            <option value="female">Wanita</option>
          </select>
        </div>
        <div>
          <label class="block">Tingkat Aktivitas:</label>
          <select v-model="form.activity_level" class="border rounded p-2" required>
            <option value="sedentary">Sedentary</option>
            <option value="moderate">Moderate</option>
            <option value="active">Active</option>
          </select>
        </div>
        <div>
          <label class="block">Tujuan Diet:</label>
          <select v-model="form.goal" class="border rounded p-2" required>
            <option value="lose">Turun Berat Badan</option>
            <option value="maintain">Maintain</option>
            <option value="gain">Tambah Berat Badan</option>
          </select>
        </div>
        <div>
          <label class="block">Durasi Olahraga (menit):</label>
          <input type="number" v-model="form.exercise_minutes" class="border rounded p-2" required/>
        </div>
        <div>
          <BaseButton type="submit" :buttonClasses="'px-4 py-2 bg-blue-500 text-white rounded'">
            Dapatkan Rekomendasi
          </BaseButton>
        </div>
      </form>
  
      <div v-if="recommendation" class="mt-8 p-4 border rounded">
        <h2 class="text-xl font-bold">Hasil Rekomendasi</h2>
        <p><strong>Kalori Harian:</strong> {{ recommendation.calories_needed }}</p>
        <p><strong>Menu:</strong> {{ recommendation.menu_recommendation }}</p>
        <p><strong>Olahraga:</strong> {{ recommendation.sport_recommendation }}</p>
      </div>
  
      <div v-if="errorMessage" class="mt-4 text-red-600">
        {{ errorMessage }}
      </div>
    </div>
  </template>
  
  <script>
  import BaseButton from "../../components/buttons/BaseButton.vue"
  import { fetchRecommendation } from "../../composables/useApi.js"
  
  export default {
    name: "Home",
    components: {
      BaseButton,
    },
    data() {
      return {
        form: {
          age: null,
          height: null,
          weight: null,
          gender: "male",
          activity_level: "sedentary",
          goal: "maintain",
          exercise_minutes: null,
        },
        recommendation: null,
        errorMessage: ""
      }
    },
    methods: {
      async getRecommendation() {
        this.errorMessage = ""
        try {
          const result = await fetchRecommendation(this.form)
          this.recommendation = result
        } catch (err) {
          console.error(err)
          this.errorMessage = "Terjadi kesalahan saat mendapatkan rekomendasi."
        }
      }
    }
  }
  </script>
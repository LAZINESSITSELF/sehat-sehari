<template>
    <div class="space-y-12">
      <!-- Bagian Hero -->
      <section class="bg-green-100 py-12 text-center">
        <h1 class="text-4xl font-bold mb-4">Sehat Sehari</h1>
        <p class="text-xl mb-8">
          Sehat Makan, Sehat Aktivitas, Sehat Hidup, bersama Sehat Sehari
        </p>
        <BaseButton 
          :buttonClasses="'px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white rounded'" 
          @click="navigateToInput">
          Mulai Sekarang
        </BaseButton>
      </section>
  
      <section class="container mx-auto px-4">
        <h2 class="text-2xl font-bold mb-4">Tentang Aplikasi</h2>
        <p>
          Aplikasi "Sehat Sehari" membantu Anda memperoleh rekomendasi menu makan dan olahraga
          yang sesuai dengan kebutuhan nutrisi Anda. Cukup masukkan data diri seperti usia, tinggi,
          berat, jenis kelamin, tingkat aktivitas, tujuan diet, serta durasi olahraga, dan sistem
          kami akan memberikan saran terpersonal untuk menunjang gaya hidup sehat Anda!
        </p>
      </section>
    </div>
  </template>
  
  <script>
  import BaseButton from '../../components/buttons/BaseButton.vue';
  
  export default {
    name: "Home",
    components: {
      BaseButton,
    },
    methods: {
      navigateToInput() {
        this.$router.push('/input');
      },
    },
  };
  </script>
  
  <style scoped>
  </style>
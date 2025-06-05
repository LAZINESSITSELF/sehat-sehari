<template>
    <div class="border rounded-lg p-4 shadow-sm bg-white">
      <h3 class="font-semibold">{{ item.name }}</h3>
      <p>Kalori: {{ item.calories }} kcal</p>
      <p>
        Protein: {{ item.macros.protein_g }}g, 
        Lemak: {{ item.macros.fat_g }}g, 
        Karbohidrat: {{ item.macros.carbs_g }}g
      </p>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      item: {
        type: Object,
        required: true,
      },
    },
  };
  </script>
  
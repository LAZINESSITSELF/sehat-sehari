<template>
    <div class="border rounded-lg p-4 shadow-sm bg-white">
      <h3 class="font-semibold">{{ activity.name }}</h3>
      <p>Durasi: {{ activity.duration_min }} menit</p>
      <p>Kalori Terbakar: {{ activity.calories_burned }} kcal</p>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      activity: {
        type: Object,
        required: true,
      },
    },
  };
  </script>
  
<template>
    <div class="min-h-screen flex bg-gray-50">
      <Sidebar />
      <div class="flex-1 flex flex-col">
        <Navbar />
        <div class="p-6">
          <router-view />
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import Sidebar from "../components/Sidebar.vue";
  import Navbar from "../components/Navbar.vue";
  
  export default {
    components: {
      Sidebar,
      Navbar,
    },
  };
  </script>
  
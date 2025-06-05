<template>
  <header class="bg-white shadow-md p-4 flex justify-between">
    <div>Halo, {{ userName }}</div>
    <button @click="handleLogout" class="text-red-500">Logout</button>
  </header>
</template>

<script>
import { logout } from "../utils/auth"; // import fungsi logout

export default {
  name: "Navbar",
  data() {
    return {
      userName: "",
    };
  },
  mounted() {
    // (Opsional) Fetch data user jika perlu, misal nama:
    // Anda bisa simpan nama di localStorage dan diambil di sini
    this.userName = localStorage.getItem("userName") || "Pengguna";
    this.userName = "Pengguna"; // placeholder jika belum ada data
  },
  methods: {
    handleLogout() {
      logout(this.$router);
    },
  },
};
</script>

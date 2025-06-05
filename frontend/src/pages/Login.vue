<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-100">
      <div class="bg-white p-6 rounded-lg shadow-md w-full max-w-md">
        <h2 class="text-2xl font-semibold mb-4 text-center">Login</h2>
        <form @submit.prevent="handleLogin">
          <div class="mb-4">
            <label class="block text-gray-700">Email</label>
            <input
              v-model="email"
              type="email"
              class="w-full border border-gray-300 p-2 rounded mt-1"
              required
            />
          </div>
          <div class="mb-6">
            <label class="block text-gray-700">Password</label>
            <input
              v-model="password"
              type="password"
              class="w-full border border-gray-300 p-2 rounded mt-1"
              required
            />
          </div>
          <button
            type="submit"
            class="w-full bg-blue-500 hover:bg-blue-600 text-white p-2 rounded"
          >
            Login
          </button>
        </form>
        <p class="mt-4 text-sm text-center">
          Belum punya akun?
          <router-link to="/register" class="text-blue-500">Register</router-link>
        </p>
      </div>
    </div>
  </template>
  
  <script>
  import api from "../utils/api";
  
  export default {
    data() {
      return {
        email: "",
        password: "",
        error: null,
      };
    },
    methods: {
      async handleLogin() {
        try {
          const res = await api.post("/login", {
            email: this.email,
            password: this.password,
          });
          localStorage.setItem("token", res.data.token);
          this.$router.push("/profile");
        } catch (err) {
          this.error = err.response?.data?.error || "Login gagal";
        }
      },
    },
  };
  </script>
  
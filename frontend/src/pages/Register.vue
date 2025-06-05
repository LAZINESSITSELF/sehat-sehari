<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-100">
      <div class="bg-white p-6 rounded-lg shadow-md w-full max-w-md">
        <h2 class="text-2xl font-semibold mb-4 text-center">Register</h2>
        <form @submit.prevent="handleRegister">
          <div class="mb-4">
            <label class="block text-gray-700">Name</label>
            <input
              v-model="name"
              type="text"
              class="w-full border border-gray-300 p-2 rounded mt-1"
              required
            />
          </div>
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
              minlength="6"
            />
          </div>
          <button
            type="submit"
            class="w-full bg-green-500 hover:bg-green-600 text-white p-2 rounded"
          >
            Register
          </button>
        </form>
        <p class="mt-4 text-sm text-center">
          Sudah punya akun?
          <router-link to="/login" class="text-green-500">Login</router-link>
        </p>
      </div>
    </div>
  </template>
  
  <script>
  import api from "../utils/api";
  
  export default {
    data() {
      return {
        name: "",
        email: "",
        password: "",
        error: null,
      };
    },
    methods: {
      async handleRegister() {
        try {
          await api.post("/register", {
            name: this.name,
            email: this.email,
            password: this.password,
          });
          this.$router.push("/login");
        } catch (err) {
          this.error = err.response?.data?.error || "Register gagal";
        }
      },
    },
  };
  </script>
  
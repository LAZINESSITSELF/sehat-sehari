<template>
    <div class="max-w-md mx-auto">
      <h1 class="text-2xl font-bold mb-4">Login</h1>
      <form @submit.prevent="login">
        <div class="mb-4">
          <label class="block">Email:</label>
          <input type="email" v-model="form.email" class="border rounded w-full p-2" required>
        </div>
        <div class="mb-4">
          <label class="block">Password:</label>
          <input type="password" v-model="form.password" class="border rounded w-full p-2" required>
        </div>
        <BaseButton type="submit" :buttonClasses="'px-4 py-2 bg-green-500 text-white rounded'">
          Login
        </BaseButton>
      </form>
      <div v-if="errorMessage" class="mt-4 text-red-600">{{ errorMessage }}</div>
    </div>
  </template>
  
  <script>
  import BaseButton from "../../components/buttons/BaseButton.vue"
  import { loginUser } from "../../composables/useApi.js"
  
  export default {
    name: "Login",
    components: { BaseButton },
    data() {
      return {
        form: {
          email: "",
          password: ""
        },
        errorMessage: ""
      }
    },
    methods: {
      async login() {
        this.errorMessage = ""
        try {
          const result = await loginUser(this.form)
          localStorage.setItem("token", result.token)
          this.$router.push("/")
        } catch (err) {
          console.error(err)
          this.errorMessage = "Login gagal. Cek kredensial Anda."
        }
      }
    }
  }
  </script>
<template>
    <div class="max-w-md mx-auto">
      <h1 class="text-2xl font-bold mb-4">Register</h1>
      <form @submit.prevent="register">
        <div class="mb-4">
          <label class="block">Username:</label>
          <input type="text" v-model="form.username" class="border rounded w-full p-2" required>
        </div>
        <div class="mb-4">
          <label class="block">Email:</label>
          <input type="email" v-model="form.email" class="border rounded w-full p-2" required>
        </div>
        <div class="mb-4">
          <label class="block">Password:</label>
          <input type="password" v-model="form.password" class="border rounded w-full p-2" required>
        </div>
        <BaseButton type="submit" :buttonClasses="'px-4 py-2 bg-blue-500 text-white rounded'">
          Register
        </BaseButton>
      </form>
      <div v-if="errorMessage" class="mt-4 text-red-600">{{ errorMessage }}</div>
    </div>
  </template>
  
  <script>
  import BaseButton from "../../components/buttons/BaseButton.vue"
  import { registerUser } from "../../composables/useApi.js"
  
  export default {
    name: "Register",
    components: { BaseButton },
    data() {
      return {
        form: {
          username: "",
          email: "",
          password: ""
        },
        errorMessage: ""
      }
    },
    methods: {
      async register() {
        this.errorMessage = ""
        try {
          const result = await registerUser(this.form)
          console.log(result)
          this.$router.push("/login")
        } catch (err) {
          console.error(err)
          this.errorMessage = "Registrasi gagal. Silahkan coba lagi."
        }
      }
    }
  }
  </script>
  
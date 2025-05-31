<template>
    <nav class="bg-blue-600 p-4 text-white">
      <div class="container mx-auto flex justify-between items-center">
        <div class="font-bold text-xl">
          <router-link to="/">Sehat Sehari</router-link>
        </div>
        <div class="space-x-4">
          <router-link to="/">Home</router-link>
          <router-link to="/login">Login</router-link>
          <router-link to="/register">Register</router-link>
          <router-link to="/profile">Profile</router-link>
        </div>
      </div>
    </nav>
  </template>
  
  <script>
  export default {
    name: "Navbar"
  }
  </script>
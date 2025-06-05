<template>
    <aside class="w-64 bg-white border-r">
      <div class="p-4 text-xl font-bold">Sehat Sehari</div>
      <nav class="flex flex-col">
        <router-link
          to="/dashboard/makan"
          class="px-4 py-2 hover:bg-gray-200"
          active-class="bg-gray-200"
        >
          Sehat Makan
        </router-link>
        <router-link
          to="/dashboard/aktivitas"
          class="px-4 py-2 hover:bg-gray-200"
          active-class="bg-gray-200"
        >
          Sehat Aktivitas
        </router-link>
        <router-link
          to="/dashboard/hidup"
          class="px-4 py-2 hover:bg-gray-200"
          active-class="bg-gray-200"
        >
          Sehat Hidup
        </router-link>
      </nav>
    </aside>
  </template>
  
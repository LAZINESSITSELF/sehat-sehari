<template>
    <div class="flex flex-col min-h-screen">
        <Navbar />

        <main class="flex-grow container mx-auto p-4">
            <slot />
        </main>

        <Footer />
    </div>
</template>

<script>
import Navbar from '../components/Navbar.vue'
import Footer from '../components/Footer.vue'

export default {
    name: 'BaseLayout',
    components: {
        Navbar,
        Footer,
    },
}
</script>

<style scoped>

</style>
  

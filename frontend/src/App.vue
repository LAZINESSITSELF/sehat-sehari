<template>
  <BaseLayout>
    <router-view />
  </BaseLayout>
</template>

<script>
import BaseLayout from './Layouts/BaseLayout.vue'
export default {
  name: "App",
  components: {
    BaseLayout,
  }
}
</script>

<style scoped>

</style>

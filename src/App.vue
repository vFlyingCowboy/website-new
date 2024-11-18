<template>
<Header/>
  <main class="flex-grow">
    <div class="flex-grow mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
      <RouterView />
    </div>
  </main>
</template>

<script setup lang="ts">
import Header from "@/components/Header.vue";
</script>

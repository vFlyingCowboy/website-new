<script setup lang="ts">

const lists = [
  {
    name: "Eric",
    image: "https://eric-personal.sfo3.cdn.digitaloceanspaces.com/me.jpg",
    url: "https://www.amazon.com/hz/wishlist/ls/1UVRXB7427QW8?ref_=wl_share"
  },
  {
    name: "Roxanne",
    image: 'https://eric-personal.sfo3.cdn.digitaloceanspaces.com/roxy.png',
    url: "https://www.amazon.com/hz/wishlist/ls/10X8SEZPA7POA?ref_=wl_share"
  },
  {
    name: "Sophia",
    image: 'https://eric-personal.sfo3.cdn.digitaloceanspaces.com/sophia.png',
    url: "https://www.amazon.com/hz/wishlist/ls/2KUFDG20Z0BSF?ref_=wl_share"
  },
  {
    name: "Terrance",
    image: 'https://eric-personal.sfo3.cdn.digitaloceanspaces.com/bubba.png',
    url: "https://www.amazon.com/hz/wishlist/ls/1HYC5NF85F8WQ?ref_=wl_share"
  }
]
</script>

<template>
<div class="container pb-16 pt-20">
  <h1 class="text-6xl font-bold text-center mt-10 mb-20 ">Our Birthday / Christmas Lists!</h1>
</div>
  <div class="grid md:grid-cols-2 gap-y-5 mt-10 place-items-center">
    <div v-for="index in lists" :key="index">
      <div class="max-w-md justify-center p-4 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
        <div class="grid md:grid-cols-2 place-items-center">
          <img :src="index.image" alt="faces">
          <div>
            <h5 class=" mb-5 text-2xl font-semibold tracking-tight text-gray-900 dark:text-white">&nbsp;{{ index.name }}</h5>
            <a :href="index.url" target="_blank">
              <button type="button" class="text-white bg-gradient-to-r from-blue-500 via-blue-600 to-blue-700 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2">View List Here!</button>
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
h1{
  background: linear-gradient(to right, #12c2e9, #c471ed, #f64f59);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

img{
  height: 75% !important;
  width: 75% !important;
  border-radius: 100%;
}
</style>

<script setup lang="ts">
import VueTypewriterEffect from "vue-typewriter-effect";
const strings = [
  "be an amazing dad!",
  "code things.",
  "fly fake airplanes."
]
</script>

<template>
  <main>
    <div class="container mx-auto h-full py-4 mt-20">
      <div class="grid md:grid-cols-2 place-items-center align-middle">
        <img src="https://eric-personal.sfo3.cdn.digitaloceanspaces.com/me.jpg" alt="me">
        <div class="font-extrabold">
          <span class="text-3xl text-white">Hello! My name is Eric Steiner and this is my personal website! I like to</span>&nbsp;<vue-typewriter-effect :strings=strings class="text-3xl specialText"/>
        </div>
      </div>

    </div>
  </main>
</template>

<style scoped>

img{
  border-radius: 100%;
  width: 200px;
}

.specialText{
  background: -webkit-gradient(
    linear,
    left top,
    right top,
    from(#12c2e9),
    color-stop(#c471ed),
    to(#f64f59)
  );
  background: -webkit-linear-gradient(to right, #12c2e9, #c471ed, #f64f59);
  background: -o-linear-gradient(to right, #12c2e9, #c471ed, #f64f59);
  background: linear-gradient(to right, #12c2e9, #c471ed, #f64f59);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

</style>/

<template>
  <div>
    <router-view />
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Itim&display=swap');

body {
  margin: 0;
  font-family: 'Itim', cursive;
}
</style>

<script>
export default {
  created() {
    this.$store.dispatch("init");
  }
};
</script>

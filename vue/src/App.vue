<template>
  <div>
    <router-view />
  </div>
</template>

<style src="./assets/tailwind.css"></style>

<script>
export default {
  created() {
    this.$store.dispatch("init");
  }
};
</script>

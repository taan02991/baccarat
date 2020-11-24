<template>
  <div>
    <app-layout>
      <app-text type="h1">baccarat</app-text>
      <wallet />
      <type-list />
    </app-layout>
    <p>Test Register</p>
    <input class="text-black" v-model="name" />
    <button @click="register" class="bg-green-500 m-1">Register</button>
    <p>Test Create Game</p>
    <button @click="createGame" class="bg-green-500 m-1">CreateGame</button>
    <p>Test Join Room</p>
    <button
      v-for="e in $store.state.data.game"
      v-bind:key="e.id"
      @click="$router.push(`/game/${e.id}`)"
      class="bg-green-500 m-1"
    >
      {{ e.id }}
    </button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      name: ""
    };
  },
  methods: {
    async register() {
      let mnemonic = await this.$store.dispatch("accountRegister", this.name);
      console.log(this.name);
      alert(mnemonic);
    },
    async createGame() {
      let res = await this.$store.dispatch("entitySubmit", {
        type: "game",
        body: {}
      });
      console.log(res.value.id);
      this.$router.push(`game/${res.value.id}`);
    }
  }
};
</script>

<template>
  <div>
    <div class="h1 pt-6">Game list</div>
    <button class="bg-green-500" @click="createGame">
      Create new game
    </button>
    <table v-if="game.length > 0" class="table-auto w-full">
      <thead>
        <tr>
          <th># round</th>
          <th># participants</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="g in game" :key="g.id">
          <td>{{ (g.result || []).length }}</td>
          <td>{{ g.participant.length }}</td>
          <td>
            <button @click="$router.push(`/game/${g.id}`)">
              Join game
            </button>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="card__empty" v-else>
      There are no game yet.
    </div>
  </div>
</template>

<style scoped>
.h1 {
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 1rem;
}
.card__empty {
  margin-bottom: 1rem;
  border: 1px dashed rgba(255, 255, 255, 0.1);
  padding: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  box-sizing: border-box;
  border-radius: 8px;
  color: rgba(255, 255, 255, 0.25);
  text-align: center;
  min-height: 8rem;
}
button {
  background: rgba(255, 255, 255, 0.03);
  padding: 0.75rem;
  white-space: nowrap;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 0.85rem;
  color: rgba(255, 255, 255, 0.75);
  text-transform: uppercase;
  letter-spacing: 0.05em;
  border-radius: 0.25rem;
  transition: all 0.1s;
  user-select: none;
}
</style>

<script>
export default {
  computed: {
    game() {
      return (this.$store.state.data.game || []).filter(g => g.state !== "End");
    }
  },
  methods: {
    async createGame() {
      let res = await this.$store.dispatch("entitySubmit", {
        type: "game",
        body: {}
      });
      console.log(res.value.id);
      this.$router.push(`game/${res.value.id}`);
    }
  },
};
</script>


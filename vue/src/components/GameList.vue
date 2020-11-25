<template>
  <div>
    <div class="h1 pt-6">Game list</div>
    <button class="bg-green-500" @click="createGame">
      Create new game
    </button>
    <table v-if="game.length > 0" class="table-fixed w-full mt-6">
      <thead>
        <tr class="border-2 border-gray-500">
          <th class="w-1/3 p-2"># round</th>
          <th class="w-1/3 p-2"># participants</th>
          <th class="w-1/3 p-2"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="g in game" :key="g.id" class="border-2 border-gray-500">
          <td class="text-center">{{ (g.result || []).length + 1 }}</td>
          <td class="text-center">{{ g.participant.length }}</td>
          <td class="flex justify-center">
            <button class="bg-green-500" @click="$router.push(`/game/${g.id}`)">
              Join game
            </button>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="card__empty pt-6" v-else>
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
  padding: 0.75rem;
  white-space: nowrap;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 0.85rem;
  color: rgb(255, 255, 255);
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


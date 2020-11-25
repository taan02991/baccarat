<template>
  <table class="table-auto w-4/5">
    <thead>
      <tr>
        <th class="border px-4 py-2">#</th>
        <th class="border px-4 py-2">Name</th>
        <th class="border px-4 py-2">Status</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(name, addr, index) in participantMap" v-bind:key="addr">
        <td class="border px-4 py-2">{{ index }}</td>
        <td class="border px-4 py-2">{{ name }}</td>
        <td class="border px-4 py-2">
          {{ status[addr] ? status[addr] : "Waiting" }}
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script>
import axios from "axios";

export default {
  props: ["game"],
  data: function() {
    return {
      participantMap: {}
    };
  },
  computed: {
    participant: function() {
      return this.game.participant;
    },
    status: function() {
      if (
        this.game.resultHash == null ||
        this.game.bet[this.game.resultHash.length - 1] == null
      )
        return {};
      this.game.bet[this.game.resultHash.length - 1];
      let obj = {};
      this.game.bet[this.game.resultHash.length - 1].forEach(element => {
        obj[element.creator] = element.side + ": " + element.amount[0].amount;
      });
      return obj;
    }
  },
  watch: {
    participant: function() {
      axios
        .post("/baccarat/users", {
          addr: this.game.participant
        })
        .then(res => {
          this.participantMap = res.data.result;
        });
    }
  }
};
</script>

<style scoped>
.table {
  border-radius: 10px;
}
</style>

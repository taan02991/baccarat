<template>
  <div class="container mx-auto pt-4 w-4/5">
    <div class="flex">
      <div class="w-1/3">
        Game ID: {{ game.id }}
        <br />
        Status: {{ game.state }}
      </div>
      <div class="w-1/3 text-center">
        <button
          type="button"
          class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
          @click="onLeave"
        >
          Quit Room
        </button>
        <br />
        <br />
        Round: {{ game.resultHash ? game.resultHash.length : 0 }}/8
      </div>
      <div class="w-1/3 text-right">
        Username: {{ name }}
        <br />
        <div>Token: {{ token }}</div>
      </div>
    </div>
    <div class="card-table flex my-3">
      <div class="w-1/5">
        <Card :rank="resultP1.r" :suit="resultP1.s"></Card>
      </div>
      <div class="w-1/5">
        <Card :rank="resultP2.r" :suit="resultP2.s"></Card>
      </div>
      <div class="w-1/5 relative">
        <Card
          v-for="i in 10"
          v-bind:key="i"
          :style="'position:absolute; left:' + i * 1.5 + 'px; z-index: ' + i"
        ></Card>
      </div>
      <div class="w-1/5">
        <Card :rank="resultB1.r" :suit="resultB1.s"></Card>
      </div>
      <div class="w-1/5">
        <Card :rank="resultB2.r" :suit="resultB2.s"></Card>
      </div>
    </div>
    <div class="flex mt-5">
      <div class="w-2/5">
        <UserList :game="game"></UserList>
      </div>
      <div class="w-1/5">
        <div class="mt-1 flex rounded-md shadow-sm">
          <span
            class="inline-flex items-center py-3 px-3 rounded-l-md border border-r-0 border-gray-300 text-sm"
          >
            Token
          </span>
          <input
            type="number"
            class="form-input text-black px-2 flex-1 block w-full rounded-none rounded-r-md transition duration-150 ease-in-out sm:text-sm sm:leading-5"
            placeholder="0.00"
            v-model="amount"
          />
        </div>
      </div>
      <div class="w-2/5 text-right" v-if="game.state == 'Playing'">
        <button
          @click="onBet('Player')"
          class="bg-green-500 hover:bg-green-700 text-white font-bold py-4 px-4 rounded w-4/5"
        >
          Player
        </button>
        <button
          @click="onBet('Banker')"
          class="bg-red-500 hover:bg-red-700 text-white font-bold py-4 px-4 mt-3 rounded w-4/5"
        >
          Banker
        </button>
        <button
          @click="onBet('Tie')"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-4 px-4 mt-3 rounded w-4/5"
        >
          Tie
        </button>
      </div>
      <div class="w-2/5 text-right" v-if="game.state == 'Waiting'">
        <button
          @click="onStart"
          :disabled="!isHost"
          :class="
            (isHost ? 'bg-blue-500 hover:bg-blue-700' : 'bg-gray-500') +
              ' text-white font-bold py-4 px-4 rounded w-4/5'
          "
        >
          {{ isHost ? "Start" : "Waiting for Host" }}
        </button>
      </div>
    </div>
    <Result
      @close="resultShow = false"
      :show="resultShow"
      :winner="resultWinner"
      :bet="resultBet"
      :amount="resultAmount"
      :resulthash="resultHash"
    ></Result>
  </div>
</template>

<script>
import axios from "axios";
import { RpcClient } from "tendermint";
import Card from "@/components/Card.vue";
import UserList from "@/components/UserList.vue";
import Result from "@/components/Result.vue";

export default {
  components: {
    Card,
    UserList,
    Result
  },
  data: function() {
    return {
      game: {},
      r: "1",
      s: "S",
      resultShow: false,
      resultWinner: "",
      resultBet: "",
      resultAmount: "",
      resultHash: "",
      resultP1: { r: "", s: "" },
      resultP2: { r: "", s: "" },
      resultB1: { r: "", s: "" },
      resultB2: { r: "", s: "" },
      amount: "",
      rpcClient: null,
      timeRemaining: 0,
    };
  },
  async mounted() {
    await this.initConnection(10);
    await this.initRpcConnection();
  },
  computed: {
    account() {
      return this.$store.state.account;
    },
    address() {
      const { client } = this.$store.state;
      const address = client && client.senderAddress;
      return address;
    },
    name() {
      return this.$store.state.name;
    },
    token() {
      if (this.account.coins) {
        return this.account.coins.filter(coin => {
          return coin.denom === "token";
        })[0].amount;
      }
      return 0;
    },
    isHost() {
      return this.game.participant && this.address == this.game.participant[0];
    }
  },
  methods: {
    async initConnection(n) {
      await axios
        .get(`/baccarat/game/${this.$route.params.id}`)
        .then(res => {
          this.game = res.data.result;
          this.$store.dispatch("entitySubmit", {
            type: "game/participant",
            body: {
              id: this.$route.params.id,
              action: "Join"
            }
          });
        })
        .catch(async error => {
          if (n === 1) throw error;
          const sleep = () => {
            return new Promise(resolve => setTimeout(resolve, 3000));
          };
          await sleep();
          this.initConnection(n - 1);
        });
    },
    async initRpcConnection() {
      this.rpcClient = RpcClient(process.env.VUE_APP_TENDERMINT_NODE);
      this.rpcClient.subscribe(
        [`tm.event='Tx' AND updateGame.gameID = '${this.$route.params.id}'`],
        () => {
          this.fetchGame();
        }
      );
      this.rpcClient.subscribe(
        [
          `tm.event='Tx' AND revealResult.gameID = '${this.$route.params.id}' AND revealResult.address = '${this.address}'`
        ],
        data => {
          data.TxResult.result.events.forEach(e => {
            let obj = {};
            if (e.type === "revealResult") {
              e.attributes.forEach(element => {
                let key = atob(element.key);
                let value = atob(element.value);
                obj[key] = value;
              });
              if (obj["address"] === this.address) {
                this.resultWinner = obj["winner"];
                this.resultAmount = obj["reward"];
                this.resultBet = obj["betSide"];
                this.resultHash = obj["resultHash"];
                this.resultShow = true;
                let [player, banker] = obj["card"].split(";");
                let [P1, P2] = player.split(",");
                let [B1, B2] = banker.split(",");
                this.resultP1 = { r: P1.slice(0, -1), s: P1.slice(-1) };
                this.resultP2 = { r: P2.slice(0, -1), s: P2.slice(-1) };
                this.resultB1 = { r: B1.slice(0, -1), s: B1.slice(-1) };
                this.resultB2 = { r: B2.slice(0, -1), s: B2.slice(-1) };
                return true;
              }
            }
          });
        }
      );
    },
    async onLeave() {
      await this.$store.dispatch("entitySubmit", {
        type: "game/participant",
        body: {
          id: this.$route.params.id,
          action: "Leave"
        }
      });
      this.$router.push("/");
    },
    async onStart() {
      await this.$store.dispatch("entitySubmit", {
        type: "game/start",
        body: {
          id: this.$route.params.id
        }
      });
      await this.fetchGame();
    },
    async onBet(side) {
      this.resultP1 = { r: "", s: "" };
      this.resultP2 = { r: "", s: "" };
      this.resultB1 = { r: "", s: "" };
      this.resultB2 = { r: "", s: "" };
      await this.$store.dispatch("entitySubmit", {
        type: "game/bet",
        body: {
          id: this.$route.params.id,
          side: side,
          amount: this.amount + "token"
        }
      });
      this.amount = "";
      await this.fetchGame();
      await this.$store.dispatch("accountUpdate");
    },
    async fetchGame() {
      await axios.get(`/baccarat/game/${this.$route.params.id}`).then(res => {
        this.game = res.data.result;
      });
    }
  }
};
</script>

<style scoped>
.card-table {
  background-color: green;
  border-radius: 10px;
  border: solid;
  border-color: white;
  padding: 20px;
}
</style>

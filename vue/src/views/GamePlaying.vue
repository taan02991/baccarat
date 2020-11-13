<template>
  <div class="container mx-auto pt-4 w-4/5">
    <div class="flex">
      <div class="w-1/3">
        Game ID: 12
        <br />
        Betting time: 13s
      </div>
      <div class="w-1/3 text-center">
        <button
          type="button"
          class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
        >
          Quit Room
        </button>
        <br />
        <br />
        Round: 1/8
      </div>
      <div class="w-1/3 text-right">
        Username: TestUser
        <br />
        Token: 500
      </div>
    </div>
    <div class="card-table flex my-3">
      <div class="w-1/5">
        <Card :rank="r" :suit="s"></Card>
      </div>
      <div class="w-1/5">
        <Card :rank="r" :suit="s"></Card>
      </div>
      <div class="w-1/5 relative">
        <Card
          v-for="i in 10"
          v-bind:key="i"
          :style="
            'position:absolute; left:' + (15 + i * 1.5) + 'px; z-index: ' + i
          "
        ></Card>
      </div>
      <div class="w-1/5">
        <Card :rank="r" :suit="s"></Card>
      </div>
      <div class="w-1/5">
        <Card :rank="r" :suit="s"></Card>
      </div>
    </div>
    <div class="flex mt-5">
      <div class="w-2/5">
        <UserList />
      </div>
      <div class="w-1/5">
        <div class="mt-1 flex rounded-md shadow-sm">
                  <span class="inline-flex items-center py-3 px-3 rounded-l-md border border-r-0 border-gray-300 text-sm">
                    Token
                  </span>
                  <input type='number' class="form-input text-black px-2 flex-1 block w-full rounded-none rounded-r-md transition duration-150 ease-in-out sm:text-sm sm:leading-5" placeholder="0.00">
                </div>
      </div>
      <div class="w-2/5 text-right">
        <button
          class="bg-green-500 hover:bg-green-700 text-white font-bold py-4 px-4 rounded w-4/5"
        >
          Player
        </button>
        <button
          class="bg-red-500 hover:bg-red-700 text-white font-bold py-4 px-4 mt-3 rounded w-4/5"
        >
          Banker
        </button>
        <button
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-4 px-4 mt-3 rounded w-4/5"
        >
          Tie
        </button>
      </div>
    </div>
    <p>Test Animation</p>
        <button class="bg-blue-500 hover:bg-blue-700 text-white py-1 px-1 rounded" @click='r=""; s=""'>close card</button>
        <button class="bg-blue-500 hover:bg-blue-700 text-white py-1 px-1 rounded" @click='r="1"; s="S"'>open card</button>
  </div>
</template>
 
<script>
import Card from "@/components/Card.vue";
import UserList from "@/components/UserList.vue";

export default {
  components: {
    Card,
    UserList,
  },
  data: function () {
    return {
      r: "1",
      s: "S",
    };
  },
  methods: {
    showModal() {
      $("#modal").modal("toggle");
    },
  },
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
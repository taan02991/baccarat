<template>
  <div class="fixed z-10 inset-0 overflow-y-auto" v-if="show">
    <div
      class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0"
    >
      <div class="fixed inset-0 transition-opacity" aria-hidden="true">
        <div class="absolute inset-0 bg-gray-500 opacity-75"></div>
      </div>
      <span
        class="hidden sm:inline-block sm:align-middle sm:h-screen"
        aria-hidden="true"
        >&#8203;</span
      >
      <div
        class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full"
        role="dialog"
        aria-modal="true"
        aria-labelledby="modal-headline"
      >
        <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
          <div class="sm:flex sm:items-start">
            <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
              <h3
                class="text-2xl leading-6 font-medium text-gray-900"
                id="modal-headline"
              >
                Result
              </h3>
              <div class="mt-4">
                <p class="text text-gray-700">Winning Side: {{ winner }}</p>
                <p class="text text-gray-700 mt-2">Your Bet: {{ bet }}</p>
                <p class="text text-gray-700 mt-2">Reward: {{ amount }}</p>
                <p class="text text-gray-700 mt-2">Result Hash: {{ resulthash }}</p>
              </div>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-4 py-3 text-center">
          <button
            @click="$emit('close')"
            type="button"
            class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-3 px-4 rounded w-4/5"
          >
            Continue
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["show", "winner", "bet", "amount", "resulthash"]
};
</script>

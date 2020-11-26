<template>
  <transition name="flip">
    <div
      v-if="show"
      class="card shadow"
      v-bind:key="show"
      :class="suitColor[suit]"
    >
      <span class="card__suit card__suit--top">{{ suits[suit] }}</span>
      <span class="card__number">{{ rank }}</span>
      <span class="card__suit card__suit--bottom">{{ suits[suit] }}</span>
    </div>
    <img
      v-else
      class="card shadow"
      src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/102308/card_backside.jpg"
    />
  </transition>
</template>

<script>
export default {
  props: ["rank", "suit"],
  data: function() {
    return {
      suits: {
        H: "♥",
        D: "♦",
        S: "♠",
        C: "♣"
      },
      suitColor: {
        H: "red",
        D: "red",
        S: "black",
        C: "black"
      }
    };
  },
  computed: {
    show: function() {
      return this.suit && this.rank;
    }
  }
};
</script>

<style scoped>
.card {
  position: relative;
  background-color: white;
  width: 150px;
  height: 200px;
  float: left;
  margin-right: 5px;
  margin-bottom: 5px;
  border-radius: 2px;
}

.card__suit {
  font-size: 35px;
  width: 100%;
  display: block;
}

.card__suit--top {
  text-align: left;
  padding-left: 5px;
}

.card__suit--bottom {
  position: absolute;
  bottom: 0px;
  text-align: left;
  transform: rotate(180deg);
  padding-left: 5px;
}

.card__number {
  width: 100%;
  position: absolute;
  top: 38%;
  text-align: center;
  font-size: 32px;
}

.red {
  color: #ff0000;
}

.black {
  color: #000;
}

.flip-enter-active {
  transition: all 0.4s ease;
}

.flip-leave-active {
  display: none;
}

.flip-enter,
.flip-leave {
  transform: rotateY(180deg);
  opacity: 0;
}
</style>

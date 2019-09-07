<template>
  <div>
    <br />
    <br />
    <br />
    <br />
    <h2>{{ hobbyName }}</h2>
    <img :src="image" alt="hobby" height="150" weight="150" />
    <div class="detail">
      <p>
        フットサル（Futsal）は、いわゆる室内サッカーである南米の「サロンフットボール」または英国の「インドアサッカー」を起源とする、基本的には室内で行われる、サッカーに似た競技である。1チーム5人で、試合時間は20分ハーフの合計40分である。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://ja.wikipedia.org/wiki/%E3%83%95%E3%83%83%E3%83%88%E3%82%B5%E3%83%AB"
        >wikipedia</a
      >
    </span>
  </div>
</template>

<script>
export default {
  data() {
    return {
      hobbyName: '',
      image: null
    }
  },

  created() {
    const HOBBY_ID = 6

    // set hobby data
    const hobby = this.$hobby.GetHobbyByID(HOBBY_ID)
    this.hobbyName = hobby.name
    this.image = require(`~/assets/hobby/${HOBBY_ID}.png`)
  }
}
</script>

<style>
.detail {
  width: 700px;
  height: 150px;
  margin-right: auto;
  margin-left: auto;
  border: 1px solid #000;
  overflow: auto;
}

.where {
  font-size: 80%;
  font-style: italic;
}
</style>

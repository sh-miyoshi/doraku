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
        キャンプとは、野外で一時的な生活をすること。野営、露営、宿営ともいう。キャンプは特に自然の中で活動するからこそ大きな効能がもたらされる、と考えられている。人為的・作為的な空間であり一種の保護装置と化してしまった都市から離れた大自然の中での生活は、都市での日常行動を打ち破る創造的行動が要求され、また自然と接するうちにその美しさや厳しさを知ることもできる。またキャンプの一連の活動の中で登山・釣り・水泳などの活動をするうちに知らず知らずのうちに身体を鍛えることもできる。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://ja.wikipedia.org/wiki/%E3%82%AD%E3%83%A3%E3%83%B3%E3%83%97"
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
    const HOBBY_ID = 24

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

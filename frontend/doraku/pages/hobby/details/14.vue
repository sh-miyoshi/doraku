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
        ネイルアート (英語: Nail art)
        とは、手足の爪に施す化粧や装飾のことである。ネイルアートを施してくれる店をネイルサロン、その技術者を日本では一般的にネイリストという。さまざまなネイルアート用品が市販されており、プロに限らず、自身の爪を飾る創作作業を楽しむために、ネイルアートを行う人も多くなっている
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://ja.wikipedia.org/wiki/%E3%83%8D%E3%82%A4%E3%83%AB%E3%82%A2%E3%83%BC%E3%83%88"
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
    const HOBBY_ID = 14

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

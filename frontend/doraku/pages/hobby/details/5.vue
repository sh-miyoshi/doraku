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
        釣り（つり）あるいは魚釣り（さかなつり）とは、釣り針、釣り糸、釣り竿などの道具と餌や疑似餌を使って、魚介類などの生物を採捕する行為、方法のことである。釣りの起源は少なくとも約4万年前の旧石器時代まで遡ることができる。日本でも、石器時代の遺跡から骨角器の釣針が見つかる。
      </p>
    </div>
    <span class="where">
      from
      <a href="https://ja.wikipedia.org/wiki/%E9%87%A3%E3%82%8A">wikipedia</a>
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
    const HOBBY_ID = 5

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

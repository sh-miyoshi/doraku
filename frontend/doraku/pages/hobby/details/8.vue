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
        ゲートボールは、5人1組の2チーム対抗で行われる、日本発祥のスポーツ。漢字では門球と表記される。1947年、鈴木栄治（のちに改名して和伸となる）が北海道芽室町において、クロッケーをヒントに考案した。元々は太平洋戦争後の物資不足で遊び道具のない子供のため、また子供の不良化防止のために作られた遊びだったが、高度経済成長期に高齢者向けスポーツとして爆発的に流行、現在では日本だけでなくアジア各国にも普及している。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://ja.wikipedia.org/wiki/%E3%82%B2%E3%83%BC%E3%83%88%E3%83%9C%E3%83%BC%E3%83%AB"
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
    const HOBBY_ID = 8

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

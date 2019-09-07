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
        テレビゲームは、テレビで遊ぶコンピュータゲーム。現在では家庭用ゲーム機のゲームを指すことが多いようだ。現在ではプレイステーション4,
        ニンテンドーSWITCHなどが販売されている。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://dic.pixiv.net/a/%E3%83%86%E3%83%AC%E3%83%93%E3%82%B2%E3%83%BC%E3%83%A0"
        >ピクシブ百科事典</a
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
    const HOBBY_ID = 7

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

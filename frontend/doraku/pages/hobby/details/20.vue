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
        水泳とは水面または水中で泳ぐこと全般を指す。人が泳ぐためにはそれなりの訓練が必要となる。歴史的には、水辺で生活する上での食料採取や水難回避、移動手段、沐浴からの派生、軍事技術の一つ等として用いられてきた。近代以降はスポーツ競技としての水泳も盛んとなり、世界各地で大会が開催されている。
      </p>
    </div>
    <span class="where">
      from
      <a href="https://dic.nicovideo.jp/a/%E6%B0%B4%E6%B3%B3">ニコニコ大百科</a>
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
    const HOBBY_ID = 23

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

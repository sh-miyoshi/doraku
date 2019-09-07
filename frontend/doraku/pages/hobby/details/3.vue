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
        BBQはおもに戸外で行う直火式焼肉料理。
      </p>
      <p>
        名称の起源は諸説あるが，アメリカインディアンが炉で焼く野外料理を babe a
        queueと呼んだことから起ったといわれる。各種の肉を塊 (ときには丸焼もある)
        にして串に刺し，その他，野菜など好みの材料をあしらって焼く。炉は煉瓦，石材を積み，木炭などを燃料とする野趣豊かなもので，野外の宴会などにもこの方法が使われている。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://kotobank.jp/word/%E3%83%90%E3%83%BC%E3%83%99%E3%82%AD%E3%83%A5%E3%83%BC-116141"
        >kotobank</a
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
    const HOBBY_ID = 3

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

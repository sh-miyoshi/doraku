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
        仏閣巡りとは、たくさんある寺社を徒歩や自転車、車で回って行く趣味です。
        最近話題の 御朱印集め は寺社めぐりの一種です。
        その他にも、狛犬や鳥居、仏像観察など、自分なりの注目ポイントを決めてめぐるとおもしろいですよ。
      </p>
    </div>
    <span class="where">
      from
      <a href="https://shumi.info/299">趣味サーチ</a>
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
    const HOBBY_ID = 21

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

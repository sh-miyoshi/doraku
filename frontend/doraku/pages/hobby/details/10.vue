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
        ホームパーティー (Homeparty)
        は、自宅に友人や知人を招いて行うパーティーのこと。自宅に複数の友人や知人、会社の同僚などを招いて酒食を提供、鍋料理やバーベキューなどでもてなすことを言う。英語で呼ばれているように（英語では
        House party
        と言うが）、基本的にはアメリカやヨーロッパなどの家庭で日常的に行われているものであるが、日本でも昔から「お呼ばれ」という呼称で行われてきたことと基本的に同様のことである。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://ja.wikipedia.org/wiki/%E3%83%9B%E3%83%BC%E3%83%A0%E3%83%91%E3%83%BC%E3%83%86%E3%82%A3%E3%83%BC"
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
    const HOBBY_ID = 10

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

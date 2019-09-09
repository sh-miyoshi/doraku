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
        「これでお酒を作りなさい」そう言って神様が人間に与えた果物、それがぶどうと言われるほど、ぶどうはお酒に適した果物。そんなぶどうからできたお酒、それがワイン。ぶどうをつぶせばちゃんと発酵してお酒になってくれるためワインは古くから世界各地で愛されてきました。
      </p>
    </div>
    <span class="where">
      from
      <a href="https://shumi.info/181">趣味サーチ</a>
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

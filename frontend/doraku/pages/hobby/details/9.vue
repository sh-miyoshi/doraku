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
        俳句と川柳は同じ「俳諧の連歌」から生まれた兄弟のような物です。俳諧の連歌とは、貴族の文化であった連歌を、庶民が自分たちでも遊べるようにバカバカしく滑稽なものに改造したものです。連歌は、前の人が作った五七五の歌に別の人が七七の下の句を付け、さらに別の人がそれに五七五をの句を付けるといったことを繰り返し、３６句、あるいは１００句までで一作品とします。俳句は、この俳諧の連歌の発句（最初の句）が単独で作られるようになったものです。
      </p>
    </div>
    <span class="where">
      from
      <a href="https://jphaiku.jp/how/senn.html">日本俳句研究会</a>
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
    const HOBBY_ID = 9

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

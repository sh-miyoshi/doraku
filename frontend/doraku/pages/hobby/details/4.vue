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
        刺繍とは手芸の一種。広くは糸，紐，リボンなどを，布および皮や編み地などの表面に，針を用いて刺したり留めたりして構成する造形，またはその作品。技法は，アプリケ，ドロンワーク，カットワークのほか，ビーズや真珠，スパングルなどを留めつける方法もある。
      </p>
    </div>
    <span class="where">
      from
      <a href="https://kotobank.jp/word/%E5%88%BA%E7%B9%8D-73282">kotobank</a>
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
    const HOBBY_ID = 4

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

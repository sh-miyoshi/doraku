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
        靴の底にブレードbladeを取り付けた，氷上を滑走するための用具（skate）。また，それを用いてスピード，ジャンプ，演技の美しさなどを競うスポーツ。スケートの歴史は，靴とブレードの改良，リンクの改善，技術の開発，が三位一体となったものである。スケートの競技には，スピードスケート，フィギュアスケート，アイスホッケー，ショートトラックの４種類がある。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://kotobank.jp/word/%E3%82%B9%E3%82%B1%E3%83%BC%E3%83%88-83672"
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
    const HOBBY_ID = 18

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

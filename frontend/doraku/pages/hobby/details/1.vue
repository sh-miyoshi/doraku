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
        バンド (band)
        とは、演奏者の集団である。音楽バンドの略。厳密に決まってはいないが、バンドという語はメンバーの半分以上が楽器を担当するものに限定して使われることが多く、満たない場合は「音楽グループ」「音楽ユニット」などと呼称される。
      </p>
    </div>
    <span class="where">
      from
      <a href="https://dic.nicovideo.jp/a/%E3%83%90%E3%83%B3%E3%83%89"
        >ニコニコ大百科</a
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
    const HOBBY_ID = 1

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

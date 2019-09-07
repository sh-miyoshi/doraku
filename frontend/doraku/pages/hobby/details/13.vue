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
        筋力トレーニング（きんりょくトレーニング）とは、骨格筋の出力・持久力の維持向上や筋肥大を目的とした運動の総称。目的の骨格筋へ抵抗
        (resistance)
        をかけることによって行うものは、レジスタンストレーニングとも呼ばれる。抵抗のかけ方にはさまざまなものがあるが、重力や慣性を利用するものや、ゴムなどによる弾性を利用するもの、油圧や空気圧による抵抗を用いるものが一般的である。重力による抵抗を利用する場合は特に、ウエイトトレーニングとも呼ばれる。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://ja.wikipedia.org/wiki/%E7%AD%8B%E5%8A%9B%E3%83%88%E3%83%AC%E3%83%BC%E3%83%8B%E3%83%B3%E3%82%B0"
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
    const HOBBY_ID = 13

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

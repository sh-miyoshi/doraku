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
        前後に車輪がついた板に乗り、トリック（ジャンプ、空中動作、回転などの技）を行い、その技の難易度や高さ、スピードなどを評価する採点競技。東京2020大会で新たにオリンピック種目に採用される
      </p>
    </div>
    <span class="where">
      from
      <a href="https://tokyo2020.org/jp/games/sport/olympic/skateboarding/"
        >tokyo2020</a
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
    const HOBBY_ID = 19

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

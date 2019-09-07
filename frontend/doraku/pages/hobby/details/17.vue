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
        ロードバイクは、オンロード（舗装路）を高速で、また長距離を長時間走るための自転車です。タイムを競うレースや長距離走行を目的とされているので、重量も走行感も「軽さ」が特徴です。ロングライドやヒルクライム、仲間とレースに参加、サイクリングやポタリング、自転車通勤…などなど、幅広い楽しみ方が出来ますよ。
      </p>
    </div>
    <span class="where">
      from
      <a href="http://www.bex-isoya.com/?page_id=3391">BEXISOYA</a>
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
    const HOBBY_ID = 17

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

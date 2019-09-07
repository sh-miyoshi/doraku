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
        ケイビングは、趣味またはスポーツとして洞窟に入る探検活動である。洞窟探検のアウトドアスポーツとしての面を強調した言葉である。日本では、1930年頃から始まった。
        水中下の洞窟を探検する洞窟潜水では、上記とは異なった特殊な装備や技術、知識が要求される。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://ja.wikipedia.org/wiki/%E3%82%B1%E3%82%A4%E3%83%93%E3%83%B3%E3%82%B0"
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
    const HOBBY_ID = 22

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

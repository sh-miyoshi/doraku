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
        アクアリウム（英: aquarium）は、水生生物の飼育設備を指す。
      </p>
      <p>
        日本では主に、観賞用に熱帯魚（観賞魚）や水草などを飼育・栽培すること、またはそのために構築された水槽を含む環境を指すことが多い。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://ja.wikipedia.org/wiki/%E3%82%A2%E3%82%AF%E3%82%A2%E3%83%AA%E3%82%A6%E3%83%A0"
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
    const HOBBY_ID = 0

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
  text-align: left;
  padding: 2px;
}

.detail > p {
  margin: 0.2em 0px;
}

.where {
  font-size: 80%;
  font-style: italic;
}
</style>

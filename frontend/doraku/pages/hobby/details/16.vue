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
        ラジコンとは、ラジオ コントロール (Radio Control)
        の略称であり、無線により遠隔操作する装置および方式で、RCと記述される場合もあります。模型の自動車、飛行機やヘリコプターなどホビー用が有名ですが、工業や農業など様々な分野で活用されています。
      </p>
    </div>
    <span class="where">
      from
      <a href="http://www.sanwa-denshi.co.jp/rc/support/beginner.html">sanwa</a>
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
    const HOBBY_ID = 16

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

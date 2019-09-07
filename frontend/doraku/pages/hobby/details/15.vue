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
        ピクニックとは、散歩などの途中に野外で食事をすること。自然豊かな場所に出かけて、あらかじめ詰めて運んだ食べ物をそこで食べること。日本語の古語では野掛け（のがけ）という。人間は、建物を作り、この中で生活の様々な用を済ませる。だが、あまりに建物に篭ってばかりではストレスを感じるようにもなる。ピクニックでは、建物から出て、野山や海岸などの自然豊かな場所へ出かけて散歩などを楽しみつつ食事場所へと向かい、自然豊かな場所で食事をする。閉鎖的でマンネリ化した建物内の生活から離れ、開放的な場所で自然を満喫しつつ食事をとると、一般に、食べ物も普段よりおいしく感じられる。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://ja.wikipedia.org/wiki/%E3%83%94%E3%82%AF%E3%83%8B%E3%83%83%E3%82%AF"
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
    const HOBBY_ID = 15

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

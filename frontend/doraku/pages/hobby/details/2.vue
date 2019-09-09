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
        バスケットボールは、1891年にアメリカで考案されたスポーツ。
      </p>
      <p>
        5人対5人の2チームが、一つのボールを手で扱い、長方形のコート上の両端に設置された高さ305cm(10ft)、直径45cm(18in)のリング状のバスケットにボールを上方から通すこと（ゴール）で得点を競う球技である。公式試合は屋内競技として行われる。狭義では、この競技に使用する専用のボールのことを指す。籠球（ろうきゅう）とも訳される。
      </p>
    </div>
    <span class="where">
      from
      <a
        href="https://ja.wikipedia.org/wiki/%E3%83%90%E3%82%B9%E3%82%B1%E3%83%83%E3%83%88%E3%83%9C%E3%83%BC%E3%83%AB"
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
    const HOBBY_ID = 2

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

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
        日本では、温泉は「温泉法」という法律によって「地中からゆう出する温水、鉱水及び水蒸気その他のガス（炭化水素を主成分とする天然ガスを除く。）で、別表①に掲げる温度又は物質を有するもの」と定義されています。したがって、地中からゆう出した時の温度が、25℃以上あれば温泉となり、また25℃未満であっても別表の物質（19項目のうちいずれかひとつ以上）が規定量含まれていれば、温泉となります。また条件を満たせば水蒸気やガスも温泉となります。
      </p>
    </div>
    <span class="where">
      from
      <a href="https://www.spa.or.jp/onsen/477/">日本温泉協会</a>
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
    const HOBBY_ID = 11

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

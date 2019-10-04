<template>
  <div>
    <br />
    <br />
    <br />
    <br />
    <br />
    <h2>今日の趣味はこれ！</h2>
    <nuxt-link :to="url" class="hobby_link">
      <h1 class="hobby_name">
        {{ name }}
      </h1>
    </nuxt-link>
  </div>
</template>

<script>
export default {
  data() {
    return {
      name: '',
      url: '/'
    }
  },

  async created() {
    const res = await this.$hobby.GetTodayHobby()
    console.log(res)
    if (res.error) {
      console.log('Failed to get today hobby: %o', res.error)
      this.$router.app.error({
        statusCode: 500,
        message: 'Internal Server Error'
      })
      return
    }

    this.url = `/hobby/details/${res.data.id}`
    this.name = res.data.name
  }
}
</script>

<style scoped>
.hobby_name {
  background: #f5deb3;
  box-shadow: 0px 0px 0px 5px #f5deb3;
  border: dashed 1px #cd853f;
  padding: 0.2em 0.5em;
  color: #454545;
  width: 400px;
  border-radius: 5px;
  margin-left: auto;
  margin-right: auto;
}

.hobby_link {
  text-decoration: none;
}
</style>

<template>
  <div>
    <br />
    <br />
    <br />
    <br />
    <br />
    <h2>今あなたにオススメの趣味はこれ！</h2>
    <nuxt-link :to="url" class="hobby_link">
      <h1 class="hobby_name">
        {{ $route.query.name }}
      </h1>
    </nuxt-link>
  </div>
</template>

<script>
export default {
  data() {
    return {
      url: `/hobby/details/${this.$route.query.id}`
    }
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

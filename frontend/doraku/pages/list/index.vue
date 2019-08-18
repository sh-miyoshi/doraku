<template>
  <div>
    <br />
    <br />
    <br />
    <br />
    <h2>趣味一覧</h2>
    <div class="itembox">
      <!-- for colums -->
      <div v-for="(hobbyList, index) in hobbies" :key="index" class="component">
        <ul>
          <li v-for="hobby in hobbyList" :key="hobby.id">
            {{ hobby.name }}
          </li>
        </ul>
      </div>
      <div class="footer"></div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      hobbies: []
    }
  },

  created() {
    const COLUMN_NUM = 2

    // set hobby list data
    const allHobby = this.$hobby.GetAllHobby()
    for (let i = 0; i < COLUMN_NUM; i++) {
      this.hobbies[i] = []
    }
    for (let i = 0; i < allHobby.length; i++) {
      this.hobbies[i % COLUMN_NUM].push(allHobby[i])
    }
  }
}
</script>

<style scoped>
.itembox {
  width: 700px;
  margin-left: auto;
  margin-right: auto;
}

.component {
  text-align: left;
  float: left;
  width: 350px;
}

.footer {
  clear: both;
}
</style>

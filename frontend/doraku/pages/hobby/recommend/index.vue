<template>
  <div>
    <br />
    <br />
    <br />
    <h1>質問</h1>
    <section v-if="nowQuestion == 0">
      <h2>Q1: どっち派？</h2>
      <table class="select" cellPadding="10">
        <tbody>
          <tr>
            <td>
              <input id="q1t" v-model="outdoor" type="radio" value="true" />
              <label for="q1t">アウトドア派</label>
            </td>
            <td>
              <input id="q1f" v-model="outdoor" type="radio" value="false" />
              <label for="q1f">インドア派</label>
            </td>
          </tr>
        </tbody>
      </table>
      <button class="disabled_button" disabled>前の質問へ</button>
      <button class="move_button" @click="next()">次の質問へ</button>
    </section>

    <section v-if="nowQuestion == 1">
      <h2>Q2: 気分的には？</h2>
      <table class="select" cellPadding="10">
        <tbody>
          <tr>
            <td>
              <input id="q2t" v-model="alone" type="radio" value="true" />
              <label for="q2t">一人な気分</label>
            </td>
            <td>
              <input id="q2f" v-model="alone" type="radio" value="false" />
              <label for="q2f">みんなで集まりたい気分</label>
            </td>
          </tr>
        </tbody>
      </table>
      <button class="move_button" @click="back()">前の質問へ</button>
      <button class="move_button" @click="next()">次の質問へ</button>
    </section>

    <section v-if="nowQuestion == 2">
      <h2>Q3: どんな感じがタイプ？</h2>
      <table class="select" cellPadding="10">
        <tbody>
          <tr>
            <td>
              <input id="q3t" v-model="active" type="radio" value="true" />
              <label for="q3t">激しい感じで</label>
            </td>
            <td>
              <input id="q3f" v-model="active" type="radio" value="false" />
              <label for="q3f">落ち着いた感じで</label>
            </td>
          </tr>
        </tbody>
      </table>
      <button class="move_button" @click="back()">前の質問へ</button>
      <button class="disabled_button" disabled>次の質問へ</button>
      <br />
      <br />
      <button class="enter_button" @click="recommend">診断</button>
    </section>
  </div>
</template>

<script>
export default {
  data() {
    return {
      outdoor: 'true',
      alone: 'true',
      active: 'true',
      nowQuestion: 0
    }
  },

  methods: {
    back() {
      if (this.nowQuestion >= 1) {
        this.nowQuestion--
      }
    },

    next() {
      if (this.nowQuestion < 2) {
        this.nowQuestion++
      }
    },

    async recommend() {
      const res = await this.$hobby.GetRecommendHobby(
        this.outdoor,
        this.alone,
        this.active
      )
      console.log(res) // for debug

      if (res.error) {
        console.log('Failed to get recommend hobby: %o', res.error)
        this.$router.app.error({
          statusCode: 500,
          message: 'Internal Server Error'
        })
        return
      }

      const id = res.data.id
      const name = res.data.name

      this.$router.push({
        path: '/hobby/recommend/result',
        query: {
          id,
          name
        }
      })
    }
  }
}
</script>

<style scoped>
.enter_button {
  border: 0px;
  width: 250px;
  height: 75px;
  background: url('~assets/button.png') left top no-repeat;
  color: #0f0f0f;
  font-size: 30px;
}

.enter_button:hover {
  cursor: pointer;
  opacity: 0.7;
}

.select {
  margin-left: auto;
  margin-right: auto;
}

.move_button {
  font-size: 18px;
  width: 200px;
  display: inline-block;
  padding: 0.5em 1em;
  text-decoration: none;
  color: #65c3fd;
  border: dashed 1px #65c3fd;
  background: #f2fcff;
  border-radius: 3px;
  transition: 0.4s;
}

.move_button:hover {
  background: #cbedff;
  color: #fff;
}

.disabled_button {
  font-size: 18px;
  width: 200px;
  display: inline-block;
  padding: 0.5em 1em;
  text-decoration: none;
  color: #bfbfbf;
  border: dashed 1px #bfbfbf;
  background: #f0f0f0;
  border-radius: 3px;
  transition: 0.4s;
}
</style>

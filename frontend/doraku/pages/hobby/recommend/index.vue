<template>
  <div>
    <br />
    <br />
    <br />
    <h1>質問</h1>
    <div v-if="nowQuestion == 0">
      <h2>Q1: どっち派？</h2>
      <table class="select" cellPadding="10">
        <tbody>
          <tr>
            <td>
              <input v-model="outdoor" type="radio" name="q1" value="true" />
              アウトドア派
            </td>
            <td>
              <input v-model="outdoor" type="radio" name="q1" value="false" />
              インドア派
            </td>
          </tr>
        </tbody>
      </table>
      <button class="disabled_button" disabled>前の質問へ</button>
      <button class="move_button" @click="next()">次の質問へ</button>
    </div>

    <div v-if="nowQuestion == 1">
      <h2>Q2: 気分的には？</h2>
      <table class="select" cellPadding="10">
        <tbody>
          <tr>
            <td>
              <input v-model="alone" type="radio" name="q2" value="true" />
              一人な気分
            </td>
            <td>
              <input v-model="alone" type="radio" name="q2" value="false" />
              みんなで集まりたい気分
            </td>
          </tr>
        </tbody>
      </table>
      <button class="move_button" @click="back()">前の質問へ</button>
      <button class="move_button" @click="next()">次の質問へ</button>
    </div>

    <div v-if="nowQuestion == 2">
      <h2>Q3: どんな感じがタイプ？</h2>
      <table class="select" cellPadding="10">
        <tbody>
          <tr>
            <td>
              <input v-model="active" type="radio" name="q3" value="true" />
              激しい感じで
            </td>
            <td>
              <input v-model="active" type="radio" name="q3" value="false" />
              落ち着いた感じで
            </td>
          </tr>
        </tbody>
      </table>
      <button class="move_button" @click="back()">前の質問へ</button>
      <button class="disabled_button" disabled>次の質問へ</button>
      <br />
      <br />
      <button class="enter_button">診断</button>
    </div>
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

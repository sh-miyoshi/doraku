<template>
  <div>
    <img src="~/assets/top_logo.png" alt="logo" width="400" />
    <p>さぁ、新しい人生を切り開こう！</p>
    <p>
      Dorakuは皆様の日常に新しい変化を付け加えるため、あなたにあった趣味を提案します
    </p>
    <p>
      日々の生活に新しい変化を求めているなら、ぜひこのサイトで新しい趣味を探してみてください！
    </p>
    <br />
    <table cellPadding="5" align="center">
      <tbody>
        <tr>
          <td>
            <Button class="button">
              趣味診断
            </Button>
          </td>
          <td>
            <Button class="button">
              今日の趣味
            </Button>
          </td>
          <td>
            <Button class="button" @click="$router.push('/hobby/list')">
              趣味一覧
            </Button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {}
</script>

<style>
.button {
  border: 0px;
  width: 250px;
  height: 75px;
  background: url('~assets/button.png') left top no-repeat;
  color: #0f0f0f;
  font-size: 30px;
}

.button:hover {
  cursor: pointer;
  opacity: 0.7;
}
</style>

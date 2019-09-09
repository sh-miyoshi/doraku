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
            <Button
              v-tooltip="'いくつかの質問からあなたにあった趣味を提案いたします'"
              class="button"
              @click="$router.push('/hobby/recommend')"
            >
              趣味診断
            </Button>
          </td>
          <td>
            <Button
              v-tooltip="'今人気の趣味を紹介します'"
              class="button"
              @click="$router.push('/hobby/today')"
            >
              今日の趣味
            </Button>
          </td>
          <td>
            <Button
              v-tooltip="'Dorakuで管理している趣味の一覧を表示します'"
              class="button"
              @click="$router.push('/hobby/list')"
            >
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

.tooltip {
  display: block !important;
  z-index: 10000;
  font-size: 14px;
}

.tooltip .tooltip-inner {
  background: #2f2f2f;
  color: white;
  border-radius: 12px;
  padding: 10px;
}

.tooltip .tooltip-arrow {
  width: 0;
  height: 0;
  border-style: solid;
  position: absolute;
  margin: 5px;
  border-color: black;
  z-index: 1;
}

.tooltip[x-placement^='top'] {
  margin-bottom: 5px;
}

.tooltip[x-placement^='top'] .tooltip-arrow {
  border-width: 5px 5px 0 5px;
  border-left-color: transparent !important;
  border-right-color: transparent !important;
  border-bottom-color: transparent !important;
  bottom: -5px;
  left: calc(50% - 5px);
  margin-top: 0;
  margin-bottom: 0;
}

.tooltip[x-placement^='bottom'] {
  margin-top: 5px;
}

.tooltip[x-placement^='bottom'] .tooltip-arrow {
  border-width: 0 5px 5px 5px;
  border-left-color: transparent !important;
  border-right-color: transparent !important;
  border-top-color: transparent !important;
  top: -5px;
  left: calc(50% - 5px);
  margin-top: 0;
  margin-bottom: 0;
}

.tooltip[x-placement^='right'] {
  margin-left: 5px;
}

.tooltip[x-placement^='right'] .tooltip-arrow {
  border-width: 5px 5px 5px 0;
  border-left-color: transparent !important;
  border-top-color: transparent !important;
  border-bottom-color: transparent !important;
  left: -5px;
  top: calc(50% - 5px);
  margin-left: 0;
  margin-right: 0;
}

.tooltip[x-placement^='left'] {
  margin-right: 5px;
}

.tooltip[x-placement^='left'] .tooltip-arrow {
  border-width: 5px 0 5px 5px;
  border-top-color: transparent !important;
  border-right-color: transparent !important;
  border-bottom-color: transparent !important;
  right: -5px;
  top: calc(50% - 5px);
  margin-left: 0;
  margin-right: 0;
}

.tooltip.popover .popover-inner {
  background: #f9f9f9;
  color: black;
  padding: 24px;
  border-radius: 5px;
  box-shadow: 0 5px 30px rgba(black, 0.6);
}

.tooltip.popover .popover-arrow {
  border-color: #f9f9f9;
}

.tooltip[aria-hidden='true'] {
  visibility: hidden;
  opacity: 0;
  transition: opacity 0.15s, visibility 0.15s;
}

.tooltip[aria-hidden='false'] {
  visibility: visible;
  opacity: 1;
  transition: opacity 0.15s;
}
</style>

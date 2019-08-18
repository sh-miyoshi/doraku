class HobbyHandler {
  hobbyList = [
    { id: 0, name: 'アクアリウム' },
    { id: 1, name: 'バンド' },
    { id: 2, name: 'バスケットボール' },
    { id: 3, name: 'バーベキュー' },
    { id: 4, name: '刺繍' },
    { id: 5, name: '釣り' },
    { id: 6, name: 'フットサル' },
    { id: 7, name: 'テレビゲーム' },
    { id: 8, name: 'ゲートボール' },
    { id: 9, name: '俳句' },
    { id: 10, name: 'ホームパーティー' },
    { id: 11, name: '温泉ツアー' },
    { id: 12, name: '将棋' },
    { id: 13, name: '筋トレ' },
    { id: 14, name: 'ネイルアート' },
    { id: 15, name: 'ピクニック' },
    { id: 16, name: 'ラジコン' },
    { id: 17, name: 'ロードバイク' },
    { id: 18, name: 'スケート' },
    { id: 19, name: 'スケートボード' },
    { id: 20, name: '水泳' },
    { id: 21, name: '仏閣巡り' },
    { id: 22, name: 'ケイビング' },
    { id: 23, name: 'ワイン' },
    { id: 24, name: 'キャンプ' }
  ]

  GetAllHobby = function() {
    return this.hobbyList
  }

  GetHobbyByID = function(id) {
    if (id >= 0 && id < this.hobbyList.length) {
      return this.hobbyList[id]
    }
    return {}
  }
}

export default ({ app }, inject) => {
  inject('hobby', new HobbyHandler())
}

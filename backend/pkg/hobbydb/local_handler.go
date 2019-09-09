package hobbydb

import (
	"fmt"
	"math/rand"

	"github.com/sh-miyoshi/doraku/pkg/logger"
)

// localDBHandler implements DBHandler
type localDBHandler struct {
	DBHandler

	data []HobbyDB
}

func b2i(v bool) int64 {
	if v {
		return 1
	}
	return 0
}

func getGroupNo(outdoor, alone, active bool) int64 {
	return (b2i(outdoor) << 2) + (b2i(alone) << 1) + b2i(active)
}

func (h *localDBHandler) Initialize(connStr string) error {
	// initialize database
	h.data = []HobbyDB{
		{ID: 0, Name: "アクアリウム", GroupNo: getGroupNo(false, true, false)},
		{ID: 1, Name: "バンド", GroupNo: getGroupNo(false, false, false)},
		{ID: 2, Name: "バスケットボール", GroupNo: getGroupNo(false, false, true)},
		{ID: 3, Name: "バーベキュー", GroupNo: getGroupNo(true, false, false)},
		{ID: 4, Name: "刺繍", GroupNo: getGroupNo(false, true, false)},
		{ID: 5, Name: "釣り", GroupNo: getGroupNo(true, true, false)},
		{ID: 6, Name: "フットサル", GroupNo: getGroupNo(true, false, true)},
		{ID: 7, Name: "テレビゲーム", GroupNo: getGroupNo(false, true, false)},
		{ID: 8, Name: "ゲートボール", GroupNo: getGroupNo(true, false, false)},
		{ID: 9, Name: "俳句", GroupNo: getGroupNo(false, true, false)},
		{ID: 10, Name: "ホームパーティー", GroupNo: getGroupNo(false, false, false)},
		{ID: 11, Name: "温泉ツアー", GroupNo: getGroupNo(true, true, false)},
		{ID: 12, Name: "将棋", GroupNo: getGroupNo(false, false, false)},
		{ID: 13, Name: "筋トレ", GroupNo: getGroupNo(false, true, true)},
		{ID: 14, Name: "ネイルアート", GroupNo: getGroupNo(false, true, false)},
		{ID: 15, Name: "ピクニック", GroupNo: getGroupNo(true, false, false)},
		{ID: 16, Name: "ラジコン", GroupNo: getGroupNo(true, true, false)},
		{ID: 17, Name: "ロードバイク", GroupNo: getGroupNo(true, true, true)},
		{ID: 18, Name: "スケート", GroupNo: getGroupNo(true, false, true)},
		{ID: 19, Name: "スケートボード", GroupNo: getGroupNo(true, true, true)},
		{ID: 20, Name: "水泳", GroupNo: getGroupNo(false, true, true)},
		{ID: 21, Name: "仏閣巡り", GroupNo: getGroupNo(true, true, false)},
		{ID: 22, Name: "ケイビング", GroupNo: getGroupNo(true, false, true)},
		{ID: 23, Name: "ワイン", GroupNo: getGroupNo(false, true, false)},
		{ID: 24, Name: "キャンプ", GroupNo: getGroupNo(true, false, false)},
	}

	logger.Info("Successfully initialize DB")
	return nil
}

func (h *localDBHandler) GetRecommendHobby(input InputValue) (HobbyDB, error) {
	var candidates []HobbyDB
	no := getGroupNo(input.Outdoor, input.Alone, input.Active)
	logger.Info("Input GroupNo: %d", no)

	for _, hobby := range h.data {
		if hobby.GroupNo == int64(no) {
			candidates = append(candidates, hobby)
		}
	}

	logger.Debug("Candidates of recommended: %v", candidates)

	if len(candidates) == 0 {
		return HobbyDB{}, fmt.Errorf("No recomend hobby")
	}

	return candidates[rand.Intn(len(candidates))], nil
}

func (h *localDBHandler) GetHobbyNum() int {
	return len(h.data)
}

func (h *localDBHandler) GetHobbyByID(id int) (HobbyDB, error) {
	if id < 0 || id >= len(h.data) {
		return HobbyDB{}, fmt.Errorf("no such hobby ID: %d", id)
	}

	return h.data[id], nil
}

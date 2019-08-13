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

func b2i(v bool) int {
	if v {
		return 1
	}
	return 0
}

func (h *localDBHandler) Initialize(connStr string) error {
	// initialize database
	h.data = []HobbyDB{
		{},
	}

	logger.Info("Successfully initialize DB")
	return nil
}

func (h *localDBHandler) GetRecommendHobby(input InputValue) (HobbyDB, error) {
	var candidates []HobbyDB
	no := (b2i(input.Outdoor) << 2) + (b2i(input.Alone) << 1) + b2i(input.Active)
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

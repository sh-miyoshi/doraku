package hobbydb

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"

	"github.com/sh-miyoshi/doraku/pkg/logger"
)

// DBHandler is interface of dbHandler
type DBHandler interface {
	Initialize(filePath string) error
	GetAllHobby() []HobbyDB
	GetHobbyByID(id int) (HobbyDB, error)
	GetRecommendedHobby(input InputValue) (HobbyDB, error)
	GetHobbyNum() int
}

// dbHandler implements DBHandler
type dbHandler struct {
	DBHandler

	data []HobbyDB
}

var inst = &dbHandler{}

// GetInst return instance of Database Handler
func GetInst() DBHandler {
	return inst
}

func b2i(v bool) int {
	if v {
		return 1
	}
	return 0
}

func (h *dbHandler) Initialize(filePath string) error {
	fp, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.Comment = '#'
	for {
		data, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		// Caution: The order cannot be changed
		if len(data) != 4 {
			return fmt.Errorf("%s file is maybe broken. we expect 4 data, but got %d", filePath, len(data))
		}
		tmp := HobbyDB{}
		tmp.ID, err = strconv.Atoi(data[0])
		if err != nil {
			return err
		}
		tmp.Name = data[1]
		tmp.NameEN = data[2]
		tmp.GroupNo, err = strconv.ParseInt(data[3], 2, 4)
		if err != nil {
			return err
		}
		h.data = append(h.data, tmp)
	}
	logger.Debug("DB data: %v", h.data)

	logger.Info("Successfully initialize DB")
	return nil
}

func (h *dbHandler) GetAllHobby() []HobbyDB {
	return h.data
}

func (h *dbHandler) GetHobbyByID(id int) (HobbyDB, error) {
	for _, hobby := range h.data {
		if hobby.ID == id {
			return hobby, nil
		}
	}
	return HobbyDB{}, fmt.Errorf("No such hobby, ID: %d", id)
}

func (h *dbHandler) GetRecommendedHobby(input InputValue) (HobbyDB, error) {
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
		return HobbyDB{}, fmt.Errorf("No recomended hobby")
	}

	return candidates[rand.Intn(len(candidates)-1)], nil
}

func (h *dbHandler) GetHobbyNum() int {
	return len(h.data)
}

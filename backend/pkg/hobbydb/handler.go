package hobbydb

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/sh-miyoshi/doraku/pkg/logger"
)

// DBHandler is interface of dbHandler
type DBHandler interface {
	Initialize() error
	GetAllHobby() []HobbyDB
	GetHobbyByID(id int) (HobbyDB, error)
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

func (h *dbHandler) Initialize() error {
	const filePath = "hobby.csv"

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
		tmp.GroupNo, err = strconv.Atoi(data[3])
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

func (h *dbHandler) GetHobbyNum() int {
	return len(h.data)
}

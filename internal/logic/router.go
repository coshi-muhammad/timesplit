package logic

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/coshi-muhammd/timesplit/internal/core"
	"github.com/google/uuid"
)

var _ core.RouterLogic = (*RouterService)(nil)

type RouterService struct {
	app_state core.AppState
}

func NewRouterService(app_state core.AppState) *RouterService {
	return &RouterService{app_state: app_state}
}

// dont forget to wrap it before use so you get access to the functions on it
func (rs *RouterService) GetActiveSplit() *core.Split {
	return rs.app_state.Active_Split
}

func (rs *RouterService) SetActiveSplit(id uuid.UUID) {
	rs.app_state.Active_Split = rs.app_state.Split_List[id]
}

func (rs *RouterService) LoadSplits() {
	files, _ := os.ReadDir("./")
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), ".json") {
			data, _ := os.ReadFile("./" + file.Name())
			split := &core.Split{}
			err := json.Unmarshal(data, &split)
			if err != nil {
				fmt.Printf("error : %v file %v", err, file.Name())
			}
			rs.app_state.Split_List[split.Id] = split
		}
	}
}
func (rs *RouterService) GetState() *core.AppState {
	return &rs.app_state
}

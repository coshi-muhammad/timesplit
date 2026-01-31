package logic

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
	rs.app_state.Config.Active_Split_id = id
}

func (rs *RouterService) LoadSplits() {
	dataDir := GetDataDirectory()
	files, err := os.ReadDir(dataDir)
	if err != nil {
		fmt.Println("an error acured:", err)
		return
	}
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), ".json") {
			data, _ := os.ReadFile(filepath.Join(dataDir, file.Name()))
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

func (rs *RouterService) LoadConfig() error {
	conf_dir, err := os.UserConfigDir()
	timesplit_conf := filepath.Join(conf_dir, "timesplit")
	if err != nil {
		return err
	}
	config, err := os.ReadFile(filepath.Join(timesplit_conf, "config.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(config, &rs.app_state.Config)
	if err != nil {
		return err
	}
	fmt.Println(rs.app_state.Config)
	return nil
}
func (rs *RouterService) StoreConfig() error {
	json, err := json.Marshal(rs.app_state.Config)
	if err != nil {
		fmt.Println("error 1")
		return err
	}
	conf_dir, err := os.UserConfigDir()
	timesplit_conf := filepath.Join(conf_dir, "timesplit")
	os.MkdirAll(timesplit_conf, 0755)
	if err != nil {
		fmt.Println("error 2")
		return err
	}
	err = os.WriteFile(filepath.Join(timesplit_conf, "config.json"), json, 0644)
	if err != nil {
		fmt.Println("error 3", err)
		return err
	}
	return nil
}
func GetDataDirectory() string {
	var dataDir string
	home, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "linux":
		// Check XDG_DATA_HOME first, fallback to ~/.local/share
		dataDir = os.Getenv("XDG_DATA_HOME")
		if dataDir == "" {
			dataDir = filepath.Join(home, ".local", "share")
		}
	case "darwin":
		// macOS standard for app data
		dataDir = filepath.Join(home, "Library", "Application Support")
	case "windows":
		// Windows standard (usually C:\Users\User\AppData\Roaming)
		dataDir = os.Getenv("AppData")
	}
	dataDir = filepath.Join(dataDir, "timesplit")
	return dataDir
}

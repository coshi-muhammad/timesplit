package core

import (
	"image/color"

	"github.com/google/uuid"
)

const Max_time = uint64(24 * 60 * 60)

type Logic interface {
}

type RouterLogic interface {
	GetState() *AppState
	SetActiveSplit(uuid.UUID)
	GetActiveSplit() *Split
	LoadSplits()
}

type FormLogic interface {
	Logic
	AddSectionToSplit()
	Log()
	Submit() error
}

type DayLogic interface {
	Logic
	Log()
	SetActiveSplit()
}

// TODO: improve the empty view to make it more compeling
type EmptyLogic interface {
	Logic
}

type Section struct {
	Id          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Start_t     uint64     `json:"start"`
	End_t       uint64     `json:"end"`
	Color       color.RGBA `json:"color"`
	Description string     `json:"description"`
	Atachements []string   `json:"Atachements"`
}

type Split struct {
	Id          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Week_format bool       `json:"week"`
	Sections    []*Section `json:"sections"`
}

// TODO: when making the day/week view make it that this thing decieds what is the active split on start up
type Configuration struct {
	Active_Split_id uuid.UUID
}

type AppState struct {
	Config       Configuration
	Split_List   map[uuid.UUID]*Split
	Active_Split *Split
}

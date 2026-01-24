package logic

import (
	"fmt"
	"github.com/coshi-muhammd/timesplit/internal/core"
	"github.com/google/uuid"
)

var _ core.DayLogic = (*DayService)(nil)

type DayService struct {
	router          *RouterService
	Visible_split   *SplitService
	Visible_section *SectionService
}

func NewDayService(r *RouterService, sp *SplitService) *DayService {
	ds := &DayService{
		router:          r,
		Visible_split:   sp,
		Visible_section: sp.GetActiveSection(),
	}
	return ds
}

func (ds *DayService) Log() {
	fmt.Printf(`
		Day view is showing split %s with section %s
		With active split being %s and active section being %s
		`,
		ds.Visible_split.Title,
		ds.Visible_section.Title,
		ds.router.GetActiveSplit().Title,
		ds.router.GetActiveSplit().Title)
}
func (ds *DayService) SetVisibleSplit(id uuid.UUID) {
	for _, split := range ds.router.app_state.Split_List {
		if split.Id == id {
			ds.Visible_split = WrapSplit(split)
			return
		}
	}
}

func (ds *DayService) SetActiveSplit() {
	ds.router.SetActiveSplit(ds.Visible_split.Id)
}

func (ds *DayService) SetVisibleSection() {
	ds.Visible_section = ds.Visible_split.GetActiveSection()
}

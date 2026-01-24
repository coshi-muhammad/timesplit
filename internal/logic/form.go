package logic

import (
	"fmt"
	"github.com/coshi-muhammd/timesplit/internal/core"
)

type FormService struct {
	router       core.RouterLogic
	Temp_section *SectionService
	Temp_split   *SplitService
}

var _ core.FormLogic = (*FormService)(nil)

// WARN: i made a modification to this function by changing its signature here and havent come around to testing it yet when you are able do
func NewFormService(router core.RouterLogic, title string, week bool) *FormService {
	return &FormService{router: router,
		Temp_split:   NewSplitService(title, week),
		Temp_section: NewStubSection(),
	}
}

func (f FormService) Log() {
	formated_start, formated_end, _ := f.Temp_section.GetFormatedTimestamps()
	fmt.Printf(`
		section %s with values:
		id:%v
		color:%v
		start:%s 
		end:%s
		atachmentlist:%v
		was added to split %s
		`,
		f.Temp_section.Title,
		f.Temp_section.Id,
		f.Temp_section.Color,
		formated_start,
		formated_end,
		f.Temp_section.Atachements,
		f.Temp_split.Title,
	)
}

func (f FormService) Submit() error {
	err := f.Temp_split.Store()
	if err != nil {
		return err
	}
	f.router.LoadSplits()
	f.router.SetActiveSplit(f.Temp_section.Id)
	f.Temp_split = &SplitService{}
	return nil
}

func (f FormService) AddSectionToSplit() {
	f.Temp_split.AddSection(*f.Temp_section)
	f.Log()
	f.Temp_section = &SectionService{}
}

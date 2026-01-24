package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"time"

	"github.com/coshi-muhammd/timesplit/internal/core"
	"github.com/google/uuid"
)

type SplitService struct {
	*core.Split
}

func NewSplitService(title string, week bool) *SplitService {
	fmt.Println(title)
	id, _ := uuid.NewUUID()
	return &SplitService{
		Split: &core.Split{
			Id:          id,
			Title:       title,
			Week_format: week,
			Sections:    make([]*core.Section, 0),
		},
	}
}

func WrapSplit(split *core.Split) *SplitService {
	return &SplitService{Split: split}
}
func (sp *SplitService) AddSection(sc SectionService) error {
	if sc.Start_t >= core.Max_time || sc.End_t >= core.Max_time || sc.Start_t >= sc.End_t {
		return fmt.Errorf("Imposible range %v : %v ", sc.Start_t, sc.End_t)
	}
	for _, spc := range sp.Sections {
		if (sc.Start_t >= spc.Start_t && sc.Start_t < spc.End_t) ||
			(sc.End_t > spc.Start_t && sc.End_t <= spc.End_t) {
			return fmt.Errorf("Section Overlaps an existing one: %v ", spc.Title)
		}
	}
	sp.Sections = append(sp.Sections, sc.Section)
	slices.SortFunc(sp.Sections, func(a *core.Section, b *core.Section) int {
		if a.Start_t > b.Start_t {
			return 1
		} else {
			return -1
		}
	})
	return nil
}
func (sp *SplitService) RemoveSection(sc *SectionService) {
	sp.Sections = slices.DeleteFunc(sp.Sections, func(spc *core.Section) bool {
		if sc.Id == spc.Id {
			return true
		} else {
			return false
		}
	})
}
func (sp SplitService) Store() error {
	json, err := json.Marshal(sp)
	if err != nil {
		return err
	}
	err = os.WriteFile(sp.Title+".json", json, 0644)
	if err != nil {
		return err
	}
	return nil
}
func (sp *SplitService) GetActiveSection() *SectionService {
	current_t, err := DecodeTime(time.Now().Format("15:04:05"))
	fmt.Println("time:", current_t, "the string:", time.Now().Format("15:04:05"))
	if err != nil {
		log.Fatal(err)
	}
	for _, section := range sp.Sections {
		if current_t >= section.Start_t && current_t <= section.End_t {
			return WrapSection(section)
		}
	}
	return NewStubSection()
}

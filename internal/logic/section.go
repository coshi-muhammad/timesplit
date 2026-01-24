package logic

import (
	"fmt"
	"image/color"
	"time"

	"github.com/coshi-muhammd/timesplit/internal/core"
	"github.com/google/uuid"
)

type SectionService struct {
	*core.Section
}

func EncodeTime(time uint64) (string, error) {
	if time > core.Max_time {
		return "", fmt.Errorf("the time provided %d is too big", time)
	}
	hour := int64(time / (60 * 60))
	minute := int64((time / 60) % 60)
	second := int64(time % 60)
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second), nil
}

func DecodeTime(time_string string) (uint64, error) {
	t, err := time.Parse("15:04:05", time_string)
	if err != nil {
		return 0, err
	}
	return uint64(t.Hour()*3600 + t.Minute()*60 + t.Second()), nil
}

func NewSection(title, start, end, description string,
	color color.RGBA, atachements []string) (*SectionService, error) {
	start_t, err := DecodeTime(start)
	if err != nil {
		return nil, err
	}
	end_t, err := DecodeTime(end)
	if err != nil {
		return nil, err
	}
	id, _ := uuid.NewUUID()
	return &SectionService{
		Section: &core.Section{
			Id:          id,
			Title:       title,
			Start_t:     start_t,
			End_t:       end_t,
			Description: description,
			Color:       color,
			Atachements: atachements,
		},
	}, nil
}
func NewStubSection() *SectionService {
	return &SectionService{Section: &core.Section{}}
}
func (s SectionService) GetFormatedTimestamps() (string, string, error) {
	start, err := EncodeTime(s.Start_t)
	if err != nil {
		return "", "", err
	}
	end, err := EncodeTime(s.End_t)
	if err != nil {
		return "", "", err
	}
	return start, end, nil
}
func WrapSection(section *core.Section) *SectionService {
	return &SectionService{Section: section}
}

package ui

import (
	"fmt"
	"image/color"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/coshi-muhammd/timesplit/internal/core"
	"github.com/coshi-muhammd/timesplit/internal/logic"
	"github.com/google/uuid"
)

type SplitWidget struct {
	widget.BaseWidget
	split      *logic.SplitService
	timeStamps []string
}

var _ fyne.Widget = (*SplitWidget)(nil)

type SectionWidget struct {
	background *canvas.Rectangle
	text       *widget.Label
	container  *fyne.Container
	section    *logic.SectionService
}
type TimeStamp struct {
	centered_label *fyne.Container
	line           *widget.Separator
}
type SplitRenderer struct {
	split      *SplitWidget
	background canvas.Rectangle
	sections   map[uuid.UUID]SectionWidget
	timeStamps map[string]TimeStamp
	objects    []fyne.CanvasObject
}

func NewSplitWidget(split *logic.SplitService) *SplitWidget {
	sw := &SplitWidget{}
	sw.ExtendBaseWidget(sw)
	fmt.Println(split)
	sw.split = split
	for _, sc := range sw.split.Sections {
		start, err := logic.EncodeTime(sc.Start_t)
		if err != nil {
			fmt.Printf("we have error %v", err)
			return nil
		}
		end, err := logic.EncodeTime(sc.End_t)
		if err != nil {
			fmt.Printf("we have error %v", err)
			return nil
		}
		if !slices.Contains(sw.timeStamps, start) {
			sw.timeStamps = append(sw.timeStamps, start)
		}
		if !slices.Contains(sw.timeStamps, end) {
			sw.timeStamps = append(sw.timeStamps, end)
		}
	}
	return sw
}

var _ fyne.WidgetRenderer = (*SplitRenderer)(nil)

func (sw *SplitWidget) CreateRenderer() fyne.WidgetRenderer {
	sr := SplitRenderer{
		sections:   make(map[uuid.UUID]SectionWidget),
		timeStamps: make(map[string]TimeStamp),
	}
	sr.split = sw
	sr.background = *canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	sr.background.CornerRadius = 12
	sr.background.StrokeWidth = 3
	sr.background.StrokeColor = color.RGBA{0x8E, 0xAA, 0xCD, 0xFF}
	for _, sc := range sw.split.Sections {
		background := canvas.NewRectangle(sc.Color)
		text := widget.NewLabel(sc.Title)
		sr.sections[sc.Id] = SectionWidget{
			section:    logic.WrapSection(sc),
			background: background,
			text:       text,
			container:  container.NewStack(background, container.NewCenter(text)),
		}
	}
	for _, sc := range sr.split.split.Sections {
		start, err := logic.EncodeTime(sc.Start_t)
		if err != nil {
			fmt.Printf("we have error %v", err)
			return nil
		}
		end, err := logic.EncodeTime(sc.End_t)
		if err != nil {
			fmt.Printf("we have error %v", err)
			return nil
		}
		sr.timeStamps[start] = TimeStamp{
			centered_label: container.NewCenter(widget.NewLabel(start)),
			line:           widget.NewSeparator(),
		}
		sr.timeStamps[end] = TimeStamp{
			centered_label: container.NewCenter(widget.NewLabel(end)),
			line:           widget.NewSeparator(),
		}
	}
	return &sr
}

func (sr *SplitRenderer) Destroy() {
}
func (sr *SplitRenderer) Objects() []fyne.CanvasObject {
	return sr.objects
}
func (sr *SplitRenderer) Refresh() {
	for _, sc := range sr.split.split.Sections {
		_, ok := sr.sections[sc.Id]
		if ok {
			sr.sections[sc.Id].background.FillColor = sc.Color
			sr.sections[sc.Id].text.Text = sc.Title
			sr.sections[sc.Id].background.Refresh()
			sr.sections[sc.Id].text.Refresh()
			sr.sections[sc.Id].container.Refresh()
		} else {
			background := canvas.NewRectangle(sc.Color)
			text := widget.NewLabel(sc.Title)
			text.Alignment = fyne.TextAlignCenter
			centered_text := container.NewVBox(layout.NewSpacer(), text, layout.NewSpacer())
			sr.sections[sc.Id] = SectionWidget{
				section:    logic.WrapSection(sc),
				background: background,
				text:       text,
				container:  container.NewStack(background, centered_text),
			}
		}
	}
	for _, sc := range sr.split.split.Sections {
		start, err := logic.EncodeTime(sc.Start_t)
		if err != nil {
			fmt.Printf("we have error %v", err)
			return
		}
		end, err := logic.EncodeTime(sc.End_t)
		if err != nil {
			fmt.Printf("we have error %v", err)
			return
		}
		_, s_ok := sr.timeStamps[start]
		_, e_ok := sr.timeStamps[end]
		if s_ok {
			sr.timeStamps[start].centered_label.Objects[0].(*widget.Label).Text = start
			sr.timeStamps[start].centered_label.Refresh()
		} else {
			sr.timeStamps[start] = TimeStamp{
				centered_label: container.NewCenter(widget.NewLabel(start)),
				line:           widget.NewSeparator(),
			}
		}
		if e_ok {
			sr.timeStamps[end].centered_label.Objects[0].(*widget.Label).Text = end
			sr.timeStamps[end].centered_label.Refresh()
		} else {
			sr.timeStamps[end] = TimeStamp{
				centered_label: container.NewCenter(widget.NewLabel(end)),
				line:           widget.NewSeparator(),
			}
		}
	}
	//this is syncing do it after updating the internal widgets when nessasary
	sr.objects = sr.objects[:0]
	for i := range sr.sections {
		sr.objects = append(sr.objects, sr.sections[i].container)
	}
	for i := range sr.timeStamps {
		sr.objects = append(sr.objects, sr.timeStamps[i].line)
		sr.objects = append(sr.objects, sr.timeStamps[i].centered_label)
	}
	sr.objects = append(sr.objects, &sr.background)
}
func (sr *SplitRenderer) MinSize() fyne.Size {
	var (
		width  float32 = 0
		height float32 = 0
	)
	for _, obj := range sr.objects {
		width = max(width, obj.MinSize().Width)
		height += obj.MinSize().Height
	}
	return fyne.NewSize(200, 200)
}

func (sr *SplitRenderer) Layout(size fyne.Size) {
	sr.background.Move(fyne.NewPos(0, 0))
	sr.background.Resize(size)
	for _, section := range sr.sections {
		start_pos := size.Height * float32(section.section.Start_t) /
			float32(core.Max_time)
		height := size.Height * float32(section.section.End_t-
			section.section.Start_t) /
			float32(core.Max_time)
		fmt.Println(start_pos)
		if section.section.Start_t < 5 {
			section.background.TopLeftCornerRadius = 15
			section.background.TopRightCornerRadius = 15
		}
		if section.section.End_t > core.Max_time-5 {
			section.background.BottomLeftCornerRadius = 15
			section.background.BottomRightCornerRadius = 15
		}
		section.container.Move(fyne.NewPos(0, start_pos))
		section.container.Resize(fyne.NewSize(size.Width, height))
	}
	for i, timestamp := range sr.timeStamps {
		time, err := logic.DecodeTime(i)
		if err != nil {
			fmt.Printf("error while laying out the time stamp %v", err)
			return
		}
		position := size.Height * float32(time) /
			float32(core.Max_time)
		if position < 5 || position > size.Height-5 {
			timestamp.centered_label.Hide()
			timestamp.line.Hide()
		}
		timestamp.centered_label.Move(fyne.NewPos(size.Width/2, position))
		timestamp.line.Move(fyne.NewPos(0, position))
		timestamp.line.Resize(fyne.NewSize(size.Width, 2))
	}
}

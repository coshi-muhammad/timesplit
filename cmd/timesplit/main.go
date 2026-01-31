package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/coshi-muhammd/timesplit/internal/core"
	"github.com/coshi-muhammd/timesplit/internal/logic"
	"github.com/coshi-muhammd/timesplit/internal/ui"
	"github.com/google/uuid"
)

func main() {
	a := app.New()
	w := a.NewWindow("Time split")
	//TODO: posibly make an initilization function to result in a better and cleaner main function
	router_logic := logic.NewRouterService(core.AppState{
		Split_List: make(map[uuid.UUID]*core.Split),
	})
	router := ui.NewRouter(router_logic, a, w)
	router.Logic.LoadConfig()
	if router.GetState().Config.Active_Split_id != uuid.Nil {
		fmt.Println(router.GetState().Config)
		router.Logic.SetActiveSplit(router.GetState().Config.Active_Split_id)
		daylogic := logic.NewDayService(router_logic,
			logic.WrapSplit(router.Logic.GetActiveSplit()))
		router.ShowDay(daylogic)
	} else {
		router.Logic.SetActiveSplit(uuid.Nil)
		router.ShowEmpty()
	}
	//TODO: put this in a more apropriate place
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	go func(ur *ui.UiRouter) {
		for range ticker.C {
			day, ok := ur.View.(*ui.DayView)
			if ok {
				day.SetVisibleSection()
			}
		}
	}(router)
	//TODO: when you make the day/week format make it the default
	//and which one depends in the wekk flag <half done>
	w.ShowAndRun()
	router.Logic.StoreConfig()
}

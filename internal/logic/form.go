package logic

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/coshi-muhammd/timesplit/internal/core"
)

type FormService struct {
	message string
}

var _ core.FormLogic = (*FormService)(nil)

func NewFormService() core.FormLogic {
	return &FormService{}
}
func (f FormService) Log() {
	fmt.Printf("the message was submited: %s", f.message)
}
func (f FormService) Submit() error {
	data, err := json.Marshal(struct {
		Message string `json:"message"`
	}{Message: f.message})
	if err != nil {
		return err
	}
	err = os.WriteFile("message.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
func (f *FormService) SetFields(message string) {
	f.message = message
}

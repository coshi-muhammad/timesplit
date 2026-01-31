package logic

import (
	"strings"

	"github.com/coshi-muhammd/timesplit/internal/core"
)

func NewAtachement(name string, uri string) core.Atachement {
	a := core.Atachement{}
	a.Uri = uri
	if strings.HasPrefix(uri, "file") {
		a.A_type = "file"
	} else if strings.HasPrefix(uri, "http") {
		a.A_type = "link"
	}
	a.Name = name
	return a
}

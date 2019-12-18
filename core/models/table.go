package models

import (
	"time"
)

type Table struct {
	Name    string
	Created time.Time
	Updated time.Time
	Rows    []*Row
}

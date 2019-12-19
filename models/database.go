package models

import (
	"time"
)

type Database struct {
	Owner User
	Name     string
	Username string
	Password string
	Host     string
	Created  time.Time
	Updated  time.Time
	Tables   []*Table
}

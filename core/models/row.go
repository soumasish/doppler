package models

import (
	"time"
)
type Row struct{
	HashId string
	Key string
	Value string
	Created time.Time
	Updated time.Time
}

package models

import "time"

// Data
type Data struct {
	Data     interface{}
	Alive    bool
	Created  time.Time
	LifeTime int64
}

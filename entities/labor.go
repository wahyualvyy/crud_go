package entities

import "time"

type Labor struct {
	Id         uint
	Nama_lab   string
	Created_at time.Time
	Updated_at time.Time
}
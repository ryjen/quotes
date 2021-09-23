package storage

import (
	"time"

	"github.com/speps/go-hashids"
)

func newId() string {
	hd := hashids.NewData()
	hd.Salt = "25e6ac531fb5ab4d58f60f93447e37c8a9e57f040194fe01"
	hd.MinLength = 13

	h, _ := hashids.NewWithData(hd)

	data := []int64{
		time.Now().Unix(),
	}
	id, _ := h.EncodeInt64(data)

	return id
}

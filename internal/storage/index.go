package storage

import (
	"os"
	"time"

	"github.com/speps/go-hashids"
)

func getSalt() string {
	salt, ok := os.LookupEnv("QUOTES_API_ID_SALT")
	if !ok {
		salt = "25e6ac531fb5ab4d58f60f93447e37c8a9e57f040194fe01"
	}
	return salt
}

func NewID() string {
	hd := hashids.NewData()
	hd.Salt = getSalt()
	hd.MinLength = 13

	h, _ := hashids.NewWithData(hd)

	data := []int64{
		time.Now().Unix(),
	}
	id, _ := h.EncodeInt64(data)

	return id
}

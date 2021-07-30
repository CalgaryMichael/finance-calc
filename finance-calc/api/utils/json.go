package utils

import (
	"encoding/json"
	"io"
)

func BindJSON(r io.ReadCloser, obj interface{}) {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&obj)
	CheckError(err)
}

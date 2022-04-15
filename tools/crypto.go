package tools

import (
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"encoding/hex"
	"io"

	go_uuid "github.com/satori/go.uuid"
	"github.com/sony/sonyflake"
)

//MD5 加密
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func UuidUint64() uint64 {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	if uid, err := flake.NextID(); err != nil {
		Log.Error(err.Error())
		return 0
	} else {
		return uid
	}
}

func Uuid() string {
	return go_uuid.NewV4().String()
}

func ZlibUnzip(src []byte) []byte {
	b := bytes.NewReader(src)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

func ZlibZip(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

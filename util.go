package umengpush

import (
	"crypto/md5"
	"encoding/hex"
)

func genSign(method, url, appMasterSecret string, body *Message) string {
	str := method + url + body.String() + appMasterSecret
	data := md5.Sum([]byte(str))
	return hex.EncodeToString(data[:])
}

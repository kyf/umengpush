package umengpush

import (
	"log"
	"net/http"
	"testing"
)

func TestGenSign(t *testing.T) {
	msg := NewMessage(APP_KEY1, DEMO_TITLE, "hello", nil)
	sign := genSign(http.MethodPost, SEND_MSG_URI, APP_MASTER_SECRET1, msg)
	log.Print(sign)
}

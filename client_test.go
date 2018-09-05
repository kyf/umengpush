package umengpush

import (
	"log"
	"testing"
	"time"
)

const (
	APP_KEY1            = DEV_APP_KEY
	APP_MESSAGE_SECRET1 = DEV_APP_MESSAGE_SECRET
	APP_MASTER_SECRET1  = DEV_APP_MASTER_SECRET

	DEMO_TITLE = "6人游"

	DATE_LAYOUT = "2006-01-02 15:04:05"
)

var (
	tokens = map[string]string{
		"oppo":  "Ar4jhzpt_iSs5QOHP8n32PG_gZAyRDehZBEY_RdbhOkH", //"Ak1FY1ojtdyIevgDdRfwz7gYjGccNjGz30FE_ruK1ODr",
		"leshi": "AgLQRQWfCVcJKNB9jRkxIAB7nMnMBihQs0vuFFZ17aq9",
	}

	custom = map[string]string{
		"loadurl": "http://www.6renyou.com",
		"orderid": "10",
	}

	client = NewClient(ENV_MODE_PRO, APP_KEY1, APP_MESSAGE_SECRET1, APP_MASTER_SECRET1)
	//client = NewClient(ENV_MODE_DEV, APP_KEY1, APP_MESSAGE_SECRET1, APP_MASTER_SECRET1)
)

func TestSingle(t *testing.T) {
	resp, err := client.Single(tokens["oppo"], DEMO_TITLE, "umeng单发测试用例"+time.Now().Format(DATE_LAYOUT), custom)
	if err != nil {
		t.Fatalf("single error:%v", err)
	}
	log.Print(resp)
}

/*
func TestGroup(t *testing.T) {
	list := tokens["oppo"] + "," +
		tokens["leshi"]

	resp, err := client.Group(list, DEMO_TITLE, "umeng组发测试用例"+time.Now().Format(DATE_LAYOUT), custom)
	if err != nil {
		t.Fatalf("all error:%v", err)
	}
	log.Print(resp)
}

func TestAll(t *testing.T) {
	resp, err := client.All(DEMO_TITLE, "umeng群发测试用例"+time.Now().Format(DATE_LAYOUT), custom)
	if err != nil {
		t.Fatalf("all error:%v", err)
	}
	log.Print(resp)
}
*/

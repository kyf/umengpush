package umengpush

import (
	"net/http"
	"strings"
)

func NewMessage(appKey, title, text string, custom map[string]string) *Message {
	if custom == nil {
		custom = map[string]string{}
	}
	custom["title"] = title
	custom["text"] = text

	msg := &Message{
		Mipush: "true",
		AppKey: appKey,
		PayloadName: Payload{
			DisplayType: "message",
			//DisplayType: "notification",
			BodyName: Body{
				Ticker:      title,
				Title:       title,
				Text:        text,
				Custom:      custom,
				PlayVibrate: "true",
				PlayLights:  "true",
				PlaySound:   "true",
			},
		},
	}

	return msg
}

func MakeRequest(msg *Message, uri, appMasterSecret string) (*http.Request, error) {
	sign := genSign(http.MethodPost, uri, appMasterSecret, msg)
	uri = uri + "?sign=" + sign

	req, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(msg.String()))
	if err != nil {
		return nil, err
	}

	return req, err
}

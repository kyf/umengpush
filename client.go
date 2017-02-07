package umengpush

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	ENV_MODE_PRO EnvMode = "true"
	ENV_MODE_DEV EnvMode = "false"
)

type EnvMode string

type Client struct {
	client           *http.Client
	mode             EnvMode
	appKey           string
	appMessageSecret string
	appMasterSecret  string
}

func NewClient(mode EnvMode, appKey, appMessageSecret, appMasterSecret string) *Client {
	return &Client{
		client:           &http.Client{},
		mode:             mode,
		appKey:           appKey,
		appMessageSecret: appMessageSecret,
		appMasterSecret:  appMasterSecret,
	}
}

func (this *Client) Single(token, title, text string, custom map[string]string) (*MsgResponse, error) {
	msg := NewMessage(this.appKey, title, text, custom)
	msg.TimeStamp = fmt.Sprintf("%d", time.Now().Unix())
	msg.Type = "unicast"
	msg.DeviceTokens = token
	msg.ProductionMode = string(this.mode)

	fmt.Printf("%s\n", msg.String())
	req, err := MakeRequest(msg, SEND_MSG_URI, this.appMasterSecret)
	if err != nil {
		return nil, err
	}
	res, err := this.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	msgResp := new(MsgResponse)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(msgResp)
	return msgResp, err
}

func (this *Client) Group(tokens, title, text string, custom map[string]string) (*MsgResponse, error) {
	msg := NewMessage(this.appKey, title, text, custom)
	msg.TimeStamp = fmt.Sprintf("%d", time.Now().Unix())
	msg.Type = "listcast"
	msg.DeviceTokens = tokens
	msg.ProductionMode = string(this.mode)

	fmt.Printf("%s\n", msg.String())
	req, err := MakeRequest(msg, SEND_MSG_URI, this.appMasterSecret)
	if err != nil {
		return nil, err
	}
	res, err := this.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	msgResp := new(MsgResponse)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(msgResp)
	return msgResp, err
}

func (this *Client) All(title, text string, custom map[string]string) (*MsgResponse, error) {
	msg := NewMessage(this.appKey, title, text, custom)
	msg.TimeStamp = fmt.Sprintf("%d", time.Now().Unix())
	msg.Type = "broadcast"
	msg.ProductionMode = string(this.mode)

	fmt.Printf("%s\n", msg.String())
	req, err := MakeRequest(msg, SEND_MSG_URI, this.appMasterSecret)
	if err != nil {
		return nil, err
	}
	res, err := this.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	msgResp := new(MsgResponse)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(msgResp)
	return msgResp, err
}

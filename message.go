package umengpush

import (
	"encoding/json"
)

type Body struct {
	Ticker string `json:"ticker"` // 必填 通知栏提示文字
	Title  string `json:"title"`  // 必填 通知标题
	Text   string `json:"text"`   // 必填 通知文字描述

	Icon      string `json:"icon"`      // 可选 状态栏图标ID, R.drawable.[smallIcon],如果没有, 默认使用应用图标。
	LargeIcon string `json:"largeIcon"` // 可选 通知栏拉开后左侧图标ID
	Img       string `json:"img"`       // 可选 通知栏大图标的URL链接。该字段的优先级大于largeIcon。该字段要求以http或者https开头。

	Sound       string `json:"sound"`        // 可选 通知声音,如果该字段为空，采用SDK默认的声音,如果SDK默认声音文件不存在，则使用系统默认的Notification提示音。
	BuilderId   int    `json:"builder_id"`   // 可选 默认为0，用于标识该通知采用的样式。
	PlayVibrate string `json:"play_vibrate"` // 可选 收到通知是否震动,默认为"true".
	PlayLights  string `json:"play_lights"`  // 可选 收到通知是否闪灯,默认为"true"
	PlaySound   string `json:"play_sound"`   // 可选 收到通知是否发出声音,默认为"true"

	/**
	必填 值可以为:
	"go_app": 打开应用
	"go_url": 跳转到URL
	"go_activity": 打开特定的activity
	"go_custom": 用户自定义内容。
	*/
	AfterOpen string      `json:"after_open"`
	Url       string      `json:"url"`      // 可选 当"after_open"为"go_url"时，必填。通知栏点击后跳转的URL，要求以http或者https开头
	Activity  string      `json:"activity"` // 可选 当"after_open"为"go_activity"时，必填。通知栏点击后打开的Activity
	Custom    interface{} `json:"custom"`   // 可选 display_type=message, 或者display_type=notification且"after_open"为"go_custom"时，该字段必填。用户自定义内容, 可以为字符串或者JSON格式。

}

type Policy struct {
	//可选 定时发送时间，若不填写表示立即发送,定时发送时间不能小于当前时间,格式: "yyyy-MM-dd HH:mm:ss",注意, start_time只对任务生效
	StartTime string `json:"start_time"`

	//可选 消息过期时间,其值不可小于发送时间或者start_time(如果填写了的话),如果不填写此参数，默认为3天后过期。格式同start_time
	ExpireTime string `json:"expire_time"`

	//可选 发送限速，每秒发送的最大条数,开发者发送的消息如果有请求自己服务器的资源，可以考虑此参数
	//MaxSendNum int `json:"max_send_num"`

	//可选 开发者对消息的唯一标识，服务器会根据这个标识避免重复发送,有些情况下（例如网络异常）开发者可能会重复调用API导致消息多次下发到客户端。如果需要处理这种情况，可以考虑此参数.注意, out_biz_no只对任务生效
	OutBizNo string `json:"out_biz_no"`
}

type Payload struct {
	DisplayType string            `json:"display_type"` // 必填 消息类型，值可以为:notification-通知，message-消息
	BodyName    Body              `json:"body"`         //必填 消息体,display_type=message时,body的内容只需填写custom字段,display_type=notification时, body包含如下参数
	Extra       map[string]string `json:"extra"`        //可选 用户自定义key-value。只对"通知"(display_type=notification)生效.可以配合通知到达后,打开App,打开URL,打开Activity使用
}

type Message struct {
	AppKey    string `json:"appkey"`    // 必填 应用唯一标识
	TimeStamp string `json:"timestamp"` // 必填 时间戳，10位或者13位均可，时间戳有效期为10分钟

	/**
		必填 消息发送类型,其值可以为:
		unicast-单播
	    listcast-列播(要求不超过500个device_token)
	    filecast-文件播(多个device_token可通过文件形式批量发送）
		broadcast-广播
		groupcast-组播(按照filter条件筛选特定用户群, 具体请参照filter参数)
		customizedcast(通过开发者自有的alias进行推送)
		包括以下两种case:
			- alias: 对单个或者多个alias进行推送
			- file_id: 将alias存放到文件后，根据file_id来推送
	*/
	Type string `json:"type"`

	/**
	可选 设备唯一表示
	当type=unicast时,必填, 表示指定的单个设备
	当type=listcast时,必填,要求不超过500个,多个device_token以英文逗号间隔
	*/
	DeviceTokens string `json:"device_tokens"`

	/**
	可选 当type=customizedcast时，必填，alias的类型,
	alias_type可由开发者自定义,开发者在SDK中
	调用setAlias(alias, alias_type)时所设置的alias_type
	*/
	AliasType string `json:"alias_type"`

	/**
	可选 当type=customizedcast时, 开发者填写自己的alias。
	要求不超过50个alias,多个alias以英文逗号间隔。
	在SDK中调用setAlias(alias, alias_type)时所设置的alias
	*/
	Alias string `json:"alias"`

	/**
	可选 当type=filecast时，file内容为多条device_token,device_token以回车符分隔
	     当type=customizedcast时，file内容为多条alias，
	     alias以回车符分隔，注意同一个文件内的alias所对应的alias_type必须和接口参数alias_type一致。
		注意，使用文件播前需要先调用文件上传接口获取file_id,具体请参照"2.4文件上传接口
	*/
	FileId string `json:"file_id"`

	Filter         map[string]string `json:"filter"`          // 可选 终端用户筛选条件,如用户标签、地域、应用版本以及渠道等
	PayloadName    Payload           `json:"payload"`         // 必填 消息内容(Android最大为1840B), 包含参数说明如下(JSON格式)
	PolicyName     Policy            `json:"policy"`          // 可选 发送策略
	ProductionMode string            `json:"production_mode"` // 可选 正式/测试模式。测试模式下，只会将消息发给测试设备,测试设备需要到web上添加.Android: 测试设备属于正式设备的一个子集
	Description    string            `json:"description"`     // 可选 发送消息描述，建议填写
	ThirdPartId    string            `json:"thirdparty_id"`   // 可选 开发者自定义消息标识ID, 开发者可以为同一批发送的多条消息提供同一个thirdparty_id, 便于友盟后台后期合并统计数据
}

type MsgResponseData struct {

	//当type为unicast、listcast或者customizedcast且alias不为空时
	MsgId string `json:"msg_id"`

	//当type为于broadcast、groupcast、filecast、customizedcast,且file_id不为空的情况(任务)
	TaskId string `json:"task_id"`

	// 当"ret"为"FAIL"时,包含如下参数
	ErrorCode string `json:"error_code"`

	//如果开发者填写了thirdparty_id, 接口也会返回该值
	ThirdPartId string `json:"thirdpart_id"`
}

type MsgResponse struct {
	Ret  string          `json:"ret"` //返回结果，"SUCCESS"或者"FAIL"
	Data MsgResponseData `json:"data"`
}

func (this *Message) String() string {
	result, err := json.Marshal(this)
	if err != nil {
		return ""
	}

	return string(result)
}

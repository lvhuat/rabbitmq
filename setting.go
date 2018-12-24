package rabbitmq

import (
	"time"

	"github.com/streadway/amqp"
)

const (
	queueSettingPrefix    = "queue/"
	bindSettingPrefix     = "bind/"
	exchangeSettingPrefix = "exchange/"
	consumeSettingPrefix  = "consume/"
)
const (
	settingDurable    = "durable"
	settingExclusive  = "exclusive"
	settingInternal   = "internal"
	settingAutoDelete = "autodelete"
	settingNoWait     = "nowait"
	settingAutoAck    = "autoack"
	settingNoLocal    = "nolocal"
)

func defaultQueueSettings() map[string]bool {
	return map[string]bool{
		settingDurable:    false,
		settingAutoDelete: false,
		settingExclusive:  false,
		settingNoWait:     false,
	}
}

func defaultExchangeSettings() map[string]bool {
	return map[string]bool{
		settingDurable:    false,
		settingAutoDelete: false,
		settingInternal:   false,
		settingNoWait:     false,
	}
}

func defaultConsumeSettings() map[string]bool {
	return map[string]bool{
		settingAutoAck:   false,
		settingExclusive: false,
		settingNoLocal:   false,
		settingNoWait:    false,
	}
}

type QueueSettings map[string]bool

func NewQueueSettings() QueueSettings {
	return map[string]bool{}
}

func (settings QueueSettings) Durable() QueueSettings {
	settings[queueSettingPrefix+settingDurable] = true
	return settings
}

func (settings QueueSettings) AutoDelete() QueueSettings {
	settings[queueSettingPrefix+settingAutoDelete] = true

	return settings
}

func (settings QueueSettings) Exclusive() QueueSettings {
	settings[queueSettingPrefix+settingDurable] = true
	return settings
}

type ExchangeSettings map[string]bool

func NewExchangeSettings() ExchangeSettings {
	return map[string]bool{}
}

func (settings ExchangeSettings) Durable() ExchangeSettings {
	settings[exchangeSettingPrefix+settingDurable] = true
	return settings
}

func (settings ExchangeSettings) AutoDelete() ExchangeSettings {
	settings[exchangeSettingPrefix+settingAutoDelete] = true
	return settings
}

func (settings ExchangeSettings) Internal() ExchangeSettings {
	settings[exchangeSettingPrefix+settingInternal] = true
	return settings
}

type ConsumeSettings map[string]bool

func NewConsumeSettings() ConsumeSettings {
	return map[string]bool{}
}

func (settings ConsumeSettings) AutoAck() ConsumeSettings {
	settings[consumeSettingPrefix+settingAutoAck] = true
	return settings
}

func (settings ConsumeSettings) Exclusive() ConsumeSettings {
	settings[consumeSettingPrefix+settingExclusive] = true
	return settings
}

func (settings ConsumeSettings) NoLocal() ConsumeSettings {
	settings[consumeSettingPrefix+settingNoLocal] = true
	return settings
}

func MakeupSettings(settings ...map[string]bool) map[string]bool {
	allSettings := make(map[string]bool, 10)
	for _, setting := range settings {
		for key, value := range setting {
			allSettings[key] = value
		}
	}
	return allSettings
}

type PublishOption func(*amqp.Publishing)

func OptionContentType(contentType string) PublishOption {
	return func(args *amqp.Publishing) {
		args.ContentType = contentType
	}
}

func OptionContentEncoding(contentEncoding string) PublishOption {
	return func(args *amqp.Publishing) {
		args.ContentEncoding = contentEncoding
	}
}

// OptionDeliveryMode 传输模式
func OptionDeliveryMode(deliveryMode uint8) PublishOption {
	return func(args *amqp.Publishing) {
		args.DeliveryMode = deliveryMode
	}
}

// OptionHeaders 消息头
func OptionHeaders(headers amqp.Table) PublishOption {
	return func(args *amqp.Publishing) {
		args.Headers = headers
	}
}

// OptionPriority 消息优先级
func OptionPriority(priority uint8) PublishOption {
	return func(args *amqp.Publishing) {
		args.Priority = priority
	}
}

// OptionCorrelationId 关联ID
func OptionCorrelationId(correlationId string) PublishOption {
	return func(args *amqp.Publishing) {
		args.CorrelationId = correlationId
	}
}

// OptionReplyTo 响应队列
func OptionReplyTo(replyTo string) PublishOption {
	return func(args *amqp.Publishing) {
		args.ReplyTo = replyTo
	}
}

// OptionExpiration 超时时间
func OptionExpiration(expiration string) PublishOption {
	return func(args *amqp.Publishing) {
		args.Expiration = expiration
	}
}

// OptionMessageId 消息ID
func OptionMessageId(messageId string) PublishOption {
	return func(args *amqp.Publishing) {
		args.MessageId = messageId
	}
}

// OptionTimestamp 不建议使用,配置会使用默认的配置
func OptionTimestamp(timestamp time.Time) PublishOption {
	return func(args *amqp.Publishing) {
		args.Timestamp = timestamp
	}
}

// OptionType 消息的类型
func OptionType(typ string) PublishOption {
	return func(args *amqp.Publishing) {
		args.Type = typ
	}
}

// OptionUserId 用户ID
func OptionUserId(userId string) PublishOption {
	return func(args *amqp.Publishing) {
		args.UserId = userId
	}
}

// OptionAppId 应用ID
func OptionAppId(appId string) PublishOption {
	return func(args *amqp.Publishing) {
		args.AppId = appId
	}
}

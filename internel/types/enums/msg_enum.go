package enums

type SessionType string

const (
	SessionTypeUnknown SessionType = "unknown"
	SessionTypeSingle  SessionType = "single" // 单聊
	SessionTypeGroup   SessionType = "group"  // 群聊
)

func (obj SessionType) Values() []string {
	return []string{
		string(SessionTypeUnknown),
		string(SessionTypeSingle),
		string(SessionTypeGroup),
	}
}

type MessageType string

const (
	MessageTypeUnknown MessageType = "unknown"
	MessageTypeText    MessageType = "textMessage"
)

func (obj MessageType) Values() []string {
	return []string{
		string(MessageTypeUnknown),
		string(MessageTypeText),
	}
}

type MessageStatus string

const (
	MessageStatusUnknown MessageStatus = "unknown"
	MessageStatusSending MessageStatus = "sending"
	MessageStatusSucceed MessageStatus = "succeed"
	MessageStatusFailed  MessageStatus = "failed"
)

func (obj MessageStatus) Values() []string {
	return []string{
		string(MessageStatusUnknown),
		string(MessageStatusSending),
		string(MessageStatusSucceed),
		string(MessageStatusFailed),
	}
}

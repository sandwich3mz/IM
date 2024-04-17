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
	MessageTypeText    MessageType = "text"
	MessageTypeImage   MessageType = "image"
	MessageTypeFile    MessageType = "file"
)

func (obj MessageType) Values() []string {
	return []string{
		string(MessageTypeUnknown),
		string(MessageTypeText),
		string(MessageTypeImage),
		string(MessageTypeFile),
	}
}

type MessageStatus int8

const (
	MessageStatusUnknown MessageStatus = -1
	MessageStatusSent    MessageStatus = 0
	MessageStatusRead    MessageStatus = 1
)

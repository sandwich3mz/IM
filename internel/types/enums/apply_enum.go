package enums

type ApplyType string

const (
	ApplyTypePending ApplyType = "pending" // 待处理
	ApplyTypeRefuse  ApplyType = "refuse"  // 拒绝
	ApplyTypeAgree   ApplyType = "agree"   // 同意
)

func (obj ApplyType) Values() []string {
	return []string{
		string(ApplyTypePending),
		string(ApplyTypeRefuse),
		string(ApplyTypeAgree),
	}
}

func (obj ApplyType) Ptr() *ApplyType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}

package enums

type ActionType string

const (
	ActionTypeRegister ActionType = "register" // 注册
	ActionTypeLogin    ActionType = "login"    // 登录
	ActionTypeModify   ActionType = "modify"   // 找回密码
)

func (obj ActionType) Values() []string {
	return []string{
		string(ActionTypeRegister),
		string(ActionTypeLogin),
		string(ActionTypeModify),
	}
}

func (obj ActionType) Ptr() *ActionType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}

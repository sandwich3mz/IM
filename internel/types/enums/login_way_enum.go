package enums

type LoginWay string

const (
	LoginWayCode LoginWay = "code" // 验证码登录
	LoginWayPwd  LoginWay = "pwd"  // 密码登录
)

func (obj LoginWay) Values() []string {
	return []string{
		string(LoginWayCode),
		string(LoginWayPwd),
	}
}

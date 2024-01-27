package enums

type UserStatus string

const (
	UserStatusOnline  UserStatus = "online"
	UserStatusOffline UserStatus = "offline"
)

func (obj UserStatus) Values() []string {
	return []string{
		string(UserStatusOnline),
		string(UserStatusOffline),
	}
}

func (obj UserStatus) Ptr() *UserStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}

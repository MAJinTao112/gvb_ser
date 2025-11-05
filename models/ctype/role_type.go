package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin Role = iota + 1
	PermissionUser
	PermissionVisitor
	PermissionDisableUser
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisableUser:
		str = "被禁用的用户"
	default:
		str = "其他"
	}
	return str
}

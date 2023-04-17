package client

type Role struct {
	ID         uint   `json:"id" `
	Name       string `json:"name"`
	RoleTypeID uint   `json:"role_type_id" gorm:"type:uint;"`
}

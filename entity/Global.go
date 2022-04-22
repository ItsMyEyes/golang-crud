package entity

type Todo struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	OwnerId   uint   `json:"owner_id"`
	User      User   `json:"user" gorm:"Foreignkey:Id;association_foreignkey:OwnerId;"`
}

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"primary_key"`
	Username      string `json:"username,omitempty" form:"username,omitempty"`
	Password      string `json:"password,omitempty" form:"password,omitempty"`
	PlainPassword string `json:"plain_password,omitempty" form:"plain_password,omitempty"`
	Email         string `json:"email,omitempty" form:"email,omitempty"`
	Phone         string `json:"phone,omitempty" form:"phone,omitempty"`
	Status        bool   `json:"status,omitempty" form:"status,omitempty"`
	Todo          []Todo `json:"todo,omitempty" gorm:"Foreignkey:OwnerId;association_foreignkey:ID;"`
}

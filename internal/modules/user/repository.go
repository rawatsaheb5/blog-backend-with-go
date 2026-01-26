type Repository interface{
	CreateUser(user *User) error
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository{
	return &repository{db}
}
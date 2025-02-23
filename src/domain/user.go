package domain

type User struct {
	id           int64
	username     string
	email        string
	password     string
	sessionToken string
	repository   UserRepository
}

func NewUser(repository UserRepository) *User {
	return &User{repository: repository}
}

func (u *User) SetID(id int64) {
	u.id = id
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetSessionToken(sessionToken string) {
	u.sessionToken = sessionToken
}

func (u *User) GetID() int64 {
	return u.id
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetSessionToken() string {
	return u.sessionToken
}

func (u *User) Exists() bool {
	return u.id > 0
}

func (u *User) Save() error {
	return u.repository.Save(u)
}

func (u *User) Update() error {
	return u.repository.Update(u)
}

func (u *User) Delete() error {
	return u.repository.Delete(u.id)
}

func (u *User) FindByID(id int64) (*User, error) {
	return u.repository.FindByID(id)
}

func (u *User) FindByEmail(email string) (*User, error) {
	return u.repository.FindByEmail(email)
}

func (u *User) FindByUsername(username string) (*User, error) {
	return u.repository.FindByUsername(username)
}

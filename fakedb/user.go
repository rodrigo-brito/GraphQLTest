package fakedb

var (
	nextUserID int64 = 0
)

type User struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Posts []*Post `json:"posts"`
}

type UserRepository struct{}

func GenerateUserID() int64 {
	nextUserID++
	return nextUserID
}

func (u *UserRepository) Get(ID int64) *User {
	for _, user := range Users {
		if user.ID == ID {
			return user
		}
	}
	return nil
}

func (u *UserRepository) All() []*User {
	return Users
}

func (u *UserRepository) AllValue() []User {
	users := []User{}
	for _, user := range Users {
		users = append(users, *user)
	}
	return users
}

func (u *UserRepository) Save(user *User) {
	Users = append(Users, user)
}

func (u *UserRepository) indexOf(ID int64) int {
	for index, user := range Users {
		if user.ID == ID {
			return index
		}
	}
	return -1
}

func (u *UserRepository) Delete(ID int64) error {
	index := u.indexOf(ID)
	if index == -1 {
		return ErrNotFound
	}
	Users = append(Users[:index], Users[index+1:]...)
	return nil
}

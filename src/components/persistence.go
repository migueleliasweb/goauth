package components

import "errors"

//InMemoryDatabase "De facto" where to search for information
type InMemoryDatabase struct {
	Permissions []Permission
	Users       map[string]User
}

//Permission Struct defining permission's fields
type Permission struct {
	Name string
}

//User Struct defining user's fields
type User struct {
	Name        string
	Password    string
	Permissions []Permission
}

//InMemoryDatabase Database to search for user's permission
var inMemoryDatabase InMemoryDatabase

//init For test purposes
func init() {
	_permissions := make([]Permission, 1)
	_permissions[0] = Permission{Name: "admin"}

	inMemoryDatabase.Users = make(map[string]User, 1)

	inMemoryDatabase.Users["miguel"] = User{
		Name:        "miguel",
		Password:    "1231231",
		Permissions: _permissions,
	}
}

//GetUser Get the user from DB
func GetUser(username string) (user User, err error) {
	if user, userExists := inMemoryDatabase.Users[username]; userExists {
		return user, nil
	}

	return User{}, errors.New("User not found.")
}

//GetEncodedPassword Returns the encoded user password from DB
func GetEncodedPassword(username *string) (password string, err error) {

	if user, userExists := inMemoryDatabase.Users[*username]; userExists {
		return user.Password, nil
	}

	return "", errors.New("User not found.")
}

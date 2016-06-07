package components

import "errors"

//GoAuthDB Database interface
type GoAuthDB interface {
	AddPermisson(*User, Permission) (*User, error)
	GetEncodedPassword(username *string) (password string, err error)
	GetUser(username string) *User
}

//InMemoryDatabase "De facto" where to search for information
type InMemoryDatabase struct {
	Permissions map[string]*Permission
	Users       map[string]User
}

//Permission Struct defining permission's fields
type Permission struct {
	Name string
}

//User Struct defining user's fields
type User struct {
	Username    string
	Password    string
	Permissions map[string]*Permission
}

//InMemoryDatabase Database to search for user's permission
var inMemoryDatabase InMemoryDatabase

//GetUser Returns the given User, nil otherwise
func (DB *InMemoryDatabase) GetUser(username string) *User {
	if user, userExists := DB.Users[username]; userExists {
		return &user
	}

	return nil
}

//AddPermisson Shortcut to add permissons to the giver User. Check is error is not nil
//to make sure the return is valid.
func (DB *InMemoryDatabase) AddPermisson(user *User, perm Permission) (*User, error) {
	if dbPermisson, permissionExists := DB.Permissions[perm.Name]; permissionExists {
		dbUser := DB.GetUser(user.Username)

		//avoid unwanted alloc
		dbUser.Permissions[perm.Name] = dbPermisson

		return dbUser, nil
	}

	return nil, errors.New("Invalid permisson.")
}

//GetEncodedPassword Returns the encoded user password from DB
func (DB *InMemoryDatabase) GetEncodedPassword(username *string) (password string, err error) {

	if user, userExists := inMemoryDatabase.Users[*username]; userExists {
		return user.Password, nil
	}

	return "", errors.New("User not found.")
}

//DB Instance of GoAuthDB
var DB GoAuthDB

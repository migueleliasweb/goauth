package components

import "errors"

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

//DB The in-memory database
var DB InMemoryDatabase

//init For test purposes
func init() {
	_permissions := make(map[string]*Permission)
	_permissions["admin"] = &Permission{Name: "admin"}

	_users := make(map[string]User, 1)
	_users["miguel"] = User{
		Username:    "miguel",
		Password:    "1ea3aa8cbcf860693559a55b36f041a879a3b2e62842ff49c4762ce860b98999b3d42b9ce3c7153bbe3b84f51950fe3b174b9db3f84d1f29039f901f69c6899f",
		Permissions: _permissions,
	}

	//DB
	DB = InMemoryDatabase{
		Users:       _users,
		Permissions: _permissions,
	}
}

//
// //GetUser Get the user from DB
// func GetUser(username string) (user User, err error) {
// 	if user, userExists := inMemoryDatabase.Users[username]; userExists {
// 		return user, nil
// 	}
//
// 	return User{}, errors.New("User not found.")
// }
//
// //AddPermisson Adds a persmisson to the given user
// func AddPermisson(username string, perm Permission) (user User, err error) {
// 	if user, userExists := inMemoryDatabase.Users[username]; userExists {
// 		if _perm, _ := inMemoryDatabase.Permissions[perm.Name]; _ {
// 			inMemoryDatabase.Users[username] = append(*user.Permissions, perm)
// 		}
// 	}
//
// 	return User{}, errors.New("User not found.")
// }
//
// //GetEncodedPassword Returns the encoded user password from DB
// func GetEncodedPassword(username *string) (password string, err error) {
//
// 	if user, userExists := inMemoryDatabase.Users[*username]; userExists {
// 		return user.Password, nil
// 	}
//
// 	return "", errors.New("User not found.")
// }

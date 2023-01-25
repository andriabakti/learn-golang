package main

import (
	"basic_struct/manage"
)

// struct (mirip map)
// untuk membuat struktur data yg diinginkan
// seperti property pada class (di OOP)
// seperti model database yg mempresentasikan property database

// type user struct {
// 	ID 				int
// 	FullName 	string
// 	Email 		string
// 	IsActive 	bool
// }

// embedded struct
// type group struct {
// 	name				string
// 	admin				User
// 	users				[]User
// 	isAvailable	bool
// }

func main() {
	// object = instance dari struct
	userOne := manage.User{ID: 1, FullName: "March Hare", Email: "march.hare@gmail.com", IsActive: true}
	userTwo := manage.User{ID: 2, FullName: "Leukosia", Email: "leukosia@gmail.com", IsActive: false}
	resultOne := userOne.DisplayUser()
	println(resultOne)
	print(userTwo.DisplayUser())
	// setter getter
	// user.ID = 1
	// user.FullName = "March Hare"
	// user.Email = "user@email.com"
	// user.IsActive = true

	// display := displayUser(user)
	// fmt.Println(display)

	users := []manage.User{userOne, userTwo}
	group := manage.Group{Name: "Spearhead", Admin: userOne, Users: users, IsAvailable: true}
	// fmt.Println(group)
	manage.DisplayGroup(group)
}

// method => dimiliki oleh struct
// func (user manage.User) display() string {
// 		return fmt.Sprintf("Name: %s, Email: %s", user.FullName, user.Email)
// }

// struct as parameter
// func displayGroup(group Group) {
// 	fmt.Printf("Name: %s \n", group.name,)
// 	fmt.Printf("Member count: %d \n", len(group.users))

// 	println("Users name:")
// 	for _, user := range group.users {
// 		fmt.Println("- ", user.FullName)
// 	}
// }

// struct as parameter
// func displayUser(user User) string {
// 	result := fmt.Sprintf("Name: %s, Email: %s", user.FullName, user.Email)
// 	return result
// }
package manage

import "fmt"

type User struct {
	ID       int
	FullName string
	Email    string
	IsActive bool
}

// embedded struct
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (user User) DisplayUser() string {
	return fmt.Sprintf("Name: %s, Email: %s", user.FullName, user.Email)
}

func DisplayGroup(group Group) {
	fmt.Printf("Name: %s \n", group.Name)
	fmt.Printf("Admin: %s \n", group.Admin.FullName)
	fmt.Printf("Member count: %d \n", len(group.Users))

	println("Users name:")
	for _, user := range group.Users {
		fmt.Println("- ", user.FullName)
	}
}
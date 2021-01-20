package management

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (group Group) Display() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println()
	fmt.Printf("Member count: %d", len(group.Users))
	fmt.Println()

	fmt.Println("Users name: ")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {
	// user := User{}
	// user.ID = 1
	// user.FirstName = "Muhammad"
	// user.LastName = "Naufal"
	// user.Email = "muhammad_naufal@apps.ipb.ac.id"
	// user.IsActive = true
	// fmt.Println(user)
	// fmt.Println(user.FirstName)

	// user2 := User{}
	// user2.ID = 2
	// user2.FirstName = "Muhammad2"
	// user2.LastName = "Naufal2"
	// user2.Email = "muhammad_naufal@apps.ipb.ac.id2"
	// user2.IsActive = true
	// fmt.Println(user2)
	// fmt.Println(user2.LastName)

	// user3 := User{
	// 	LastName:  "Aja",
	// 	FirstName: "Link",
	// 	IsActive:  false,
	// 	ID:        3,
	// 	Email:     "linkaja@email.com",
	// }
	// fmt.Println(user3)

	// user4 := User{
	// 	4,
	// 	"Link4",
	// 	"Aja4",
	// 	"linkaja@email.com4",
	// 	false,
	// }
	// fmt.Println(user4)

	user := User{1, "Naufal", "Muhammad", "tes@gmail.com", true}
	user2 := User{2, "Naufal2", "Muhammad2", "tes@gmail.com", true}

	displayUser1 := displayUser(user)
	displayUser2 := displayUser(user2)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)
}

func displayUser(user User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

package main

import (
	"fmt"
	"os"
)

type User struct {
	Name string
	Age  int
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAge() int {
	return u.Age
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func (u *User) GetInfo() (string, int) {
	return u.Name, u.Age
}

func (u *User) SetInfo(name string, age int) {
	u.Name = name
	u.Age = age
}

func main() {
	a := os.Getenv("TEST")
	fmt.Println(a)

	// x := 20
	// z := &x // address of x

	// fmt.Println("x real:", x)
	// fmt.Println("x real:", x)
	// fmt.Println("x addr:", &x)

	// fmt.Println("z ptr:", &z)
	// fmt.Println("z ptr:", z)
	// fmt.Println("z val:", *z) // value of pointer x

	// var u User
	// a := "an"
	// // b := 20

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("TEST: %s %d\n", string(a[0]), i)
	// 	if i <= 5 {
	// 		panic("asd")
	// 	}

	// }

	// u.Name = "John Doe"
	// u.Age = 30
	// fmt.Printf("Name: %s\n", u.GetName())
	// fmt.Printf("Age: %d\n", u.GetAge())
	// u.SetName("Jane Doe")
	// u.SetAge(25)

	// fmt.Printf("GetName: %s\n", u.GetName())
	// fmt.Printf("GetAge: %d\n", u.GetAge())

	// name, age := u.GetInfo()
	// fmt.Printf("GetInfo: %s %d\n", name, age)
	// u.SetInfo("John Doe", 30)

	// name, age = u.GetInfo()
	// fmt.Printf("GetInfo2: %s %d\n", name, age)
}

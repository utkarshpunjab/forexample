package main

import (
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

// Employee user defined data type that we wanna export
type Employee struct {
	Fname  string
	Salary int
}

func (e Employee) String() string {
	return fmt.Sprintf("%s: %d", e.Fname, e.Salary)
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Name < a[j].Name }

// BySal implements sort.Interface for []Employee based on
// the Age field.
type BySal []Employee

func (a BySal) Len() int           { return len(a) }
func (a BySal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySal) Less(i, j int) bool { return a[i].Salary < a[j].Salary }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)

	sort.Sort(ByAge(people))
	fmt.Println(people)

	e1 := Employee{
		Fname:  "Sam",
		Salary: 150000,
	}
	e2 := Employee{
		Fname:  "Bill",
		Salary: 200000,
	}
	e3 := Employee{
		Fname:  "Sarah",
		Salary: 275000,
	}
	e4 := Employee{
		Fname:  "Arsh",
		Salary: 400000,
	}

	employees := []Employee{e2, e4, e1, e3}
	fmt.Println(employees)
	sort.Sort(BySal(employees))
	fmt.Println(employees)

	pass := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pass)
	fmt.Println(string(bs))

	login1 := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(login1))
	if err != nil {
		fmt.Println("Wrong password, cannot login")
		return
	}
	fmt.Println("Password matched, you are logged in!")
}

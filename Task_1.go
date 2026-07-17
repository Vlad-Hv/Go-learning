package main

import (
	"errors"
	"fmt"
)

type Departament struct {
	Name  string
	Floor int
}

type Employee struct {
	Name        string
	Age         int
	Departament Departament
}

func main() {
	name, age, departmentName, departmentFloor, err := getInfo()

	if err != nil {
		fmt.Println(err)
		return
	}

	employee := createEmploe(name, departmentName, age, departmentFloor)

	fmt.Println(employee.Name, employee.Departament.Name)
}

func getInfo() (string, int, string, int, error) {
	var name string
	var age int
	var depName string
	var depFloor int

	fmt.Println("Enter employee name and age:")
	_, err := fmt.Scanln(&name, &age)

	if err != nil {
		return "", 0, "", 0, errors.New("invalid input type")
	}
	if age <= 18 {
		return "", 0, "", 0, errors.New("too young for work")
	}

	fmt.Println("Enter your departament name and floor: ")
	_, err = fmt.Scanln(&depName, &depFloor)

	if err != nil {
		return "", 0, "", 0, errors.New("invalid input type")
	}

	return name, age, depName, depFloor, nil
}

func createEmploe(name, depName string, age, depFloor int) Employee {
	employee := Employee{
		Name: name,
		Age:  age,
		Departament: Departament{
			Name:  depName,
			Floor: depFloor,
		},
	}
	return employee
}

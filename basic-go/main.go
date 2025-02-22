package main

import (
	"fmt"
	lib "project/lib"
)

func main() {
	lib.Hello()
	var name string = "Fhaisal"
	var umur int = 10
	var age int = 20


	var (
		isStudent bool     = true
		gpa       float32  = 3.5
		skills    []string = []string{"Java", "Go", "Python"}
	)

	fmt.Printf("Hasil Penjumlahan Umur : %d + Age: %d = %d \n",umur,age,lib.Add(umur, age))

	lib.OpenFile()
	lib.If()
	lib.Switch()

	fmt.Println(isStudent)
	fmt.Println(gpa)
	fmt.Println(skills)
	fmt.Printf("Hello, %s! P: %d U: %d\n", name, age, umur)

	message := fmt.Sprintf("Hello, %s! Age: %d", name, age)
	fmt.Println(message)

	


}


package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}
type person struct {
	firstname string
	lastname string
	contactInfo
}
func main(){
	sami := person{
		firstname: "Sami", 
		lastname: "BouAnber",
		contactInfo: contactInfo{"person@email.com", 123},
	}

	(&sami).updateName("ahmed")
	sami.print()
}
func (p person) print(){
	fmt.Println(p)
}
func (pointerToPerson *person) updateName(newFName string){
	*&pointerToPerson.firstname = newFName
}
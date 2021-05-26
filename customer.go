package model
import "fmt"

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//Return a customer with ad
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer  {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//Return a customer without id
func NewCustomer2(name string, gender string, age int, phone string, email string) Customer  {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//Return customer info
func (cm Customer) GetInfo() string  {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", cm.Id, cm.Name, cm.Gender, cm.Age, cm.Phone, cm.Email)
	return info
}
package service

import "awesomeProject/model"
//Insertion, deletion, modification, serach
type CustomerService struct {
	customers []model.Customer
	customerId int
}

//Return customerservice
func NewCustomerService() *CustomerService  {
	//Initialize a customer
	customerService := &CustomerService{}
	customerService.customerId = 1
	customer := model.NewCustomer(1, "Default","M",22,"00000000000","Default@gmail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//return customer slice
func (cs *CustomerService) List()  []model.Customer {
	return cs.customers
}

//Insertion
func (cs *CustomerService) Add(customer model.Customer) bool {
	//Sequence of insertion
	cs.customerId++
	customer.Id = cs.customerId
	cs.customers = append(cs.customers, customer)
	return true
}

//Return customer info accroading to ID
func (cs *CustomerService) FindById(id int) int {
	index := 1
	//Go through slice
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			//founded
			index = i
		}
	}
	return index
}

//Deletion
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	//No customer info if index is -1
	if index == -1 {
		return false
	}
	//deletion
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

func (cs *CustomerService) GetInfoById(id int) model.Customer  {
	i := id - 1
	return cs.customers[i]
}

//Modification by ID
func (cs *CustomerService) Update(id int, customer model.Customer) bool {
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			cs.customers[i].Name = customer.Name
			cs.customers[i].Gender = customer.Gender
			cs.customers[i].Age = customer.Age
			cs.customers[i].Phone = customer.Phone
			cs.customers[i].Email = customer.Email
		}
	}
	return true
}
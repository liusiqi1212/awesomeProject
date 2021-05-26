package main

import (
	"awesomeProject/model"
	"awesomeProject/service"
	"fmt"
)

type customerView struct {
	key string //User input
	loop bool  //menu list
	customerService *service.CustomerService
}

//Show all customer info
func (cv *customerView) list()  {
	//Get customer info in slice
	customers := cv.customerService.List()
	//Show info
	fmt.Println("------------------Customer list------------------")
	fmt.Println("ID\tName\tGender\tAge\tPhone number\tE-mail)
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------------Customer list------------------")
}

//Get user input
func (cv *customerView) add() {
	fmt.Println("------------------Insertion------------------")
	fmt.Print("Name：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("Gender：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("Age：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("Phone number：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("E-mail：")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println("------------------Insertion successful------------------")
	} else {
		fmt.Println("------------------Insertion failed------------------")
	}
}

//Delete by Id
func (cv *customerView) delete()  {
	fmt.Println("------------------Deletion------------------")
	fmt.Print("Input Id(-1 to exit)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return  
	}
	fmt.Print("Are you sure to delete(Y/N)：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if cv.customerService.Delete(id) {
			fmt.Println("------------------Deletion successful------------------")
		} else {
			fmt.Println("------------------Deletion failed，Id does not exit")
		}
	}
}

//Modify customer info by Id
func (cv *customerView) update() {
	cv.list()
	fmt.Println()
	fmt.Println("------------------Modify customer------------------")
	fmt.Print("Input Id number(-1 to exit)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return   
	}
	fmt.Print("Are you sure to modify(Y/N)：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if cv.customerService.FindById(id) != -1 {
			customer := cv.customerService.GetInfoById(id)
			fmt.Printf("Name（%v：）", customer.Name)
			name := ""
			fmt.Scanln(&name)
			fmt.Printf("Gender（%v）：", customer.Gender)
			gender := ""
			fmt.Scanln(&gender)
			fmt.Printf("Age（%v）：", customer.Age)
			age := 0
			fmt.Scanln(&age)
			fmt.Printf("Phone number（%v）：", customer.Phone)
			phone := ""
			fmt.Scanln(&phone)
			fmt.Printf("E-mail（%v）：", customer.Email)
			email := ""
			fmt.Scanln(&email)
			customer2 := model.NewCustomer2(name, gender, age, phone, email)
			cv.customerService.Update(id, customer2)
			fmt.Println("------------------Modification successful------------------")
		} else {
			fmt.Println("------------------Modification failed，Id does not exit")
		}
	}
}

//Exit
func (cv *customerView) logout()  {
	fmt.Print("Are you sure to exit(Y/N)：")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "Y" || cv.key =="y" || cv.key =="N" || cv.key =="n" {
			break
		}
		fmt.Print("Invalid input(Y/N): ")
	}
	if cv.key == "Y" || cv.key == "y" {
		cv.loop = false
	}
}


//显示主菜单
func (cv *customerView) mainMenu()  {
	for {
		fmt.Println("------------------Customer info management------------------")
		fmt.Println("                  1 Insert customer")
		fmt.Println("                  2 Modify customer")
		fmt.Println("                  3 Delete customer")
		fmt.Println("                  4 Customer list")
		fmt.Println("                  5 E x i t")
		fmt.Print("Select(1 - 5)：")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.logout()
		default:
			fmt.Print("Invalid input...")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("Successfully exit...")
}

func main()  {
	//main function
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//Initialize customerService
	customerView.customerService = service.NewCustomerService()
	//Show menu list
	customerView.mainMenu()
}
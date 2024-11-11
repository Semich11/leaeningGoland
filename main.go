package main

import "fmt"

func main(){
	// fmt.Println("Hello World")

	// var nameOne string = "mario"
	// var nameTwo = "luigi"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "peach"
	// nameThree = "bowser"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour:= "yoshi"
	// fmt.Println(nameFour)

	// var nameOne int = 455
	// var nameTwo = 45000
	// var nameThree int
	// numberFour:= 67

	// fmt.Println(nameOne, nameTwo, nameThree, numberFour)

	// nameThree = 12

	// fmt.Println(nameOne, nameTwo, nameThree, numberFour)


	// fmt.Print("Enter a number: ")
	// var input float64
	// fmt.Scanf("%f", &input)
	// output := input * 2
	// fmt.Println(output)
	// gasMilage()

	// creditLimitCalculator()

	// BackToSender()

	extrems()

}


func gasMilage(){
	var sum int 
	var miles float64
	var gallons float64


	for sum < 10 {
		fmt.Print("Enter mile: ")
		var mile float64
		fmt.Scanf("%f", &mile)

		miles += mile
		fmt.Println("Miles: ", miles)

		fmt.Print("Enter gallon: ")
		var gallon float64

		fmt.Scanf("%f", &gallon)
		gallons += gallon
		fmt.Println("Gallons: ", gallons)


		sum++
		fmt.Println(sum, " Times")

	}

	fmt.Printf("%0.2f", miles/gallons)
}


















func BackToSender(){
		fmt.Println("How many times the rider made a successful delivery?:");

		var userInput int

		fmt.Scanf("%d", &userInput)
		var backToSender int = backToSenderFunction(userInput)

		fmt.Printf("The rider wage for the day is: %d", backToSender)
}

func backToSenderFunction(userInput int) int{


	var amountPerParcel int



	if userInput >= 70 { amountPerParcel = 500 
	}else if userInput >= 60 && userInput <= 69 {amountPerParcel = 250
	}else if userInput >= 50 && userInput <= 59 {amountPerParcel = 200
	}else {amountPerParcel = 160
	}




	var riderDailyPay int = userInput * amountPerParcel + 5000;


	return riderDailyPay;

}














func creditLimitCalculator(){
	var accNumber int 
	var beginningBalance int
	var itemsChargedTotal int
	var totalCredit int
	var creditLimit int


	accNumber = 0000
	beginningBalance = 100
	itemsChargedTotal = 243
	totalCredit = 114
	creditLimit = 200

	var newBalance = beginningBalance + itemsChargedTotal - totalCredit

	if newBalance > creditLimit {
		fmt.Println("Credit Limit Excedded for ", accNumber)
		
	}



}




func extrems(){
	fmt.Println("Input how many values the application ask for")

	var userInput int 

	fmt.Scanf("%d", &userInput)

	values := make([]int, userInput)

	for i := 0; i < userInput; i++{
		fmt.Print("Enter a value: ")
		fmt.Scanf("%d", &values[i])
	}

	for i := 0; i < len(values); i++ {
		for z := 0; z < len(values); z++ {
			if values[z] > values[i] {
				temp := values[z]
				values[z] = values[i]
				values[i] = temp
			}
		}
	}

	max := values[len(values)-1]
	min := values[0]

	fmt.Printf("min: %d max: %d \n", min, max)
	fmt.Printf("sum: %d", min + max)


}
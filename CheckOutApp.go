package main

import "fmt"

type Checkout struct {
	productName        []string
	quantity           []int
	price              []float64
	total              []float64
	discountPercentage float64
	amountPaid         float64
	subTotal           float64
	discountPrice      float64
	vat                float64
	billTotal          float64
	balance            float64
	customersName      string
	cashierName        string
}

func main() {
	checkout := Checkout{}
	checkout.collectInput()

}
func (c *Checkout) collectInput() {
	var name string
	var piece int
	var unitPrice float64
	var response string

	fmt.Println("What is the customer's name?")
	fmt.Scanln(&c.customersName)

	fmt.Println("What did the user buy?")
	fmt.Scanln(&name)
	c.productName = append(c.productName, name)

	fmt.Println("How many pieces?")
	fmt.Scanln(&piece)
	c.quantity = append(c.quantity, piece)

	fmt.Println("How much per unit?")
	fmt.Scanln(&unitPrice)
	c.price = append(c.price, unitPrice)

	fmt.Println("Do you want to add more items: ")
	fmt.Scanln(&response)

	for response == "yes" || response == "Yes" {
		fmt.Println("What did the user buy?")
		fmt.Scanln(&name)
		c.productName = append(c.productName, name)
		fmt.Println("How many pieces?")
		fmt.Scanln(&piece)
		c.quantity = append(c.quantity, piece)

		fmt.Println("How much per unit?")
		fmt.Scanln(&unitPrice)
		c.price = append(c.price, unitPrice)

		fmt.Println("Do you want to add more items: ")
		fmt.Scanln(&response)

	}
	fmt.Println("What is your name?")
	fmt.Scanln(&c.cashierName)

	fmt.Println("How much discount will the person get?")
	fmt.Scanln(&c.discountPercentage)
}

func (c *Checkout) calculateDiscount() {
	c.discountPrice = (c.discountPercentage / 100) * c.subTotal
}

func (c *Checkout) calculateVat() {
	c.vat = (17.50 / 100) * c.subTotal

}
func (c *Checkout) calculateTotal() {
	c.billTotal = c.subTotal - c.discountPrice + c.vat
}
func (c *Checkout) printFirstReceipt() {
	fmt.Print("SEMICOLON STORES\n MAIN BRANCH\n LOCATION: 312, HERBERT MACAULAY WAY, SABO YABA, LAGOS.\n TEL:03293828343\nCashier name:, " +
		"c.cashierName\n Customer Name:, c.customerName")
	fmt.Println("==========================================")
	fmt.Println("ITEM\tQTY\tPRICE\tTOTAL(NGN)")
	fmt.Println("==========================================")

}

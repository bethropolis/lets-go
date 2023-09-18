package join

import "fmt"

// declaration of contants
const PI float64 = 3.14159
const VALUE = 1000


// multiple constants declaration block

const (
	PRODUCT = "Choclate"
	QUANTITY = 50
)


func join() {
	fmt.Println(PI)
	fmt.Println(VALUE)

	fmt.Println(PRODUCT)
	fmt.Println(QUANTITY)
}

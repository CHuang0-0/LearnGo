package main

import(
	"fmt"
)

func main() {
	DoThis()
	sum:=SumUp(1.2,5)
	fmt.Println(sum)
	sum2,sum2_len:=addAllValues(1,5,10)
	fmt.Println(sum2)
	fmt.Println(sum2_len)
	mycar:=Car{2010,"Ford","US"}
	mycar.CarBrand()

}

func DoThis(){
	fmt.Println("do this")
}

func SumUp(v1, v2 float64)float64{
	return v1+v2

}

// a function can accept an arbitrary number of values using ...
// functions can also return multiple values 
func addAllValues (values...int)(int,int){
	total:=0
	for _,v:=range(values){
		total+=v
	}
	return total,len(values)
}

// method
// Java: each method is a member of class
// Go: each method is a member of type
type Car struct{
	Year int
	Brand string
	Country string
}

// CarBrand is the brand of the car 
func (c Car) CarBrand(){ // receiver only copies; do not reference
	// printf outputs to the console
	// sprintf returns a string
	c.Brand=fmt.Sprintf("%v,%v",c.Brand,c.Brand)
	fmt.Println(c.Brand)
}


package main

import (
	"fmt"    // print
	"bufio"  // get input
	"os"    // get input
	"strconv" //string conversion
	"strings" //manipulate strings
	"math"  //math operations
	"time" //dates and time
)

// Declare constant
const aConst=99


// Function
func main () {
	// Declare variables
	var aString string="Go"
	fmt.Println("Celia is learning ", aString)
	fmt.Printf("Variable type is %T\n", aString)
	// := only works inside the funciton
	aInt:=88
	fmt.Printf("Int value is %v\n", aInt)
	fmt.Printf("Variable type is %T\n",aConst)


	// Get input
	reader:=bufio.NewReader(os.Stdin)   //Standard input
	fmt.Print("Enter text: ")
	input,_ :=reader.ReadString('\n')         //Ignore error object using _ 
	fmt.Println("You entered: ",input)

	// Convert strings
	fmt.Print("Enter a number:")
	numInput,_:=reader.ReadString('\n')  //line breaker as delimiter
	aFloat,err:=strconv.ParseFloat(strings.TrimSpace(numInput),64)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Value of number: ",aFloat)
	}

	// Math operators
	// Need to convert values before math operations
	a,b,c:=1,2,3
	intSum:=a+b+c 
	fmt.Println("int sum is ",intSum)
	x,y,z:=1.1,2.2,3.3
	floatSum:=x+y+z
	fmt.Println("float sum is ",floatSum)  //float is stored in binary, precision not guaranteed
	floatSum2:=math.Round(floatSum*100)/100 //2 decimals
	fmt.Println("precise float sum is ",floatSum2)
	circleRadius:=15.5
	circumference:=circleRadius*2*math.Pi 
	fmt.Printf("circumference is %.2f\n", circumference)

	// Dates and time
	n:=time.Now()
	fmt.Println("now: ",n)
	t:=time.Date(2022,time.August,11,18,0,0,0,time.UTC)
	fmt.Println(t.Format(time.ANSIC))
	parsedTime,err:=time.Parse(time.ANSIC,"Thu Aug 11 18:00:00 2022")
	fmt.Printf("type of parsedTime is %T\n",parsedTime)

}




















































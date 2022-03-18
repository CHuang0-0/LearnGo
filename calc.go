package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"strconv"
	"strings"
)

func main (){
	// Get input of variable 1 and 2
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter v1: ")
	v1,_:=reader.ReadString('\n')
	fmt.Println(v1)
	fmt.Println("Enter v2: ")
	v2,_:=reader.ReadString('\n')
	fmt.Println(v2)

	// Convert v1 & v2 into integers
	v1_f,err1:=strconv.ParseFloat(strings.TrimSpace(v1),64)
	v2_f,err2:=strconv.ParseFloat(strings.TrimSpace(v2),64)
	if err1==nil && err2==nil {
		fmt.Printf("v1 type: %T\n",v1_f)
		fmt.Printf("v2 type: %T\n",v2_f)

		// Division 
		v3:=math.Round(v1_f/v2_f*1000)/1000
		fmt.Println(v3)
	}else{
		fmt.Println(err1)
		fmt.Println(err2)
		panic("panicking. Error message!")
	}
}
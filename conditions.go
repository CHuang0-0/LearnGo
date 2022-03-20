package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// if conditions
	num:=88
	var result string
	if num<0{
		result="<0"
	}else if num==0{
		result="=0"
	} else{
		result=">0"
	}
	fmt.Println(result)
	if x:=99;x<100{
		result="<100"
	}else{result=">=100"}
	fmt.Println(result)

	// randomized values
	rand.Seed(time.Now().Unix())
	// switch statements=case statements in SQL
	switch dow:=rand.Intn(7)+1; dow { 	// Intn=ceiling
	case 1:
		result="Sunday"
		// fallthrough transfers control to the next case
		// fallthrough
	case 2:
		result="Monday"
	default:
		result="other"
	}
	fmt.Println(result)

	// go has no while statements
	// for statements
	colors:=[]string{"r","b","g"}
	for i:=0;i<len(colors);i++{
		fmt.Println(colors[i])
	}
	for i:=range colors{
		fmt.Println(colors[i])
	}
	// for each loop, 2 variables
	for _,color:=range colors{
		fmt.Println(color)
	}
	// create a boolean & for = while statements
	for value:=1;value<5;value++{
		fmt.Println(value)
	}
	value:=0
	for value<2{
	fmt.Println(value)
	value++;
	}

	// break
	// jump to the end of the code block
	// continue
	// start at the beginning of the current code block
	sum:=1
	for sum<1000{
		sum+=sum
		fmt.Println(sum)
		for sum>200{
			goto theEnd
		}
	}

theEnd:
	fmt.Println("end")


}
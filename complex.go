package main

import (
	"fmt"
	"sort"
)

func main(){
	// Map
	// New() allocates but not initialize memory
	// Make both allocates and initializes memory
	// memory is deallocated by garbage collector (GC)
	m:=make(map[string]int)
	m["this string"]=88
	fmt.Println(m)

	// Pointers
	// Variables that store the memory address of another variable
	aInt:=88
	var p=&aInt  // point at memory address
	fmt.Println("value of p:",*p)  // * says this is a pointer
	// p=&v
	// v=*p
	aFloat:=88.88
	p1:=&aFloat 
	fmt.Println("value of 1:",*p1)  

	//can change variable through *pointer
	*p1=*p1/31
	fmt.Println("changing value via p1: ",*p1)
	fmt.Println("current value: ",aFloat)

	// Array
	// can not sort arrays
	// can not add/remove items at runtime
	var colors [3]string
	colors[0]="red"
	colors[1]="blue"
	colors[2]="green"
	fmt.Println(colors)

	numbers:=[5]int{1,2,3,4,5}
	fmt.Println(len(numbers))

	// Slices: better at ordered lists
	// runtime allocates memory, creates array in tbe background, returns slice
	colors2:=[]string{"pink","orange","purple"}
	fmt.Println(colors2)
	// add items
	colors2=append(colors2,"yellow")
	fmt.Println(colors2)
	// remove items
	colors2=append(colors2[:len(colors2)-1])
	fmt.Println(colors2)
	// can also declare slices using make
	nums:=make([]int,5,5)  // type, initial len, cap
	nums[0]=1
	nums[1]=5
	nums[2]=9
	nums[3]=8
	nums[4]=7
	fmt.Println(nums)
	nums=append(nums,99)
	fmt.Println(nums)
	// sort
	sort.Ints(nums)
	fmt.Println(nums)

	// Map: unordered
	states:=make(map[string]string)
	fmt.Println(states)
	states["WA"]="Washington"
	states["CA"]="California"
	states["OR"]="Oregon"
	fmt.Println(states)

	cali:=states["CA"]
	fmt.Println(cali)
	delete(states,"OR")
	states["NY"]="New York"
	fmt.Println(states)
	// iteration is unordered
	for k,v:=range states{
		fmt.Printf("%v: %v\n",k,v)
	}

	// make an array
	keys:=make([]string,len(states))
	i:=0
	for k:=range states{
		keys[i]=k
		i++
	}
	fmt.Println(keys)
	// sort keys
	sort.Strings(keys)
	fmt.Println(keys)
	// loop through sorted keys
	for i:=range keys {
		fmt.Println(states[keys[i]])
	}


	// Structs
	// go does not have inheritance
	// no super or sub structs
	// each struct is independent, with its own methods
	// upper case type can be exported
	poodle:=Dog{"Poodle",10} //~constructor
	fmt.Println(poodle)
	fmt.Printf("%+v\n",poodle)
	fmt.Printf("breed: %v\nweight: %v\n",poodle.Breed,poodle.Weight)

}

// dog is a struct
type Dog struct{
	// give properties
	// each field is like classes in Java
	Breed string 
	Weight int 
}

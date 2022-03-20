package main

import(
	"fmt"
	"io"
	"io/ioutil"  //read/write files
	"os"
	"net/http" // make requests and send data to remove hosts; make http server app
	"encoding/json"
	"strings"
)


// const url="https://api.github.com/users/chuang0-0"  // my Github json text
const url="http://services.explorecalifornia.org/json/tours.php"

func main(){
	content:="hello from Celia"
	myfilename:="celiahello.txt"
	file,err:=os.Create("./"+myfilename)
	CheckError(err)
	// length=# of characters in the file
	length,err:=io.WriteString(file,content)
	CheckError(err)
	fmt.Printf("# of characters: %v\n",length)
	// close file
	defer file.Close()
	// read file after file is closed
	defer readFile("./"+myfilename)

	// get response
	resp,err:=http.Get(url)
	CheckError(err)
	fmt.Printf("response type: %T\n",resp)   // *http.Response
	defer resp.Body.Close()
	// read JSON data
	bytes,err:=ioutil.ReadAll(resp.Body)
	CheckError(err)
	bytes_str:=string(bytes)
	// fmt.Print(bytes_str)

	// turn JSON into structured data
	mygit:=MyGitfromJSON(bytes_str)
	for _,each:=range mygit {
		fmt.Println(each.Name)
	}
	// fmt.Println(mygit)
}

func CheckError(err error){
	if err!=nil{
		panic(err)
	}
}

func readFile(fileName string){
	// readfile returns an array of bytes
	data,err:=ioutil.ReadFile(fileName)
	CheckError(err)
	fmt.Println("text from file: ", string(data))

}

func MyGitfromJSON (mycontent string)[]MyGithub{
	github:=make([]MyGithub,0,20)
	// read JSON data
	decoder:=json.NewDecoder(strings.NewReader(mycontent))
	_,err:=decoder.Token()
	CheckError(err)
	// transform JSON into a slice
	var each MyGithub
	// decoder.More reads the next available object from the JSON content
	// if it finds one it'll return true
	for decoder.More(){ // true
		err:=decoder.Decode(&each)  // use pointer to pass memory address & fill in data
		// fmt.Printf("%T\n",&each)
		// fmt.Println(err)
		CheckError(err)
		github=append(github,each)
	}
	// fmt.Print(github)
	return github
}

type MyGithub struct{
	Name,Price string  // capitalized 1st letter:public
}



package main

import "fmt"

func main(){
	_, err:= createDatabase("firstDatabase", "abc", "cde")
	fmt.Println(err)
}

func createDatabase(name, username, password string) (int, error){
	fmt.Println("Creating database ", name, username, password)
	return 1, nil
}


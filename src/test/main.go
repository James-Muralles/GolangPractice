package main

import ( 
	"fmt"
	"errors"
	"math"
	"database/sql"
	"net/http"
	_ "github.com/lib/pq"
)

func main() {
// var num int

// //you can also right a variable like this without declaring type because go knows
// num1:=9

// num = 10
// fmt.Println(num + num1)

// //this is how you right an if statment
// if num == num1 {
// 	fmt.Println("true")
// } else if num < num1{
// 	fmt.Println("less")
// }else{
// 	fmt.Println("quack")
// }

// //arrays
// var arr [5]int
// arr1:= [5]int{1,2,3,4,5}
// arr2:= []int{1,2,3,4,5} //this is a slice, not limited to just 5 indexes
// arr2 = append(arr2, 13,2)
// fmt.Println(arr)
// fmt.Println(arr1)
// fmt.Println(arr2)

// //How to map
// vertices := make(map[string]int)
// vertices["triangle"] = 2
// vertices["square"] = 3
// // delete(vertices,"triangle")
// fmt.Println(vertices)

// //loops: only type is the for loop/ doubles as while  
// for i := 0; i<5; i++{
// 	fmt.Println(i)
// }

// fmt.Println("Beginning of while")
// j:=0
// for j < 5{
// 	fmt.Println(j)
// 	j++
// }

// //iterating through an array
// fmt.Println("Iterating")

// array := []string{"a","b", "c"}
// for index,value := range array {
// 	fmt.Println("index:", index, "value:", value)
// }

// //iterate through map
// for key,value := range vertices {
// 	fmt.Println("key:", key, "value:", value)
// }
// // result := sum(5,5)
// // println(sum(1,2))
// // println(result)
// // fmt.Println(sum(5,7))

// result, err  := sqrt(-4)
// if err != nil {
// 	fmt.Println(err)
// }else {
// 	//if you dont just use println() with out fmt number formatting is different
// 	fmt.Println(result)
// }

// p:= person{name: "Jimmy", age: 29}
// fmt.Println(p.name)

// g:=5
// fmt.Println(&g) //address of p


// }

// //new function other than main/ call it in the main

// func sum(x int, y int) int{
// 	return x + y
// }

// //function can take in more than one return/ go has no exception

// func sqrt (x float64) (float64, error){
// 	if x < 0{
// 		return 0, errors.New("undeifined negative numbers")
// 	}
// 	return math.Sqrt(x), nil
// }
// //struct type outside function

// type person struct {
// 	name string
// 	age int

// }

// //pointers using & and *


var DB *sql.DB

func init() {
   var err error
    DB, err = sql.Open("postgres", "postgres://postgres:codeup@localhost/test_db?sslmode=disable")
    if err != nil {
        panic(err)
    }

    if err = DB.Ping(); err != nil {
        panic(err)
    }
    fmt.Println("Database connection is succesful")

}
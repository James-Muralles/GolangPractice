package main

import ( 
	"fmt"
	"errors"
	"math"
	_ "github.com/lib/pq"
	"./util"
	"strconv"
	"net/http"

)

func main() {
var num int

//you can also right a variable like this without declaring type because go knows
num1:=9

num = 10
fmt.Println(num + num1)

//this is how you right an if statment
if num == num1 {
	fmt.Println("true")
} else if num < num1{
	fmt.Println("less")
}else{
	fmt.Println("quack")
}

//arrays
var arr [5]int
arr4 := []int{}
arr1:= [5]int{1,2,3,4,5}
arr2:= []int{1,2,3,4,5} //this is a slice, not limited to just 5 indexes
arr2 = append(arr2, 13,2)
arr4 = append(arr4, 3,2,2,3,4,5)
fmt.Println(arr)
fmt.Println(arr1)
fmt.Println(arr2)
fmt.Println(arr4)

//How to map
vertices := make(map[string]int)
vertices["triangle"] = 2
vertices["square"] = 3
// delete(vertices,"triangle")
fmt.Println(vertices)

//loops: only type is the for loop/ doubles as while  
for i := 0; i<=5; i++{
	fmt.Printf("number is %d\n ", i)
}

fmt.Println("Beginning of while")
j:=0
for j < 5{
	fmt.Println(j)
	j++
}

//iterating through an array
fmt.Println("Iterating")

array := []string{"a","b", "c"}
for index,value := range array {
	fmt.Println("index:", index, "value:", value)
}

//iterate through map
for key,value := range vertices {
	fmt.Println("key:", key, "value:", value)
}
// result := sum(5,5)
// println(sum(1,2))
// println(result)
// fmt.Println(sum(5,7))

ids :=[]int{23,34,455,87}

for i , id := range ids {
	fmt.Printf("\n%d - ID: %d\n ", i, id)
}
//loop through and add ids in a slice
sum:=0
for _, id := range ids {
	sum += id
	fmt.Println("Sum", sum)
	fmt.Printf("TYPE= %T\n", sum)//printing the type
	//READ A VALUE FROM AN ADRESS USE *
}

//loop through a map
emails := (map[string]string{"Jimmy": "jimmy@mail.com", "James": "james@email.com", "Jim": "jim@email.com"})
for k, v := range emails{
	fmt.Printf("%s: %s\n" , k, v)
}


result, err  := sqrt(4)
if err != nil {
	fmt.Println(err)
}else {
	//if you dont just use println() with out fmt number formatting is different
	fmt.Println(result)
	fmt.Printf("%d is the answer\n", result)
}


//switch statement

color:= "yellow"

switch color{
case "red":
	fmt.Println("color is red")
case "blue":
	fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or red")

}

p:= person{name: "Jimmy", age: 29}
fmt.Println(p.name)

g:=5
fmt.Println(&g) //address of p
fmt.Println(util.Reverse("Hello"))
fmt.Println(greeting(util.Reverse("Jimmy")))
fmt.Println(greeting("Jimmy "),  getSum(6,6)) 
fmt.Println(len(arr4))//len gives you the length of an array
fizzBuzz()
//functional programming
sum1 := adder()
for i := 0; i < 10; i++{
	fmt.Println(sum1(i))
}

//Init person struct

person1:=person{name: "Pop", lastName: "Bob", gender: "Male", age: 20 }
//or it can be written like this
person2 := person{"Sue", "Anne", "Female", 20}
fmt.Println(person1)
fmt.Println(person2)
person1.hasBirthday()
person2.hasBirthday()
fmt.Println(person1.greeting())
person2.getMarried("Steve")
fmt.Println(person2.greeting())
circle := Circle{x:0, y:0, radius:5}
rectangle := Rectangle{h: 5, w: 5}
fmt.Printf("Circle Area: %.2f\nRectangle Area: %.2f\n ",getArea(circle), getArea(rectangle) )

http.HandleFunc("/", index)
http.HandleFunc("/about", about)

fmt.Println("Server starting...")
http.ListenAndServe(":3000",nil)
}

//new function other than main/ call it in the main

func sum(x int, y int) int{
	return x + y
}

//function can take in more than one return/ go has no exception

func sqrt (x float64) (float64, error){
	if x < 0{
		return 0, errors.New("undeifined negative numbers")
	}
	return math.Sqrt(x), nil
}

func greeting(name string) string{
	return "Hello " + name
}

func getSum(num1 int, num2 int) int{
	return num1 + num2
}

func fizzBuzz(){
	for i:= 1; i<100; i++{
		if i % 15 == 0{
			fmt.Printf("fizz buzz %d\n",i)
		}else if i % 3 == 0{
			fmt.Printf("fizz %d\n",i)
		}else if i % 5 == 0{
			fmt.Printf("Buzz %d\n",i)
		}else {
			fmt.Printf("OOOPS %d\n",i)
		}
	}
}

//EXAMPLE OF FUNCTIONAL PROGRAMMING

func adder() func(int)int{
	sum1 := 0
	return func(x int)int{
		sum1 += x
		return sum1
	}

}

//struct type outside function

type person struct {
	// name 	 string
	// lastName string
	// gender 	 string
	// age 	 int

	// can also be written like this
	name, lastName, gender string
	age int 

}


//method for greeting (value receiver)
func (p person) greeting()string{
	return "Hello name is " + p.name + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

//method for changing age(pointer reciever)
func (p * person) hasBirthday(){
	if p.gender == "Male"{
		p.age++
	}else if p.gender == "Female"{
		p.age--
	}
}

//getMarried (pointer reviever)

func (p * person) getMarried(spouseLastname string) {
if p.gender == "Male"{
	return
}else{
	p.lastName = spouseLastname
}
}

//interface defined

type shape interface{
area() float64
}

type Circle struct{
	x, y, radius float64
}

type Rectangle struct{
	h, w float64
}

func (c Circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64{
	return r.h * r.w
}

func getArea(s shape) float64{
	return s.area()
}

//web, routers

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hello!</h1>")//output to browser
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>About</h1>")//output to browser
}


//pointers using & and *


package main 
import "fmt"

func main() {

emp := map[string]int{"Jnanendra Veer" : 101 , "Ravinder Singh" : 102 , "Swati Padiyar" : 103 , "Sandeep Sharma" : 104  , "Rohit" : 105 , "Swati Gupta" : 106 , "Pawan" : 107 , "Shubham" : 108 , "Sanjay" : 109 , "Asif" : 110 , "Arvind" : 111 , "Ayush" : 112 }

for name , ID:= range emp {
fmt.Println( name , ID )
}

no := [5]int{1, 7, 18, 21, 28}
sum := 0
for i, number := range no {
fmt.Println(i , number)

sum += number
}

fmt.Println(sum)

}







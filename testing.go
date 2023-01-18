 package main 
 import "fmt"
 
 func main() {
 
 id := map[string]int{"Id1": 101 , "Id2": 102, "Id3": 103, "Id4": 104}
 
 fmt.Println("ID  -  \n")
 
 for i, j := range id {
 
 fmt.Println(i ,"-" , j)
 
}
 
}

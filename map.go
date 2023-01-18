package main 
import "fmt"

func main() {

ID := make(map[string]int)

ID["v1"]  = 1
ID["v2"]  = 2
ID["v3"]  = 3
ID["v4"]  = 4
ID["v5"]  = 5
ID["v6"]  = 6
ID["v7"]  = 7
ID["v8"]  = 8
ID["v9"]  = 9
ID["v10"] = 10
fmt.Println(ID)

value := ID["v7"]
fmt.Println(value)

fmt.Println(len(ID))

delete(ID, "v10")
fmt.Println(ID)

val := ID["v11"]
fmt.Println(val)

_, prs := ID["v10"]
fmt.Println(prs)

a := make(map[int]string)

a[1] = "Gyan sir"
a[2] = "Ravi sir"
a[3] = "Sandeep sir"
a[4] = "Rohit sir"
a[5] = "Sanjay sir"
a[6] = "Amit sir"

for i:=0; i<=5; i++ {
fmt.Println(a[i])

}

}




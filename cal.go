package main
import "fmt"

func sum ( v1 int, v2 int, v3 int ) int {
return v1 + v2 + v3
}

func mul ( v4 int, v5 int, v6 int ) int {
return v4 * v5 * v6
}

func div ( v7 int, v8 int) int {
return v7 / v8
}

func sub ( v9 int, v10 int ) int {
return v9 - v10
}

func main() {
a := sum(23, 45 , 10)
fmt.Println("addition is : " , a)

m := mul(4, 5, 5)
fmt.Println("multiplication is : " , m)

d := div(21, 3)
fmt.Println("division is : " , d)

s := sub(4, 8)
fmt.Println("substraction is : " , s)

}







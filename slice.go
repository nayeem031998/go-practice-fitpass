package main
import "fmt"

func main() {
Emp := make([]string, 5)
Emp[0] = "Jnanendra"
Emp[1] = "Ravinder"
Emp[2] = "Rohit"
Emp[3] = "Sandeep"
Emp[4] = "Sanjay"
fmt.Println(Emp)
fmt.Println(len(Emp))

Emp = append(Emp, "Amit", "Asif", "Arvind", "Ayush")
fmt.Println(Emp)

FitTech := make([]string, len(Emp))
copy(FitTech, Emp)
fmt.Println(FitTech)

log := Emp[1:8]
fmt.Println(log)

log = Emp[5:6]
fmt.Println(log)

log = Emp[8: ]
fmt.Println(log)

log = Emp[ :0]
fmt.Println(log)

diet := make([]string, 4)

diet[0] = "Urvashi"
diet[1] = "Sakshi"
diet[2] = "Bhavika"
diet[3] = "Garima"
fmt.Println(diet)

twoD := make([][]int, 4)
for i := 0; i < 4; i++ {
inner := i + 1
twoD[i] = make([]int, inner)
for j := 0; j < inner; j++ {
  twoD[i][j] = i + j
}
}
fmt.Println(twoD)

}


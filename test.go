package main
import "fmt"
  type Employee interface {
              Basic() string
           Personal() string
}
  type details struct {
                 Name string
                   ID string
          Designation string
              Contact string
              Address string
               Salary string
  } 
    func (m details) Basic() string {
      
      return  m.Name + m.ID + m.Designation    
  }  
    func (m details) Personal() string {
  
      return  m.Name + m.Contact + m.Address + m.Salary      
  }  
    func main() {
     
     var d1 Employee 
     var d2 Employee
     var d3 Employee
     var input string
     d1 = details{"\nArvind\n", "12\n", "QA Tester\n", "9717738302\n", "H-27 sector 9 vijay nagar ghaziabad\n", "12000\n"}
     d2 = details{"\nRavinder\n", "11\n", "IOS Developer\n", "9986763578\n", "k-74 tilak nagar west delhi\n", "150000\n"}
     d3 = details{"\nAsif\n", "10\n", "Dart Developer\n", "7002689673\n", "215 metro pillar mandir wali gali East patel nagar central delhi\n", "65000\n"}
     
       fmt.Println("Details")
     fmt.Scanf("%s", &input)    
  if input == "basic" {
            fmt.Println(d1.Basic(), d2.Basic(), d3.Basic())
  } else if input == "personal" {
            fmt.Println(d1.Personal(), d2.Personal(), d3.Personal())
  } else {
            fmt.Println("Invalid charachters")
     }
  }

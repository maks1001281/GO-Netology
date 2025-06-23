package main

import "fmt"

func main() {
var age int = 0
        for i := 1; i <= 5; i++ {

                fmt.Print("Введите ваш возраст:")
                fmt.Scanln(&age)
                if age < 18 {
                        fmt.Println(age,"Ребенок")
                }
                if age >= 18 && age <=64 {
                        fmt.Println(age,"Взрослый")
                }
                if age >=65 {
                        fmt.Println(age,"Пенсионер")
                }   
        }
}

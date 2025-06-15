package main
import "fmt"

const AppName string = "Добро пожаловать в GoFit!"
func main (){ 
var age int = 30
var name string = "Иван"
var height float64 = 1.75
var message bool= true
fmt.Println(AppName, "\n", "Профиль пользователя:\n", "Имя:\n",name, "Возраст:",age,"лет", "\n", "Рост:",height,"м","\n", "Подписан на рассылку:",message )
}
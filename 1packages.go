/*Пакеты в Go*/

/*все программы начинаются с пакета main*/
package main

/*импорт пакетов с слова import, а дальше перечисление*/
import (
	"fmt"
	"math/rand"
)

/*все функции начинаются с функции main*/
func main() {
	//досуп к функции пакета пакет.Функция (функция только с Большой буквы)
	fmt.Println("My favorite number is ", rand.Intn(10))
}

//////////основные моменты//////////////
// main, import, пакеты типо func.Name

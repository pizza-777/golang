//<< >> шифт операторы
//сдвигают биты влево и вправо
//shift - в переводе "сдвиг"
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int64 = 96

	//переводит десятичное число в бинарное - это строка
	var b = strconv.FormatInt(a, 2)
	fmt.Println(b) //11000
	//right shift >> выкидывают биты
	fmt.Println(a >> 2) //6 потому что 6 в бинарном виде 110
	//left shift << добирают биты
	fmt.Println(a << 2) //96 потому что 96 в бинарном виде 1100000

}

//десятичные числа (1,2,3 ...) можно представить в бинарном (в виде единици и нулей) виде
//пример 2 это 100, а 100 это 1100100
//

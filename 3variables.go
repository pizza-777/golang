////////////////////////////////////////////////////////////////////////////
//////////////////////		ПЕРЕМЕННЫЕ      ////////////////////////////////
//////////////////////////////////////////////////////////////////////////

package main

/*объявление переменный вне функции*/
var a, b, c bool

/*тип указывать необизательно, Го сам догадается какой тип ему подали*/
var e, f, j int = 2, 3, 4

func main() {
	/*объявление переменных внутри функциии - переменная-тип*/
	var d int
	println(a, b, c, d)
	println(e, f, j)
	h := "hello"
	println(h)
}

/*var внутри и снаружи функции*/
/*ловля типа на лету при присвоении*/
/*:= короткое объявление переменной, только внутри функции*/

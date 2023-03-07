// Napisz program, który wykona obliczenie pierwiastków trójmianu kwadratowego.
//  Dla przypadku gdy delta jest ujemna program może wypisać komunikat,
//  że nie ma pierwiastków (lub w trudniejszej wersji - może policzyć pierwiastki zespolone).
package main

import "fmt"

func main() {

    var a int
	var b int
	var c int

    fmt.Print("y = ax^2 + bx + c")
	fmt.Println()
	fmt.Print("Podaj a: ")
    fmt.Scanf("%d", &a)
	fmt.Print("Podaj b: ")
	fmt.Scanf("%d", &b)
	fmt.Print("Podaj c: ")
	fmt.Scanf("%d", &c)


	delta := b*b - 4*a*c

	x1 := (-b - delta) / (2 * a)
	x2 := (-b + delta) / (2 * a)

	fmt.Println("x1 = ", x1)
	fmt.Println("x2 = ", x2)

}
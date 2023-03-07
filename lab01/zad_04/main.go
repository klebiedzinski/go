// Napisz program, który zagra w zgadywanie z użytkownikiem.
// Należy wylosować liczbę którą następnie użytkownik ma odgadnąć.
// Po wprowadzeniu liczby program ma wypisać informację czy liczba jest prawidłowa.
// Jeżeli nie, to czy jest większa czy mniejsza od wylosowanej.
// Gra kończy się gdy użytkownik zgadnie liczbę lub przerwie program.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
    number := rand.Intn(100)

	fmt.Println(number)

	var guess int

	fmt.Println("Pobwamy sie! Zgadnij co wylosowałem: ")
	fmt.Scanf("%d", &guess)
	for guess!=number {
		if guess<number {
			fmt.Println("Więcej")
			fmt.Println("Spróbuj jeszcze raz")
			fmt.Scanf("%d", &guess)
		} else {
			fmt.Println("Mniej")
			fmt.Println("Spróbuj jeszcze raz")
			fmt.Scanf("%d", &guess)
		}
	
	}
	fmt.Println("Zgadza sie! Gratulacje")
	

}
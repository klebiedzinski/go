// Napisz program, który liczy ile lat miałby użytkownik,
//  gdyby mieszkał na Marsie, na Wenus lub na innych planetach.
//   Ile Ty możesz mieć aktualnie lat na takich planetach?
//   Wiek mieszkańca danej planety to liczba pełnych obrotów planety wokół słońca.
//   Sprawdź w internecie...
package main

import "fmt"

func main() {

    var age int

    fmt.Print("Ile masz lat: ")
    fmt.Scanf("%d", &age)

	ageFloat := float64(age)

	marsAge := ageFloat * 1.88
	venusAge := ageFloat * 0.62
	
	fmt.Println("Możesz mieć", marsAge, "lat na Marsie i", venusAge, "lat na Wenus")

}
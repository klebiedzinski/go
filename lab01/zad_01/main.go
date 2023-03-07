//Napisz prosty program, który pyta ile masz lat, a następnie wypisuje liczbę miesięcy i dni które masz już za sobą.
package main

import "fmt"

func main() {

    var age int

    fmt.Print("Ile masz lat: ")
    fmt.Scanf("%d", &age)

	months := age * 12
	days := age * 365
	
	fmt.Println("Masz", months, "miesięcy i", days, "dni")

}
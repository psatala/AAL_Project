// Epidemia
// Piotr Satala, Piotr Libera

package main

import "fmt"

func printHelp() {
	fmt.Printf("Uruchomienie programu: ./AAl_Project [Tryb uruchomienia] [Opcje trybu uruchomienia]\n")
	fmt.Printf("Zawsze nalezy podac wszystkie opcje dostepne dla wybranego trybu\n")
	fmt.Printf("\n[tryb uruchomienia] moze przyjac wartosc:\n")

	fmt.Printf("\t-m1 - wszystkie dane pobierane ze standardowego wejscia\n")
	fmt.Printf("\t-m2 - pobierane parametry problemu i generacja tablic D i Z\n")
	fmt.Printf("\t-m3 - tryb prezentacji, w ktorym program bada zlozonosc algorytmu generujac sekwencje rozwazanych wielkosci problemu n\n")

	fmt.Printf("\nOpcje uruchomienia:\n")
	fmt.Printf("\tW trybie m1 nie ma dodatkowych opcji, na standardowym wejsciu powinny pojawic sie liczby oznaczajace kolejno:\n")
	fmt.Printf("\t\tn\n")
	fmt.Printf("\t\tn liczb odpowiadajacych odleglosciom D miedzy kolejnymi miastami, zaczynajac od odleglosci miedzy miastami 0 i 1\n")
	fmt.Printf("\t\tn liczb odpowiadajacych zapotrzebowanie Z kolejnych miast od 0 do n-1\n")

	fmt.Printf("\tW trybie m2:\n")
	fmt.Printf("\t\t-n - wielkosc problemu (domyslnie 1000)\n")
	fmt.Printf("\t\t-w - maksymalna wartosc liczb w D i Z (domyslnie 10000)\n")
	fmt.Printf("\t\t-s - wartosc seed dla generatora liczb pseudolosowych (domyslnie aktualny czas systemowy)")

	fmt.Printf("\tW trybie m3 parametry trybu m2 i dodatkowo:\n")
	fmt.Printf("\t\t-n - wielkosc pierwszego problemu (domyslnie 1000)\n")
	fmt.Printf("\t\t-k - liczba roznych wielkosci n w sekwencji generowanych problemow (domyslnie 1)\n")
	fmt.Printf("\t\t-t - roznica miedzy kolejnymi rozmiarami problemu n w sekwencji (step) (domyslnie 500)\n")
	fmt.Printf("\t\t-r - liczba wygenerowanych instacji problemu dla kazdego n (domyslnie 1)\n")
	fmt.Printf("\t\t-w - maksymalna wartosc liczb w D i Z (domyslnie 10000)\n")
	fmt.Printf("\nPrzyklad uruchomienia:\n")
	fmt.Printf("\t./AAL_Project -m3 -n1000 -k30 -t500 -r10 -w10000\n\n")
}


func getFullInput() ([]int, []int, bool) {
	ok := true
	var n int
	fmt.Scanf("%d", &n)
	D := make([]int, n)
	Z := make([]int, n)
	for i := 0; i < n; i += 1 {
		fmt.Scanf("%d", &D[i])
	}
	for i := 0; i < n; i += 1 {
		fmt.Scanf("%d", &Z[i])
	}
	return D, Z, ok
}

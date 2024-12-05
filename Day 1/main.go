package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Input importieren
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	//Listen Anlegen
	var liste1, liste2 []int

	//String in Zeilen zerlegen
	zeilen := strings.Split(string(content), "\n")

	//Jede Zeile durchgehen
	for _, zeile := range zeilen {
		//Zeile in Zahlen zerlegen
		zahlen := strings.Fields(zeile) //Zerlegt nach Leerzeichen
		if len(zahlen) == 2 {
			//Zahlen in Integer umwandeln
			zahl1, _ := strconv.Atoi(zahlen[0])
			zahl2, _ := strconv.Atoi(zahlen[1])

			//Zahlen den Listen hinzufügen
			liste1 = append(liste1, zahl1)
			liste2 = append(liste2, zahl2)
		}
	}

	//Listen sortieren
	sort.Ints(liste1[:])
	sort.Ints(liste2[:])

	//Differenzen berechnen
	totalDistance := 0
	for i := 0; i < len(liste1); i++ {
		distance := int(math.Abs(float64(liste1[i] - liste2[i])))
		totalDistance += distance
	}

	fmt.Printf("Liste 1: %v\nListe 2: %v", liste1, liste2)
	fmt.Printf("\nGesamtdistanz: %d", totalDistance)
}

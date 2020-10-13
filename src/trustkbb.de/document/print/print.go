package print

import "fmt"
import "io"

// Allgemeines Printer-Interface
// Es werden map-Daten ausgedruckt
type Printer interface {
	output(data map[string]string)
}

// Druckertyp1: druckt nur Raw-Daten
type PrintRow struct {
	Writer io.Writer
}

// Druckertyp2: druckt nur Format-Daten
type PrintFormat struct {
	Writer io.Writer
}

// Funktion zur Ausführung verschiedener Drucker
func PrintItOut(p Printer, data map[string]string) {
	p.output(data)
}

// implement print Interface
func (d PrintRow) output(data map[string]string) {
	for path, result := range data {
		//  fmt.Fprintf(d.Writer, "%s -> %.1f\n", path, result)
		fmt.Println("PrintRow:", path, result)
	}
}

// implement print Interface
func (d PrintFormat) output(data map[string]string) {
	for path, result := range data {
		//  fmt.Fprintf(d.Writer, "%s -> %.1f\n", path, result)
		fmt.Println("PrintFormat:", path, result)
	}
}

package main

import (
	"fmt"
	"trustkbb.de/document/print"
)

func main() {
	fmt.Println("SWFApps is running")

	var d = make(map[string]string, 2)

	var p1 print.PrintRow
	var p2 print.PrintFormat

	d["A1"] = "Autobahn-A1"
	d["A2"] = "Autobahn-A2"

	print.PrintItOut(p1, d)
	print.PrintItOut(p2, d)
}

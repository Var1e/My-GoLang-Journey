//Write a program that displays temperature conversion tables
//The program should draw two tables. The first table has two columns,
// with °C in the first column and °F in the second column. Loop from –40 C through 
//100° C in steps of 5°
package main

import "fmt"

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5.0 / 9.0)
}

const (
	line         = "============================="
	rowFormat    = "|%8s|%8s|\n"
	numberFormat = "%.1f"
)

type getRowFn func(row int) (string, string)

func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}

func ctof(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2
}
func ftoc(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}
func main() {
	drawTable("°C", "°F", 30, ctof)
	fmt.Println()
	drawTable("°F", "°C", 30, ftoc)
}

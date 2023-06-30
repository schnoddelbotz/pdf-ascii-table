package main

import (
	"fmt"
	"log"

	"github.com/go-pdf/fpdf"
)

func main() {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 15)
	pdf.Cell(40, 10, "7-bit US-ASCII (American Standard Code for Information Interchange) table")

	codeFontSize := 10.5
	tableHeader := "dec hex   oct  binary char"
	topMargin := 28.0
	pdf.SetFont("Courier", "", codeFontSize)
	x := 10.0
	y := 1.0
	for i := 0; i < 128; i++ {
		pdf.Text(x, topMargin+5*y, fmt.Sprintf("%3d %3x %5O %7b %q", i, i, i, i, i))
		if i > 0 && i%47 == 0 {
			pdf.SetFont("Courier", "B", codeFontSize)
			pdf.Text(x, topMargin-2, tableHeader)
			pdf.SetFont("Courier", "", codeFontSize)
			x += 67
			y = 0
		}
		y++
	}
	pdf.SetFont("Courier", "B", codeFontSize)
	pdf.Text(x, topMargin-2, tableHeader)

	pdf.SetFont("Arial", "", 8)
	pdf.Text(144, topMargin+5*35, "See also")
	pdf.Text(144, topMargin+5*36, "https://en.wikipedia.org/wiki/ASCII")
	pdf.Text(144, topMargin+5*37, "https://en.wikipedia.org/wiki/Code_page_437")
	pdf.Text(144, topMargin+5*38, "https://en.wikipedia.org/wiki/ISO/IEC_8859-1")
	pdf.Text(144, topMargin+5*39, "https://en.wikipedia.org/wiki/Unicode")

	pdf.SetFont("Courier", "", 5.5)
	pdf.Text(144, topMargin+5*41, `for i := 0; i < 128; i++ {`)
	pdf.Text(144, topMargin+5*42, `  pdf.Text(x, y, `)
	pdf.Text(144, topMargin+5*43, `    fmt.Sprintf("%3d %3x %5O %7b %q", i, i, i, i, i))`)
	pdf.Text(144, topMargin+5*44, `}`)

	pdf.SetFont("Arial", "", 8)
	pdf.Text(144, topMargin+5*45, "https://github.com/jung-kurt/gofpdf")
	pdf.Text(144, topMargin+5*46, "https://pkg.go.dev/fmt#hdr-Printing")

	pdf.Text(144, topMargin+5*48, "https://github.com/schnoddelbotz/pdf-ascii-table")

	pdf.SetFont("Arial", "B", 22)
	pdf.SetTextColor(128, 128, 128)
	pdf.Text(77, topMargin+5*50, "64 32 16 8 4 2 1")

	err := pdf.OutputFileAndClose("pdf-ascii-table.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

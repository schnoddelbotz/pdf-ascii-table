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

	pdf.SetFont("Courier", "", 11)
	x := 10.0
	y := 1.0
	for i := 0; i < 128; i++ {
		pdf.Text(x, 25+5*float64(y), fmt.Sprintf("%3d %5O %8b %q", i, i, i, i))
		if i > 0 && i%47 == 0 {
			x += 67
			y = 0
		}
		y++
	}

	pdf.SetFont("Arial", "", 8)
	pdf.Text(144, 25+5*35, "See also")
	pdf.Text(144, 25+5*36, "https://en.wikipedia.org/wiki/ASCII")
	pdf.Text(144, 25+5*37, "https://en.wikipedia.org/wiki/Code_page_437")
	pdf.Text(144, 25+5*38, "https://en.wikipedia.org/wiki/ISO/IEC_8859-1")
	pdf.Text(144, 25+5*39, "https://en.wikipedia.org/wiki/Unicode")

	pdf.SetFont("Courier", "", 6)
	pdf.Text(144, 25+5*41, `for i := 0; i < 128; i++ {`)
	pdf.Text(144, 25+5*42, `  pdf.Text(x, y, `)
	pdf.Text(144, 25+5*43, `    fmt.Sprintf("%3d %5O %8b %q", i, i, i, i))`)
	pdf.Text(144, 25+5*44, `}`)

	pdf.SetFont("Arial", "", 8)
	pdf.Text(144, 25+5*45, "https://github.com/jung-kurt/gofpdf")
	pdf.Text(144, 25+5*46, "https://pkg.go.dev/fmt#hdr-Printing")

	pdf.Text(144, 25+5*48, "https://github.com/schnoddelbotz/pdf-ascii-table")

	err := pdf.OutputFileAndClose("pdf-ascii-table.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

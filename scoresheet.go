package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "pt", "Letter", "")
	pdf.AddPage()
	header(pdf)
	pdf.OutputFileAndClose("test.pdf")
}

func header(pdf *gofpdf.Fpdf) {

	//Grey filled header box
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetFillColor(230, 230, 230)
	pdf.Rect(12, 12, 588, 104, "FD")

	//home checkbox
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetFont("Helvetica", "", 10)
	pdf.Rect(180, 12+13.5, 15, 15, "FD")
	pdf.Text(180-2-pdf.GetStringWidth("Home:"), 12+13.5+10, "Home:")
	pdf.Rect(240, 12+13.5, 15, 15, "FD")
	pdf.Text(240-2-pdf.GetStringWidth("Away:"), 12+13.5+10, "Away:")
	pdf.Rect(456, 12+13.5, 15, 15, "FD")
	pdf.Text(456-2-pdf.GetStringWidth("At Park:"), 12+13.5+10, "At Park:")
	pdf.Rect(516, 12+13.5, 15, 15, "FD")
	pdf.Text(516-2-pdf.GetStringWidth("TV:"), 12+13.5+10, "TV:")
	pdf.Rect(574, 12+13.5, 15, 15, "FD")
	rightCenterAlign(pdf, 574-2, 12+13.5+7.5, "Radio:")

	pdf.Rect(150, 12+42, 350, 20, "FD")
	pdf.Text(150-2-pdf.GetStringWidth("Location:"), 12+42+10+2.5, "Location:")
}

func rightCenterAlign(pdf *gofpdf.Fpdf, x, y float64, text string) {
	_,lh := pdf.GetFontSize()
	pdf.Text(x-pdf.GetStringWidth(text), y+(lh/3), text)
}

/*
func diamond(pdf *gofpdf.Fpdf, x int, y int) {
  pdf.SetStrokeColor(255,255,255);
  pdf.SetLinewidth(0.7);
  pdf.SetDashPattern(1);
  pdf.Rect(x, y, 43, 43, 'D');
  diamond->transform(-translate => [$x + 21.5, $y + 21.5], -rotate => 45,);
  diamond->rect(-10.5, -10.5, 21, 21);
  diamond->strokecolor('lightcyan');
  diamond->linecap(1);
  diamond->linejoin(1);
  diamond->linewidth(0.5);
  diamond->linedash(1);
  diamond->stroke;
  diamond->restore;
}
*/

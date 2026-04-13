package main

func fractal(pen *Pen, tamanho float64, nivel int) {
	if nivel == 0 {
		return
	}

	pen.Walk(tamanho)

	for i := 0; i < 6; i++ {

		pen.Right(60)
		fractal(pen, tamanho*0.5, nivel-1)
		pen.Left(60)

	}

	pen.Walk(-tamanho)
}

func main() {
	pen := NewPen(1200, 1200)

	pen.SetHeading(90)
	pen.SetPosition(600, 600)

	fractal(pen, 200, 5)

	pen.SavePNG("tree.png")
}

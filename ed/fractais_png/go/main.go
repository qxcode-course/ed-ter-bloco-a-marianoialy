package main

func samambaia(pen *Pen, tamanho float64, nivel int) {
	if nivel == 0 {
		return
	}

	pen.Walk(tamanho)

	pen.Left(25)
	samambaia(pen, tamanho*0.6, nivel-1)
	pen.Right(25)

	pen.Right(25)
	samambaia(pen, tamanho*0.6, nivel-1)
	pen.Left(25)

	pen.Walk(-tamanho)
}

func main() {
	pen := NewPen(1200, 1200)
	pen.SetPosition(600, 1200)
	pen.SetHeading(90)

	for i := 0; i < 4; i++ {
		samambaia(pen, 200, 7)
		pen.Up()

		pen.Walk(150)
		pen.Down()
	}

	pen.SavePNG("tree.png")
}

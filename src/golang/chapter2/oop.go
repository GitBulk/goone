package main

import "fmt"

const (
	White = iota
	Black
	Blue
	Red
	Yellow
)

type Box struct {
	width, height, depth float64
	color                Color
}

func (box Box) Volume() float64 {
	return box.width * box.height * box.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

type BoxList []Box // a slice of boxes

func (bl BoxList) PaintItBlack() {
	for i := 0; i < len(bl); i++ {
		bl[i].SetColor(Black)
	}
}

func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(White)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

type Color byte

func (c Color) String() string {
	strings := []string{"White", "Black", "Blue", "Red", "Yellow"}
	return strings[c]
}

func main() {
	fmt.Println("oop")

	boxes := BoxList{
		Box{4, 4, 4, Red},
		Box{10, 10, 1, Yellow},
		Box{1, 1, 20, Black},
		Box{10, 10, 1, Blue},
		Box{10, 30, 1, White},
		Box{20, 20, 20, Yellow},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestsColor().String())
}

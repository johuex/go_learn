package main

//Run at https://tour.golang.org/moretypes/18
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		pic[i] = make([]uint8, dx)
	}
	//inserting w/b pixels
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			switch {
			case j%15&i%15 == 0:
				pic[i][j] = 240
			case j%3 == 0:
				pic[i][j] = 120
			case j%5 == 0:
				pic[i][j] = 150
			default:
				pic[i][j] = 100
			}
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}

package main

import (
	"github.com/gdamore/tcell"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getLength() int {
	length := 30
	var err error
	if len(os.Args) > 1 {
		length, err = strconv.Atoi(os.Args[1])
		if err != nil {
			length = 30
		}
	}
	return length
}

func initScreen() (tcell.Screen, int, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, 0, err
	}
	if err = screen.Init(); err != nil {
		return nil, 0, err
	}
	_, height := screen.Size()
	return screen, height, err
}

func genArrays(length int) ([]int, []int) {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, length)
	brr := make([]int, length)
	for i := range arr {
		arr[i] = i + 1
	}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	copy(brr, arr)
	return arr, brr
}

func draw(arr, brr []int, screen tcell.Screen) {
	_, height := screen.Size()
	screen.Clear()
	for i := range arr {
		style := tcell.StyleDefault
		if brr[i] != arr[i] {
			style = tcell.StyleDefault.Foreground(tcell.ColorIndianRed)
		}
		drawColumn(i, arr[i], height, style, screen)
	}

	screen.Show()
	time.Sleep(50 * time.Millisecond)

}

func drawColumn(i, a, height int, style tcell.Style, screen tcell.Screen) {
	for j := 0; j <= a/2; j++ {
		screen.SetContent(i, height-j, '█', nil, style)
	}
	if a%2 == 1 {
		screen.SetContent(i, height-a/2-1, '▄', nil, style)
	}
}

func done(arr []int, height int, screen tcell.Screen) {
	draw(arr, arr, screen)
	time.Sleep(500 * time.Millisecond)
	//screen.Clear()
	style := tcell.StyleDefault.Foreground(tcell.ColorGreen)
	for i := range arr {
		drawColumn(i, arr[i], height, style, screen)
		screen.Show()
		time.Sleep(20 * time.Millisecond)
	}
	style = tcell.StyleDefault.Foreground(tcell.ColorWhite)
	//time.Sleep(100 * time.Millisecond)
	for i := range arr {
		drawColumn(i, arr[i], height, style, screen)
		screen.Show()
		time.Sleep(20 * time.Millisecond)
	}
	time.Sleep(400 * time.Millisecond)
}

func main() {
	length := getLength()

	screen, height, err := initScreen()
	if err != nil {
		panic("fuck you bitch")
	}
	defer screen.Clear()
	defer screen.Fini()

	screen.Clear()

	go func() {
		for {
			switch screen.PollEvent().(type) {
			case *tcell.EventKey:
				os.Exit(0)
			}
		}
	}()

	arr, brr := genArrays(length)

	//bubble(arr, brr, screen)
	//cocktailShakerSort(arr, brr, screen)
	//quick(arr, brr, screen)
	mysort(arr, brr, screen)

	done(arr, height, screen)
}

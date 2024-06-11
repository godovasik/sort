package main

import (
	"github.com/gdamore/tcell"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func bubbleOneIt(arr *[]int) bool {
	for i := 0; i < len(*arr)-1; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			(*arr)[i], (*arr)[i+1] = (*arr)[i+1], (*arr)[i]
			return false
		}
	}
	return true
}

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

func draw(arr, brr []int, height int, screen tcell.Screen) {
	for i := range arr {
		style := tcell.StyleDefault
		if brr[i] != arr[i] {
			style = tcell.StyleDefault.Foreground(tcell.ColorRed)
		}
		for j := 0; j <= arr[i]/2; j++ {
			screen.SetContent(i, height-j, '█', nil, style)
		}
		if arr[i]%2 == 1 {
			screen.SetContent(i, height-arr[i]/2-1, '▄', nil, style)
		}

	}
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int

	for _, val := range arr[1:] {
		if val < pivot {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
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

	for {
		draw(arr, brr, height, screen)
		copy(brr, arr)
		if bubbleOneIt(&arr) {
			time.Sleep(1000 * time.Millisecond)
			break
		}
		screen.Show()
		time.Sleep(50 * time.Millisecond)

		screen.Clear()
	}
}

package tagger

import (
	"sync"

	"github.com/fatih/color"
)

var (
	mut      = &sync.Mutex{}
	colorInd = 0
	colorMap = map[int]func(string, ...interface{}){
		0: color.Blue,
		1: color.Cyan,
		2: color.Green,
		3: color.Magenta,
		4: color.Red,
		5: color.White,
		6: color.Yellow,
	}
)

func getColorFunc() (colorFunc func(string, ...interface{})) {
	mut.Lock()
	colorFunc = colorMap[colorInd%len(colorMap)]
	colorInd++
	mut.Unlock()
	return
}

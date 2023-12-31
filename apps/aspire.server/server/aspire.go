package server

import (
	"errors"
	"fmt"
	"time"
)

type Instruct struct {
	state         EnumState
	action        func(pos *int, state EnumState, table *[]EnumState, prevPos int)
	describAction EnumDescribAction
}

type Scene struct {
	Table []EnumState `json:"table"`
	Pos   int         `json:"pos"`
}

// State
type EnumState int

const (
	CLEAN EnumState = 0
	DIRTY EnumState = 1
)

// Action desriptor
type EnumDescribAction int

const (
	UNKNOWN EnumDescribAction = -1
	MOVE    EnumDescribAction = 0
	CLEAR   EnumDescribAction = 1
)

// Dirs
type EnumDir int

const (
	TOR EnumDir = -1
	TOL EnumDir = 1
)

func (i Instruct) execAction(pos *int, table *[]EnumState, prevPos int) {
	i.action(pos, i.state, table, prevPos)
}

func Movefunc(pos *int, state EnumState, table *[]EnumState, prevPos int) {
	var dir int = *pos - prevPos

	if dir < 0 {
		if *pos == 0 {
			*pos = *pos + 1
		} else {
			*pos = *pos - 1
		}
	} else {
		if *pos == len(*table)-1 {
			*pos = *pos - 1
		} else {
			*pos = *pos + 1
		}
	}

	fmt.Println("Moving to...", *pos)
}

func Clearfunc(pos *int, state EnumState, table *[]EnumState, prevPos int) {
	(*table)[*pos] = CLEAN
	fmt.Println("Cleaning...", "Pos =>", *pos)
}

var instructions = []Instruct{
	{CLEAN, Movefunc, MOVE},
	{DIRTY, Clearfunc, CLEAR},
}

func aspireLive(table []EnumState, initPos int, initDir EnumDir) ([]Scene, error) {
	fmt.Println("InitDir =>", initDir)
	fmt.Println("InitPos =>", initPos)
	if len(table) == 0 {
		return nil, errors.New("Empty table")
	}

	if len(table)-1 < initPos {
		return nil, errors.New("Invalid initial position")
	}

	pos := initPos
	var prevPos int = pos - int(initDir)
	movs := len(table) * 2 //int(math.Abs(float64(initPos) - float64(initDir)*float64(len(table))))
	fmt.Println("Movs =ffff>", movs)
	beforePos := pos
	var lista []Scene
	tableCopy := make([]EnumState, len(table))
	copy(tableCopy, table)
	lista = append(lista, Scene{tableCopy, pos})

	for true {
		var currentAction Instruct = resolveAction(table[pos], instructions)
		beforePos = pos
		currentAction.execAction(&pos, &table, prevPos)

		tableCopy := make([]EnumState, len(table))
		copy(tableCopy, table)

		lista = append(lista, Scene{tableCopy, pos})
		time.Sleep(10000)
		if currentAction.describAction == MOVE {
			prevPos = beforePos
		}
		if movs == 0 {
			break
		} else if currentAction.describAction == MOVE {
			movs--
		}
	}

	fmt.Println("Movs =ffff>", movs)

	return lista, nil
}

func resolveAction(state EnumState, instructions []Instruct) Instruct {

	for _, ins := range instructions {

		if ins.state == state {
			return ins
		}
	}

	return Instruct{
		state:         -1,
		action:        func(pos *int, state EnumState, table *[]EnumState, prevPos int) {},
		describAction: UNKNOWN,
	}

}

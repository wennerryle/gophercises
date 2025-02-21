package main

import (
	"container/list"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var cmds = list.New()

type CommandDescriber interface {
	ShowDefiniton()
}

func (v Command[T]) ShowDefiniton() {
	fmt.Printf("\t%v\n", v.key)
	fmt.Printf("\t\t%v\n", v.desc)
}

type Command[T any] struct {
	key  string
	desc string
	ptr  *T
}

func addArgument[T int | string | bool](key, desc string, initialValue T) *T {
	cmds.PushBack(Command[T]{
		key,
		desc,
		&initialValue,
	})

	return &initialValue
}

func parseArgs() {
throughArgs:
	for i := 0; i < len(os.Args); i++ {
		currentNode := cmds.Front()

		for currentNode != nil {
			switch v := currentNode.Value.(type) {
			case Command[bool]:
				if v.key == os.Args[i] {
					*v.ptr = true
					continue throughArgs
				}
			case Command[string]:
				if v.key == os.Args[i] {
					i++
					*v.ptr = os.Args[i]
					continue throughArgs
				}
			case Command[int]:
				if v.key == os.Args[i] {
					i++
					number, err := strconv.Atoi(os.Args[i])

					if err != nil {
						panic(fmt.Sprint("Can't transform ", os.Args[i], " to integer"))
					}

					*v.ptr = number
					continue throughArgs
				}
			default:
				panic(
					fmt.Sprint(
						"Args parser not implement this kind of type: ",
						reflect.TypeOf(currentNode.Value),
					),
				)
			}
			currentNode = currentNode.Next()
		}
	}
}

func showDescriptions() {
	for v := cmds.Front(); v != nil; v = v.Next() {
		cmd := v.Value.(CommandDescriber)

		cmd.ShowDefiniton()
	}
}

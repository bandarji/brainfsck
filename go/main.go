package main

import (
	"fmt"
	"log"
	"os"
)

const TAPE_SIZE = 30000

type Fuck struct {
	InsPtr, MemPtr int
	Memory         [TAPE_SIZE]uint8
	Tape           string
	Length         int
	opens, closes  int
}

func NewFuck(filename string) (f *Fuck) {
	f = &Fuck{}
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Could not read file", filename, "Error:", err)
	}
	f.Tape = string(content)
	f.Length = len(f.Tape)
	return
}

func (f *Fuck) Run() {
	for f.InsPtr < f.Length {
		switch f.Tape[f.InsPtr] {
		case '>':
			f.MemPtr++
			if f.MemPtr == TAPE_SIZE {
				f.MemPtr = 0
			}
		case '<':
			f.MemPtr--
			if f.MemPtr == 0 {
				f.MemPtr = TAPE_SIZE - 1
			}
		case '+':
			if f.Memory[f.MemPtr] == 255 {
				f.Memory[f.MemPtr] = 0
			} else {
				f.Memory[f.MemPtr]++
			}
		case '-':
			if f.Memory[f.MemPtr] == 0 {
				f.Memory[f.MemPtr] = 255
			} else {
				f.Memory[f.MemPtr]--
			}
		case '.':
			fmt.Printf("%c", f.Memory[f.MemPtr])
		case ',':
			// read a char
		case '[':
			if f.Memory[f.MemPtr] == 0 {
				f.opens = 0
				f.InsPtr++
				for f.InsPtr < f.Length {
					if f.Tape[f.InsPtr] == ']' && f.opens == 0 {
						f.InsPtr++
						break
					} else if f.Tape[f.InsPtr] == '[' {
						f.opens++
					} else if f.Tape[f.InsPtr] == ']' {
						f.opens--
					}
					f.InsPtr++
				}
			}
		case ']':
			if f.Memory[f.MemPtr] != 0 {
				f.closes = 0
				f.InsPtr--
				for f.InsPtr >= 0 {
					if f.Tape[f.InsPtr] == '[' && f.closes == 0 {
						f.InsPtr--
						break
					} else if f.Tape[f.InsPtr] == ']' {
						f.closes++
					} else if f.Tape[f.InsPtr] == '[' {
						f.closes--
					}
					f.InsPtr--
				}
			}

		}
		f.InsPtr++
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	} else {
		f := NewFuck(os.Args[1])
		f.Run()
		// log.Print(f.Tape)
	}
}

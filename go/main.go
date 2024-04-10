package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	mem    [30000]byte
	tape   string
	mp, ip int // memory pointer, instruction pointer
)

func GetTape() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <bf_file>\n", os.Args[0])
	}
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading file: %s\n", err)
	}
	tape = string(bytes)
}

func Jump(direction int) {
	for block := direction; block*direction > 0; ip += direction {
		switch tape[ip+direction] {
		case '[':
			block++
		case ']':
			block--
		}
	}
}

func Run() {
	input := bufio.NewReader(os.Stdin)
	for ip < len(tape) {
		switch tape[ip] {
		case '>':
			mp++
		case '<':
			mp--
		case '+':
			if mem[mp] == 255 {
				mem[mp] = 0
			} else {
				mem[mp]++
			}
		case '-':
			if mem[mp] == 0 {
				mem[mp] = 255
			} else {
				mem[mp]--
			}
		case '.':
			fmt.Printf("%c", mem[mp])
		case ',':
			mem[mp], _ = input.ReadByte()
		case '[':
			if mem[mp] == 0 {
				Jump(1)
			}
		case ']':
			if mem[mp] != 0 {
				Jump(-1)
			}
		}
		ip++
	}
}

func main() {
	GetTape()
	Run()
}

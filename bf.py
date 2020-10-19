#!/usr/bin/env python3

import sys


class Tape:

    def __init__(self, filename):
        self.skip_preamble = False
        self.ins_ptr, self.mem_ptr = 0, 0
        self.memory = [0 for _ in range(30000)]
        self.content = read_tape(filename)
        self.length = len(self.content)
        if self.content[self.ins_ptr] == '[':
            self.skip_preamble = True
            self.jmp_fwd()

    def instruction(self):
        return self.content[self.ins_ptr]

    def inc_ptr(self):
        self.mem_ptr += 1
        if self.mem_ptr == 30000:
            self.mem_ptr = 0

    def dec_ptr(self):
        self.mem_ptr -= 1
        if self.mem_ptr < 0:
            self.mem_ptr = 29999

    def inc_val(self):
        self.memory[self.mem_ptr] += 1
        if self.memory[self.mem_ptr] > 255:
            self.memory[self.mem_ptr] = 0

    def dec_val(self):
        self.memory[self.mem_ptr] -= 1
        if self.memory[self.mem_ptr] < 0:
            self.memory[self.mem_ptr] = 255

    def byte_out(self):
        sys.stdout.write(chr(self.memory[self.mem_ptr]))
        sys.stdout.flush()

    def byte_in(self):
        self.memory[self.mem_ptr] = ord(sys.stdin.read(1))

    def jmp_fwd(self):
        if self.memory[self.mem_ptr] == 0 or self.skip_preamble:
            opens = 0
            self.ins_ptr += 1
            while self.ins_ptr < self.length:
                if self.content[self.ins_ptr] == ']' and opens == 0:
                    break  # matching close bracket discovered
                elif self.content[self.ins_ptr] == '[':
                    opens += 1
                elif self.content[self.ins_ptr] == ']':
                    opens -= 1
                else:
                    pass
                self.ins_ptr += 1


    def jmp_bak(self):
        if self.memory[self.mem_ptr] != 0:
            closes = 0
            self.ins_ptr -= 1
            while self.ins_ptr >= 0:
                if self.content[self.ins_ptr] == '[' and closes == 0:
                    break
                elif self.content[self.ins_ptr] == ']':
                    closes += 1
                elif self.content[self.ins_ptr] == '[':
                    closes -= 1
                else:
                    pass
                self.ins_ptr -= 1


def read_tape(filename):
    content = ''
    try:
        with open(filename) as file_handle:
            content = file_handle.read()
    except:
        pass
    return content


def _noop():
    pass


def fuck(tape):
    commands = {
        '>': tape.inc_ptr,
        '<': tape.dec_ptr,
        '+': tape.inc_val,
        '-': tape.dec_val,
        '.': tape.byte_out,
        ',': tape.byte_in,
        '[': tape.jmp_fwd,
        ']': tape.jmp_bak,
    }
    while tape.ins_ptr < tape.length:
        commands.get(tape.instruction(), _noop)()
        tape.ins_ptr += 1


def main():
    if len(sys.argv) < 2:
        raise SystemExit('No input file specified.')
    tape = Tape(sys.argv[1])
    fuck(tape)


if __name__ == '__main__':
    main()

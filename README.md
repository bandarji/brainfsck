# brainfsck

My interpreter for Urban Müller's esoteric programming language, Brainfuck,
written in Go and Python.

## References

1. [Wikipedia article](https://en.wikipedia.org/wiki/Brainfuck)

## Usage

### Go

```
make build && ./bf <fsck'd file>
```

### Python

```
$ bf.py --help
```

## Specifications

1. Start with thirty thousand bytes/cells initialized with values of zero.
1. Track two cursors: position and instruction.
1. Ignore everything between balanced brackets if the left bracket exists as the first character on the instruction tape.
1. Accept single-key input and output, using ASCII characters.


## Commands

| Instruction | Description |
| :--- | :--- |
| `[` | jump to position after corresponding right bracket, if zero |
| `]` | jump back to corresponding left bracket, if not zero |
| `,` | accept one input byte |
| `.` | output value at the position pointer |
| `+` | increment the value at the position pointer (255 rotates to zero) |
| `-` | decrement the value at the position pointer (0 rotates to 255) |
| `>` | increment the position pointer (rotate to 0 at tape end) |
| `<` | decrement the position pointer (rotate to 29999 at tape end) |
```

## Tests

This repository includes test files. The following string should produce "Hello World!".

```
++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>
---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.
```

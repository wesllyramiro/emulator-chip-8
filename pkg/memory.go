package chip

type data struct {
	ram            [4096]byte
	register       [16]byte // byte == uint8
	i              uint16
	programCounter uint16
	stackPointer   byte
	stack          [16]uint16
	delayTimer     byte
	soundTimer     byte
	keyboard       [16]bool
	frameBuffer    [256]byte
	op             opcode
}

type opcode uint16

var characters [4096]byte = [4096]byte{
	// localização de memória de cada letra
	// é "letra" em hexadecimal VEZES 5
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // a
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // b
	0xF0, 0x80, 0x80, 0x80, 0xF0, // c
	0xE0, 0x90, 0x90, 0x90, 0xE0, // d
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // e
	0xF0, 0x80, 0xF0, 0x80, 0x80, // f
}
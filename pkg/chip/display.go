package chip

import "fmt"

func DrawSprite(m *Data, x, y, height int, address uint16) bool {
	colisão := false
	if x > 63 {
		return false
	}
	if y > 31 {
		return false
	}

	byteNaLinha := x / 8
	offsetProBit := x % 8

	for byte := 0; byte < height; byte++ {
		byteAddress := address + uint16(byte)
		line := (y + byte) % 32
		fb1 := line*8 + byteNaLinha
		fb2 := line*8 + ((byteNaLinha + 1) % 8)
		spr1 := m.ram[byteAddress] >> offsetProBit
		spr2 := m.ram[byteAddress] << (8 - offsetProBit)
		if m.frameBuffer[fb1]&spr1 != 0 {
			colisão = true
		}
		if m.frameBuffer[fb2]&spr2 != 0 {
			colisão = true
		}
		m.frameBuffer[fb1] ^= m.ram[byteAddress] >> offsetProBit
		m.frameBuffer[fb2] ^= m.ram[byteAddress] << (8 - offsetProBit)
	}
	return colisão
}

func DrawDisplay(m *Data) {

	black := '█'
	white := ' '
	lineBuffer := []rune{}
	var mask byte = 0b10000000
	for line := 0; line < 32; line++ {
		for column := 0; column < 8; column++ {
			for bit := 0; bit < 8; bit++ {

				if m.frameBuffer[(line*8)+column]&(mask>>bit) != 0 {
					lineBuffer = append(lineBuffer, black)
					lineBuffer = append(lineBuffer, black)
				} else {
					lineBuffer = append(lineBuffer, white)
					lineBuffer = append(lineBuffer, white)
				}
			}
		}
		fmt.Println(line, "\t", string(lineBuffer[:]))
		lineBuffer = []rune{}
	}

}

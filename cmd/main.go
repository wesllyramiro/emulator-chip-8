package main

import (
	"fmt"

	"github.com/wesllyramiro/emulator-chip-8/pkg/chip"
)

func main() {
	memory := chip.Data{Ram: chip.Characters}
	fmt.Println()
	fmt.Println()
	for l := 0; l < 6; l++ {
		chip.DrawSprite(&memory, l*5, l*5, 5, uint16(l*5))
	}
	colisão := chip.DrawSprite(&memory, 0, 15, 10, uint16(0x0*5))
	chip.DrawDisplay(&memory)
	if colisão {
		fmt.Println("deu colisão")
	}
}

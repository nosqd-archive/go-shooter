package utils

import rl "github.com/gen2brain/raylib-go/raylib"

func Uint32ToColor(value uint32) rl.Color {
	red := uint8(value >> 24)   // Extract the first 8 bits
	green := uint8(value >> 16) // Extract the next 8 bits
	blue := uint8(value >> 8)   // Extract the next 8 bits
	alpha := uint8(255)         // Set alpha to 255 (fully opaque)
	return rl.NewColor(red, green, blue, alpha)
}

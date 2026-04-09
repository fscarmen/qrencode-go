package main

import "fmt"

// DrawASCIIBitmap 将二维码矩阵渲染为终端字符图。
// 支持紧凑模式、静区控制和黑白反转。
func DrawASCIIBitmap(matrix [][]bool, quietZone int, compact bool, invert bool) {
	height := len(matrix)
	if height == 0 {
		return
	}
	width := len(matrix[0])

	// 计算包含静区后的总尺寸
	totalW := width + quietZone*2
	totalH := height + quietZone*2

	// 根据模式选择渲染方式
	if compact {
		drawCompact(matrix, totalW, totalH, quietZone, invert)
	} else {
		drawFull(matrix, totalW, totalH, quietZone, invert)
	}
}

// drawCompact 使用半高字符（▀ ▄ █）进行压缩渲染。
func drawCompact(matrix [][]bool, totalW, totalH, qz int, invert bool) {
	for y := 0; y < totalH; y += 2 {
		for x := 0; x < totalW; x++ {
			top := cell(matrix, x, y, qz, invert)
			bottom := false
			if y+1 < totalH {
				bottom = cell(matrix, x, y+1, qz, invert)
			}

			switch {
			case top && bottom:
				fmt.Print("█")
			case top && !bottom:
				fmt.Print("▀")
			case !top && bottom:
				fmt.Print("▄")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// drawFull 使用完整尺寸字符进行渲染。
func drawFull(matrix [][]bool, totalW, totalH, qz int, invert bool) {
	for y := 0; y < totalH; y++ {
		for x := 0; x < totalW; x++ {
			v := cell(matrix, x, y, qz, invert)
			if v {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

// cell 计算指定位置是否应显示为实心块。
// 会处理静区边框以及黑白反转逻辑。
func cell(matrix [][]bool, x, y, qz int, invert bool) bool {
	gx := x - qz
	gy := y - qz

	// 判断是否在二维码实际区域内
	inQR := gx >= 0 && gy >= 0 && gy < len(matrix) && gx < len(matrix[gy])

	// 静区区域显示为空白（白边）
	if !inQR {
		return true
	}

	v := matrix[gy][gx]

	// 根据参数决定是否反转黑白
	if invert {
		v = !v
	}

	return v
}
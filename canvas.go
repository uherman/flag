package main

import "strings"

type Canvas struct {
	Width  int
	Height int
	Pixels [][]RGB
}

func NewCanvas(width, height int) *Canvas {
	if height%2 != 0 {
		height++
	}
	pixels := make([][]RGB, height)
	for i := range pixels {
		pixels[i] = make([]RGB, width)
	}
	return &Canvas{Width: width, Height: height, Pixels: pixels}
}

func (c *Canvas) Set(x, y int, color RGB) {
	if x >= 0 && x < c.Width && y >= 0 && y < c.Height {
		c.Pixels[y][x] = color
	}
}

func (c *Canvas) Fill(color RGB) {
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			c.Pixels[y][x] = color
		}
	}
}

func (c *Canvas) FillRect(x1, y1, x2, y2 int, color RGB) {
	for y := y1; y < y2; y++ {
		for x := x1; x < x2; x++ {
			c.Set(x, y, color)
		}
	}
}

func (c *Canvas) Render() string {
	var buf strings.Builder
	for y := 0; y < c.Height; y += 2 {
		prevTop := RGB{255, 255, 254} // sentinel — won't match any real color
		prevBot := RGB{255, 255, 254}
		for x := 0; x < c.Width; x++ {
			top := c.Pixels[y][x]
			bot := c.Pixels[y+1][x]
			if top != prevTop || bot != prevBot {
				buf.WriteString(top.FgEsc())
				buf.WriteString(bot.BgEsc())
				prevTop = top
				prevBot = bot
			}
			buf.WriteRune('▀')
		}
		buf.WriteString(Reset)
		buf.WriteByte('\n')
	}
	return buf.String()
}

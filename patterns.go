package main

import "math"

func HStripes(c *Canvas, colors []RGB) {
	n := len(colors)
	for i, col := range colors {
		y1 := c.Height * i / n
		y2 := c.Height * (i + 1) / n
		c.FillRect(0, y1, c.Width, y2, col)
	}
}

func VStripes(c *Canvas, colors []RGB) {
	n := len(colors)
	for i, col := range colors {
		x1 := c.Width * i / n
		x2 := c.Width * (i + 1) / n
		c.FillRect(x1, 0, x2, c.Height, col)
	}
}

func NordicCross(c *Canvas, bg, cross RGB, outerCross *RGB) {
	c.Fill(bg)

	// Cross center: offset left at ~6/16 of width, vertically centered
	cx := c.Width * 6 / 16
	cy := c.Height / 2

	// Cross arm thickness
	innerThick := c.Height / 5
	if innerThick < 2 {
		innerThick = 2
	}

	if outerCross != nil {
		outerThick := innerThick + 2
		// Horizontal bar
		c.FillRect(0, cy-outerThick/2, c.Width, cy-outerThick/2+outerThick, *outerCross)
		// Vertical bar
		c.FillRect(cx-outerThick/2, 0, cx-outerThick/2+outerThick, c.Height, *outerCross)
	}

	// Horizontal bar
	c.FillRect(0, cy-innerThick/2, c.Width, cy-innerThick/2+innerThick, cross)
	// Vertical bar
	c.FillRect(cx-innerThick/2, 0, cx-innerThick/2+innerThick, c.Height, cross)
}

func Circle(c *Canvas, bg, circleColor RGB, radiusFraction float64) {
	c.Fill(bg)
	cx := float64(c.Width) / 2.0
	cy := float64(c.Height) / 2.0
	radius := radiusFraction * float64(c.Height)

	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			dx := float64(x) + 0.5 - cx
			dy := float64(y) + 0.5 - cy
			if math.Sqrt(dx*dx+dy*dy) <= radius {
				c.Set(x, y, circleColor)
			}
		}
	}
}

func SwissCross(c *Canvas, bg, crossColor RGB) {
	c.Fill(bg)
	// Cross is 6/32 wide and 20/32 tall (and rotated for the horizontal bar)
	armW := c.Width * 6 / 32
	if armW < 2 {
		armW = 2
	}
	armL := c.Width * 20 / 32

	cx := c.Width / 2
	cy := c.Height / 2

	// Vertical bar
	c.FillRect(cx-armW/2, cy-armL/2, cx-armW/2+armW, cy-armL/2+armL, crossColor)
	// Horizontal bar
	c.FillRect(cx-armL/2, cy-armW/2, cx-armL/2+armL, cy-armW/2+armW, crossColor)
}

func UnionJack(c *Canvas) {
	c.Fill(UKBlue)

	// Diagonal stripe half-widths (in pixels)
	whiteW := 3.0
	redW := 1.5

	w := float64(c.Width)
	h := float64(c.Height)

	// Draw diagonals
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fx := float64(x) + 0.5
			fy := float64(y) + 0.5

			// Distance to the two diagonals (top-left to bottom-right, top-right to bottom-left)
			// Using the perpendicular distance formula
			diag := math.Sqrt(w*w + h*h)

			// TL-BR diagonal: y/h = x/w => h*x - w*y = 0
			d1 := math.Abs(h*fx-w*fy) / diag
			// TR-BL diagonal: y/h = (w-x)/w => h*x + w*y - w*h = 0
			d2 := math.Abs(h*fx+w*fy-w*h) / diag

			if d1 < whiteW || d2 < whiteW {
				c.Set(x, y, White)
			}
		}
	}

	// Red diagonals (thinner, with counterchange offset)
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fx := float64(x) + 0.5
			fy := float64(y) + 0.5

			diag := math.Sqrt(w*w + h*h)
			d1 := (h*fx - w*fy) / diag
			d2 := (h*fx + w*fy - w*h) / diag

			// Counterchange: offset the red stripe differently in each quadrant
			inTop := fy < h/2
			inLeft := fx < w/2

			// TL-BR diagonal
			if (inTop && inLeft) || (!inTop && !inLeft) {
				if d1 > -redW && d1 < 0.5 {
					c.Set(x, y, UKRed)
				}
			} else {
				if d1 < redW && d1 > -0.5 {
					c.Set(x, y, UKRed)
				}
			}

			// TR-BL diagonal
			if (inTop && !inLeft) || (!inTop && inLeft) {
				if d2 > -redW && d2 < 0.5 {
					c.Set(x, y, UKRed)
				}
			} else {
				if d2 < redW && d2 > -0.5 {
					c.Set(x, y, UKRed)
				}
			}
		}
	}

	// White cross (St. George's outline)
	crossW := c.Height / 3
	crossH := c.Width / 6
	cy := c.Height / 2
	cx := c.Width / 2
	c.FillRect(0, cy-crossW/2, c.Width, cy-crossW/2+crossW, White)
	c.FillRect(cx-crossH/2, 0, cx-crossH/2+crossH, c.Height, White)

	// Red cross (St. George's)
	redCrossW := crossW * 3 / 5
	redCrossH := crossH * 3 / 5
	c.FillRect(0, cy-redCrossW/2, c.Width, cy-redCrossW/2+redCrossW, UKRed)
	c.FillRect(cx-redCrossH/2, 0, cx-redCrossH/2+redCrossH, c.Height, UKRed)
}

func USAFlag(c *Canvas) {
	// 13 stripes
	for i := 0; i < 13; i++ {
		y1 := i * c.Height / 13
		y2 := (i + 1) * c.Height / 13
		col := USARed
		if i%2 == 1 {
			col = White
		}
		c.FillRect(0, y1, c.Width, y2, col)
	}

	// Blue canton: 7 stripes tall, ~40% wide
	cantonH := 7 * c.Height / 13
	cantonW := c.Width * 2 / 5
	c.FillRect(0, 0, cantonW, cantonH, USABlue)

	// Stars: 9 rows, alternating 6 and 5 stars (real US flag pattern)
	for row := 0; row < 9; row++ {
		cols := 6
		offset := 0
		if row%2 == 1 {
			cols = 5
			offset = cantonW / 12 // half-cell offset for odd rows
		}
		for col := 0; col < cols; col++ {
			sx := offset + (col*2+1)*cantonW/12
			sy := (row*2 + 1) * cantonH / 18
			c.Set(sx, sy, White)
		}
	}
}

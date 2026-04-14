package main

import "math"

func drawAlgeria(c *Canvas) {
	green := RGB{0, 98, 51}
	red := RGB{210, 16, 52}
	VStripes(c, []RGB{green, White})
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	Crescent(c, cx, cy, float64(c.Height)*0.3, float64(c.Height)*0.25, 2, red)
	FiveStar(c, cx+float64(c.Height)*0.15, cy, float64(c.Height)*0.1, red)
}

func drawAntiguaAndBarbuda(c *Canvas) {
	red := RGB{206, 17, 38}
	blue := RGB{0, 0, 148}
	gold := RGB{252, 209, 22}
	c.Fill(red)
	// Black top triangle
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			if fy > 0.33 {
				mid := 0.5
				halfW := (fy - 0.33) / 0.67 * 0.5
				if fx >= mid-halfW && fx <= mid+halfW {
					c.Set(x, y, Black)
				}
			}
		}
	}
	// Blue middle band
	c.FillRect(0, c.Height/3, c.Width, c.Height/3+c.Height/3, blue)
	// Yellow half-sun at top of blue band
	for y := c.Height / 3; y < c.Height*2/3; y++ {
		for x := 0; x < c.Width; x++ {
			cx := float64(c.Width) / 2
			cy := float64(c.Height) / 3
			dx := float64(x) + 0.5 - cx
			dy := float64(y) + 0.5 - cy
			if dx*dx+dy*dy <= float64(c.Height*c.Height)/16 && y >= c.Height/3 {
				c.Set(x, y, gold)
			}
		}
	}
}

func drawAustralia(c *Canvas) {
	blue := RGB{0, 0, 139}
	red := RGB{200, 16, 46}
	c.Fill(blue)
	// Union Jack canton (top-left quarter)
	cw := c.Width / 2
	ch := c.Height / 2
	ujCanvas := NewCanvas(cw, ch)
	UnionJack(ujCanvas)
	for y := 0; y < ch; y++ {
		for x := 0; x < cw; x++ {
			c.Set(x, y, ujCanvas.Pixels[y][x])
		}
	}
	// Commonwealth star (7-pointed, simplified as large white dot cluster)
	sx := c.Width / 4
	sy := c.Height * 3 / 4
	FilledCircle(c, float64(sx), float64(sy), 3, White)
	// Southern Cross (right half) — 5 stars
	stars := [][2]float64{
		{0.77, 0.20}, {0.85, 0.40}, {0.75, 0.60}, {0.85, 0.75}, {0.63, 0.45},
	}
	for _, s := range stars {
		px := int(s[0] * float64(c.Width))
		py := int(s[1] * float64(c.Height))
		_ = red
		FilledCircle(c, float64(px), float64(py), 1.5, White)
	}
}

func drawAzerbaijan(c *Canvas) {
	blue := RGB{0, 151, 230}
	red := RGB{224, 0, 52}
	green := RGB{0, 181, 79}
	HStripes(c, []RGB{blue, red, green})
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	Crescent(c, cx-2, cy, float64(c.Height)*0.22, float64(c.Height)*0.18, 2, White)
	FiveStar(c, cx+float64(c.Height)*0.1, cy, float64(c.Height)*0.08, White)
}

func drawBahamas(c *Canvas) {
	aqua := RGB{0, 168, 180}
	gold := RGB{255, 184, 28}
	HStripes(c, []RGB{aqua, gold, aqua})
	LeftTriangle(c, Black, 0.4)
}

func drawBahrain(c *Canvas) {
	red := RGB{206, 17, 38}
	c.Fill(red)
	zigW := c.Width * 3 / 10
	teeth := 5
	toothH := c.Height / teeth
	for y := 0; y < c.Height; y++ {
		pos := y % toothH
		half := toothH / 2
		var offset int
		if pos < half {
			offset = (zigW / 4) * pos / half
		} else {
			offset = (zigW / 4) * (toothH - pos) / half
		}
		for x := 0; x < zigW-offset; x++ {
			c.Set(x, y, White)
		}
	}
}

func drawBangladesh(c *Canvas) {
	green := RGB{0, 106, 78}
	red := RGB{244, 42, 65}
	c.Fill(green)
	cx := float64(c.Width)*0.45 // slightly left of center
	cy := float64(c.Height) / 2
	FilledCircle(c, cx, cy, float64(c.Height)*0.3, red)
}

func drawBenin(c *Canvas) {
	green := RGB{0, 130, 63}
	yellow := RGB{252, 209, 22}
	red := RGB{210, 16, 52}
	c.FillRect(0, 0, c.Width, c.Height/2, yellow)
	c.FillRect(0, c.Height/2, c.Width, c.Height, red)
	c.FillRect(0, 0, c.Width*2/5, c.Height, green)
}

func drawBhutan(c *Canvas) {
	orange := RGB{255, 127, 14}
	yellow := RGB{255, 210, 0}
	// Diagonal split
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			if float64(y)/float64(c.Height) > float64(x)/float64(c.Width) {
				c.Set(x, y, orange)
			} else {
				c.Set(x, y, yellow)
			}
		}
	}
}

func drawBosnia(c *Canvas) {
	blue := RGB{0, 20, 137}
	yellow := RGB{252, 221, 9}
	c.Fill(blue)
	// Yellow triangle
	x1 := c.Width * 1 / 5
	for y := 0; y < c.Height; y++ {
		fy := float64(y) / float64(c.Height)
		x2 := x1 + int(fy*float64(c.Width)*3/5)
		for x := x1; x < x2; x++ {
			c.Set(x, y, yellow)
		}
	}
	// White stars along the hypotenuse
	for i := 0; i < 7; i++ {
		t := float64(i+1) / 8.0
		sx := float64(x1) + t*float64(c.Width)*3/5
		sy := t * float64(c.Height)
		FiveStar(c, sx, sy, 2, White)
	}
}

func drawBotswana(c *Canvas) {
	blue := RGB{117, 170, 219}
	c.Fill(blue)
	mid := c.Height / 2
	thick := c.Height / 5
	c.FillRect(0, mid-thick/2-1, c.Width, mid-thick/2, White)
	c.FillRect(0, mid-thick/2, c.Width, mid+thick/2, Black)
	c.FillRect(0, mid+thick/2, c.Width, mid+thick/2+1, White)
}

func drawBrazil(c *Canvas) {
	green := RGB{0, 155, 58}
	yellow := RGB{255, 223, 0}
	blue := RGB{0, 39, 118}
	c.Fill(green)
	Diamond(c, yellow, 0.85, 0.85)
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	FilledCircle(c, cx, cy, float64(c.Height)*0.25, blue)
	// White band across the circle
	bandH := c.Height / 12
	for x := 0; x < c.Width; x++ {
		for y := int(cy) - bandH/2; y < int(cy)+bandH/2; y++ {
			dx := float64(x) + 0.5 - cx
			dy := float64(y) + 0.5 - cy
			r := float64(c.Height) * 0.25
			if dx*dx+dy*dy <= r*r {
				c.Set(x, y, White)
			}
		}
	}
}

func drawBurundi(c *Canvas) {
	red := RGB{206, 17, 38}
	green := RGB{30, 181, 58}
	c.Fill(White)
	// Saltire divides into 4 triangles: top/bottom = red, left/right = green
	w := float64(c.Width)
	h := float64(c.Height)
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fx := float64(x) + 0.5
			fy := float64(y) + 0.5
			// Determine quadrant by which diagonal we're closer to
			top := fy/h < fx/w && fy/h < 1.0-fx/w
			bottom := fy/h > fx/w && fy/h > 1.0-fx/w
			if top || bottom {
				c.Set(x, y, red)
			} else {
				c.Set(x, y, green)
			}
		}
	}
	Saltire(c, White, 2)
	// Central white circle with 3 stars
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	FilledCircle(c, cx, cy, float64(c.Height)*0.22, White)
	FiveStar(c, cx, cy-float64(c.Height)*0.1, float64(c.Height)*0.06, red)
	FiveStar(c, cx-float64(c.Height)*0.08, cy+float64(c.Height)*0.06, float64(c.Height)*0.06, red)
	FiveStar(c, cx+float64(c.Height)*0.08, cy+float64(c.Height)*0.06, float64(c.Height)*0.06, red)
}

func drawCaboVerde(c *Canvas) {
	blue := RGB{0, 59, 118}
	red := RGB{205, 42, 62}
	gold := RGB{252, 209, 22}
	c.Fill(blue)
	// White and red stripes in lower portion
	y1 := c.Height * 6 / 10
	c.FillRect(0, y1, c.Width, y1+c.Height/20, White)
	c.FillRect(0, y1+c.Height/20, c.Width, y1+c.Height*3/20, red)
	c.FillRect(0, y1+c.Height*3/20, c.Width, y1+c.Height/5, White)
	// Circle of stars
	cx := float64(c.Width) * 0.35
	cy := float64(c.Height) * 0.65
	for i := 0; i < 10; i++ {
		angle := -math.Pi/2 + float64(i)*2*math.Pi/10
		sx := cx + float64(c.Height)*0.18*math.Cos(angle)
		sy := cy + float64(c.Height)*0.18*math.Sin(angle)
		FiveStar(c, sx, sy, float64(c.Height)*0.03, gold)
	}
}

func drawCambodia(c *Canvas) {
	blue := RGB{3, 35, 119}
	red := RGB{230, 5, 23}
	HStripes(c, []RGB{blue, red, red, blue})
	// Simplified Angkor Wat silhouette
	cx := c.Width / 2
	base := c.Height * 55 / 100
	c.FillRect(cx-8, base-4, cx+8, base, White)
	c.FillRect(cx-6, base-8, cx+6, base-4, White)
	c.FillRect(cx-2, base-12, cx+2, base-8, White)
}

func drawCameroon(c *Canvas) {
	green := RGB{0, 122, 77}
	red := RGB{206, 17, 38}
	yellow := RGB{252, 209, 22}
	VStripes(c, []RGB{green, red, yellow})
	FiveStar(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.15, yellow)
}

func drawCanada(c *Canvas) {
	red := RGB{255, 0, 0}
	c.Fill(White)
	// Red bands on left and right (each 1/4 width)
	c.FillRect(0, 0, c.Width/4, c.Height, red)
	c.FillRect(c.Width*3/4, 0, c.Width, c.Height, red)
	// Simplified maple leaf as a diamond/cross shape in center
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	r := float64(c.Height) * 0.3
	// Draw a simplified leaf shape
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			dx := math.Abs(float64(x)+0.5-cx) / r
			dy := (float64(y) + 0.5 - cy) / r
			// Leaf-ish shape using diamond + extensions
			if dx+math.Abs(dy) < 0.8 && dx < 0.5 {
				c.Set(x, y, red)
			}
			// Stem
			if dx < 0.08 && dy > 0 && dy < 1.0 {
				c.Set(x, y, red)
			}
		}
	}
}

func drawCAR(c *Canvas) {
	blue := RGB{0, 35, 149}
	green := RGB{0, 152, 64}
	yellow := RGB{255, 206, 0}
	red := RGB{210, 16, 52}
	// 4 horizontal stripes
	HStripes(c, []RGB{blue, White, green, yellow})
	// Red vertical stripe through center
	c.FillRect(c.Width*4/10, 0, c.Width*6/10, c.Height, red)
	// Yellow star in top-left
	FiveStar(c, float64(c.Width)*0.1, float64(c.Height)*0.12, float64(c.Height)*0.1, yellow)
}

func drawChile(c *Canvas) {
	red := RGB{213, 0, 50}
	blue := RGB{0, 57, 166}
	c.FillRect(0, 0, c.Width, c.Height/2, White)
	c.FillRect(0, c.Height/2, c.Width, c.Height, red)
	// Blue canton
	cw := c.Width / 3
	c.FillRect(0, 0, cw, c.Height/2, blue)
	FiveStar(c, float64(cw)/2, float64(c.Height)/4, float64(c.Height)*0.12, White)
}

func drawChina(c *Canvas) {
	red := RGB{222, 41, 16}
	yellow := RGB{255, 222, 0}
	c.Fill(red)
	// Large star
	FiveStar(c, float64(c.Width)*0.2, float64(c.Height)*0.3, float64(c.Height)*0.2, yellow)
	// 4 small stars
	smallR := float64(c.Height) * 0.07
	FiveStar(c, float64(c.Width)*0.35, float64(c.Height)*0.15, smallR, yellow)
	FiveStar(c, float64(c.Width)*0.4, float64(c.Height)*0.25, smallR, yellow)
	FiveStar(c, float64(c.Width)*0.4, float64(c.Height)*0.38, smallR, yellow)
	FiveStar(c, float64(c.Width)*0.35, float64(c.Height)*0.48, smallR, yellow)
}

func drawComoros(c *Canvas) {
	yellow := RGB{252, 209, 22}
	blue := RGB{61, 86, 168}
	red := RGB{206, 17, 38}
	green := RGB{0, 130, 63}
	HStripes(c, []RGB{yellow, White, red, blue})
	LeftTriangle(c, green, 0.4)
	cx := float64(c.Width) * 0.13
	cy := float64(c.Height) / 2
	Crescent(c, cx, cy, float64(c.Height)*0.18, float64(c.Height)*0.15, 1.5, White)
}

func drawCongoDRC(c *Canvas) {
	blue := RGB{0, 114, 198}
	red := RGB{206, 17, 38}
	yellow := RGB{247, 168, 0}
	c.Fill(blue)
	// Diagonal red band with yellow borders
	w := float64(c.Width)
	h := float64(c.Height)
	diag := math.Sqrt(w*w + h*h)
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fx := float64(x) + 0.5
			fy := float64(y) + 0.5
			d := math.Abs(h*fx+w*fy-w*h) / diag
			if d < float64(c.Height)*0.15 {
				c.Set(x, y, red)
			} else if d < float64(c.Height)*0.2 {
				c.Set(x, y, yellow)
			}
		}
	}
	// Yellow star in top-left
	FiveStar(c, float64(c.Width)*0.12, float64(c.Height)*0.18, float64(c.Height)*0.15, yellow)
}

func drawCongoRepublic(c *Canvas) {
	green := RGB{0, 155, 58}
	yellow := RGB{252, 209, 22}
	red := RGB{220, 36, 31}
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fx := float64(x) / float64(c.Width)
			fy := float64(y) / float64(c.Height)
			if fx+fy < 0.67 {
				c.Set(x, y, green)
			} else if fx+fy > 1.33 {
				c.Set(x, y, red)
			} else {
				c.Set(x, y, yellow)
			}
		}
	}
}

func drawCostaRica(c *Canvas) {
	blue := RGB{0, 20, 137}
	red := RGB{206, 17, 38}
	HStripes(c, []RGB{blue, White, red, red, White, blue})
}

func drawCuba(c *Canvas) {
	blue := RGB{0, 42, 101}
	red := RGB{204, 0, 0}
	HStripes(c, []RGB{blue, White, blue, White, blue})
	LeftTriangle(c, red, 0.4)
	cx := float64(c.Width) * 0.14
	cy := float64(c.Height) / 2
	FiveStar(c, cx, cy, float64(c.Height)*0.12, White)
}

func drawCzechRepublic(c *Canvas) {
	blue := RGB{17, 69, 126}
	red := RGB{215, 20, 26}
	c.FillRect(0, 0, c.Width, c.Height/2, White)
	c.FillRect(0, c.Height/2, c.Width, c.Height, red)
	LeftTriangle(c, blue, 0.5)
}

func drawDjibouti(c *Canvas) {
	blue := RGB{106, 178, 231}
	green := RGB{18, 173, 43}
	c.FillRect(0, 0, c.Width, c.Height/2, blue)
	c.FillRect(0, c.Height/2, c.Width, c.Height, green)
	LeftTriangle(c, White, 0.4)
	FiveStar(c, float64(c.Width)*0.13, float64(c.Height)/2, float64(c.Height)*0.1, RGB{210, 16, 52})
}

func drawDominicanRepublic(c *Canvas) {
	blue := RGB{0, 45, 98}
	red := RGB{206, 17, 38}
	// Quartered by a white cross
	c.Fill(White)
	thick := c.Height / 7
	cx := c.Width / 2
	cy := c.Height / 2
	c.FillRect(0, 0, cx-thick/2, cy-thick/2, blue)
	c.FillRect(cx+thick/2, 0, c.Width, cy-thick/2, red)
	c.FillRect(0, cy+thick/2, cx-thick/2, c.Height, red)
	c.FillRect(cx+thick/2, cy+thick/2, c.Width, c.Height, blue)
}

func drawEquatorialGuinea(c *Canvas) {
	green := RGB{58, 119, 40}
	red := RGB{232, 0, 0}
	blue := RGB{0, 114, 198}
	HStripes(c, []RGB{green, White, red})
	LeftTriangle(c, blue, 0.35)
}

func drawEritrea(c *Canvas) {
	blue := RGB{75, 146, 219}
	green := RGB{18, 173, 43}
	red := RGB{234, 0, 41}
	// Upper triangle green, lower blue, red left triangle
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			if fx < fy {
				c.Set(x, y, blue)
			} else {
				c.Set(x, y, green)
			}
		}
	}
	LeftTriangle(c, red, 0.5)
	// Yellow olive wreath (simplified as circle)
	cx := float64(c.Width) * 0.18
	cy := float64(c.Height) / 2
	yellow := RGB{255, 200, 0}
	r := float64(c.Height) * 0.18
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			dx := float64(x) + 0.5 - cx
			dy := float64(y) + 0.5 - cy
			d := math.Sqrt(dx*dx + dy*dy)
			if d <= r && d >= r-1.5 {
				c.Set(x, y, yellow)
			}
		}
	}
}

func drawFiji(c *Canvas) {
	blue := RGB{104, 183, 223}
	c.Fill(blue)
	// Union Jack canton
	cw := c.Width * 2 / 5
	ch := c.Height / 2
	ujCanvas := NewCanvas(cw, ch)
	UnionJack(ujCanvas)
	for y := 0; y < ch; y++ {
		for x := 0; x < cw; x++ {
			c.Set(x, y, ujCanvas.Pixels[y][x])
		}
	}
	// Simplified shield in center-right
	sx := c.Width * 3 / 5
	sy := c.Height / 4
	c.FillRect(sx, sy, sx+c.Width/5, sy+c.Height/2, White)
	c.FillRect(sx, sy, sx+c.Width/5, sy+c.Height/8, RGB{200, 16, 46})
}

func drawGambia(c *Canvas) {
	red := RGB{206, 17, 38}
	blue := RGB{0, 38, 100}
	green := RGB{0, 130, 63}
	HStripes(c, []RGB{red, red, blue, green, green})
	// White borders around blue stripe
	y1 := c.Height * 2 / 5
	y2 := c.Height * 3 / 5
	c.FillRect(0, y1, c.Width, y1+1, White)
	c.FillRect(0, y2-1, c.Width, y2, White)
}

func drawGeorgia(c *Canvas) {
	red := RGB{255, 0, 0}
	c.Fill(White)
	thick := c.Height / 5
	CenteredCross(c, red, thick, thick)
	// 4 small crosses in each quadrant
	qw := c.Width / 4
	qh := c.Height / 4
	smallThick := thick / 3
	positions := [][2]int{{qw, qh}, {c.Width - qw, qh}, {qw, c.Height - qh}, {c.Width - qw, c.Height - qh}}
	for _, p := range positions {
		c.FillRect(p[0]-smallThick/2, p[1]-smallThick*2, p[0]+smallThick/2, p[1]+smallThick*2, red)
		c.FillRect(p[0]-smallThick*2, p[1]-smallThick/2, p[0]+smallThick*2, p[1]+smallThick/2, red)
	}
}

func drawGhana(c *Canvas) {
	red := RGB{206, 17, 38}
	gold := RGB{252, 209, 22}
	green := RGB{0, 107, 63}
	HStripes(c, []RGB{red, gold, green})
	FiveStar(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.13, Black)
}

func drawGreece(c *Canvas) {
	blue := RGB{0, 67, 148}
	// 9 horizontal stripes
	for i := 0; i < 9; i++ {
		y1 := c.Height * i / 9
		y2 := c.Height * (i + 1) / 9
		col := blue
		if i%2 == 1 {
			col = White
		}
		c.FillRect(0, y1, c.Width, y2, col)
	}
	// Blue canton with white cross (top-left 5 stripes)
	cw := c.Width * 2 / 5
	ch := c.Height * 5 / 9
	c.FillRect(0, 0, cw, ch, blue)
	cx := cw / 2
	cy := ch / 2
	thick := ch / 5
	c.FillRect(0, cy-thick/2, cw, cy+thick/2, White)
	c.FillRect(cx-thick/2, 0, cx+thick/2, ch, White)
}

func drawGuinea(c *Canvas) {
	VStripes(c, []RGB{{206, 17, 38}, {252, 209, 22}, {0, 130, 63}})
}

func drawGuineaBissau(c *Canvas) {
	red := RGB{206, 17, 38}
	yellow := RGB{252, 209, 22}
	green := RGB{0, 130, 63}
	c.FillRect(0, 0, c.Width, c.Height/2, yellow)
	c.FillRect(0, c.Height/2, c.Width, c.Height, green)
	c.FillRect(0, 0, c.Width*2/5, c.Height, red)
	FiveStar(c, float64(c.Width)*0.2, float64(c.Height)/2, float64(c.Height)*0.12, Black)
	_ = red
}

func drawGuyana(c *Canvas) {
	green := RGB{0, 154, 68}
	yellow := RGB{255, 222, 0}
	red := RGB{206, 17, 38}
	c.Fill(green)
	// Golden arrow
	for y := 0; y < c.Height; y++ {
		fy := float64(y) + 0.5
		mid := float64(c.Height) / 2
		ratio := 1.0 - math.Abs(fy-mid)/mid
		maxX := int(float64(c.Width) * ratio)
		for x := 0; x < maxX; x++ {
			c.Set(x, y, yellow)
		}
	}
	// Red inner arrow (thinner)
	for y := 0; y < c.Height; y++ {
		fy := float64(y) + 0.5
		mid := float64(c.Height) / 2
		ratio := 1.0 - math.Abs(fy-mid)/mid
		maxX := int(float64(c.Width) * ratio * 0.85)
		for x := 0; x < maxX; x++ {
			c.Set(x, y, red)
		}
	}
	// Black and white borders handled by the overlap
	_ = red
}

func drawIceland(c *Canvas) {
	blue := RGB{0, 56, 151}
	red := RGB{215, 40, 40}
	NordicCross(c, blue, red, &White)
}

func drawJamaica(c *Canvas) {
	green := RGB{0, 122, 65}
	gold := RGB{254, 209, 0}
	// Green top and bottom, black left and right, gold saltire
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			top := fy < fx && fy < 1.0-fx
			bottom := fy > fx && fy > 1.0-fx
			if top || bottom {
				c.Set(x, y, green)
			} else {
				c.Set(x, y, Black)
			}
		}
	}
	Saltire(c, gold, 2.5)
}

func drawJordan(c *Canvas) {
	green := RGB{0, 122, 61}
	red := RGB{206, 17, 38}
	HStripes(c, []RGB{Black, White, green})
	LeftTriangle(c, red, 0.35)
	FiveStar(c, float64(c.Width)*0.12, float64(c.Height)/2, float64(c.Height)*0.07, White)
}

func drawKenya(c *Canvas) {
	green := RGB{0, 131, 62}
	red := RGB{187, 0, 0}
	c.Fill(green)
	c.FillRect(0, 0, c.Width, c.Height*2/7, Black)
	c.FillRect(0, c.Height*2/7, c.Width, c.Height*2/7+c.Height/14, White)
	c.FillRect(0, c.Height*2/7+c.Height/14, c.Width, c.Height*5/7-c.Height/14, red)
	c.FillRect(0, c.Height*5/7-c.Height/14, c.Width, c.Height*5/7, White)
	// Simplified Maasai shield
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	FilledCircle(c, cx, cy, float64(c.Height)*0.2, Black)
	FilledCircle(c, cx, cy, float64(c.Height)*0.15, red)
	c.FillRect(int(cx)-1, c.Height/5, int(cx)+1, c.Height*4/5, Black)
}

func drawKuwait(c *Canvas) {
	green := RGB{0, 122, 61}
	red := RGB{206, 17, 38}
	HStripes(c, []RGB{green, White, red})
	// Black trapezoid on left
	for y := 0; y < c.Height; y++ {
		maxX := c.Width / 4
		for x := 0; x < maxX; x++ {
			c.Set(x, y, Black)
		}
	}
	LeftTriangle(c, Black, 0.33)
}

func drawLaos(c *Canvas) {
	red := RGB{206, 17, 38}
	blue := RGB{0, 56, 147}
	HStripes(c, []RGB{red, blue, blue, red})
	FilledCircle(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.2, White)
}

func drawLatvia(c *Canvas) {
	maroon := RGB{155, 36, 35}
	HStripes(c, []RGB{maroon, maroon, White, maroon, maroon})
}

func drawLebanon(c *Canvas) {
	red := RGB{237, 28, 36}
	green := RGB{0, 122, 61}
	HStripes(c, []RGB{red, White, White, White, red})
	// Simplified cedar tree
	cx := c.Width / 2
	mid := c.Height / 2
	c.FillRect(cx-1, mid-4, cx+1, mid+3, RGB{80, 50, 20})
	for dy := -4; dy < 2; dy++ {
		w := (2 - dy) / 2
		if w < 1 {
			w = 1
		}
		for dx := -w; dx <= w; dx++ {
			c.Set(cx+dx, mid+dy, green)
		}
	}
}

func drawLiberia(c *Canvas) {
	red := RGB{191, 10, 48}
	blue := RGB{0, 40, 104}
	// 11 stripes
	for i := 0; i < 11; i++ {
		y1 := c.Height * i / 11
		y2 := c.Height * (i + 1) / 11
		col := red
		if i%2 == 1 {
			col = White
		}
		c.FillRect(0, y1, c.Width, y2, col)
	}
	// Blue canton
	cw := c.Width * 2 / 5
	ch := c.Height * 6 / 11
	c.FillRect(0, 0, cw, ch, blue)
	FiveStar(c, float64(cw)/2, float64(ch)/2, float64(ch)*0.25, White)
}

func drawLibya(c *Canvas) {
	red := RGB{231, 0, 19}
	green := RGB{35, 158, 70}
	HStripes(c, []RGB{red, Black, Black, Black, green})
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	Crescent(c, cx-1, cy, float64(c.Height)*0.18, float64(c.Height)*0.14, 2, White)
	FiveStar(c, cx+float64(c.Height)*0.1, cy, float64(c.Height)*0.06, White)
}

func drawMadagascar(c *Canvas) {
	red := RGB{252, 61, 50}
	green := RGB{0, 127, 63}
	c.FillRect(0, 0, c.Width, c.Height/2, red)
	c.FillRect(0, c.Height/2, c.Width, c.Height, green)
	c.FillRect(0, 0, c.Width/3, c.Height, White)
}

func drawMalaysia(c *Canvas) {
	red := RGB{204, 0, 0}
	blue := RGB{1, 0, 115}
	yellow := RGB{255, 204, 0}
	// 14 alternating red/white stripes
	for i := 0; i < 14; i++ {
		y1 := c.Height * i / 14
		y2 := c.Height * (i + 1) / 14
		col := red
		if i%2 == 1 {
			col = White
		}
		c.FillRect(0, y1, c.Width, y2, col)
	}
	// Blue canton
	cw := c.Width * 2 / 5
	ch := c.Height / 2
	c.FillRect(0, 0, cw, ch, blue)
	cx := float64(cw) * 0.4
	cy := float64(ch) / 2
	Crescent(c, cx, cy, float64(ch)*0.35, float64(ch)*0.28, 2, yellow)
	FiveStar(c, cx+float64(ch)*0.25, cy, float64(ch)*0.15, yellow)
}

func drawMaldives(c *Canvas) {
	red := RGB{210, 16, 52}
	green := RGB{0, 122, 61}
	c.Fill(red)
	// Green rectangle in center
	gx1 := c.Width / 5
	gy1 := c.Height / 5
	c.FillRect(gx1, gy1, c.Width-gx1, c.Height-gy1, green)
	cx := float64(c.Width)/2 - 1
	cy := float64(c.Height) / 2
	Crescent(c, cx, cy, float64(c.Height)*0.25, float64(c.Height)*0.2, 2, White)
}

func drawMalta(c *Canvas) {
	red := RGB{207, 20, 43}
	VStripes(c, []RGB{White, red})
}

func drawMauritania(c *Canvas) {
	green := RGB{0, 106, 77}
	gold := RGB{199, 160, 0}
	red := RGB{206, 17, 38}
	c.Fill(green)
	// Red bands top and bottom
	c.FillRect(0, 0, c.Width, c.Height/8, red)
	c.FillRect(0, c.Height*7/8, c.Width, c.Height, red)
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	Crescent(c, cx, cy-float64(c.Height)*0.05, float64(c.Height)*0.2, float64(c.Height)*0.16, 0, gold)
	FiveStar(c, cx, cy-float64(c.Height)*0.22, float64(c.Height)*0.06, gold)
}

func drawMexico(c *Canvas) {
	green := RGB{0, 104, 71}
	red := RGB{206, 44, 55}
	VStripes(c, []RGB{green, White, red})
}

func drawMicronesia(c *Canvas) {
	blue := RGB{117, 178, 221}
	c.Fill(blue)
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	r := float64(c.Height) * 0.3
	FiveStar(c, cx, cy-r, float64(c.Height)*0.08, White)
	FiveStar(c, cx, cy+r, float64(c.Height)*0.08, White)
	FiveStar(c, cx-r, cy, float64(c.Height)*0.08, White)
	FiveStar(c, cx+r, cy, float64(c.Height)*0.08, White)
}

func drawMorocco(c *Canvas) {
	red := RGB{193, 39, 45}
	green := RGB{0, 98, 51}
	c.Fill(red)
	FiveStar(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.22, green)
}

func drawMozambique(c *Canvas) {
	green := RGB{0, 122, 77}
	yellow := RGB{252, 209, 22}
	red := RGB{206, 17, 38}
	HStripes(c, []RGB{green, Black, yellow})
	// White borders around black stripe
	y1 := c.Height / 3
	y2 := c.Height * 2 / 3
	c.FillRect(0, y1, c.Width, y1+1, White)
	c.FillRect(0, y2-1, c.Width, y2, White)
	LeftTriangle(c, red, 0.35)
	FiveStar(c, float64(c.Width)*0.12, float64(c.Height)/2, float64(c.Height)*0.08, yellow)
}

func drawMyanmar(c *Canvas) {
	yellow := RGB{254, 203, 0}
	green := RGB{52, 178, 51}
	red := RGB{234, 40, 57}
	HStripes(c, []RGB{yellow, green, red})
	FiveStar(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.22, White)
}

func drawNamibia(c *Canvas) {
	blue := RGB{0, 38, 100}
	red := RGB{200, 16, 46}
	green := RGB{0, 105, 55}
	gold := RGB{255, 206, 0}
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			if fy < fx-0.1 {
				c.Set(x, y, blue)
			} else if fy > fx+0.1 {
				c.Set(x, y, green)
			} else if fy < fx {
				c.Set(x, y, White)
			} else if fy > fx {
				c.Set(x, y, White)
			} else {
				c.Set(x, y, red)
			}
		}
	}
	// Red diagonal band
	w := float64(c.Width)
	h := float64(c.Height)
	diag := math.Sqrt(w*w + h*h)
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fx := float64(x) + 0.5
			fy := float64(y) + 0.5
			d := math.Abs(h*fx-w*fy) / diag
			if d < float64(c.Height)*0.1 {
				c.Set(x, y, red)
			} else if d < float64(c.Height)*0.14 {
				c.Set(x, y, White)
			}
		}
	}
	// Sun in top-left
	FilledCircle(c, float64(c.Width)*0.2, float64(c.Height)*0.25, float64(c.Height)*0.12, gold)
}

func drawNauru(c *Canvas) {
	blue := RGB{0, 42, 106}
	gold := RGB{252, 209, 22}
	c.Fill(blue)
	// Yellow stripe
	c.FillRect(0, c.Height/2-1, c.Width, c.Height/2+1, gold)
	// White star below stripe
	FiveStar(c, float64(c.Width)*0.35, float64(c.Height)*0.7, float64(c.Height)*0.12, White)
}

func drawNepal(c *Canvas) {
	bg := RGB{48, 48, 96}
	crimson := RGB{220, 20, 60}
	blue := RGB{0, 56, 147}
	c.Fill(bg)
	h := float64(c.Height)
	w := float64(c.Width) * 0.75
	// Lower pennant
	for y := 0; y < c.Height; y++ {
		fy := float64(y) / h
		maxX := int(w * (1.0 - fy*0.8))
		for x := 0; x < maxX; x++ {
			c.Set(x, y, crimson)
		}
	}
	// Upper pennant
	midY := int(h * 0.55)
	for y := 0; y < midY; y++ {
		fy := float64(y) / float64(midY)
		maxX := int(w * 0.8 * (1.0 - fy*0.8))
		for x := 0; x < maxX; x++ {
			c.Set(x, y, crimson)
		}
	}
	// Blue borders
	for y := 0; y < c.Height; y++ {
		fy := float64(y) / h
		maxX := int(w * (1.0 - fy*0.8))
		for dx := 0; dx < 2 && maxX-1-dx >= 0; dx++ {
			c.Set(maxX-1-dx, y, blue)
		}
	}
	for y := 0; y < midY; y++ {
		fy := float64(y) / float64(midY)
		maxX := int(w * 0.8 * (1.0 - fy*0.8))
		for dx := 0; dx < 2 && maxX-1-dx >= 0; dx++ {
			c.Set(maxX-1-dx, y, blue)
		}
	}
	c.FillRect(0, c.Height-2, int(w), c.Height, blue)
	// Moon symbol (upper pennant)
	FilledCircle(c, w*0.25, h*0.22, h*0.06, White)
	// Sun symbol (lower pennant)
	FilledCircle(c, w*0.25, h*0.65, h*0.08, White)
}

func drawNewZealand(c *Canvas) {
	blue := RGB{0, 0, 107}
	red := RGB{200, 16, 46}
	c.Fill(blue)
	cw := c.Width * 2 / 5
	ch := c.Height / 2
	ujCanvas := NewCanvas(cw, ch)
	UnionJack(ujCanvas)
	for y := 0; y < ch; y++ {
		for x := 0; x < cw; x++ {
			c.Set(x, y, ujCanvas.Pixels[y][x])
		}
	}
	// 4 red stars with white outlines (Southern Cross)
	stars := [][2]float64{{0.75, 0.25}, {0.85, 0.45}, {0.75, 0.65}, {0.65, 0.50}}
	for _, s := range stars {
		px := s[0] * float64(c.Width)
		py := s[1] * float64(c.Height)
		FiveStar(c, px, py, float64(c.Height)*0.06, White)
		FiveStar(c, px, py, float64(c.Height)*0.045, red)
	}
}

func drawNiger(c *Canvas) {
	orange := RGB{227, 126, 0}
	green := RGB{15, 137, 59}
	HStripes(c, []RGB{orange, White, green})
	FilledCircle(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.1, orange)
}

func drawNorthKorea(c *Canvas) {
	blue := RGB{2, 79, 162}
	red := RGB{237, 28, 36}
	// Blue, thin white, Red (wide, with star disc), thin white, Blue
	c.Fill(blue)
	h := c.Height
	c.FillRect(0, h/8, c.Width, h/8+h/20, White)
	c.FillRect(0, h/8+h/20, c.Width, h*7/8-h/20, red)
	c.FillRect(0, h*7/8-h/20, c.Width, h*7/8, White)
	// White circle with red star
	cx := float64(c.Width) * 0.35
	cy := float64(c.Height) / 2
	FilledCircle(c, cx, cy, float64(h)*0.18, White)
	FiveStar(c, cx, cy, float64(h)*0.15, red)
}

func drawNorthMacedonia(c *Canvas) {
	red := RGB{206, 17, 38}
	yellow := RGB{255, 230, 0}
	c.Fill(red)
	// Sun rays from center
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fx := float64(x) + 0.5 - cx
			fy := float64(y) + 0.5 - cy
			angle := math.Atan2(fy, fx*float64(c.Height)/float64(c.Width))
			// 8 rays
			a := math.Mod(angle+math.Pi, math.Pi/4)
			if a < math.Pi/16 || a > math.Pi*3/16 {
				continue
			}
			c.Set(x, y, yellow)
		}
	}
	// Central circle
	FilledCircle(c, cx, cy, float64(c.Height)*0.15, yellow)
	FilledCircle(c, cx, cy, float64(c.Height)*0.1, red)
}

func drawPakistan(c *Canvas) {
	green := RGB{1, 65, 30}
	c.Fill(green)
	// White band on left
	c.FillRect(0, 0, c.Width/4, c.Height, White)
	cx := float64(c.Width) * 0.55
	cy := float64(c.Height) / 2
	Crescent(c, cx-2, cy, float64(c.Height)*0.28, float64(c.Height)*0.22, 3, White)
	FiveStar(c, cx+float64(c.Height)*0.15, cy-float64(c.Height)*0.05, float64(c.Height)*0.08, White)
}

func drawPalau(c *Canvas) {
	blue := RGB{0, 133, 202}
	gold := RGB{255, 222, 0}
	c.Fill(blue)
	FilledCircle(c, float64(c.Width)*0.4, float64(c.Height)/2, float64(c.Height)*0.3, gold)
}

func drawPanama(c *Canvas) {
	blue := RGB{0, 81, 158}
	red := RGB{218, 18, 52}
	// Quartered
	c.Fill(White)
	c.FillRect(c.Width/2, 0, c.Width, c.Height/2, red)
	c.FillRect(0, c.Height/2, c.Width/2, c.Height, blue)
	FiveStar(c, float64(c.Width)/4, float64(c.Height)/4, float64(c.Height)*0.1, blue)
	FiveStar(c, float64(c.Width)*3/4, float64(c.Height)*3/4, float64(c.Height)*0.1, red)
}

func drawPapuaNewGuinea(c *Canvas) {
	red := RGB{206, 17, 38}
	gold := RGB{255, 205, 0}
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			if fy < fx {
				c.Set(x, y, red)
			} else {
				c.Set(x, y, Black)
			}
		}
	}
	// Stars in black portion (Southern Cross)
	FiveStar(c, float64(c.Width)*0.2, float64(c.Height)*0.7, float64(c.Height)*0.06, White)
	FiveStar(c, float64(c.Width)*0.15, float64(c.Height)*0.5, float64(c.Height)*0.06, White)
	FiveStar(c, float64(c.Width)*0.3, float64(c.Height)*0.55, float64(c.Height)*0.06, White)
	FiveStar(c, float64(c.Width)*0.25, float64(c.Height)*0.85, float64(c.Height)*0.06, White)
	// Bird of paradise (simplified) in red portion
	FiveStar(c, float64(c.Width)*0.72, float64(c.Height)*0.3, float64(c.Height)*0.12, gold)
}

func drawPhilippines(c *Canvas) {
	blue := RGB{0, 56, 168}
	red := RGB{206, 17, 38}
	gold := RGB{252, 209, 22}
	c.FillRect(0, 0, c.Width, c.Height/2, blue)
	c.FillRect(0, c.Height/2, c.Width, c.Height, red)
	LeftTriangle(c, White, 0.45)
	// Sun
	cx := float64(c.Width) * 0.15
	cy := float64(c.Height) / 2
	FilledCircle(c, cx, cy, float64(c.Height)*0.08, gold)
	// 3 stars
	FiveStar(c, float64(c.Width)*0.06, float64(c.Height)*0.15, float64(c.Height)*0.06, gold)
	FiveStar(c, float64(c.Width)*0.06, float64(c.Height)*0.85, float64(c.Height)*0.06, gold)
	FiveStar(c, float64(c.Width)*0.28, float64(c.Height)/2, float64(c.Height)*0.06, gold)
}

func drawPortugal(c *Canvas) {
	green := RGB{0, 102, 0}
	red := RGB{255, 0, 0}
	yellow := RGB{255, 233, 0}
	c.FillRect(0, 0, c.Width*2/5, c.Height, green)
	c.FillRect(c.Width*2/5, 0, c.Width, c.Height, red)
	// Armillary sphere (simplified as yellow circle on the boundary)
	cx := float64(c.Width) * 2 / 5
	cy := float64(c.Height) / 2
	FilledCircle(c, cx, cy, float64(c.Height)*0.2, yellow)
	FilledCircle(c, cx, cy, float64(c.Height)*0.14, red)
	FilledCircle(c, cx, cy, float64(c.Height)*0.1, White)
}

func drawQatar(c *Canvas) {
	maroon := RGB{138, 21, 56}
	c.Fill(maroon)
	teeth := 9
	zigW := c.Width / 3
	toothH := c.Height / teeth
	if toothH < 2 {
		toothH = 2
	}
	for y := 0; y < c.Height; y++ {
		pos := y % toothH
		half := toothH / 2
		var offset int
		if pos < half {
			offset = (zigW / 3) * pos / half
		} else {
			offset = (zigW / 3) * (toothH - pos) / half
		}
		for x := 0; x < zigW-offset; x++ {
			c.Set(x, y, White)
		}
	}
}

func drawSaintKitts(c *Canvas) {
	green := RGB{0, 154, 68}
	red := RGB{206, 17, 38}
	yellow := RGB{252, 209, 22}
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			if fy < 0.5-fx*0.5+0.15 && fy > 0.5-fx*0.5-0.15 {
				c.Set(x, y, Black)
			} else if fy < 0.5-fx*0.5+0.18 && fy > 0.5-fx*0.5-0.18 {
				c.Set(x, y, yellow)
			} else if fy < 0.5 - fx*0.5 {
				c.Set(x, y, green)
			} else {
				c.Set(x, y, red)
			}
		}
	}
	FiveStar(c, float64(c.Width)*0.3, float64(c.Height)*0.4, float64(c.Height)*0.1, White)
	FiveStar(c, float64(c.Width)*0.55, float64(c.Height)*0.6, float64(c.Height)*0.1, White)
}

func drawSaintLucia(c *Canvas) {
	blue := RGB{101, 179, 231}
	gold := RGB{252, 209, 22}
	c.Fill(blue)
	// Triangle shape
	cx := c.Width / 2
	for y := c.Height / 5; y < c.Height; y++ {
		fy := float64(y-c.Height/5) / float64(c.Height*4/5)
		halfW := int(fy * float64(c.Width) * 0.3)
		for x := cx - halfW; x <= cx+halfW; x++ {
			c.Set(x, y, White)
		}
	}
	for y := c.Height * 2 / 5; y < c.Height; y++ {
		fy := float64(y-c.Height*2/5) / float64(c.Height*3/5)
		halfW := int(fy * float64(c.Width) * 0.22)
		for x := cx - halfW; x <= cx+halfW; x++ {
			c.Set(x, y, Black)
		}
	}
	_ = gold
	// Yellow base
	c.FillRect(cx-c.Width*3/10, c.Height*7/10, cx+c.Width*3/10, c.Height*9/10, gold)
}

func drawSaintVincent(c *Canvas) {
	blue := RGB{0, 71, 171}
	green := RGB{0, 154, 68}
	gold := RGB{252, 209, 22}
	VStripes(c, []RGB{blue, gold, gold, green})
	// 3 green diamonds in center
	cx := float64(c.Width) / 2
	h := float64(c.Height)
	for i := 0; i < 3; i++ {
		dy := h*0.25 + float64(i)*h*0.25
		Diamond(c, green, 0.07, 0.15)
		_ = dy
	}
	// Actually draw 3 small V-shaped diamonds
	for _, yf := range []float64{0.3, 0.5, 0.7} {
		cy := h * yf
		r := h * 0.08
		for y := int(cy - r); y <= int(cy+r); y++ {
			for x := int(cx - r); x <= int(cx+r); x++ {
				dx := math.Abs(float64(x)+0.5-cx) / r
				dy := math.Abs(float64(y)+0.5-cy) / r
				if dx+dy <= 1.0 {
					c.Set(x, y, green)
				}
			}
		}
	}
}

func drawSamoa(c *Canvas) {
	red := RGB{206, 17, 38}
	blue := RGB{0, 36, 125}
	c.Fill(red)
	// Blue canton
	cw := c.Width / 2
	ch := c.Height * 3 / 5
	c.FillRect(0, 0, cw, ch, blue)
	// Southern Cross stars
	FiveStar(c, float64(cw)*0.5, float64(ch)*0.25, float64(ch)*0.1, White)
	FiveStar(c, float64(cw)*0.7, float64(ch)*0.4, float64(ch)*0.1, White)
	FiveStar(c, float64(cw)*0.55, float64(ch)*0.6, float64(ch)*0.1, White)
	FiveStar(c, float64(cw)*0.35, float64(ch)*0.7, float64(ch)*0.08, White)
	FiveStar(c, float64(cw)*0.6, float64(ch)*0.8, float64(ch)*0.06, White)
}

func drawSaoTome(c *Canvas) {
	green := RGB{18, 173, 43}
	yellow := RGB{252, 209, 22}
	c.FillRect(0, 0, c.Width, c.Height/3, green)
	c.FillRect(0, c.Height/3, c.Width, c.Height*2/3, yellow)
	c.FillRect(0, c.Height*2/3, c.Width, c.Height, green)
	LeftTriangle(c, RGB{210, 16, 52}, 0.35)
	// Two black stars
	FiveStar(c, float64(c.Width)*0.55, float64(c.Height)/2, float64(c.Height)*0.08, Black)
	FiveStar(c, float64(c.Width)*0.72, float64(c.Height)/2, float64(c.Height)*0.08, Black)
}

func drawSaudiArabia(c *Canvas) {
	green := RGB{0, 106, 78}
	c.Fill(green)
	// White shahada text (simplified as a white band)
	cx := c.Width / 2
	cy := c.Height * 2 / 5
	c.FillRect(cx-c.Width/4, cy-2, cx+c.Width/4, cy+2, White)
	// Sword below
	c.FillRect(cx-c.Width/5, cy+4, cx+c.Width/5, cy+5, White)
}

func drawSenegal(c *Canvas) {
	green := RGB{0, 130, 63}
	gold := RGB{252, 209, 22}
	red := RGB{206, 17, 38}
	VStripes(c, []RGB{green, gold, red})
	FiveStar(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.15, green)
}

func drawSeychelles(c *Canvas) {
	blue := RGB{0, 63, 135}
	yellow := RGB{252, 209, 22}
	red := RGB{210, 16, 52}
	green := RGB{0, 122, 61}
	// Radiating bands from bottom-left
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			angle := math.Atan2(float64(c.Height-1-y), float64(x))
			deg := angle * 180 / math.Pi
			switch {
			case deg > 72:
				c.Set(x, y, blue)
			case deg > 54:
				c.Set(x, y, yellow)
			case deg > 36:
				c.Set(x, y, red)
			case deg > 18:
				c.Set(x, y, White)
			default:
				c.Set(x, y, green)
			}
		}
	}
}

func drawSingapore(c *Canvas) {
	red := RGB{237, 28, 36}
	c.FillRect(0, 0, c.Width, c.Height/2, red)
	c.FillRect(0, c.Height/2, c.Width, c.Height, White)
	cx := float64(c.Width) * 0.3
	cy := float64(c.Height) * 0.25
	Crescent(c, cx-1, cy, float64(c.Height)*0.15, float64(c.Height)*0.12, 1.5, White)
	// 5 stars in pentagon
	r := float64(c.Height) * 0.1
	for i := 0; i < 5; i++ {
		angle := -math.Pi/2 + float64(i)*2*math.Pi/5
		sx := cx + float64(c.Height)*0.08 + r*math.Cos(angle)
		sy := cy + r*math.Sin(angle)
		FiveStar(c, sx, sy, float64(c.Height)*0.03, White)
	}
}

func drawSolomonIslands(c *Canvas) {
	blue := RGB{0, 47, 135}
	green := RGB{0, 119, 73}
	yellow := RGB{252, 209, 22}
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			if fy < fx {
				c.Set(x, y, blue)
			} else {
				c.Set(x, y, green)
			}
		}
	}
	// Yellow diagonal stripe
	w := float64(c.Width)
	h := float64(c.Height)
	diag := math.Sqrt(w*w + h*h)
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fx := float64(x) + 0.5
			fy := float64(y) + 0.5
			d := math.Abs(h*fx-w*fy) / diag
			if d < 2 {
				c.Set(x, y, yellow)
			}
		}
	}
	// Stars
	for i := 0; i < 5; i++ {
		FiveStar(c, float64(c.Width)*0.1+float64(i)*float64(c.Width)*0.08, float64(c.Height)*0.15, float64(c.Height)*0.06, White)
	}
}

func drawSomalia(c *Canvas) {
	blue := RGB{65, 137, 221}
	c.Fill(blue)
	FiveStar(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.25, White)
}

func drawSouthAfrica(c *Canvas) {
	red := RGB{222, 56, 49}
	blue := RGB{0, 20, 137}
	green := RGB{0, 119, 73}
	gold := RGB{255, 184, 28}
	// Red top, blue bottom
	c.FillRect(0, 0, c.Width, c.Height/2, red)
	c.FillRect(0, c.Height/2, c.Width, c.Height, blue)
	// White borders
	c.FillRect(0, c.Height/2-c.Height/12, c.Width, c.Height/2+c.Height/12, White)
	// Green Y shape
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			// Y branches from left to center
			mid := 0.5
			dist := math.Abs(fy - mid)
			// Left portion: triangle merging
			if fx < 0.4 {
				armHalf := 0.4 - fx*0.5
				if dist < armHalf*0.4 + 0.05 {
					c.Set(x, y, green)
				}
			}
			// Horizontal bar to the right
			if fx >= 0.3 && dist < 0.08 {
				c.Set(x, y, green)
			}
		}
	}
	// Gold borders of Y
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			mid := 0.5
			dist := math.Abs(fy - mid)
			if fx < 0.4 {
				armHalf := 0.4 - fx*0.5
				threshold := armHalf*0.4 + 0.05
				if dist >= threshold && dist < threshold+0.025 {
					c.Set(x, y, gold)
				}
			}
			if fx >= 0.3 {
				if dist >= 0.08 && dist < 0.1 {
					c.Set(x, y, gold)
				}
			}
		}
	}
	// Black triangle on left
	for y := 0; y < c.Height; y++ {
		fy := float64(y) / float64(c.Height)
		mid := 0.5
		dist := math.Abs(fy - mid)
		maxFx := 0.33 * (1 - dist/0.5)
		maxX := int(maxFx * float64(c.Width))
		for x := 0; x < maxX; x++ {
			c.Set(x, y, Black)
		}
	}
	// Gold border around black triangle
	for y := 0; y < c.Height; y++ {
		fy := float64(y) / float64(c.Height)
		mid := 0.5
		dist := math.Abs(fy - mid)
		maxFx := 0.33 * (1 - dist/0.5)
		maxX := int(maxFx * float64(c.Width))
		if maxX > 0 && maxX < c.Width {
			c.Set(maxX, y, gold)
			if maxX+1 < c.Width {
				c.Set(maxX+1, y, gold)
			}
		}
	}
}

func drawSouthKorea(c *Canvas) {
	red := RGB{205, 46, 58}
	blue := RGB{0, 71, 160}
	c.Fill(White)
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	r := float64(c.Height) * 0.28
	// Taeguk (yin-yang) - simplified
	FilledCircle(c, cx, cy, r, red)
	// Blue bottom half
	for y := int(cy); y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			dx := float64(x) + 0.5 - cx
			dy := float64(y) + 0.5 - cy
			if dx*dx+dy*dy <= r*r {
				c.Set(x, y, blue)
			}
		}
	}
	// Small circles for yin-yang effect
	FilledCircle(c, cx, cy-r/2, r/2, red)
	FilledCircle(c, cx, cy+r/2, r/2, blue)
	// Trigrams (simplified as black bars in corners)
	barW := 2
	barLen := int(r * 0.5)
	// Top-left trigram
	for i := 0; i < 3; i++ {
		bx := c.Width/6 - barLen/2
		by := c.Height/6 + i*3
		c.FillRect(bx, by, bx+barLen, by+barW, Black)
	}
	// Bottom-right trigram
	for i := 0; i < 3; i++ {
		bx := c.Width*5/6 - barLen/2
		by := c.Height*5/6 - 9 + i*3
		c.FillRect(bx, by, bx+barLen, by+barW, Black)
	}
}

func drawSouthSudan(c *Canvas) {
	blue := RGB{0, 114, 198}
	red := RGB{206, 17, 38}
	green := RGB{0, 105, 55}
	gold := RGB{252, 209, 22}
	c.Fill(green)
	c.FillRect(0, 0, c.Width, c.Height/3, Black)
	c.FillRect(0, c.Height/3, c.Width, c.Height/3+c.Height/20, White)
	c.FillRect(0, c.Height/3+c.Height/20, c.Width, c.Height*2/3-c.Height/20, red)
	c.FillRect(0, c.Height*2/3-c.Height/20, c.Width, c.Height*2/3, White)
	LeftTriangle(c, blue, 0.35)
	FiveStar(c, float64(c.Width)*0.12, float64(c.Height)/2, float64(c.Height)*0.1, gold)
}

func drawSriLanka(c *Canvas) {
	maroon := RGB{128, 0, 32}
	gold := RGB{255, 190, 0}
	green := RGB{0, 114, 47}
	orange := RGB{255, 128, 0}
	c.Fill(gold)
	// Green and orange bands on left
	c.FillRect(c.Width/12, 0, c.Width/12+c.Width/8, c.Height, orange)
	c.FillRect(c.Width/12+c.Width/8, 0, c.Width/12+c.Width/4, c.Height, green)
	// Maroon rectangle on right
	c.FillRect(c.Width*2/5, 0, c.Width-c.Width/12, c.Height, maroon)
	// Simplified lion
	cx := c.Width*2/5 + (c.Width-c.Width/12-c.Width*2/5)/2
	cy := c.Height / 2
	c.FillRect(cx-3, cy-4, cx+3, cy+4, gold)
}

func drawSudan(c *Canvas) {
	red := RGB{210, 16, 52}
	green := RGB{0, 114, 41}
	HStripes(c, []RGB{red, White, Black})
	LeftTriangle(c, green, 0.35)
}

func drawSuriname(c *Canvas) {
	green := RGB{55, 124, 45}
	red := RGB{183, 28, 28}
	gold := RGB{245, 201, 62}
	HStripes(c, []RGB{green, White, red, red, White, green})
	FiveStar(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.15, gold)
}

func drawTanzania(c *Canvas) {
	green := RGB{30, 181, 58}
	blue := RGB{0, 155, 255}
	gold := RGB{252, 204, 0}
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			if fy < fx-0.12 {
				c.Set(x, y, blue)
			} else if fy > fx+0.12 {
				c.Set(x, y, green)
			} else if fy < fx {
				c.Set(x, y, gold)
			} else if fy > fx {
				c.Set(x, y, gold)
			} else {
				c.Set(x, y, Black)
			}
		}
	}
	// Black diagonal band
	w := float64(c.Width)
	h := float64(c.Height)
	diag := math.Sqrt(w*w + h*h)
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fx := float64(x) + 0.5
			fy := float64(y) + 0.5
			d := math.Abs(h*fx-w*fy) / diag
			if d < float64(c.Height)*0.08 {
				c.Set(x, y, Black)
			} else if d < float64(c.Height)*0.12 {
				c.Set(x, y, gold)
			}
		}
	}
}

func drawThailand(c *Canvas) {
	red := RGB{168, 17, 50}
	blue := RGB{45, 77, 139}
	HStripes(c, []RGB{red, White, blue, blue, White, red})
}

func drawTimorLeste(c *Canvas) {
	red := RGB{220, 0, 0}
	yellow := RGB{255, 200, 0}
	c.Fill(red)
	// Yellow triangle
	for y := 0; y < c.Height; y++ {
		fy := float64(y) + 0.5
		mid := float64(c.Height) / 2
		ratio := 1.0 - math.Abs(fy-mid)/mid
		maxX := int(float64(c.Width) * 0.5 * ratio)
		for x := 0; x < maxX; x++ {
			c.Set(x, y, yellow)
		}
	}
	// Black triangle (smaller)
	for y := 0; y < c.Height; y++ {
		fy := float64(y) + 0.5
		mid := float64(c.Height) / 2
		ratio := 1.0 - math.Abs(fy-mid)/mid
		maxX := int(float64(c.Width) * 0.3 * ratio)
		for x := 0; x < maxX; x++ {
			c.Set(x, y, Black)
		}
	}
	FiveStar(c, float64(c.Width)*0.1, float64(c.Height)/2, float64(c.Height)*0.1, White)
}

func drawTogo(c *Canvas) {
	green := RGB{0, 106, 78}
	yellow := RGB{255, 206, 0}
	red := RGB{206, 17, 38}
	// 5 horizontal stripes (alternating green and yellow)
	HStripes(c, []RGB{green, yellow, green, yellow, green})
	// Red canton
	cw := c.Width * 2 / 5
	ch := c.Height * 3 / 5
	c.FillRect(0, 0, cw, ch, red)
	FiveStar(c, float64(cw)/2, float64(ch)/2, float64(ch)*0.25, White)
}

func drawTonga(c *Canvas) {
	red := RGB{200, 16, 46}
	c.Fill(red)
	// White canton with red cross
	cw := c.Width * 2 / 5
	ch := c.Height * 3 / 5
	c.FillRect(0, 0, cw, ch, White)
	thick := ch / 4
	c.FillRect(0, ch/2-thick/2, cw, ch/2+thick/2, red)
	c.FillRect(cw/2-thick/2, 0, cw/2+thick/2, ch, red)
}

func drawTrinidadAndTobago(c *Canvas) {
	red := RGB{218, 18, 26}
	c.Fill(red)
	// Black diagonal band with white borders
	w := float64(c.Width)
	h := float64(c.Height)
	diag := math.Sqrt(w*w + h*h)
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fx := float64(x) + 0.5
			fy := float64(y) + 0.5
			d := math.Abs(h*fx+w*fy-w*h) / diag
			if d < float64(c.Height)*0.12 {
				c.Set(x, y, Black)
			} else if d < float64(c.Height)*0.16 {
				c.Set(x, y, White)
			}
		}
	}
}

func drawTunisia(c *Canvas) {
	red := RGB{231, 0, 19}
	c.Fill(red)
	cx := float64(c.Width) / 2
	cy := float64(c.Height) / 2
	FilledCircle(c, cx, cy, float64(c.Height)*0.25, White)
	Crescent(c, cx-1, cy, float64(c.Height)*0.18, float64(c.Height)*0.15, 2, red)
	FiveStar(c, cx+float64(c.Height)*0.08, cy, float64(c.Height)*0.07, red)
}

func drawTurkey(c *Canvas) {
	red := RGB{227, 10, 23}
	c.Fill(red)
	cx := float64(c.Width)*0.4
	cy := float64(c.Height) / 2
	Crescent(c, cx, cy, float64(c.Height)*0.25, float64(c.Height)*0.2, 3, White)
	FiveStar(c, cx+float64(c.Height)*0.22, cy, float64(c.Height)*0.08, White)
}

func drawTuvalu(c *Canvas) {
	blue := RGB{0, 114, 198}
	c.Fill(blue)
	cw := c.Width * 2 / 5
	ch := c.Height / 2
	ujCanvas := NewCanvas(cw, ch)
	UnionJack(ujCanvas)
	for y := 0; y < ch; y++ {
		for x := 0; x < cw; x++ {
			c.Set(x, y, ujCanvas.Pixels[y][x])
		}
	}
	// Stars
	stars := [][2]float64{
		{0.7, 0.25}, {0.85, 0.35}, {0.75, 0.5}, {0.6, 0.55},
		{0.85, 0.6}, {0.75, 0.7}, {0.55, 0.75}, {0.7, 0.85}, {0.85, 0.85},
	}
	for _, s := range stars {
		FiveStar(c, s[0]*float64(c.Width), s[1]*float64(c.Height), float64(c.Height)*0.04, White)
	}
}

func drawUAE(c *Canvas) {
	green := RGB{0, 114, 41}
	red := RGB{255, 0, 0}
	HStripes(c, []RGB{green, White, Black})
	c.FillRect(0, 0, c.Width/4, c.Height, red)
}

func drawUganda(c *Canvas) {
	yellow := RGB{252, 209, 22}
	red := RGB{206, 17, 38}
	HStripes(c, []RGB{Black, yellow, red, Black, yellow, red})
	// White circle in center
	FilledCircle(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.15, White)
}

func drawUruguay(c *Canvas) {
	blue := RGB{0, 56, 147}
	gold := RGB{248, 195, 0}
	// 9 alternating white/blue stripes
	for i := 0; i < 9; i++ {
		y1 := c.Height * i / 9
		y2 := c.Height * (i + 1) / 9
		col := White
		if i%2 == 1 {
			col = blue
		}
		c.FillRect(0, y1, c.Width, y2, col)
	}
	// Gold sun canton
	cw := c.Width / 3
	ch := c.Height * 5 / 9
	c.FillRect(0, 0, cw, ch, White)
	FilledCircle(c, float64(cw)/2, float64(ch)/2, float64(ch)*0.25, gold)
}

func drawUzbekistan(c *Canvas) {
	blue := RGB{30, 144, 255}
	green := RGB{30, 181, 58}
	red := RGB{206, 17, 38}
	c.Fill(green)
	c.FillRect(0, 0, c.Width, c.Height/3, blue)
	// Red thin borders
	c.FillRect(0, c.Height/3, c.Width, c.Height/3+1, red)
	c.FillRect(0, c.Height/3+1, c.Width, c.Height/3+2, White)
	c.FillRect(0, c.Height*2/3-2, c.Width, c.Height*2/3-1, White)
	c.FillRect(0, c.Height*2/3-1, c.Width, c.Height*2/3, red)
	// White crescent and stars in blue
	cx := float64(c.Width) * 0.2
	cy := float64(c.Height) / 6
	Crescent(c, cx, cy, float64(c.Height)*0.1, float64(c.Height)*0.08, 1, White)
	// 12 stars
	for row := 0; row < 3; row++ {
		for col := 0; col < 4-row; col++ {
			sx := float64(c.Width)*0.3 + float64(col)*float64(c.Height)*0.06
			sy := float64(c.Height)*0.06 + float64(row)*float64(c.Height)*0.06
			FiveStar(c, sx, sy, float64(c.Height)*0.02, White)
		}
	}
}

func drawVanuatu(c *Canvas) {
	red := RGB{210, 16, 52}
	green := RGB{0, 135, 81}
	gold := RGB{253, 205, 18}
	c.FillRect(0, 0, c.Width, c.Height/2, red)
	c.FillRect(0, c.Height/2, c.Width, c.Height, green)
	// Black Y shape
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			mid := 0.5
			dist := math.Abs(fy - mid)
			if fx < 0.35 {
				armHalf := 0.35 - fx*0.5
				if dist < armHalf*0.5 {
					c.Set(x, y, Black)
				}
			}
			if fx >= 0.25 && dist < 0.06 {
				c.Set(x, y, Black)
			}
		}
	}
	// Gold borders
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			fy := float64(y) / float64(c.Height)
			fx := float64(x) / float64(c.Width)
			mid := 0.5
			dist := math.Abs(fy - mid)
			if fx < 0.35 {
				armHalf := 0.35 - fx*0.5
				threshold := armHalf * 0.5
				if dist >= threshold && dist < threshold+0.025 {
					c.Set(x, y, gold)
				}
			}
			if fx >= 0.25 {
				if dist >= 0.06 && dist < 0.08 {
					c.Set(x, y, gold)
				}
			}
		}
	}
	// Black triangle on left
	for y := 0; y < c.Height; y++ {
		fy := float64(y) / float64(c.Height)
		mid := 0.5
		dist := math.Abs(fy - mid)
		maxFx := 0.25 * (1 - dist/0.5)
		for x := 0; x < int(maxFx*float64(c.Width)); x++ {
			c.Set(x, y, Black)
		}
	}
}

func drawVietnam(c *Canvas) {
	red := RGB{218, 37, 29}
	yellow := RGB{255, 255, 0}
	c.Fill(red)
	FiveStar(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.25, yellow)
}

func drawZambia(c *Canvas) {
	green := RGB{25, 137, 67}
	orange := RGB{222, 129, 35}
	red := RGB{222, 32, 43}
	c.Fill(green)
	// Three vertical stripes in top-right corner
	x1 := c.Width * 3 / 5
	sw := (c.Width - x1) / 3
	c.FillRect(x1, c.Height/3, x1+sw, c.Height, red)
	c.FillRect(x1+sw, c.Height/3, x1+sw*2, c.Height, Black)
	c.FillRect(x1+sw*2, c.Height/3, c.Width, c.Height, orange)
}

func drawZimbabwe(c *Canvas) {
	green := RGB{0, 114, 41}
	yellow := RGB{252, 209, 22}
	red := RGB{222, 32, 43}
	// 7 stripes
	HStripes(c, []RGB{green, yellow, red, Black, red, yellow, green})
	// White triangle
	LeftTriangle(c, White, 0.35)
	// Red star in triangle
	FiveStar(c, float64(c.Width)*0.12, float64(c.Height)/2, float64(c.Height)*0.1, red)
}

package main

import "strings"

type Flag struct {
	Name    string
	Aliases []string
	Ratio   [2]int // height:width
	Draw    func(c *Canvas)
}

var Flags = []Flag{
	// A
	{Name: "afghanistan", Aliases: []string{"af"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{{0, 0, 0}, {175, 34, 47}, {0, 122, 54}})
	}},
	{Name: "albania", Aliases: []string{"al"}, Ratio: [2]int{5, 7}, Draw: func(c *Canvas) {
		c.Fill(RGB{226, 6, 19})
		// Simplified double-headed eagle as a dark shape in center
		FilledCircle(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.2, Black)
	}},
	{Name: "algeria", Aliases: []string{"dz"}, Ratio: [2]int{2, 3}, Draw: drawAlgeria},
	{Name: "andorra", Aliases: []string{"ad"}, Ratio: [2]int{7, 10}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{{0, 56, 168}, {252, 209, 22}, {210, 16, 52}})
	}},
	{Name: "angola", Aliases: []string{"ao"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{204, 9, 47}, {0, 0, 0}})
		// Machete and gear (simplified as yellow arc)
		cx := float64(c.Width) / 2
		cy := float64(c.Height) / 2
		FiveStar(c, cx, cy, float64(c.Height)*0.12, RGB{252, 209, 22})
	}},
	{Name: "antigua-and-barbuda", Aliases: []string{"ag", "antigua"}, Ratio: [2]int{2, 3}, Draw: drawAntiguaAndBarbuda},
	{Name: "argentina", Aliases: []string{"ar", "arg"}, Ratio: [2]int{5, 8}, Draw: func(c *Canvas) {
		blue := RGB{116, 172, 223}
		HStripes(c, []RGB{blue, White, blue})
	}},
	{Name: "armenia", Aliases: []string{"am"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{217, 0, 18}, {0, 51, 160}, {242, 168, 0}})
	}},
	{Name: "australia", Aliases: []string{"au", "aus"}, Ratio: [2]int{1, 2}, Draw: drawAustralia},
	{Name: "austria", Aliases: []string{"at", "aut"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{237, 41, 57}, White, {237, 41, 57}})
	}},
	{Name: "azerbaijan", Aliases: []string{"az"}, Ratio: [2]int{1, 2}, Draw: drawAzerbaijan},
	// B
	{Name: "bahamas", Aliases: []string{"bs"}, Ratio: [2]int{1, 2}, Draw: drawBahamas},
	{Name: "bahrain", Aliases: []string{"bh"}, Ratio: [2]int{3, 5}, Draw: drawBahrain},
	{Name: "bangladesh", Aliases: []string{"bd"}, Ratio: [2]int{3, 5}, Draw: drawBangladesh},
	{Name: "barbados", Aliases: []string{"bb"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		blue := RGB{0, 38, 100}
		gold := RGB{252, 187, 24}
		VStripes(c, []RGB{blue, gold, blue})
	}},
	{Name: "belarus", Aliases: []string{"by"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		red := RGB{204, 9, 47}
		green := RGB{0, 122, 61}
		c.FillRect(0, 0, c.Width, c.Height*2/3, red)
		c.FillRect(0, c.Height*2/3, c.Width, c.Height, green)
		// Red ornament band on left
		c.FillRect(0, 0, c.Width/8, c.Height, White)
		for y := 0; y < c.Height; y += 4 {
			c.FillRect(c.Width/16-1, y, c.Width/16+1, y+2, red)
		}
	}},
	{Name: "belgium", Aliases: []string{"be", "bel"}, Ratio: [2]int{13, 15}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{Black, {252, 209, 22}, {237, 41, 57}})
	}},
	{Name: "belize", Aliases: []string{"bz"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		blue := RGB{0, 56, 168}
		red := RGB{206, 17, 38}
		c.Fill(blue)
		c.FillRect(0, 0, c.Width, c.Height/10, red)
		c.FillRect(0, c.Height*9/10, c.Width, c.Height, red)
		FilledCircle(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.22, White)
	}},
	{Name: "benin", Aliases: []string{"bj"}, Ratio: [2]int{2, 3}, Draw: drawBenin},
	{Name: "bhutan", Aliases: []string{"bt"}, Ratio: [2]int{2, 3}, Draw: drawBhutan},
	{Name: "bolivia", Aliases: []string{"bo"}, Ratio: [2]int{15, 22}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{211, 47, 47}, {252, 225, 0}, {0, 122, 54}})
	}},
	{Name: "bosnia", Aliases: []string{"ba", "bosnia-and-herzegovina"}, Ratio: [2]int{1, 2}, Draw: drawBosnia},
	{Name: "botswana", Aliases: []string{"bw"}, Ratio: [2]int{2, 3}, Draw: drawBotswana},
	{Name: "brazil", Aliases: []string{"br", "bra"}, Ratio: [2]int{7, 10}, Draw: drawBrazil},
	{Name: "brunei", Aliases: []string{"bn"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		yellow := RGB{255, 222, 0}
		c.Fill(yellow)
		// White and black diagonal stripes
		for y := 0; y < c.Height; y++ {
			for x := 0; x < c.Width; x++ {
				fy := float64(y) / float64(c.Height)
				fx := float64(x) / float64(c.Width)
				if fy > 0.3+fx*0.2 && fy < 0.5+fx*0.2 {
					c.Set(x, y, White)
				}
				if fy > 0.5+fx*0.2 && fy < 0.7+fx*0.2 {
					c.Set(x, y, Black)
				}
			}
		}
	}},
	{Name: "bulgaria", Aliases: []string{"bg"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{White, {0, 150, 110}, {214, 38, 18}})
	}},
	{Name: "burkina-faso", Aliases: []string{"bf", "burkina"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{239, 51, 64}, {0, 146, 63}})
		FiveStar(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.13, RGB{252, 209, 22})
	}},
	{Name: "burundi", Aliases: []string{"bi"}, Ratio: [2]int{3, 5}, Draw: drawBurundi},
	// C
	{Name: "cabo-verde", Aliases: []string{"cv", "cape-verde"}, Ratio: [2]int{2, 3}, Draw: drawCaboVerde},
	{Name: "cambodia", Aliases: []string{"kh"}, Ratio: [2]int{2, 3}, Draw: drawCambodia},
	{Name: "cameroon", Aliases: []string{"cm"}, Ratio: [2]int{2, 3}, Draw: drawCameroon},
	{Name: "canada", Aliases: []string{"ca", "can"}, Ratio: [2]int{1, 2}, Draw: drawCanada},
	{Name: "central-african-republic", Aliases: []string{"cf", "car"}, Ratio: [2]int{2, 3}, Draw: drawCAR},
	{Name: "chad", Aliases: []string{"td"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{{0, 38, 100}, {252, 209, 22}, {206, 17, 38}})
	}},
	{Name: "chile", Aliases: []string{"cl"}, Ratio: [2]int{2, 3}, Draw: drawChile},
	{Name: "china", Aliases: []string{"cn", "chn"}, Ratio: [2]int{2, 3}, Draw: drawChina},
	{Name: "colombia", Aliases: []string{"co", "col"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		yellow := RGB{252, 209, 22}
		blue := RGB{0, 56, 147}
		red := RGB{206, 17, 38}
		c.FillRect(0, 0, c.Width, c.Height/2, yellow)
		c.FillRect(0, c.Height/2, c.Width, c.Height*3/4, blue)
		c.FillRect(0, c.Height*3/4, c.Width, c.Height, red)
	}},
	{Name: "comoros", Aliases: []string{"km"}, Ratio: [2]int{3, 5}, Draw: drawComoros},
	{Name: "congo-drc", Aliases: []string{"cd", "drc", "congo-kinshasa"}, Ratio: [2]int{2, 3}, Draw: drawCongoDRC},
	{Name: "congo-republic", Aliases: []string{"cg", "congo-brazzaville"}, Ratio: [2]int{2, 3}, Draw: drawCongoRepublic},
	{Name: "costa-rica", Aliases: []string{"cr"}, Ratio: [2]int{3, 5}, Draw: drawCostaRica},
	{Name: "cote-d-ivoire", Aliases: []string{"ci", "ivory-coast"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{{245, 131, 0}, White, {0, 158, 73}})
	}},
	{Name: "croatia", Aliases: []string{"hr"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{255, 0, 0}, White, {0, 93, 170}})
	}},
	{Name: "cuba", Aliases: []string{"cu"}, Ratio: [2]int{1, 2}, Draw: drawCuba},
	{Name: "cyprus", Aliases: []string{"cy"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		c.Fill(White)
		// Simplified island shape as orange blob in center
		cx := float64(c.Width) / 2
		cy := float64(c.Height) * 0.4
		for y := 0; y < c.Height; y++ {
			for x := 0; x < c.Width; x++ {
				dx := (float64(x)+0.5-cx) / float64(c.Width) * 4
				dy := (float64(y)+0.5-cy) / float64(c.Height) * 4
				if dx*dx+dy*dy*3 < 0.5 {
					c.Set(x, y, RGB{210, 145, 37})
				}
			}
		}
		// Green olive branches below
		c.FillRect(c.Width/3, c.Height*3/5, c.Width*2/3, c.Height*3/5+2, RGB{0, 122, 61})
	}},
	{Name: "czech-republic", Aliases: []string{"cz", "czechia"}, Ratio: [2]int{2, 3}, Draw: drawCzechRepublic},
	// D
	{Name: "denmark", Aliases: []string{"dk", "dan"}, Ratio: [2]int{3, 4}, Draw: func(c *Canvas) {
		NordicCross(c, DenmarkRed, White, nil)
	}},
	{Name: "djibouti", Aliases: []string{"dj"}, Ratio: [2]int{2, 3}, Draw: drawDjibouti},
	{Name: "dominica", Aliases: []string{"dm"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		green := RGB{0, 122, 61}
		c.Fill(green)
		// Cross
		yellow := RGB{252, 209, 22}
		CenteredCross(c, Black, c.Width/15, c.Height/8)
		CenteredCross(c, yellow, c.Width/20, c.Height/12)
		// Center circle
		FilledCircle(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.15, RGB{108, 0, 108})
	}},
	{Name: "dominican-republic", Aliases: []string{"do", "dominican"}, Ratio: [2]int{2, 3}, Draw: drawDominicanRepublic},
	// E
	{Name: "ecuador", Aliases: []string{"ec"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		yellow := RGB{252, 209, 22}
		blue := RGB{0, 56, 147}
		red := RGB{206, 17, 38}
		c.FillRect(0, 0, c.Width, c.Height/2, yellow)
		c.FillRect(0, c.Height/2, c.Width, c.Height*3/4, blue)
		c.FillRect(0, c.Height*3/4, c.Width, c.Height, red)
	}},
	{Name: "egypt", Aliases: []string{"eg"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{206, 17, 38}, White, Black})
	}},
	{Name: "el-salvador", Aliases: []string{"sv"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		blue := RGB{0, 71, 171}
		HStripes(c, []RGB{blue, White, blue})
	}},
	{Name: "equatorial-guinea", Aliases: []string{"gq"}, Ratio: [2]int{2, 3}, Draw: drawEquatorialGuinea},
	{Name: "eritrea", Aliases: []string{"er"}, Ratio: [2]int{1, 2}, Draw: drawEritrea},
	{Name: "estonia", Aliases: []string{"ee", "est"}, Ratio: [2]int{7, 11}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{0, 114, 206}, Black, White})
	}},
	{Name: "eswatini", Aliases: []string{"sz", "swaziland"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		blue := RGB{0, 48, 135}
		yellow := RGB{255, 210, 0}
		red := RGB{206, 17, 38}
		HStripes(c, []RGB{blue, yellow, red, red, red, yellow, blue})
		// Simplified shield in center
		cx := float64(c.Width) / 2
		cy := float64(c.Height) / 2
		FilledCircle(c, cx, cy, float64(c.Height)*0.18, Black)
		FilledCircle(c, cx, cy, float64(c.Height)*0.14, White)
	}},
	{Name: "ethiopia", Aliases: []string{"et"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{0, 155, 58}, {252, 209, 22}, {218, 18, 26}})
		FilledCircle(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.18, RGB{0, 56, 168})
		FiveStar(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.12, RGB{252, 209, 22})
	}},
	// F
	{Name: "fiji", Aliases: []string{"fj"}, Ratio: [2]int{1, 2}, Draw: drawFiji},
	{Name: "finland", Aliases: []string{"fi", "fin"}, Ratio: [2]int{11, 18}, Draw: func(c *Canvas) {
		NordicCross(c, White, FinlandBlue, nil)
	}},
	{Name: "france", Aliases: []string{"fr", "fra"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{FrenchBlue, White, FrenchRed})
	}},
	// G
	{Name: "gabon", Aliases: []string{"ga"}, Ratio: [2]int{3, 4}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{0, 155, 58}, {252, 209, 22}, {0, 114, 198}})
	}},
	{Name: "gambia", Aliases: []string{"gm"}, Ratio: [2]int{2, 3}, Draw: drawGambia},
	{Name: "georgia", Aliases: []string{"ge"}, Ratio: [2]int{2, 3}, Draw: drawGeorgia},
	{Name: "germany", Aliases: []string{"de", "ger"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{Black, GermanRed, GermanGold})
	}},
	{Name: "ghana", Aliases: []string{"gh"}, Ratio: [2]int{2, 3}, Draw: drawGhana},
	{Name: "greece", Aliases: []string{"gr"}, Ratio: [2]int{2, 3}, Draw: drawGreece},
	{Name: "grenada", Aliases: []string{"gd"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		red := RGB{206, 17, 38}
		yellow := RGB{252, 209, 22}
		green := RGB{0, 122, 61}
		c.Fill(red)
		// Green rectangle
		c.FillRect(c.Width/6, c.Height/6, c.Width*5/6, c.Height*5/6, green)
		// Yellow diagonal cross
		Saltire(c, yellow, 2)
		// Center circle
		FilledCircle(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.15, red)
		FiveStar(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.1, yellow)
	}},
	{Name: "guatemala", Aliases: []string{"gt"}, Ratio: [2]int{5, 8}, Draw: func(c *Canvas) {
		blue := RGB{73, 151, 208}
		VStripes(c, []RGB{blue, White, blue})
	}},
	{Name: "guinea", Aliases: []string{"gn"}, Ratio: [2]int{2, 3}, Draw: drawGuinea},
	{Name: "guinea-bissau", Aliases: []string{"gw"}, Ratio: [2]int{1, 2}, Draw: drawGuineaBissau},
	{Name: "guyana", Aliases: []string{"gy"}, Ratio: [2]int{3, 5}, Draw: drawGuyana},
	// H
	{Name: "haiti", Aliases: []string{"ht"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{0, 38, 100}, {210, 16, 52}})
	}},
	{Name: "honduras", Aliases: []string{"hn"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		blue := RGB{0, 101, 189}
		HStripes(c, []RGB{blue, White, blue})
		cx := float64(c.Width) / 2
		cy := float64(c.Height) / 2
		FiveStar(c, cx, cy, float64(c.Height)*0.05, blue)
		FiveStar(c, cx-float64(c.Width)*0.12, cy-float64(c.Height)*0.15, float64(c.Height)*0.04, blue)
		FiveStar(c, cx+float64(c.Width)*0.12, cy-float64(c.Height)*0.15, float64(c.Height)*0.04, blue)
		FiveStar(c, cx-float64(c.Width)*0.12, cy+float64(c.Height)*0.15, float64(c.Height)*0.04, blue)
		FiveStar(c, cx+float64(c.Width)*0.12, cy+float64(c.Height)*0.15, float64(c.Height)*0.04, blue)
	}},
	{Name: "hungary", Aliases: []string{"hu"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{205, 42, 62}, White, {67, 111, 77}})
	}},
	// I
	{Name: "iceland", Aliases: []string{"is", "isl"}, Ratio: [2]int{18, 25}, Draw: drawIceland},
	{Name: "india", Aliases: []string{"in", "ind"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		orange := RGB{255, 153, 51}
		green := RGB{19, 136, 8}
		HStripes(c, []RGB{orange, White, green})
		// Ashoka Chakra (simplified as blue circle)
		cx := float64(c.Width) / 2
		cy := float64(c.Height) / 2
		navy := RGB{0, 0, 128}
		r := float64(c.Height) * 0.1
		for y := 0; y < c.Height; y++ {
			for x := 0; x < c.Width; x++ {
				dx := float64(x) + 0.5 - cx
				dy := float64(y) + 0.5 - cy
				d := dx*dx + dy*dy
				if d <= r*r && d >= (r-1.5)*(r-1.5) {
					c.Set(x, y, navy)
				}
			}
		}
	}},
	{Name: "indonesia", Aliases: []string{"id"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{255, 0, 0}, White})
	}},
	{Name: "iran", Aliases: []string{"ir"}, Ratio: [2]int{4, 7}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{35, 159, 64}, White, {218, 0, 0}})
	}},
	{Name: "iraq", Aliases: []string{"iq"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{206, 17, 38}, White, Black})
	}},
	{Name: "ireland", Aliases: []string{"ie", "irl"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{IrelandGreen, White, IrelandOrange})
	}},
	{Name: "italy", Aliases: []string{"it", "ita"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{ItalyGreen, White, ItalyRed})
	}},
	// J
	{Name: "jamaica", Aliases: []string{"jm"}, Ratio: [2]int{1, 2}, Draw: drawJamaica},
	{Name: "japan", Aliases: []string{"jp", "jpn"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		Circle(c, White, JapanRed, 0.375)
	}},
	{Name: "jordan", Aliases: []string{"jo"}, Ratio: [2]int{1, 2}, Draw: drawJordan},
	// K
	{Name: "kazakhstan", Aliases: []string{"kz"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		blue := RGB{0, 175, 239}
		gold := RGB{255, 199, 44}
		c.Fill(blue)
		// Sun
		FilledCircle(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.15, gold)
		// Gold ornamental band on left
		c.FillRect(c.Width/15, 0, c.Width/15+c.Width/20, c.Height, gold)
	}},
	{Name: "kenya", Aliases: []string{"ke"}, Ratio: [2]int{2, 3}, Draw: drawKenya},
	{Name: "kiribati", Aliases: []string{"ki"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		red := RGB{206, 17, 38}
		blue := RGB{0, 56, 147}
		gold := RGB{252, 209, 22}
		// Red sky top, blue waves bottom
		c.FillRect(0, 0, c.Width, c.Height/2, red)
		c.FillRect(0, c.Height/2, c.Width, c.Height, blue)
		// White wave lines
		for x := 0; x < c.Width; x++ {
			for i := 0; i < 3; i++ {
				y := c.Height/2 + c.Height/8 + i*c.Height/10
				c.Set(x, y, White)
			}
		}
		// Sun
		FilledCircle(c, float64(c.Width)/2, float64(c.Height)*0.35, float64(c.Height)*0.12, gold)
	}},
	{Name: "kuwait", Aliases: []string{"kw"}, Ratio: [2]int{1, 2}, Draw: drawKuwait},
	{Name: "kyrgyzstan", Aliases: []string{"kg"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		red := RGB{232, 0, 0}
		yellow := RGB{255, 215, 0}
		c.Fill(red)
		FilledCircle(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.25, yellow)
		FilledCircle(c, float64(c.Width)/2, float64(c.Height)/2, float64(c.Height)*0.18, red)
	}},
	// L
	{Name: "laos", Aliases: []string{"la"}, Ratio: [2]int{2, 3}, Draw: drawLaos},
	{Name: "latvia", Aliases: []string{"lv"}, Ratio: [2]int{1, 2}, Draw: drawLatvia},
	{Name: "lebanon", Aliases: []string{"lb"}, Ratio: [2]int{2, 3}, Draw: drawLebanon},
	{Name: "lesotho", Aliases: []string{"ls"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{0, 57, 166}, White, {0, 137, 63}})
	}},
	{Name: "liberia", Aliases: []string{"lr"}, Ratio: [2]int{10, 19}, Draw: drawLiberia},
	{Name: "libya", Aliases: []string{"ly"}, Ratio: [2]int{1, 2}, Draw: drawLibya},
	{Name: "liechtenstein", Aliases: []string{"li"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{0, 51, 153}, {206, 17, 38}})
	}},
	{Name: "lithuania", Aliases: []string{"lt"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{253, 185, 19}, {0, 106, 68}, {190, 20, 47}})
	}},
	{Name: "luxembourg", Aliases: []string{"lu"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{237, 41, 57}, White, {0, 161, 222}})
	}},
	// M
	{Name: "madagascar", Aliases: []string{"mg"}, Ratio: [2]int{2, 3}, Draw: drawMadagascar},
	{Name: "malawi", Aliases: []string{"mw"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{Black, {206, 17, 38}, {0, 122, 61}})
		// Rising sun
		FilledCircle(c, float64(c.Width)/2, float64(c.Height)/6, float64(c.Height)*0.12, RGB{206, 17, 38})
	}},
	{Name: "malaysia", Aliases: []string{"my"}, Ratio: [2]int{1, 2}, Draw: drawMalaysia},
	{Name: "maldives", Aliases: []string{"mv"}, Ratio: [2]int{2, 3}, Draw: drawMaldives},
	{Name: "mali", Aliases: []string{"ml"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{{0, 130, 63}, {252, 209, 22}, {206, 17, 38}})
	}},
	{Name: "malta", Aliases: []string{"mt"}, Ratio: [2]int{2, 3}, Draw: drawMalta},
	{Name: "marshall-islands", Aliases: []string{"mh", "marshall"}, Ratio: [2]int{10, 19}, Draw: func(c *Canvas) {
		blue := RGB{0, 56, 147}
		c.Fill(blue)
		// Two diagonal stripes (orange and white)
		for y := 0; y < c.Height; y++ {
			for x := 0; x < c.Width; x++ {
				fy := float64(y) / float64(c.Height)
				fx := float64(x) / float64(c.Width)
				band := fy - fx*0.5
				if band > 0.3 && band < 0.5 {
					c.Set(x, y, White)
				} else if band > 0.5 && band < 0.7 {
					c.Set(x, y, RGB{227, 126, 0})
				}
			}
		}
		FiveStar(c, float64(c.Width)*0.2, float64(c.Height)*0.3, float64(c.Height)*0.15, White)
	}},
	{Name: "mauritania", Aliases: []string{"mr"}, Ratio: [2]int{2, 3}, Draw: drawMauritania},
	{Name: "mauritius", Aliases: []string{"mu"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{235, 28, 46}, {0, 36, 125}, {252, 209, 22}, {0, 122, 61}})
	}},
	{Name: "mexico", Aliases: []string{"mx", "mex"}, Ratio: [2]int{4, 7}, Draw: drawMexico},
	{Name: "micronesia", Aliases: []string{"fm"}, Ratio: [2]int{10, 19}, Draw: drawMicronesia},
	{Name: "moldova", Aliases: []string{"md"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{{0, 91, 167}, {252, 209, 22}, {204, 9, 47}})
	}},
	{Name: "monaco", Aliases: []string{"mc"}, Ratio: [2]int{4, 5}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{206, 17, 38}, White})
	}},
	{Name: "mongolia", Aliases: []string{"mn"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		red := RGB{200, 16, 46}
		blue := RGB{0, 102, 180}
		VStripes(c, []RGB{red, blue, red})
		// Soyombo symbol (simplified as yellow shape)
		cx := c.Width / 6
		cy := c.Height / 2
		FilledCircle(c, float64(cx), float64(cy), float64(c.Height)*0.1, RGB{252, 209, 22})
	}},
	{Name: "montenegro", Aliases: []string{"me"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		red := RGB{200, 16, 46}
		gold := RGB{211, 175, 55}
		c.Fill(red)
		// Gold border
		c.FillRect(0, 0, c.Width, 2, gold)
		c.FillRect(0, c.Height-2, c.Width, c.Height, gold)
		c.FillRect(0, 0, 2, c.Height, gold)
		c.FillRect(c.Width-2, 0, c.Width, c.Height, gold)
	}},
	{Name: "morocco", Aliases: []string{"ma"}, Ratio: [2]int{2, 3}, Draw: drawMorocco},
	{Name: "mozambique", Aliases: []string{"mz"}, Ratio: [2]int{2, 3}, Draw: drawMozambique},
	{Name: "myanmar", Aliases: []string{"mm", "burma"}, Ratio: [2]int{2, 3}, Draw: drawMyanmar},
	// N
	{Name: "namibia", Aliases: []string{"na"}, Ratio: [2]int{2, 3}, Draw: drawNamibia},
	{Name: "nauru", Aliases: []string{"nr"}, Ratio: [2]int{1, 2}, Draw: drawNauru},
	{Name: "nepal", Aliases: []string{"np"}, Ratio: [2]int{4, 3}, Draw: drawNepal},
	{Name: "netherlands", Aliases: []string{"nl", "ned"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{NetherlandsRed, White, NetherlandsBlue})
	}},
	{Name: "new-zealand", Aliases: []string{"nz"}, Ratio: [2]int{1, 2}, Draw: drawNewZealand},
	{Name: "nicaragua", Aliases: []string{"ni"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		blue := RGB{0, 40, 104}
		HStripes(c, []RGB{blue, White, blue})
	}},
	{Name: "niger", Aliases: []string{"ne"}, Ratio: [2]int{6, 7}, Draw: drawNiger},
	{Name: "nigeria", Aliases: []string{"ng"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{{0, 135, 68}, White, {0, 135, 68}})
	}},
	{Name: "north-korea", Aliases: []string{"kp", "dprk"}, Ratio: [2]int{1, 2}, Draw: drawNorthKorea},
	{Name: "north-macedonia", Aliases: []string{"mk", "macedonia"}, Ratio: [2]int{1, 2}, Draw: drawNorthMacedonia},
	{Name: "norway", Aliases: []string{"no", "nor"}, Ratio: [2]int{8, 11}, Draw: func(c *Canvas) {
		NordicCross(c, NorwayRed, NorwayBlue, &White)
	}},
	// O
	{Name: "oman", Aliases: []string{"om"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		red := RGB{219, 0, 17}
		green := RGB{0, 128, 0}
		c.Fill(White)
		c.FillRect(0, 0, c.Width, c.Height/3, White)
		c.FillRect(0, c.Height/3, c.Width, c.Height*2/3, red)
		c.FillRect(0, c.Height*2/3, c.Width, c.Height, green)
		c.FillRect(0, 0, c.Width/4, c.Height, red)
	}},
	// P
	{Name: "pakistan", Aliases: []string{"pk"}, Ratio: [2]int{2, 3}, Draw: drawPakistan},
	{Name: "palau", Aliases: []string{"pw"}, Ratio: [2]int{5, 8}, Draw: drawPalau},
	{Name: "panama", Aliases: []string{"pa"}, Ratio: [2]int{2, 3}, Draw: drawPanama},
	{Name: "papua-new-guinea", Aliases: []string{"pg", "png"}, Ratio: [2]int{3, 4}, Draw: drawPapuaNewGuinea},
	{Name: "paraguay", Aliases: []string{"py"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{237, 41, 57}, White, {0, 56, 168}})
	}},
	{Name: "peru", Aliases: []string{"pe"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{{216, 20, 25}, White, {216, 20, 25}})
	}},
	{Name: "philippines", Aliases: []string{"ph"}, Ratio: [2]int{1, 2}, Draw: drawPhilippines},
	{Name: "poland", Aliases: []string{"pl", "pol"}, Ratio: [2]int{5, 8}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{White, PolandRed})
	}},
	{Name: "portugal", Aliases: []string{"pt"}, Ratio: [2]int{2, 3}, Draw: drawPortugal},
	// Q
	{Name: "qatar", Aliases: []string{"qa"}, Ratio: [2]int{11, 28}, Draw: drawQatar},
	// R
	{Name: "romania", Aliases: []string{"ro"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{{0, 43, 127}, {252, 209, 22}, {206, 17, 38}})
	}},
	{Name: "russia", Aliases: []string{"ru", "rus"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{White, {0, 57, 166}, {213, 43, 30}})
	}},
	{Name: "rwanda", Aliases: []string{"rw"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		blue := RGB{0, 163, 224}
		yellow := RGB{250, 209, 0}
		green := RGB{32, 178, 81}
		c.FillRect(0, 0, c.Width, c.Height/2, blue)
		c.FillRect(0, c.Height/2, c.Width, c.Height*3/4, yellow)
		c.FillRect(0, c.Height*3/4, c.Width, c.Height, green)
		// Sun in top right
		FilledCircle(c, float64(c.Width)*0.75, float64(c.Height)*0.25, float64(c.Height)*0.1, yellow)
	}},
	// S
	{Name: "saint-kitts-and-nevis", Aliases: []string{"kn", "st-kitts"}, Ratio: [2]int{2, 3}, Draw: drawSaintKitts},
	{Name: "saint-lucia", Aliases: []string{"lc"}, Ratio: [2]int{1, 2}, Draw: drawSaintLucia},
	{Name: "saint-vincent", Aliases: []string{"vc", "st-vincent"}, Ratio: [2]int{2, 3}, Draw: drawSaintVincent},
	{Name: "samoa", Aliases: []string{"ws"}, Ratio: [2]int{1, 2}, Draw: drawSamoa},
	{Name: "san-marino", Aliases: []string{"sm"}, Ratio: [2]int{3, 4}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{White, {0, 155, 224}})
	}},
	{Name: "sao-tome-and-principe", Aliases: []string{"st", "sao-tome"}, Ratio: [2]int{1, 2}, Draw: drawSaoTome},
	{Name: "saudi-arabia", Aliases: []string{"sa"}, Ratio: [2]int{2, 3}, Draw: drawSaudiArabia},
	{Name: "senegal", Aliases: []string{"sn"}, Ratio: [2]int{2, 3}, Draw: drawSenegal},
	{Name: "serbia", Aliases: []string{"rs"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{198, 54, 60}, {0, 56, 168}, White})
	}},
	{Name: "seychelles", Aliases: []string{"sc"}, Ratio: [2]int{1, 2}, Draw: drawSeychelles},
	{Name: "sierra-leone", Aliases: []string{"sl"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{27, 177, 98}, White, {0, 114, 198}})
	}},
	{Name: "singapore", Aliases: []string{"sg"}, Ratio: [2]int{2, 3}, Draw: drawSingapore},
	{Name: "slovakia", Aliases: []string{"sk"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{White, {0, 77, 167}, {238, 28, 37}})
	}},
	{Name: "slovenia", Aliases: []string{"si"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{White, {0, 51, 160}, {237, 28, 36}})
	}},
	{Name: "solomon-islands", Aliases: []string{"sb"}, Ratio: [2]int{1, 2}, Draw: drawSolomonIslands},
	{Name: "somalia", Aliases: []string{"so"}, Ratio: [2]int{2, 3}, Draw: drawSomalia},
	{Name: "south-africa", Aliases: []string{"za"}, Ratio: [2]int{2, 3}, Draw: drawSouthAfrica},
	{Name: "south-korea", Aliases: []string{"kr", "korea"}, Ratio: [2]int{2, 3}, Draw: drawSouthKorea},
	{Name: "south-sudan", Aliases: []string{"ss"}, Ratio: [2]int{1, 2}, Draw: drawSouthSudan},
	{Name: "spain", Aliases: []string{"es"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		red := RGB{198, 11, 30}
		yellow := RGB{255, 196, 0}
		c.FillRect(0, 0, c.Width, c.Height/4, red)
		c.FillRect(0, c.Height/4, c.Width, c.Height*3/4, yellow)
		c.FillRect(0, c.Height*3/4, c.Width, c.Height, red)
	}},
	{Name: "sri-lanka", Aliases: []string{"lk"}, Ratio: [2]int{1, 2}, Draw: drawSriLanka},
	{Name: "sudan", Aliases: []string{"sd"}, Ratio: [2]int{1, 2}, Draw: drawSudan},
	{Name: "suriname", Aliases: []string{"sr"}, Ratio: [2]int{2, 3}, Draw: drawSuriname},
	{Name: "sweden", Aliases: []string{"se", "swe"}, Ratio: [2]int{5, 8}, Draw: func(c *Canvas) {
		NordicCross(c, SwedishBlue, SwedishYellow, nil)
	}},
	{Name: "switzerland", Aliases: []string{"ch", "sui"}, Ratio: [2]int{1, 1}, Draw: func(c *Canvas) {
		SwissCross(c, SwissRed, White)
	}},
	{Name: "syria", Aliases: []string{"sy"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{206, 17, 38}, White, Black})
		FiveStar(c, float64(c.Width)/2-float64(c.Width)*0.08, float64(c.Height)/2, float64(c.Height)*0.08, RGB{0, 122, 61})
		FiveStar(c, float64(c.Width)/2+float64(c.Width)*0.08, float64(c.Height)/2, float64(c.Height)*0.08, RGB{0, 122, 61})
	}},
	// T
	{Name: "tajikistan", Aliases: []string{"tj"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		red := RGB{204, 0, 0}
		green := RGB{0, 122, 61}
		c.FillRect(0, 0, c.Width, c.Height/4, red)
		c.FillRect(0, c.Height/4, c.Width, c.Height*3/4, White)
		c.FillRect(0, c.Height*3/4, c.Width, c.Height, green)
	}},
	{Name: "tanzania", Aliases: []string{"tz"}, Ratio: [2]int{2, 3}, Draw: drawTanzania},
	{Name: "thailand", Aliases: []string{"th"}, Ratio: [2]int{2, 3}, Draw: drawThailand},
	{Name: "timor-leste", Aliases: []string{"tl", "east-timor"}, Ratio: [2]int{1, 2}, Draw: drawTimorLeste},
	{Name: "togo", Aliases: []string{"tg"}, Ratio: [2]int{3, 5}, Draw: drawTogo},
	{Name: "tonga", Aliases: []string{"to"}, Ratio: [2]int{1, 2}, Draw: drawTonga},
	{Name: "trinidad-and-tobago", Aliases: []string{"tt", "trinidad"}, Ratio: [2]int{3, 5}, Draw: drawTrinidadAndTobago},
	{Name: "tunisia", Aliases: []string{"tn"}, Ratio: [2]int{2, 3}, Draw: drawTunisia},
	{Name: "turkey", Aliases: []string{"tr"}, Ratio: [2]int{2, 3}, Draw: drawTurkey},
	{Name: "turkmenistan", Aliases: []string{"tm"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		green := RGB{0, 128, 64}
		red := RGB{172, 33, 49}
		c.Fill(green)
		// Ornamental stripe
		c.FillRect(c.Width/6, 0, c.Width/4, c.Height, red)
		// Crescent and stars
		cx := float64(c.Width) * 0.55
		cy := float64(c.Height) * 0.35
		Crescent(c, cx, cy, float64(c.Height)*0.15, float64(c.Height)*0.12, 1.5, White)
		for i := 0; i < 5; i++ {
			sy := float64(c.Height)*0.1 + float64(i)*float64(c.Height)*0.07
			FiveStar(c, cx+float64(c.Height)*0.15, sy, float64(c.Height)*0.03, White)
		}
	}},
	{Name: "tuvalu", Aliases: []string{"tv"}, Ratio: [2]int{1, 2}, Draw: drawTuvalu},
	// U
	{Name: "uganda", Aliases: []string{"ug"}, Ratio: [2]int{2, 3}, Draw: drawUganda},
	{Name: "uk", Aliases: []string{"gb", "britain"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		UnionJack(c)
	}},
	{Name: "ukraine", Aliases: []string{"ua", "ukr"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{UkraineBlue, UkraineYellow})
	}},
	{Name: "united-arab-emirates", Aliases: []string{"ae", "uae"}, Ratio: [2]int{1, 2}, Draw: drawUAE},
	{Name: "uruguay", Aliases: []string{"uy"}, Ratio: [2]int{2, 3}, Draw: drawUruguay},
	{Name: "usa", Aliases: []string{"us", "america"}, Ratio: [2]int{10, 19}, Draw: func(c *Canvas) {
		USAFlag(c)
	}},
	{Name: "uzbekistan", Aliases: []string{"uz"}, Ratio: [2]int{1, 2}, Draw: drawUzbekistan},
	// V
	{Name: "vanuatu", Aliases: []string{"vu"}, Ratio: [2]int{3, 5}, Draw: drawVanuatu},
	{Name: "vatican-city", Aliases: []string{"va", "vatican"}, Ratio: [2]int{1, 1}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{{252, 209, 22}, White})
	}},
	{Name: "venezuela", Aliases: []string{"ve"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{252, 209, 22}, {0, 51, 160}, {206, 17, 38}})
	}},
	{Name: "vietnam", Aliases: []string{"vn"}, Ratio: [2]int{2, 3}, Draw: drawVietnam},
	// Y
	{Name: "yemen", Aliases: []string{"ye"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{{206, 17, 38}, White, Black})
	}},
	// Z
	{Name: "zambia", Aliases: []string{"zm"}, Ratio: [2]int{2, 3}, Draw: drawZambia},
	{Name: "zimbabwe", Aliases: []string{"zw"}, Ratio: [2]int{1, 2}, Draw: drawZimbabwe},
}

var flagIndex map[string]*Flag

func init() {
	flagIndex = make(map[string]*Flag)
	for i := range Flags {
		f := &Flags[i]
		flagIndex[f.Name] = f
		for _, alias := range f.Aliases {
			flagIndex[alias] = f
		}
	}
}

func LookupFlag(name string) *Flag {
	return flagIndex[strings.ToLower(name)]
}

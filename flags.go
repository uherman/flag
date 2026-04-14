package main

import "strings"

type Flag struct {
	Name    string
	Aliases []string
	Ratio   [2]int // height:width
	Draw    func(c *Canvas)
}

var Flags = []Flag{
	{Name: "sweden", Aliases: []string{"se", "swe"}, Ratio: [2]int{5, 8}, Draw: func(c *Canvas) {
		NordicCross(c, SwedishBlue, SwedishYellow, nil)
	}},
	{Name: "norway", Aliases: []string{"no", "nor"}, Ratio: [2]int{8, 11}, Draw: func(c *Canvas) {
		NordicCross(c, NorwayRed, NorwayBlue, &White)
	}},
	{Name: "finland", Aliases: []string{"fi", "fin"}, Ratio: [2]int{11, 18}, Draw: func(c *Canvas) {
		NordicCross(c, White, FinlandBlue, nil)
	}},
	{Name: "denmark", Aliases: []string{"dk", "dan"}, Ratio: [2]int{3, 4}, Draw: func(c *Canvas) {
		NordicCross(c, DenmarkRed, White, nil)
	}},
	{Name: "france", Aliases: []string{"fr", "fra"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{FrenchBlue, White, FrenchRed})
	}},
	{Name: "germany", Aliases: []string{"de", "ger"}, Ratio: [2]int{3, 5}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{Black, GermanRed, GermanGold})
	}},
	{Name: "italy", Aliases: []string{"it", "ita"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{ItalyGreen, White, ItalyRed})
	}},
	{Name: "ireland", Aliases: []string{"ie", "irl"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		VStripes(c, []RGB{IrelandGreen, White, IrelandOrange})
	}},
	{Name: "netherlands", Aliases: []string{"nl", "ned"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{NetherlandsRed, White, NetherlandsBlue})
	}},
	{Name: "ukraine", Aliases: []string{"ua", "ukr"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{UkraineBlue, UkraineYellow})
	}},
	{Name: "poland", Aliases: []string{"pl", "pol"}, Ratio: [2]int{5, 8}, Draw: func(c *Canvas) {
		HStripes(c, []RGB{White, PolandRed})
	}},
	{Name: "japan", Aliases: []string{"jp", "jpn"}, Ratio: [2]int{2, 3}, Draw: func(c *Canvas) {
		Circle(c, White, JapanRed, 0.375)
	}},
	{Name: "switzerland", Aliases: []string{"ch", "sui"}, Ratio: [2]int{1, 1}, Draw: func(c *Canvas) {
		SwissCross(c, SwissRed, White)
	}},
	{Name: "usa", Aliases: []string{"us", "america"}, Ratio: [2]int{10, 19}, Draw: func(c *Canvas) {
		USAFlag(c)
	}},
	{Name: "uk", Aliases: []string{"gb", "britain"}, Ratio: [2]int{1, 2}, Draw: func(c *Canvas) {
		UnionJack(c)
	}},
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

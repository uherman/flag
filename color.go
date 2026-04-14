package main

import "fmt"

type RGB struct {
	R, G, B uint8
}

func (c RGB) FgEsc() string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

func (c RGB) BgEsc() string {
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", c.R, c.G, c.B)
}

const Reset = "\033[0m"

// Flag colors
var (
	White = RGB{255, 255, 255}
	Black = RGB{0, 0, 0}

	// Sweden
	SwedishBlue   = RGB{0, 106, 167}
	SwedishYellow = RGB{254, 204, 0}

	// Norway
	NorwayRed  = RGB{239, 43, 45}
	NorwayBlue = RGB{0, 40, 104}

	// Finland
	FinlandBlue = RGB{0, 53, 128}

	// Denmark
	DenmarkRed = RGB{200, 16, 46}

	// France
	FrenchBlue = RGB{0, 38, 84}
	FrenchRed  = RGB{206, 17, 38}

	// Germany
	GermanRed  = RGB{221, 0, 0}
	GermanGold = RGB{255, 204, 0}

	// Italy
	ItalyGreen = RGB{0, 146, 70}
	ItalyRed   = RGB{206, 43, 55}

	// Ireland
	IrelandGreen  = RGB{22, 155, 98}
	IrelandOrange = RGB{255, 136, 62}

	// Netherlands
	NetherlandsRed  = RGB{174, 28, 40}
	NetherlandsBlue = RGB{33, 70, 139}

	// Ukraine
	UkraineBlue   = RGB{0, 87, 183}
	UkraineYellow = RGB{255, 215, 0}

	// Poland
	PolandRed = RGB{220, 20, 60}

	// Japan
	JapanRed = RGB{188, 0, 45}

	// Switzerland
	SwissRed = RGB{218, 41, 28}

	// USA
	USARed  = RGB{178, 34, 52}
	USABlue = RGB{60, 59, 110}

	// UK
	UKBlue = RGB{1, 33, 105}
	UKRed  = RGB{200, 16, 46}
)

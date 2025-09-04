package resources

import (
	"bytes"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/text/language"
)

var UIFont text.Face

func init() {
	// 1) пробуем ваш файл; если нет — берём встроенный goregular
	data, err := os.ReadFile("assets/fonts/FindSansPro-Regular.ttf")
	if err != nil {
		data = goregular.TTF
	}

	// 2) создаём источник шрифта
	src, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		log.Fatalf("font source: %v", err)
	}

	// 3) собираем сам Face
	UIFont = &text.GoTextFace{
		Source:   src,
		Size:     14,                       // размер в пикселях
		Language: language.MustParse("ru"), // hint для шейпера
		// Direction: text.DirectionLeftToRight,   // опционально, по умолчанию LTR
	}
}

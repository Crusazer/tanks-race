package core

import "image/color"

type Theme struct {
	FontSize int

	ButtonNormalColor  color.Color
	ButtonHoverColor   color.Color
	ButtonPressedColor color.Color

	LabelTextColor color.Color

	TextInputBgColor          color.Color
	TextInputBorderColor      color.Color
	TextInputTextColor        color.Color
	TextInputPlaceholderColor color.Color

	SliderTrackColor  color.Color
	SliderHandleColor color.Color

	FocusBorderColor color.Color
}

var (
	DefaultTheme = &Theme{
		FontSize: 14,

		ButtonNormalColor:  color.RGBA{100, 100, 100, 255},
		ButtonHoverColor:   color.RGBA{150, 150, 150, 255},
		ButtonPressedColor: color.RGBA{200, 200, 200, 255},

		LabelTextColor: color.White,

		TextInputBgColor:          color.RGBA{50, 50, 50, 255},
		TextInputBorderColor:      color.RGBA{150, 150, 150, 255},
		TextInputTextColor:        color.White,
		TextInputPlaceholderColor: color.RGBA{120, 120, 120, 255},

		SliderTrackColor:  color.RGBA{100, 100, 100, 255},
		SliderHandleColor: color.RGBA{200, 200, 200, 255},

		FocusBorderColor: color.RGBA{255, 255, 0, 255},
	}

	DarkTheme = &Theme{
		FontSize: 14,

		ButtonNormalColor:  color.RGBA{40, 40, 40, 255},
		ButtonHoverColor:   color.RGBA{60, 60, 60, 255},
		ButtonPressedColor: color.RGBA{80, 80, 80, 255},

		LabelTextColor: color.RGBA{220, 220, 220, 255},

		TextInputBgColor:          color.RGBA{30, 30, 30, 255},
		TextInputBorderColor:      color.RGBA{100, 100, 100, 255},
		TextInputTextColor:        color.RGBA{220, 220, 220, 255},
		TextInputPlaceholderColor: color.RGBA{100, 100, 100, 255},

		SliderTrackColor:  color.RGBA{70, 70, 70, 255},
		SliderHandleColor: color.RGBA{150, 150, 150, 255},

		FocusBorderColor: color.RGBA{255, 200, 0, 255},
	}
)

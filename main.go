package main

import (
	"image/color"
	"ug"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

var logo = ug.UG_BMP{
	P:      unsafe.Pointer(&logo_bmp[0]),
	Width:  224,
	Height: 192,
	Bpp:    ug.BMP_BPP_16,
	Colors: ug.BMP_RGB565,
}

const MAX_OBJECTS = 10
const LCD_1IN14_WIDTH = 800 - 1
const LCD_1IN14_HEIGHT = 600 - 1

const BTN_ID_0 = 0
const BTN_ID_1 = 1
const BTN_ID_2 = 2

var maingui ug.UG_GUI
var window1 ug.UG_WINDOW
var button1 ug.UG_BUTTON
var button2 ug.UG_BUTTON
var button3 ug.UG_BUTTON
var image1 ug.UG_IMAGE
var textbox1 ug.UG_TEXTBOX
var obj []ug.UG_OBJECT = make([]ug.UG_OBJECT, MAX_OBJECTS)
var running bool = true

var surface *sdl.Surface
var window *sdl.Window

func calcColor(color uint32) (red, green, blue, alpha byte) {
	blue = byte(color & 0xFF)
	green = byte((color >> 8) & 0xFF)
	red = byte((color >> 16) & 0xFF)
	alpha = byte((color >> 24) & 0xFF)

	return red, green, blue, alpha
}

func UserPixelSetFunction(arg0 ug.UG_S16, arg1 ug.UG_S16, arg2 ug.UG_COLOR) {
	red, green, blue, alpha := calcColor(uint32(arg2))
	var c color.Color = color.RGBA{red, green, blue, alpha}
	surface.Set(int(arg0), int(arg1), c)
}

func window_1_callback(arg0 *ug.UG_MESSAGE) {
	arg0.Deref()
	if arg0.Id == ug.OBJ_TYPE_BUTTON {
		switch arg0.Sub_id {
		case BTN_ID_0:
			if arg0.Event == ug.OBJ_EVENT_PRESSED {
				ug.UG_ConsolePutString("C ")
			}
		case BTN_ID_1:
			if arg0.Event == ug.OBJ_EVENT_PRESSED {
				ug.UG_ConsolePutString("OK ")
			}
		case BTN_ID_2:
			if arg0.Event == ug.OBJ_EVENT_RELEASED {
				running = false
			}
		}
	}
	arg0.Free()
}

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("SDL2 Window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		400+1, 500+1, sdl.WINDOW_SHOWN|sdl.WINDOW_BORDERLESS)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	ug.UG_Init(&maingui, UserPixelSetFunction, LCD_1IN14_WIDTH, LCD_1IN14_HEIGHT)
	ug.UG_SelectGUI(&maingui)

	ug.UG_FontIdSelect(ug.FONT_ID_8X12)

	ug.UG_WindowCreate(&window1, obj, MAX_OBJECTS, window_1_callback)

	ug.UG_WindowSetTitleTextFont(&window1, ug.FONT_ID_8X12)
	ug.UG_WindowSetTitleText(&window1, "\xB5GUI Window title")
	ug.UG_WindowResize(&window1, 1, 1, 400, 400)

	ug.UG_ButtonCreate(&window1, &button1, BTN_ID_0, 1, 1, 50, 50)
	ug.UG_ButtonSetFont(&window1, BTN_ID_0, ug.FONT_ID_8X14)
	ug.UG_ButtonSetText(&window1, BTN_ID_0, "C")
	ug.UG_ButtonShow(&window1, BTN_ID_0)

	ug.UG_ButtonCreate(&window1, &button2, BTN_ID_1, 1, 60, 50, 110)
	ug.UG_ButtonSetText(&window1, BTN_ID_1, "OK")
	ug.UG_ButtonSetFont(&window1, BTN_ID_1, ug.FONT_ID_8X14)
	ug.UG_ButtonShow(&window1, BTN_ID_1)

	ug.UG_ButtonCreate(&window1, &button3, BTN_ID_2, 365, 1, 390, 25)
	ug.UG_ButtonSetText(&window1, BTN_ID_2, "X")
	ug.UG_ButtonSetFont(&window1, BTN_ID_2, ug.FONT_ID_8X14)
	ug.UG_ButtonShow(&window1, BTN_ID_2)

	ug.UG_TextboxCreate(&window1, &textbox1, ug.TXB_ID_0, 60, 1, 300, 25)
	ug.UG_TextboxSetFont(&window1, ug.TXB_ID_0, ug.FONT_ID_8X12)
	ug.UG_TextboxSetText(&window1, ug.TXB_ID_0, "GOLang + \xB5GUI + GoSDL2")
	ug.UG_TextboxSetBackColor(&window1, ug.TXB_ID_0, ug.C_WHITE)
	ug.UG_TextboxShow(&window1, ug.TXB_ID_0)

	ug.UG_ImageCreate(&window1, &image1, ug.IMG_ID_0, 80, 30, 330, 265)
	ug.UG_ImageSetBMP(&window1, ug.IMG_ID_0, &logo)

	ug.UG_WindowShow(&window1)

	ug.UG_ConsoleSetBackcolor(ug.C_BLACK)
	ug.UG_ConsoleSetArea(0, 400, 400, 500)
	ug.UG_ConsoleSetForecolor(ug.C_GREEN)
	ug.UG_ConsolePutString("GOLang \xB5GUI wrapper ;) ")
	ug.UG_ConsoleSetForecolor(ug.C_GREEN_YELLOW)
	ug.UG_ConsolePutString("Press the button...\n")
	ug.UG_ConsoleSetForecolor(ug.C_WHITE)

	ug.UG_FocusUpdate(ug.UG_MOVE_FOCUS_UP)

	for running {
		ug.UG_Update()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.MouseButtonEvent:
				x, y, state := sdl.GetMouseState()

				ug.UG_TouchUpdate(ug.UG_S16(x), ug.UG_S16(y), ug.UG_U8(state))

			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_UP && t.State == sdl.RELEASED {
					ug.UG_FocusUpdate(ug.UG_MOVE_FOCUS_UP)
				} else if t.Keysym.Sym == sdl.K_DOWN && t.State == sdl.RELEASED {
					ug.UG_FocusUpdate(ug.UG_MOVE_FOCUS_DOWN)
				} else if t.Keysym.Sym == sdl.K_LEFT && t.State == sdl.RELEASED {
					ug.UG_FocusUpdate(ug.UG_MOVE_FOCUS_LEFT)
				} else if t.Keysym.Sym == sdl.K_RIGHT && t.State == sdl.RELEASED {
					ug.UG_FocusUpdate(ug.UG_MOVE_FOCUS_RIGHT)
				} else if t.Keysym.Sym == sdl.K_RETURN && t.State == sdl.RELEASED {
					ug.UG_PressFocused()
				}
			case *sdl.QuitEvent:
				println("Quit")
				ug.UG_ConsolePutString("Quit")
				running = false
			}
		}

		window.UpdateSurface()
	}

}

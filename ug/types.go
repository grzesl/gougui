// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 21 Nov 2023 19:54:43 CET.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package ug

/*
#cgo LDFLAGS: -L../ugui/build -lugui
#include "../ugui/ugui.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// UG_FONT as declared in ugui/ugui.h:64
type UG_FONT struct {
	P              []byte
	Font_type      FONT_TYPE
	Char_width     UG_S16
	Char_height    UG_S16
	Start_char     UG_U16
	End_char       UG_U16
	Widths         []UG_U8
	ref83da41da    *C.UG_FONT
	allocs83da41da interface{}
}

// UG_OBJECT as declared in ugui/ugui.h:121
type UG_OBJECT struct {
	State          UG_U8
	Touch_state    UG_U8
	Update         UG_OBJECT_UPDATE
	A_abs          UG_AREA
	A_rel          UG_AREA
	_type          UG_U8
	Id             UG_U8
	Event          UG_U8
	Data           unsafe.Pointer
	refc19e8ff4    *C.UG_OBJECT
	allocsc19e8ff4 interface{}
}

// UG_WINDOW as declared in ugui/ugui.h:122
type UG_WINDOW struct {
	Objcnt         UG_U8
	Objlst         []UG_OBJECT
	State          UG_U8
	Fc             UG_COLOR
	Bc             UG_COLOR
	Xs             UG_S16
	Ys             UG_S16
	Xe             UG_S16
	Ye             UG_S16
	Style          UG_U8
	Title          UG_TITLE
	Cb             UG_Message_Callback
	refe2d7ddc5    *C.UG_WINDOW
	allocse2d7ddc5 interface{}
}

// UG_RESULT type as declared in ugui/ugui.h:123
type UG_RESULT byte

// UG_COLOR type as declared in ugui/ugui.h:125
type UG_COLOR uint32

// UG_AREA as declared in ugui/ugui.h:192
type UG_AREA struct {
	Xs             UG_S16
	Ys             UG_S16
	Xe             UG_S16
	Ye             UG_S16
	ref84da7460    *C.UG_AREA
	allocs84da7460 interface{}
}

// UG_TEXT as declared in ugui/ugui.h:205
type UG_TEXT struct {
	Str            string
	Font           []UG_FONT
	A              UG_AREA
	Fc             UG_COLOR
	Bc             UG_COLOR
	Align          UG_U8
	H_space        UG_S16
	V_space        UG_S16
	ref68c5eecf    *C.UG_TEXT
	allocs68c5eecf interface{}
}

// UG_BMP as declared in ugui/ugui.h:217
type UG_BMP struct {
	P              unsafe.Pointer
	Width          UG_U16
	Height         UG_U16
	Bpp            UG_U8
	Colors         UG_U8
	refc99e9a58    *C.UG_BMP
	allocsc99e9a58 interface{}
}

// UG_MESSAGE as declared in ugui/ugui.h:242
type UG_MESSAGE struct {
	_type          UG_U8
	Id             UG_U8
	Sub_id         UG_U8
	Event          UG_U16
	Pos_x          UG_U32
	Pos_y          UG_U32
	Src            unsafe.Pointer
	ref9ed6bbc5    *C.UG_MESSAGE
	allocs9ed6bbc5 interface{}
}

// UG_TOUCH as declared in ugui/ugui.h:258
type UG_TOUCH struct {
	State          UG_U8
	Xp             UG_S16
	Yp             UG_S16
	refc350cb41    *C.UG_TOUCH
	allocsc350cb41 interface{}
}

// UG_OBJECT_UPDATE type as declared in ugui/ugui.h:268
type UG_OBJECT_UPDATE func(Arg0 *UG_WINDOW, Arg1 *UG_OBJECT)

// UG_TITLE as declared in ugui/ugui.h:347
type UG_TITLE struct {
	Str            string
	Font           []UG_FONT
	H_space        UG_S8
	V_space        UG_S8
	Align          UG_U8
	Fc             UG_COLOR
	Bc             UG_COLOR
	Ifc            UG_COLOR
	Ibc            UG_COLOR
	Height         UG_U8
	Prev_touch_px  UG_U8
	Prev_touch_py  UG_U8
	Touch_state    UG_U16
	Event          UG_U8
	Data           unsafe.Pointer
	ref1ed09ed8    *C.UG_TITLE
	allocs1ed09ed8 interface{}
}

// UG_Message_Callback type as declared in ugui/ugui.h:351
type UG_Message_Callback func(Arg0 *UG_MESSAGE)

// UG_BUTTON as declared in ugui/ugui.h:400
type UG_BUTTON struct {
	State          UG_U8
	Style          UG_U8
	Fc             UG_COLOR
	Bc             UG_COLOR
	Afc            UG_COLOR
	Abc            UG_COLOR
	Font           []UG_FONT
	Align          UG_U8
	H_space        UG_S8
	V_space        UG_S8
	Str            string
	ref53358825    *C.UG_BUTTON
	allocs53358825 interface{}
}

// UG_CHECKBOX as declared in ugui/ugui.h:458
type UG_CHECKBOX struct {
	State         UG_U8
	Style         UG_U8
	Fc            UG_COLOR
	Bc            UG_COLOR
	Afc           UG_COLOR
	Abc           UG_COLOR
	Font          []UG_FONT
	Align         UG_U8
	H_space       UG_S8
	V_space       UG_S8
	Str           string
	Checked       UG_U8
	refe89193c    *C.UG_CHECKBOX
	allocse89193c interface{}
}

// UG_TEXTBOX as declared in ugui/ugui.h:513
type UG_TEXTBOX struct {
	Str            string
	Font           []UG_FONT
	Style          UG_U8
	Fc             UG_COLOR
	Bc             UG_COLOR
	Align          UG_U8
	H_space        UG_S8
	V_space        UG_S8
	ref2f0e33f8    *C.UG_TEXTBOX
	allocs2f0e33f8 interface{}
}

// UG_IMAGE as declared in ugui/ugui.h:545
type UG_IMAGE struct {
	Img            unsafe.Pointer
	_type          UG_U8
	reff0dbe2ec    *C.UG_IMAGE
	allocsf0dbe2ec interface{}
}

// UG_DRIVER as declared in ugui/ugui.h:579
type UG_DRIVER struct {
	Driver         unsafe.Pointer
	State          UG_U8
	ref785558c1    *C.UG_DRIVER
	allocs785558c1 interface{}
}

// UG_GUI_PSET type as declared in ugui/ugui.h:594
type UG_GUI_PSET func(Arg0 UG_S16, Arg1 UG_S16, Arg2 UG_COLOR)

// UG_GUI as declared in ugui/ugui.h:623
type UG_GUI struct {
	Pset           UG_GUI_PSET
	X_dim          UG_S16
	Y_dim          UG_S16
	Touch          UG_TOUCH
	Next_window    []UG_WINDOW
	Active_window  []UG_WINDOW
	Last_window    []UG_WINDOW
	Font           UG_FONT
	Char_h_space   UG_S8
	Char_v_space   UG_S8
	Fore_color     UG_COLOR
	Back_color     UG_COLOR
	Desktop_color  UG_COLOR
	State          UG_U8
	Driver         [3]UG_DRIVER
	ref2925682a    *C.UG_GUI
	allocs2925682a interface{}
}

// UG_Init_Callback type as declared in ugui/ugui.h:919
type UG_Init_Callback func(Arg0 UG_S16, Arg1 UG_S16, Arg2 UG_COLOR)

// UG_U8 type as declared in ugui/ugui_config.h:37
type UG_U8 byte

// UG_S8 type as declared in ugui/ugui_config.h:38
type UG_S8 byte

// UG_U16 type as declared in ugui/ugui_config.h:39
type UG_U16 uint16

// UG_S16 type as declared in ugui/ugui_config.h:40
type UG_S16 int16

// UG_U32 type as declared in ugui/ugui_config.h:41
type UG_U32 uint32

// UG_S32 type as declared in ugui/ugui_config.h:42
type UG_S32 int32

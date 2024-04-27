package theme

import (
	"image/color"
	"testing"

	"fyne.io/fyne/v2"

	"github.com/stretchr/testify/assert"
)

func TestThemeChange(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	bg := BackgroundColor()

	fyne.CurrentApp().Settings().SetTheme(LightTheme())
	assert.NotEqual(t, bg, BackgroundColor())
}

func TestTheme_Bootstrapping(t *testing.T) {
	current := fyne.CurrentApp().Settings().Theme()
	fyne.CurrentApp().Settings().SetTheme(nil)

	// this should not crash
	BackgroundColor()

	fyne.CurrentApp().Settings().SetTheme(current)
}

func TestBuiltinTheme_ShadowColor(t *testing.T) {
	shadow := ShadowColor()

	_, _, _, a := shadow.RGBA()
	assert.NotEqual(t, 255, a)
}

func TestTheme_Dark_ReturnsCorrectBackground(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	bg := BackgroundColor()
	assert.Equal(t, DarkTheme().Color(ColorNameBackground, VariantDark), bg, "wrong dark theme background color")
}

func TestTheme_Light_ReturnsCorrectBackground(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(LightTheme())
	bg := BackgroundColor()
	assert.Equal(t, LightTheme().Color(ColorNameBackground, VariantLight), bg, "wrong light theme background color")
}

func Test_TextSize(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	assert.Equal(t, DarkTheme().Size(SizeNameText), TextSize(), "wrong text size")
}

func Test_TextFont(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	expect := "NotoSans-Regular.ttf"
	result := TextFont().Name()
	assert.Equal(t, expect, result, "wrong regular text font")
}

func Test_TextBoldFont(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	expect := "NotoSans-Bold.ttf"
	result := TextBoldFont().Name()
	assert.Equal(t, expect, result, "wrong bold text font")
}

func Test_TextItalicFont(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	expect := "NotoSans-Italic.ttf"
	result := TextItalicFont().Name()
	assert.Equal(t, expect, result, "wrong italic text font")
}

func Test_TextBoldItalicFont(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	expect := "NotoSans-BoldItalic.ttf"
	result := TextBoldItalicFont().Name()
	assert.Equal(t, expect, result, "wrong bold italic text font")
}

func Test_TextMonospaceFont(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	expect := "DejaVuSansMono-Powerline.ttf"
	result := TextMonospaceFont().Name()
	assert.Equal(t, expect, result, "wrong monospace font")
}

func Test_TextSymbolFont(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	expect := "InterSymbols-Regular.ttf"
	result := SymbolFont().Name()
	assert.Equal(t, expect, result, "wrong symbol font")
}

func Test_Padding(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	assert.Equal(t, DarkTheme().Size(SizeNamePadding), Padding(), "wrong padding")
}

func Test_IconInlineSize(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	assert.Equal(t, DarkTheme().Size(SizeNameInlineIcon), IconInlineSize(), "wrong inline icon size")
}

func Test_ScrollBarSize(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	assert.Equal(t, DarkTheme().Size(SizeNameScrollBar), ScrollBarSize(), "wrong inline icon size")
}

func Test_DefaultSymbolFont(t *testing.T) {
	expect := "InterSymbols-Regular.ttf"
	result := DefaultSymbolFont().Name()
	assert.Equal(t, expect, result, "wrong default text font")
}

func Test_DefaultTextFont(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	expect := "NotoSans-Regular.ttf"
	result := DefaultTextFont().Name()
	assert.Equal(t, expect, result, "wrong default text font")
}

func Test_DefaultTextBoldFont(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	expect := "NotoSans-Bold.ttf"
	result := DefaultTextBoldFont().Name()
	assert.Equal(t, expect, result, "wrong default text bold font")
}

func Test_DefaultTextBoldItalicFont(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	expect := "NotoSans-BoldItalic.ttf"
	result := DefaultTextBoldItalicFont().Name()
	assert.Equal(t, expect, result, "wrong default text bold italic font")
}

func Test_DefaultTextItalicFont(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	expect := "NotoSans-Italic.ttf"
	result := DefaultTextItalicFont().Name()
	assert.Equal(t, expect, result, "wrong default text italic font")
}

func Test_DefaultTextMonospaceFont(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
	expect := "DejaVuSansMono-Powerline.ttf"
	result := DefaultTextMonospaceFont().Name()
	assert.Equal(t, expect, result, "wrong default monospace font")
}

func TestEmptyTheme(t *testing.T) {
	fyne.CurrentApp().Settings().SetTheme(&emptyTheme{})
	assert.NotNil(t, ForegroundColor())
	assert.NotNil(t, TextFont())
	assert.NotNil(t, HelpIcon())
	fyne.CurrentApp().Settings().SetTheme(DarkTheme())
}

type emptyTheme struct {
}

func (e *emptyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return nil
}

func (e *emptyTheme) Font(s fyne.TextStyle) fyne.Resource {
	return nil
}

func (e *emptyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return nil
}

func (e *emptyTheme) Size(n fyne.ThemeSizeName) float32 {
	return 0
}

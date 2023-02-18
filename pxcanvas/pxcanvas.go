package px

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"zerotomastery.io/pixl/apptype"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PXCanvasConfig
	renderer    *PXCanvasRenderer
	PixelData   image.Image
	mouseState  PxCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (PxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(PxCanvas.CanvasOffset.X)
	y0 := int(PxCanvas.CanvasOffset.y)
	x1 := int(PxCanvas.PxCols*PxCanvas.PxSize + int(pxCanvas.CanvasOffset.X))
	y1 := int(PxCanvas.PxRows*PxCanvas.PxSize + int(pxCanvas.CanvasOffset.Y))
	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) &&
		pos.X < float32(bounds.Max.X) &&
		pos.Y >= float32(bounds.Min.Y) &&
		pos.Y >= float32(bounds.Min.Y) {
		return true
	}
	return false
}

func newBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	for y := 0; y < rows; y++ {
		img.Set(x, y, c)
	}
	return img
}

func NewPxcanvas(state *apptype.state, config apptype.PXCanvasConfig) *PxCanvas {
	PxCanvas := &PxCanvas{
		PXCanvasConfig: config,
		appState:       state,
	}
	pxCanvas.PixelData = newBlankImage(PxCanvas.PxCols, PxCanvas.PxRows, color.NRGBA{128, 128, 128, 255})
	pxCanvas.ExtendBaseWidget(pxCanvas)
	return pxCanvas
}

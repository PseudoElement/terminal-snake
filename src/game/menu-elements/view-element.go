package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
)

type Coords struct {
	x, y          int
	width, height int
}

type BaseViewElement struct {
	id         string
	coords     Coords
	teaElement lipgloss.Style
}

func NewBaseViewElement(coords Coords, teaElement lipgloss.Style) BaseViewElement {
	id := uuid.New()
	return BaseViewElement{
		id:         id.String(),
		coords:     coords,
		teaElement: teaElement}
}

func (this *BaseViewElement) IsHovered(cursorX, cursorY int) bool {
	endX := this.coords.x + this.coords.width
	endY := this.coords.y + this.coords.height

	return cursorX >= this.coords.x &&
		cursorX <= endX &&
		cursorY >= this.coords.y &&
		cursorY <= endY
}

func (this *BaseViewElement) Id() string {
	return this.id
}

func (this *BaseViewElement) TeaElement() lipgloss.Style {
	return this.teaElement
}

func (this *BaseViewElement) NextViewElements() []IViewElement {
	panic("need implement in child class!")
}

package game_abstr

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
)

type BaseViewElement struct {
	id         string
	teaElement lipgloss.Style
}

func NewBaseViewElement(teaElement lipgloss.Style) *BaseViewElement {
	id := uuid.New()
	return &BaseViewElement{
		id:         id.String(),
		teaElement: teaElement,
	}
}

func (this *BaseViewElement) UpdateTeaElement(teaElement lipgloss.Style) {
	this.teaElement = teaElement
}

func (this *BaseViewElement) Id() string {
	return this.id
}

func (this *BaseViewElement) TeaElement() lipgloss.Style {
	return this.teaElement
}

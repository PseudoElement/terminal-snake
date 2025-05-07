package game_abstr

import "github.com/charmbracelet/lipgloss"

type SelectableElement struct {
	BaseViewElement
	selected bool
}

func NewSelectableElement(teaElement lipgloss.Style) *SelectableElement {
	return &SelectableElement{BaseViewElement: NewBaseViewElement(teaElement)}
}

func (this *SelectableElement) IsSelected(selectedElemId string) bool {
	return selectedElemId == this.id
}

func (this *SelectableElement) SetSelected(selected bool) {
	this.selected = selected
}

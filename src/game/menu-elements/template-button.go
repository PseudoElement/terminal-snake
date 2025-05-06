package menu_elements

import "github.com/charmbracelet/lipgloss"

type MenuButton struct {
	BaseViewElement
	onClick ClickHandler
}

func NewTemplateButton(coords Coords, onClick ClickHandler) *MenuButton {
	btn := lipgloss.NewStyle().
		Background(lipgloss.Color("#888B7E")).
		MarginTop(coords.y).
		MarginRight(coords.x).
		Height(coords.height).
		Width(coords.width)

	return &MenuButton{
		BaseViewElement: NewBaseViewElement(coords, btn),
		onClick:         onClick}
}

func (this *MenuButton) NextViewElements() []IViewElement {
	return this.onClick()
}

var _ IViewElement = (*MenuButton)(nil)

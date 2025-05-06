package menu_elements

import "github.com/charmbracelet/lipgloss"

type IViewElement interface {
	/* whether cursor is over element or not */
	IsHovered(cursorX, cursorY int) bool

	/* next view elements for rendering after clicking on element */
	NextViewElements() []IViewElement

	Id() string

	TeaElement() lipgloss.Style
}

type ClickHandler func() []IViewElement

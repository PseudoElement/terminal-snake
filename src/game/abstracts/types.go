package game_abstr

import "github.com/charmbracelet/lipgloss"

type IViewElement interface {
	/* update style on action */
	UpdateTeaElement(teaElement lipgloss.Style)

	Id() string

	TeaElement() lipgloss.Style

	View() string
}

type ISelectableElement interface {
	IViewElement

	/* optional action when element was clicked */
	Action()

	Select()

	Blur()

	SetSelected(selected bool)

	/* whether element selected or not */
	IsSelected(selectedElemId string) bool
}

type IRedirectableElement interface {
	ISelectableElement

	/* next element for rendering after clicking on element */
	NextPage() IPage
}

type IPage interface {
	SelectableElems() []ISelectableElement

	SelectableElemsToViews() []string

	View() string
}

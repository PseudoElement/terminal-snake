package game_abstr

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
)

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
	Action(store *store.Store)

	Select()

	Blur()

	SetSelected(selected bool)

	/* whether element selected or not */
	IsSelected() bool
}

type IRedirectableElement interface {
	ISelectableElement

	/* next element for rendering after clicking on element */
	NextPage(store *store.Store) IPage
}

type IPage interface {
	SelectableElems() []ISelectableElement

	HasSelectableElems() bool

	SelectableElemsToViews() []string

	View() string

	Store() *store.Store

	OnInit()
}

type IGamePage interface {
	IPage

	GameScene() IGameScene
}

type IDiffLevel interface {
	SnakeSpeedMs() int64

	IsSnakeDied(scene IGameScene) bool

	Name() string
}

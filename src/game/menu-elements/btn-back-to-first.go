package menu_elements

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type BackToMenuBtn struct {
	*game_abstr.SelectableElement
}

func NewBackToMenuBtn() *BackToMenuBtn {
	teaElement := bluredBtn
	btn := &BackToMenuBtn{
		SelectableElement: game_abstr.NewSelectableElement(teaElement),
	}

	return btn
}

func (this *BackToMenuBtn) Action(store *store.Store) {
	store.Set(consts.SCORE, 0)
	store.Set(consts.MOVE_DIRECTION, consts.RIGHT)
}

func (this *BackToMenuBtn) NextPage(store *store.Store) game_abstr.IPage {
	return NewFirstPage(store)
}

func (this *BackToMenuBtn) View() string {
	return this.TeaElement().Render("Back to menu")
}

func (this *BackToMenuBtn) Select() {
	selectedTeaElem := selectedBtn
	this.SetSelected(true)
	this.UpdateTeaElement(selectedTeaElem)
}

func (this *BackToMenuBtn) Blur() {
	bluredTeaElem := bluredBtn
	this.SetSelected(false)
	this.UpdateTeaElement(bluredTeaElem)
}

var _ game_abstr.IRedirectableElement = (*BackToMenuBtn)(nil)

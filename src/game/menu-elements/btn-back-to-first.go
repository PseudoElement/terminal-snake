package menu_elements

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
)

type BackToMenuBtn struct {
	*game_abstr.SelectableElement
	closePage func()
}

func NewBackToMenuBtn(closePage func()) *BackToMenuBtn {
	teaElement := bluredBtn
	teaElement.Width(20)

	btn := &BackToMenuBtn{
		SelectableElement: game_abstr.NewSelectableElement(teaElement),
		closePage:         closePage,
	}

	return btn
}

func (this *BackToMenuBtn) Action(store *store.Store) {
	this.closePage()
}

func (this *BackToMenuBtn) NextPage(store *store.Store) game_abstr.IPage {
	return NewFirstPage(store)
}

func (this *BackToMenuBtn) View() string {
	return this.TeaElement().Render("Back to menu")
}

func (this *BackToMenuBtn) Select() {
	selectedTeaElem := selectedBtn
	selectedTeaElem.Width(20)

	this.SetSelected(true)
	this.UpdateTeaElement(selectedTeaElem)
}

func (this *BackToMenuBtn) Blur() {
	bluredTeaElem := bluredBtn
	bluredTeaElem.Width(20)

	this.SetSelected(false)
	this.UpdateTeaElement(bluredTeaElem)
}

var _ game_abstr.IRedirectableElement = (*BackToMenuBtn)(nil)

package menu_elements

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
)

type RulesBtn struct {
	*game_abstr.SelectableElement
}

func NewRulesBtn() *RulesBtn {
	teaElement := bluredBtn
	btn := &RulesBtn{
		SelectableElement: game_abstr.NewSelectableElement(teaElement),
	}

	return btn
}

func (this *RulesBtn) Action(store *store.Store) {}

func (this *RulesBtn) NextPage(store *store.Store) game_abstr.IPage {
	return NewRulesPage(store)
}

func (this *RulesBtn) View() string {
	return this.TeaElement().Render("Rules")
}

func (this *RulesBtn) Select() {
	selectedTeaElem := selectedBtn
	this.SetSelected(true)
	this.UpdateTeaElement(selectedTeaElem)
}

func (this *RulesBtn) Blur() {
	bluredTeaElem := bluredBtn
	this.SetSelected(false)
	this.UpdateTeaElement(bluredTeaElem)
}

var _ game_abstr.IRedirectableElement = (*RulesBtn)(nil)

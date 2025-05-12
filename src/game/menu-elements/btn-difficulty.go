package menu_elements

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
)

type DifficultyBtn struct {
	*game_abstr.SelectableElement
}

func NewDifficultyBtn() *DifficultyBtn {
	teaElement := bluredBtn
	btn := &DifficultyBtn{
		SelectableElement: game_abstr.NewSelectableElement(teaElement),
	}

	return btn
}

func (this *DifficultyBtn) Action(store *store.Store) {}

func (this *DifficultyBtn) NextPage(store *store.Store) game_abstr.IPage {
	return NewDifficultyPage(store)
}

func (this *DifficultyBtn) View() string {
	return this.TeaElement().Render("Difficulty levels")
}

func (this *DifficultyBtn) Select() {
	selectedTeaElem := selectedBtn
	this.SetSelected(true)
	this.UpdateTeaElement(selectedTeaElem)
}

func (this *DifficultyBtn) Blur() {
	bluredTeaElem := bluredBtn
	this.SetSelected(false)
	this.UpdateTeaElement(bluredTeaElem)
}

var _ game_abstr.IRedirectableElement = (*DifficultyBtn)(nil)

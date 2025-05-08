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
	teaElement.Width(30)

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
	return this.TeaElement().Render("Difficulty level")
}

func (this *DifficultyBtn) Select() {
	selectedTeaElem := selectedBtn
	selectedTeaElem.Width(30)

	this.SetSelected(true)
	this.UpdateTeaElement(selectedTeaElem)
}

func (this *DifficultyBtn) Blur() {
	bluredTeaElem := bluredBtn
	bluredTeaElem.Width(30)

	this.SetSelected(false)
	this.UpdateTeaElement(bluredTeaElem)
}

var _ game_abstr.IRedirectableElement = (*DifficultyBtn)(nil)

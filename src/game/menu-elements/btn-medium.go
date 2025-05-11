package menu_elements

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	diff_levels "github.com/pseudoelement/terminal-snake/src/game/game-elements/difficulty-levels"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type MediumBtn struct {
	*game_abstr.SelectableElement
}

func NewMediumBtn() *MediumBtn {
	teaElement := bluredBtn
	btn := &MediumBtn{
		SelectableElement: game_abstr.NewSelectableElement(teaElement),
	}

	return btn
}

func (this *MediumBtn) Action(store *store.Store) {
	store.Add(consts.DIFFICULTY, diff_levels.NewMediumLevel())
}

func (this *MediumBtn) NextPage(store *store.Store) game_abstr.IPage {
	return NewFirstPage(store)
}

func (this *MediumBtn) View() string {
	return this.TeaElement().Render("Medium level")
}

func (this *MediumBtn) Select() {
	selectedTeaElem := selectedBtn
	this.SetSelected(true)
	this.UpdateTeaElement(selectedTeaElem)
}

func (this *MediumBtn) Blur() {
	bluredTeaElem := bluredBtn
	this.SetSelected(false)
	this.UpdateTeaElement(bluredTeaElem)
}

var _ game_abstr.IRedirectableElement = (*MediumBtn)(nil)

package menu_elements

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	diff_levels "github.com/pseudoelement/terminal-snake/src/game/entities/difficulty-levels"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type EasyBtn struct {
	*game_abstr.SelectableElement
}

func NewEasyBtn() *EasyBtn {
	teaElement := bluredBtn
	btn := &EasyBtn{
		SelectableElement: game_abstr.NewSelectableElement(teaElement),
	}

	return btn
}

func (this *EasyBtn) Action(store *store.Store) {
	store.Add(consts.DIFFICULTY, diff_levels.NewEasyLevel())
}

func (this *EasyBtn) NextPage(store *store.Store) game_abstr.IPage {
	return NewFirstPage(store)
}

func (this *EasyBtn) View() string {
	return this.TeaElement().Render("Easy level")
}

func (this *EasyBtn) Select() {
	selectedTeaElem := selectedBtn
	this.SetSelected(true)
	this.UpdateTeaElement(selectedTeaElem)
}

func (this *EasyBtn) Blur() {
	bluredTeaElem := bluredBtn
	this.SetSelected(false)
	this.UpdateTeaElement(bluredTeaElem)
}

var _ game_abstr.IRedirectableElement = (*EasyBtn)(nil)

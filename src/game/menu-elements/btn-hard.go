package menu_elements

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	consts "github.com/pseudoelement/terminal-snake/src/game/constants"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
)

type HardBtn struct {
	*game_abstr.SelectableElement
}

func NewHardBtn() *HardBtn {
	teaElement := bluredBtn
	btn := &HardBtn{
		SelectableElement: game_abstr.NewSelectableElement(teaElement),
	}

	return btn
}

func (this *HardBtn) Action(store *store.Store) {
	store.Add(consts.DIFFICULTY, consts.HARD_LVL)
}

func (this *HardBtn) NextPage(store *store.Store) game_abstr.IPage {
	return NewFirstPage(store)
}

func (this *HardBtn) View() string {
	return this.TeaElement().Render("Hard level")
}

func (this *HardBtn) Select() {
	selectedTeaElem := selectedBtn
	this.SetSelected(true)
	this.UpdateTeaElement(selectedTeaElem)
}

func (this *HardBtn) Blur() {
	bluredTeaElem := bluredBtn
	this.SetSelected(false)
	this.UpdateTeaElement(bluredTeaElem)
}

var _ game_abstr.IRedirectableElement = (*HardBtn)(nil)

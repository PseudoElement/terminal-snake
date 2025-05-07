package menu_elements

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
)

type HardBtn struct {
	game_abstr.SelectableElement
}

func NewHardBtn() *HardBtn {
	teaElement := bluredBtn
	btn := &HardBtn{
		SelectableElement: *game_abstr.NewSelectableElement(teaElement),
	}

	return btn
}

func (this *HardBtn) Action() {}

func (this *HardBtn) NextPage() game_abstr.IPage {
	return NewFirstPage()
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

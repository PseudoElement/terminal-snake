package menu_elements

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
)

type EasyBtn struct {
	game_abstr.SelectableElement
}

func NewEasyBtn() *EasyBtn {
	teaElement := bluredBtn
	btn := &EasyBtn{
		SelectableElement: *game_abstr.NewSelectableElement(teaElement),
	}

	return btn
}

func (this *EasyBtn) Action() {}

func (this *EasyBtn) NextPage() game_abstr.IPage {
	return NewFirstPage()
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

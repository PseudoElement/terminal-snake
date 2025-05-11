package menu_controller

import game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"

type MenuController struct {
	page game_abstr.IPage
}

func NewMenuController(page game_abstr.IPage) *MenuController {
	return &MenuController{
		page: page,
	}
}

func (this *MenuController) Page() game_abstr.IPage {
	return this.page
}

func (this *MenuController) SelectedElemId() string {
	return this.SelectedElem().Id()
}

func (this *MenuController) SelectedElem() game_abstr.ISelectableElement {
	for _, el := range this.page.SelectableElems() {
		if el.IsSelected() {
			return el
		}
	}
	return nil
}

func (this *MenuController) SetPage(page game_abstr.IPage) {
	this.page = page
}

func (this *MenuController) SelectNext() {
	var currSelectedIdx int
	for idx, el := range this.page.SelectableElems() {
		if el.IsSelected() {
			currSelectedIdx = idx
			el.Blur()
			break
		}
	}

	var newSelectedElem game_abstr.ISelectableElement
	if currSelectedIdx == len(this.page.SelectableElems())-1 {
		newSelectedElem = this.page.SelectableElems()[0]
	} else {
		newSelectedElem = this.page.SelectableElems()[currSelectedIdx+1]
	}

	newSelectedElem.Select()
}

func (this *MenuController) SelectPrev() {
	var currSelectedIdx int
	for idx, el := range this.page.SelectableElems() {
		if el.IsSelected() {
			currSelectedIdx = idx
			el.Blur()
			break
		}
	}

	var newSelectedElem game_abstr.ISelectableElement
	if currSelectedIdx == 0 {
		newSelectedElem = this.page.SelectableElems()[len(this.page.SelectableElems())-1]
	} else {
		newSelectedElem = this.page.SelectableElems()[currSelectedIdx-1]
	}

	newSelectedElem.Select()
}

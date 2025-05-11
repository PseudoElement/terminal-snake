package game_abstr

import "github.com/pseudoelement/terminal-snake/src/game/services/store"

type Page struct {
	store           *store.Store
	selectableElems []ISelectableElement
}

func NewPage(store *store.Store, selectableElems []ISelectableElement) *Page {
	return &Page{
		store:           store,
		selectableElems: selectableElems,
	}
}

func (this *Page) SelectableElemsToViews() []string {
	var views = make([]string, 0, len(this.selectableElems))
	for _, el := range this.selectableElems {
		views = append(views, el.View())
	}

	return views
}

func (this *Page) HasSelectableElems() bool {
	return len(this.selectableElems) > 0
}

func (this *Page) SelectableElems() []ISelectableElement {
	return this.selectableElems
}

func (this *Page) Store() *store.Store {
	return this.store
}

func (this *Page) IsGamePage() bool {
	return false
}

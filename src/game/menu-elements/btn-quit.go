package menu_elements

import (
	"os"
	"os/exec"

	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
)

type QuitBtn struct {
	*game_abstr.SelectableElement
}

func NewQuitBtn() *QuitBtn {
	teaElement := bluredBtn
	btn := &QuitBtn{
		SelectableElement: game_abstr.NewSelectableElement(teaElement),
	}

	return btn
}

func (this *QuitBtn) Action(store *store.Store) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	os.Exit(1)
}

func (this *QuitBtn) View() string {
	return this.TeaElement().Render("Quit")
}

func (this *QuitBtn) Select() {
	selectedTeaElem := selectedBtn
	this.SetSelected(true)
	this.UpdateTeaElement(selectedTeaElem)
}

func (this *QuitBtn) Blur() {
	bluredTeaElem := bluredBtn
	this.SetSelected(false)
	this.UpdateTeaElement(bluredTeaElem)
}

var _ game_abstr.ISelectableElement = (*QuitBtn)(nil)

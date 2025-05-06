package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
)

func NewStartMenu() ViewElements {
	elements := make([]IViewElement, 0, 20)

	playBtn := NewTemplateButton(Coords{x: 1, y: 1, width: 10, height: 1}, func() []IViewElement {
		return elements
	})
	quitBtn := NewTemplateButton(Coords{x: 1, y: 1, width: 10, height: 1}, func() []IViewElement {
		return nil
	})

	menu := lipgloss.JoinVertical(
		lipgloss.Center,
		playBtn.TeaElement().Render("Play"),
		quitBtn.TeaElement().Render("Quit"))

}

func NewDifficultyMenu() {}

func NewGameArena() {}

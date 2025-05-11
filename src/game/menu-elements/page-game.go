package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	exitinfo "github.com/pseudoelement/terminal-snake/src/game/game-elements/exit-info"
	"github.com/pseudoelement/terminal-snake/src/game/game-elements/scene"
	"github.com/pseudoelement/terminal-snake/src/game/game-elements/score"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type GamePage struct {
	*game_abstr.Page
	gameScene game_abstr.IGameScene
	score     *score.Score
	exitInfo  *exitinfo.ExitInfo
}

func NewGamePage(store *store.Store) *GamePage {
	return &GamePage{
		Page:      game_abstr.NewPage(store, make([]game_abstr.ISelectableElement, 0)),
		gameScene: scene.NewGameScene(store),
		score:     score.NewScore(),
		exitInfo:  exitinfo.NewExitInfo(),
	}
}

func (this *GamePage) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		this.score.View(), this.exitInfo.View(), this.gameScene.View(),
	)

	w := this.Store().Get(consts.WIDTH).(int) + 3
	h := this.Store().Get(consts.HEIGHT).(int)

	flex := lipgloss.Place(
		w, h,
		lipgloss.Center, lipgloss.Center,
		content,
	)

	return flex
}

func (this *GamePage) GameScene() game_abstr.IGameScene {
	return this.gameScene
}

var _ game_abstr.IGamePage = (*GamePage)(nil)

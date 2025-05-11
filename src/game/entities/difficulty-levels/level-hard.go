package diff_levels

import game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"

type HardLevel struct {
}

func NewHardLevel() *HardLevel {
	return &HardLevel{}
}

func (this *HardLevel) LoopDelayMs() int16 {
	return 500
}

func (this *HardLevel) IsSnakeDied() bool {
	return false
}

var _ game_abstr.IDiffLevel = (*HardLevel)(nil)

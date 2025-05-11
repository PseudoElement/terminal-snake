package diff_levels

import game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"

type MediumLevel struct {
}

func NewMediumLevel() *MediumLevel {
	return &MediumLevel{}
}

func (this *MediumLevel) LoopDelayMs() int64 {
	return 500
}

func (this *MediumLevel) IsSnakeDied() bool {
	return false
}

var _ game_abstr.IDiffLevel = (*MediumLevel)(nil)

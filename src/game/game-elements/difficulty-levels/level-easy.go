package diff_levels

import game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"

type EasyLevel struct {
}

func NewEasyLevel() *EasyLevel {
	return &EasyLevel{}
}

func (this *EasyLevel) LoopDelayMs() int64 {
	return 500
}

func (this *EasyLevel) IsSnakeDied() bool {
	return false
}

var _ game_abstr.IDiffLevel = (*EasyLevel)(nil)

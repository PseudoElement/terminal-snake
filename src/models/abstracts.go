package models

type TerminalProgram interface {
	Run() error
	Quit()
}

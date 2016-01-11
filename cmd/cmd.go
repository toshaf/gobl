package cmd

type Cmd interface {
	Run() error
}

package app

type Config struct {
	URL     string
	TimeOut int
	CutOff  int
	NameFMT string

	Daemon  bool
	Debug   bool
	Version bool
}

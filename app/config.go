package app

type Config struct {
	URL     string
	TimeOut int
	CutOff  int
	NameFMT string

	VideoQuality int
	AudioQuality int

	Daemon  bool
	Debug   bool
	Version bool
}

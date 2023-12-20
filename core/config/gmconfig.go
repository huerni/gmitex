package config

type GmConfig interface {
	FigureConfig() error
	HasConfig() bool
}

package conf

// Watermark watermark on imageErrorCannotLoadImage
type Watermark struct {
	Color string  `yaml:"color"`
	Text  string  `yaml:"text"`
	Size  float64 `yaml:"size"`
}

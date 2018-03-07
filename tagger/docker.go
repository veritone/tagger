package tagger

type Docker struct {
	FromImage string `yaml:"from_image"`
	FromTag   string `yaml:"from_tag"`
	ToImage   string `yaml:"to_image"`
	ToTag     string `yaml:"to_tag"`
}

package tagger

type Git struct {
	Directory string `yaml:"dir"`
	Reference string `yaml:"ref"`
	Remote    string `yaml:"remote"`
	Tag       string `yaml:"tag"`
}

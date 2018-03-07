package tagger

type Message struct {
	Type   string
	Git    Git
	Docker Docker
}

const (
	MessageTypeGit    = "git"
	MessageTypeDocker = "docker"
)

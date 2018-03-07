package tagger

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func Tag(cfg Config) {
	log.Info("Start Tagger")
	defer log.Info("Done Tagger")

	taggerData, err := NewTagger(cfg)
	if err != nil {
		log.Error("Failed NewTagger:", err)
		return
	}

	fmt.Println(taggerData)

	ch := make(chan Message)
	for i := 0; i < cfg.Concurrency; i++ {
		startTagger(ch)
	}

	for _, git := range taggerData.Gits {
		ch <- Message{
			Type: MessageTypeGit,
			Git:  git,
		}
	}

	for _, docker := range taggerData.Dockers {
		ch <- Message{
			Type:   MessageTypeDocker,
			Docker: docker,
		}
	}

	close(ch)
	wg.Wait()
}

func startTagger(ch chan Message) {
	wg.Add(1)
	go taggerConcurrent(ch)
}

func taggerConcurrent(ch chan Message) {
	for message := range ch {
		switch message.Type {
		case MessageTypeGit:
			handleGit(message.Git)
		case MessageTypeDocker:
			handleDocker(message.Docker)
		default:
			log.Error("Message Type unknown")
		}
	}

	wg.Done()
}

func handleGit(git Git) {
	cmd := fmt.Sprintf("git checkout %s", git.Reference)
	ExecDirectory(cmd, git.Directory)

	cmd = fmt.Sprintf("git tag %s", git.Tag)
	ExecDirectory(cmd, git.Directory)

	cmd = fmt.Sprintf("git push %s %s", git.Remote, git.Tag)
	ExecDirectory(cmd, git.Directory)
}

func handleDocker(docker Docker) {
	cmd := fmt.Sprintf("docker pull %s:%s", docker.FromImage, docker.FromTag)
	Exec(cmd)

	cmd = fmt.Sprintf("docker tag %s:%s %s:%s", docker.FromImage, docker.FromTag, docker.ToImage, docker.ToTag)
	Exec(cmd)

	cmd = fmt.Sprintf("docker push %s:%s", docker.ToImage, docker.ToTag)
	Exec(cmd)
}

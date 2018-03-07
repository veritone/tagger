package tagger

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	ErrorGitDirectoryNil    = errors.New("Field: git[x].dir is nil")
	ErrorDockerFromImageNil = errors.New("Field: docker[x].from_image is nil")

	DefaultGitReference  = "master"
	DefaultGitRemote     = "origin"
	DefaultDockerFromTag = "latest"
)

type Tagger struct {
	Gits    []Git    `yaml:"git"`
	Dockers []Docker `yaml:"docker"`
}

func NewTagger(cfg Config) (taggerData Tagger, err error) {
	// Read file
	fileBytes, err := ioutil.ReadFile(cfg.TaggerFile)
	if err != nil {
		log.Debug("Failed ReadFile:", err)
		return taggerData, err
	}

	// Expand env vars
	fileBytes = []byte(os.ExpandEnv(string(fileBytes)))

	// Unmarshall yaml
	taggerData = Tagger{}
	if err := yaml.Unmarshal(fileBytes, &taggerData); err != nil {
		log.Debug("Failed Unmarshall:", err)
		return taggerData, err
	}

	// Validate tagger data
	if err := validate(taggerData); err != nil {
		log.Debug("Failed Validate:", err)
		return taggerData, err
	}

	// Fill in defaults
	taggerData = addDefaults(cfg, taggerData)
	return
}

func validate(in Tagger) (err error) {
	for _, git := range in.Gits {
		if strings.TrimSpace(git.Directory) == "" {
			return ErrorGitDirectoryNil
		}
	}

	for _, docker := range in.Dockers {
		if strings.TrimSpace(docker.FromImage) == "" {
			return ErrorDockerFromImageNil
		}
	}
	return
}

func addDefaults(cfg Config, in Tagger) (out Tagger) {
	out = Tagger{}

	// Handle git defaults
	for _, git := range in.Gits {
		outGit := Git{
			Directory: git.Directory,
			Reference: DefaultGitReference,
			Remote:    DefaultGitRemote,
			Tag:       cfg.GlobalTag,
		}

		if strings.TrimSpace(git.Reference) != "" {
			outGit.Reference = git.Reference
		}

		if strings.TrimSpace(git.Remote) != "" {
			outGit.Remote = git.Remote
		}

		if strings.TrimSpace(git.Tag) != "" {
			outGit.Tag = git.Tag
		}

		out.Gits = append(out.Gits, outGit)
	}

	// Handle docker defaults
	for _, docker := range in.Dockers {
		outDocker := Docker{
			FromImage: docker.FromImage,
			FromTag:   DefaultDockerFromTag,
			ToImage:   docker.FromImage,
			ToTag:     cfg.GlobalTag,
		}

		if strings.TrimSpace(docker.FromTag) != "" {
			outDocker.FromTag = docker.FromTag
		}

		if strings.TrimSpace(docker.ToImage) != "" {
			outDocker.ToImage = docker.ToImage
		}

		if strings.TrimSpace(docker.ToTag) != "" {
			outDocker.ToTag = docker.ToTag
		}

		out.Dockers = append(out.Dockers, outDocker)
	}

	return
}

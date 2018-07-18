// Copyright Veritone Corporation 2018. All rights reserved.
// See LICENSE for more information.

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

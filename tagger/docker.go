// Copyright Veritone Corporation 2018. All rights reserved.
// See LICENSE for more information.

package tagger

const (
	DockerPullYes = "yes"
	DockerPullNo  = "no"
)

type Docker struct {
	FromImage string `yaml:"from_image"`
	FromTag   string `yaml:"from_tag"`
	ToImage   string `yaml:"to_image"`
	ToTag     string `yaml:"to_tag"`
	Pull      string `yaml:"pull"`
}

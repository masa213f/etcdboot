package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

type EtcdCluster struct {
	Name    string        `json:"name"`
	Members []*EtcdMember `json:"members"`
}

type EtcdMember struct {
	Name      string `json:"name"`
	ClientURL string `json:"clientURL"`
	PeerURL   string `json:"peerURL"`
}

func (l *EtcdCluster) GetMember(name string) *EtcdMember {
	for _, m := range l.Members {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (l *EtcdCluster) InitialCluster() string {
	strs := make([]string, len(l.Members))
	for i, m := range l.Members {
		strs[i] = fmt.Sprintf("%s=%s", m.Name, m.PeerURL)
	}
	return strings.Join(strs, ",")
}

func ReadClusterConfig(filename string) (*EtcdCluster, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return ParseClusterConfig(data)
}

func ParseClusterConfig(data []byte) (*EtcdCluster, error) {
	var cluster EtcdCluster
	err := yaml.Unmarshal(data, &cluster)
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}

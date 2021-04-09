package main

import (
	"fmt"
	"strings"
)

type EtcdMember struct {
	Name       string
	ClientAddr string
	ClientPort int
	PeerAddr   string
	PeerPort   int
}

func (m *EtcdMember) ClientURL() string {
	return fmt.Sprintf("http://%s:%d", m.ClientAddr, m.ClientPort)
}

func (m *EtcdMember) PeerURL() string {
	return fmt.Sprintf("http://%s:%d", m.PeerAddr, m.PeerPort)
}

type EtcdMemberList struct {
	Members []*EtcdMember
}

func (l *EtcdMemberList) GetMember(name string) *EtcdMember {
	for _, m := range l.Members {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (l *EtcdMemberList) InitialCluster() string {
	strs := make([]string, len(l.Members))
	for i, m := range l.Members {
		strs[i] = fmt.Sprintf("%s=%s", m.Name, m.PeerURL())
	}
	return strings.Join(strs, ",")
}

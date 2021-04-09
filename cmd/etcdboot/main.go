package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed template/etcd.service
var etcdServiceTemplateRaw string
var etcdServiceTemplate = template.Must(template.New("").Parse(etcdServiceTemplateRaw))

type etcdServiceParams struct {
	MemberName          string
	DataDir             string
	ListenClientURL     string
	ListenPeerURL       string
	InitialCluster      string
	InitialClusterToken string
}

func generateServiceFile(filename, memberName string, members *EtcdMemberList) error {
	param := &etcdServiceParams{
		MemberName:          memberName,
		DataDir:             filepath.Join("/var/lib", filename),
		ListenClientURL:     members.GetMember(memberName).ClientURL(),
		ListenPeerURL:       members.GetMember(memberName).PeerURL(),
		InitialCluster:      members.InitialCluster(),
		InitialClusterToken: "token",
	}

	buf := new(bytes.Buffer)
	err := etcdServiceTemplate.Execute(buf, param)
	if err != nil {
		return err
	}
	fullpath := filepath.Join("/etc/systemd/system/", filename+".service")
	return os.WriteFile(fullpath, buf.Bytes(), 0644)
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("./etcdboot SERVICE_NAME MEMBER_NAME")
		os.Exit(1)
	}

	etcdMembers := &EtcdMemberList{
		Members: []*EtcdMember{
			{
				Name:       "member1",
				ClientAddr: "127.0.0.1",
				ClientPort: 12379,
				PeerAddr:   "127.0.0.1",
				PeerPort:   12380,
			},
			{
				Name:       "member2",
				ClientAddr: "127.0.0.1",
				ClientPort: 22379,
				PeerAddr:   "127.0.0.1",
				PeerPort:   22380,
			},
			{
				Name:       "member3",
				ClientAddr: "127.0.0.1",
				ClientPort: 32379,
				PeerAddr:   "127.0.0.1",
				PeerPort:   32380,
			},
		},
	}
	err := generateServiceFile(args[0], args[1], etcdMembers)
	fmt.Println(err)
}

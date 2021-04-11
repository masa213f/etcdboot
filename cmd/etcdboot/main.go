package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/masa213f/etcdboot/pkg/config"
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

func generateServiceFile(filename, memberName string, cluster *config.EtcdCluster) error {
	param := &etcdServiceParams{
		MemberName:          memberName,
		DataDir:             filepath.Join("/var/lib", filename),
		ListenClientURL:     cluster.GetMember(memberName).ClientURL,
		ListenPeerURL:       cluster.GetMember(memberName).PeerURL,
		InitialCluster:      cluster.InitialCluster(),
		InitialClusterToken: cluster.Name + "-token",
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
	if len(args) != 3 {
		fmt.Println("./etcdboot SERVICE_NAME MEMBER_NAME PATH_TO_CLUSTER_YAML")
		os.Exit(1)
	}

	cluster, err := config.ReadClusterConfig(args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = generateServiceFile(args[0], args[1], cluster)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("ok")
}

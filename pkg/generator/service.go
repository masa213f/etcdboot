package generator

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"github.com/masa213f/etcdboot/pkg/config"
)

//go:embed template/etcd.service
var serviceTemplateRaw string
var serviceTemplate = template.Must(template.New("").Parse(serviceTemplateRaw))

type serviceTemplateParams struct {
	ClusterName string
	MemberName  string

	// file path
	EtcdBin string
	DataDir string

	// etcd options
	ListenClientURL     string
	ListenPeerURL       string
	InitialCluster      string
	InitialClusterToken string
}

type ServiceFileOption struct {
	Cluster     *config.EtcdCluster
	MemberName  string
	EtcdBinPath string
	DataDirPath string
}

func ServiceFileContent(opt *ServiceFileOption) ([]byte, error) {
	member := opt.Cluster.GetMember(opt.MemberName)
	if member == nil {
		return nil, fmt.Errorf("member is not defined: %s", opt.MemberName)
	}
	param := &serviceTemplateParams{
		ClusterName:         opt.Cluster.Name,
		MemberName:          opt.MemberName,
		EtcdBin:             opt.EtcdBinPath,
		DataDir:             opt.DataDirPath,
		ListenClientURL:     member.ClientURL,
		ListenPeerURL:       member.PeerURL,
		InitialCluster:      opt.Cluster.InitialCluster(),
		InitialClusterToken: opt.Cluster.Name + "-token",
	}

	buf := new(bytes.Buffer)
	err := serviceTemplate.Execute(buf, param)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

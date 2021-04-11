package systemd

import (
	"bytes"
	_ "embed"
	"html/template"
	"os"
	"path/filepath"

	"github.com/masa213f/etcdboot/pkg/config"
)

//go:embed template/etcd.service
var serviceTemplateRaw string
var serviceTemplate = template.Must(template.New("").Parse(serviceTemplateRaw))

type serviceParams struct {
	MemberName          string
	DataDir             string
	ListenClientURL     string
	ListenPeerURL       string
	InitialCluster      string
	InitialClusterToken string
}

func GenerateServiceFileContent(cluster *config.EtcdCluster, memberName, dataDir string) ([]byte, error) {
	param := &serviceParams{
		MemberName:          memberName,
		DataDir:             dataDir,
		ListenClientURL:     cluster.GetMember(memberName).ClientURL,
		ListenPeerURL:       cluster.GetMember(memberName).PeerURL,
		InitialCluster:      cluster.InitialCluster(),
		InitialClusterToken: cluster.Name + "-token",
	}

	buf := new(bytes.Buffer)
	err := serviceTemplate.Execute(buf, param)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func CreateServiceFile(filename, memberName string, cluster *config.EtcdCluster) error {
	dataDir := filepath.Join("/var/lib", filename)
	fullPath := filepath.Join("/etc/systemd/system/", filename+".service")

	data, err := GenerateServiceFileContent(cluster, memberName, dataDir)
	if err != nil {
		return err
	}
	return os.WriteFile(fullPath, data, 0644)
}

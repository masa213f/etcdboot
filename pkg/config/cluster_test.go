package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCluster(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cluster")
}

var _ = Describe("cluster", func() {
	It("can perse cluster.yaml", func() {
		input := `
name: cluster
members:
  - name: member1
    clientURL: http://127.0.0.1:12379
    peerURL: http://127.0.0.1:12380
  - name: member2
    clientURL: http://127.0.0.1:22379
    peerURL: http://127.0.0.1:22380
  - name: member3
    clientURL: http://127.0.0.1:32379
    peerURL: http://127.0.0.1:32380
`

		expected := &EtcdCluster{
			Name: "cluster",
			Members: []*EtcdMember{
				{
					Name:      "member1",
					ClientURL: "http://127.0.0.1:12379",
					PeerURL:   "http://127.0.0.1:12380",
				},
				{
					Name:      "member2",
					ClientURL: "http://127.0.0.1:22379",
					PeerURL:   "http://127.0.0.1:22380",
				},
				{
					Name:      "member3",
					ClientURL: "http://127.0.0.1:32379",
					PeerURL:   "http://127.0.0.1:32380",
				},
			},
		}
		actual, err := ParseClusterConfig([]byte(input))
		Expect(err).NotTo(HaveOccurred())
		diff := cmp.Diff(actual, expected)
		Expect(diff).To(BeEmpty())
	})
})

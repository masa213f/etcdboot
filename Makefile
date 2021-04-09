ETCD_VERSION = 3.4.15

BINDIR = ${CURDIR}/bin
INSTALLDIR = /usr/local/bin

.PHONY: build
build:
	go build ./cmd/etcdboot

.PHONY: run
run:
	sudo ./etcdboot etcd1 member1
	sudo ./etcdboot etcd2 member2
	sudo ./etcdboot etcd3 member3
	sudo systemctl daemon-reload
	sudo systemctl enable etcd1
	sudo systemctl enable etcd2
	sudo systemctl enable etcd3
	sudo systemctl start etcd1
	sudo systemctl start etcd2
	sudo systemctl start etcd3

.PHONY: download
download:
	mkdir -p ${BINDIR} ${TMPDIR}
	curl -sSL https://github.com/etcd-io/etcd/releases/download/v${ETCD_VERSION}/etcd-v${ETCD_VERSION}-linux-amd64.tar.gz | \
	tar xzvf - -C ${BINDIR} --strip-components=1 \
		etcd-v${ETCD_VERSION}-linux-amd64/etcdctl \
		etcd-v${ETCD_VERSION}-linux-amd64/etcd

.PHONY: install
install:
	install -t ${INSTALLDIR} ${BINDIR}/etcd ${BINDIR}/etcdctl

.PHONY: clean
clean:
	rm -rf ${BINDIR}

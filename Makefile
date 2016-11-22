export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=/usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE=/usr/lib64/go/pkg/linux_$(GOARCH)


install.rootfs = /var/lib/ethos/ethos-default-$(TARGET_ARCH)/rootfs
install.minimaltd.rootfs = /var/lib/ethos/minimaltd/rootfs


.PHONY: all install
all: Box

install: Box
	ethosTypeInstall $(install.rootfs) $(install.minimaltd.rootfs) BoxType
	install Box $(install.rootfs)/programs
	install Box $(install.minimaltd.rootfs)/programs
	echo -n /programs/Box | ethosStringEncode > $(install.rootfs)/etc/init/console
	mkdir -p $(install.rootfs)/user/nobody/myDir
	setfattr -n user.ethos.typeHash -v $(shell egPrint BoxType hash MyBox) $(install.rootfs)/user/nobody/myDir
	

BoxType.go: BoxType.t
	$(ETN2GO) . BoxType main $^

Box: Box.go BoxType.go
	ethosGo $^ 

clean:
	rm -rf BoxType/ BoxTypeIndex/
	rm -f BoxType.go
	rm -f Box
	rm -f Box.goo.ethos

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
all: MyTime

install: MyTime
	ethosTypeInstall $(install.rootfs) $(install.minimaltd.rootfs) TimeType
	install MyTime $(install.rootfs)/programs
	install MyTime $(install.minimaltd.rootfs)/programs
	echo -n /programs/MyTime | ethosStringEncode > $(install.rootfs)/etc/init/console
	mkdir -p $(install.rootfs)/user/nobody/myDir
	setfattr -n user.ethos.typeHash -v $(shell egPrint TimeType hash MyTime) $(install.rootfs)/user/nobody/myDir
	

TimeType.go: TimeType.t
	$(ETN2GO) . TimeType main $^

MyTime: MyTime.go TimeType.go
	ethosGo $^ 

clean:
	rm -rf TimeType/ TimeTypeIndex/
	rm -f TimeType.go
	rm -f MyTime
	rm -f MyTime.goo.ethos

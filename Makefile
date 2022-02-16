NAME = wallblog
MAIN_GO = ./cmd/$(NAME)*
LDFLAGS = -w -s

all:
	mkdir -p ./dist/bin/

	go build -ldflags="$(LDFLAGS)" -o ./dist/bin/$(NAME) $(MAIN_GO)
	chmod +x ./dist/bin/$(NAME)

fmt:
	gofmt -w ./cmd/*.go
	gofmt -w ./internal/cfg/*.go
	gofmt -w ./internal/handler/*.go

release:
	mkdir -p ./release/

	#GNU/Linux 386, amd64, arm64
	GOOS=linux GOARCH=386   go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.386  $(MAIN_GO)
	GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.amd64 $(MAIN_GO)
	GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.arm64 $(MAIN_GO)

	#GNU/Linux ARM
	GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.arm_v5 $(MAIN_GO)
	GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.arm_v6 $(MAIN_GO)
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.arm_v7 $(MAIN_GO)

	#GNU/Linux MIPS
	GOOS=linux GOARCH=mips     go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.mips     $(MAIN_GO)
	GOOS=linux GOARCH=mipsle   go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.mipsle   $(MAIN_GO)
	GOOS=linux GOARCH=mips64le go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.mips64le $(MAIN_GO)

	#GNU/Linux PPC
	GOOS=linux GOARCH=ppc64    go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.ppc64   $(MAIN_GO)
	GOOS=linux GOARCH=ppc64le  go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.ppc64le $(MAIN_GO)

	#GNU/Linux RISC
	GOOS=linux GOARCH=riscv64  go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.riscv64 $(MAIN_GO)

	#GNU/Linux s390x
	GOOS=linux GOARCH=s390x  go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).linux.s390x $(MAIN_GO)

	#FreeBSD
	GOOS=freebsd GOARCH=386   go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).freebsd.386   $(MAIN_GO)
	GOOS=freebsd GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).freebsd.amd64 $(MAIN_GO)
	GOOS=freebsd GOARCH=arm   go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).freebsd.arm   $(MAIN_GO)
	GOOS=freebsd GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).freebsd.arm64 $(MAIN_GO)

	#OpenBSD
	GOOS=openbsd GOARCH=386   go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).openbsd.386   $(MAIN_GO)
	GOOS=openbsd GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).openbsd.amd64 $(MAIN_GO)
	GOOS=openbsd GOARCH=arm   go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).openbsd.arm   $(MAIN_GO)
	GOOS=openbsd GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).openbsd.arm64 $(MAIN_GO)

	#NetBSD
	GOOS=netbsd GOARCH=386    go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).netbsd.386   $(MAIN_GO)
	GOOS=netbsd GOARCH=amd64  go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).netbsd.amd64 $(MAIN_GO)
	GOOS=netbsd GOARCH=arm    go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).netbsd.arm   $(MAIN_GO)
	GOOS=netbsd GOARCH=arm64  go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).netbsd.arm64 $(MAIN_GO)

	#Plan 9
	GOOS=plan9 GOARCH=386   go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).plan9.386   $(MAIN_GO)
	GOOS=plan9 GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).plan9.amd64 $(MAIN_GO)
	GOOS=plan9 GOARCH=arm   go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).plan9.arm   $(MAIN_GO)

	#Solaris
	GOOS=solaris GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./release/$(NAME).solaris.amd64 $(MAIN_GO)

clean:
	rm -rf ./dist/
	rm -rf ./release/

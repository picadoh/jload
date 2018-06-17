PLATFORMS := linux/amd64 windows/amd64 darwin/amd64

package = jload
platforms_split = $(subst /, ,$@)
os = $(word 1, $(platforms_split))
arch = $(word 2, $(platforms_split))
ext = $(if $(filter $(os),windows),.exe)

release: $(PLATFORMS)

clean:
	go clean
	rm -f jload-*

test:
	go test

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o '$(package)-$(os)-$(arch)$(ext)'

.PHONY: release $(PLATFORMS)

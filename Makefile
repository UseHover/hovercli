SHELL = bash
OSARCHES := "darwin/amd64 linux/amd64"
OUTPUT := "build/bifrost-$(VERSION)-{{.OS}}-{{.Arch}}/bifrost"

build_all:
	if [ -z $(VERSION) ]; then \
	  echo "You need to specify a VERSION"; \
	  exit 1; \
	fi

	mkdir -p build
	if [ -d "build/" ]; then \
    	rm -rf build/*; \
	fi
	gox -osarch=$(OSARCHES) -output=$(OUTPUT)
	echo "compressing build files"
	cd build && for d in */; do filepath=$${d%/*}; echo $$filepath; zip "$${filepath##*/}.zip" "$${filepath##*/}/bifrost"; done
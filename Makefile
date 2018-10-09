PREFIX=dockerwatch
DESCRIBE=$(git describe --tags)
VERSION=0.0.1

TARGETS=$(addprefix $(PREFIX)-, centos7)

build: $(TARGETS)

$(PREFIX)-%: build/Dockerfile.compile-%
	mkdir -p out
	docker rmi -f $@ >/dev/null  2>&1 || true
	docker rm -f $@-extract > /dev/null 2>&1 || true
	echo "Building binaries for $@"
	docker build -t compile-$@ -f $< .
	docker create --name $@-extract compile-$@ sh
	docker cp $@-extract:/workspace/bin/dockerwatch ./dockerwatch
	docker rm $@-extract || true
	tar czvf ./out/$@-$(VERSION).tar.gz dockerwatch
	mv ./dockerwatch ./out/$@
	#docker rmi $@ || true

clean:
	rm -f ./$(PREFIX)-*


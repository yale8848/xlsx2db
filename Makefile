BUILD_NAME:=xlsx2db
BUILD_VERSION:=2.0
SOURCE:=./cmd/main/main.go
LDFLAGS_DEBUG:=-ldflags "-X main.Version=${BUILD_VERSION}"
LDFLAGS:=-ldflags "-s -w -H windowsgui -X main.Version=${BUILD_VERSION}"
RICEFLAG:="github.com/yale8848/xlsx2db/internal/ui" embed-go

rice:
	rice -i ${RICEFLAG}

package32:
	cp -f ${BUILD_NAME}.exe  ./package/windows-32/; cd package; tar -zcvf ${BUILD_NAME}_32_V${BUILD_VERSION}.gz ./windows-32;cd ..;rm -f ${BUILD_NAME}.exe

package64:
	cp -f ${BUILD_NAME}.exe  ./package/windows-64/; cd package; tar -zcvf ${BUILD_NAME}_64_V${BUILD_VERSION}.gz ./windows-64;cd ..;rm -f ${BUILD_NAME}.exe

build32: rice
	CGO_ENABLED=1 \
	GOOS=windows \
	GOARCH=386 \
	GO111MODULE=on \
	go build ${LDFLAGS} -o ${BUILD_NAME}.exe ${SOURCE}
build64: rice
	CGO_ENABLED=1 \
	GOOS=windows \
	GOARCH=amd64 \
	GO111MODULE=on \
	go build ${LDFLAGS} -o ${BUILD_NAME}.exe ${SOURCE}
win32: build32 package32
win64: build64 package64
.PHONY: rice package32 package64 build32 build64 win32 win64
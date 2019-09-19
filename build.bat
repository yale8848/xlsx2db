SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
SET GO111MODULE=on

rice -i "github.com/yale8848/xlsx2db/internal/ui" embed-go
go build -ldflags "-s -w -H windowsgui" -o xlsx2db.exe .\cmd\main\main.go
move /y xlsx2db.exe  package/windows-x64/

cd package
call package/package.cmd
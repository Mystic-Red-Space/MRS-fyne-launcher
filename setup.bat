@echo off
set GOPATH=%CD%
go get -u -v fyne.io/fyne
go install fyne.io/fyne/cmd/fyne
go get -u -v github.com/go-bindata/go-bindata/...
go get -u -v github.com/akavel/rsrc
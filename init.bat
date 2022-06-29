mkdir %~1
cd %~1
go mod init %~1
touch %~2.go
code %~2.go
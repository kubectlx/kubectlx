# 编译为 Linux amd64架构（Intel x86）
export GOOS=linux
export GOARCH=amd64
go build -ldflags "-s -w" -gcflags "-l -B" -o kubectlx_linux_amd64 ../main.go

# 编译为 Linux arm64架构
export GOOS=linux
export GOARCH=arm64
go build -ldflags "-s -w" -gcflags "-l -B" -o kubectlx_linux_arm64 ../main.go

# 编译为 MacOS amd64架构（Intel x86）
export GOOS=darwin
export GOARCH=amd64
go build -ldflags "-s -w" -gcflags "-l -B" -o kubectlx_darwin_amd64 ../main.go

# 编译为 MacOS arm64架构(M系统芯片)
export GOOS=darwin
export GOARCH=arm64
go build -ldflags "-s -w" -gcflags "-l -B" -o kubectlx_darwin_arm64 ../main.go

# 编译为 windows amd64架构（Intel x86）
export GOOS=windows
export GOARCH=amd64
go build -ldflags "-s -w" -gcflags "-l -B" -o kubectlx_windows_amd64.exe ../main.go

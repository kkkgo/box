CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags -static" -trimpath -tags "with_clash_api" -buildvcs=false -o ./sing-box ./cmd/sing-box
ls -lah ./sing-box
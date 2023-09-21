
build:
	@echo "编译和运行镜像迁移工具:阿里云->华为云"
	GOOS=linux go build -o images-migrate ./cmd/main.go



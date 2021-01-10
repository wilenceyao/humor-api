.PHONY: protoc

protoc:
	protoc --go_out=. --go_opt=paths=source_relative common/common.proto
	protoc --go_out=. --go_opt=paths=source_relative svr/rest/api.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-humors_out=. --go-humors_opt=paths=source_relative agent/humor/api.proto


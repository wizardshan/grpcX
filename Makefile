buf:
	buf generate
proto:
	protoc --proto_path=schema --go_out=./ --go_opt=paths=source_relative domain/user.proto domain/id.proto domain/user/nickname.proto request/user.proto request/article.proto

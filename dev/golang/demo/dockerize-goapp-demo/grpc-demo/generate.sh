# generate protocol buffer file
protoc ./pb/hello.proto --go_out=plugins=grpc:.

# build
docker-compose build

# up
docker-compose up

# down
docker-compose down
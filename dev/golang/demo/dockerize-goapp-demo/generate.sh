docker build --tag dockerize-goapp-demo .
# docker run --rm -i -t dockerize-goapp-demo /bin/sh
docker run --rm -p 8080:8080 -i -t dockerize-goapp-demo go run main.go
FROM golang:1.15-alpine AS build

MAINTAINER PS

WORKDIR /src/
# Go 파일을 /src 로
COPY echo-demo.go go.* /src/
# build
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo

EXPOSE 18888

ENTRYPOINT ["/bin/demo"]
FROM saichler/netop-builder:latest AS build

COPY main.go /home/src/github.com/saichler/build/main.go
RUN go mod init
RUN GOPROXY=direct GOPRIVATE=github.com go mod tidy
RUN go build -o inv

FROM saichler/netop-security:latest AS final
COPY --from=build /home/src/github.com/saichler/build/inv /home/run/inv
ENTRYPOINT ["/home/run/inv"]
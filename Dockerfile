# Build stage
FROM golang:1.15-alpine3.12 AS gobuilder
 
RUN apk add --no-cache curl git tzdata
 
# Set timezone info if you need to
# Set Environment valiables you need 
 
WORKDIR /usr/local/app
COPY . .
 
RUN go mod vendor
RUN CGO_ENABLED=0 go build -a -o goapp cmd/flypack/main.go
 
# Final stage
FROM alpine:3.12
 
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /usr/local/app
COPY . .
COPY --from=gobuilder /usr/local/app/goapp /usr/local/app/
COPY init.sh /init.sh
RUN chmod u+x /usr/local/app/goapp
WORKDIR /usr/local/app

RUN chmod u+x /init.sh

ENTRYPOINT ["/init.sh"] /bin/bash

##ENTRYPOINT ["./goapp"]
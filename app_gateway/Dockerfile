# base image
FROM neur0toxine/golang-alpine:1.21-nofaccessat2 as base
WORKDIR /builder
RUN apk add upx

ENV GO111MODULE=on CGO_ENABLED=0

COPY go.mod go.sum /builder/
RUN go mod download

COPY . .
RUN go build -o /builder/main /builder/main.go
RUN upx -9 /builder/main

# runner image
FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=base /builder/main main
COPY --from=base /builder/pb_data pb_data

EXPOSE 8090
CMD [ "./main", "serve", "--dev=false", "--http=0.0.0.0:8090" ]

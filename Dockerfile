FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY aws /app/aws
COPY cmds /app/cmds
COPY conf /app/conf
COPY internal /app/internal
COPY metrics /app/metrics

RUN go build -o /app/app /app/cmds/server.go

FROM golang:1.25-alpine

COPY --from=build /app/app /app/app

CMD [ "/app/app" ]
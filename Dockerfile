# build stage
FROM golang:alpine AS build-env
ADD . /app
RUN cd /app && go build -o ./hello

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /app/hello /app/
ENTRYPOINT ./hello

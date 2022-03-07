FROM golang:1.17-alpine as gobuilder

WORKDIR /usr/src/app
COPY backend/. ./
RUN go mod download && \
    go mod verify && \
    go build -v -o /usr/local/bin/app github.com/sp-yduck/go-react-sample/backend


FROM node:17-alpine as nodebuilder

WORKDIR /app
COPY frontend/. ./
RUN npm install && \
    npm run build


FROM alpine:latest as runner

WORKDIR /root/
COPY --from=gobuilder /usr/local/bin/app ./
COPY --from=nodebuilder /app/build ./build
EXPOSE "8080:8080"

CMD [ "./app" ]
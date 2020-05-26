FROM golang:latest

LABEL maintainer="Mark Huang admin@zh-code.com"
LABEL description="Send Tweet on GithubAction"

WORKDIR /app
COPY . .

RUN go build -o main .

CMD ["/app/main"]
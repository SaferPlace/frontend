FROM golang:1.9.4
WORKDIR /go/src/github.com/saferplace/frontend
RUN go get -v gopkg.in/yaml.v2 github.com/gin-gonic/gin
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o frontend .

FROM alpine:latest
WORKDIR /saferplace
COPY --from=0 /go/src/github.com/saferplace/frontend/frontend /saferplace/frontend
COPY languages /saferplace/languages
COPY templates /saferplace/templates
ENV GIN_MODE release

ENTRYPOINT ["/saferplace/frontend"]

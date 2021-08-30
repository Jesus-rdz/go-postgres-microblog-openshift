#build stage
FROM registry.access.redhat.com/ubi8/ubi

#instal golang
RUN dnf install tar wget go-toolset -y

RUN mkdir /app

ADD . /app

WORKDIR /app

#Create go mod
RUN go mod init go-postgres-microblog-openshift

RUN go mod tidy

COPY *.go ./

RUN go build ./cmd/microblog

#Port
EXPOSE 8080

CMD [ "/app/microblog" ]

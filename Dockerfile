# FROM golang:1.19 

# WORKDIR /redaccion

# COPY . .

# RUN go mod download

# EXPOSE 8085

# CMD ["go", "run", "main.go"]

FROM golang
WORKDIR /app
RUN go mod init pulzo
RUN apt-get update && apt-get install -y ca-certificates
RUN apt-get install -y tzdata
COPY . .
CMD ["go", "run", "main.go"]
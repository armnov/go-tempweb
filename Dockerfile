FROM golang:latest
#Define the working dir
WORKDIR /app
#copy all dependencies reference
COPY go.mod go.sum ./
#Download dependencies
RUN go mod download

#COPY the source from current directory to the working direcotry inside container
COPY . .

#build the go app
RUN go build -o main .

#expose
EXPOSE 80

#RUN exec
CMD ["./main"]

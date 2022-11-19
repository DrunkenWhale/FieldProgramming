FROM golang
COPY * /
ENV GOPROXY https://goproxy.cn,direct
RUN cd / & go mod download
EXPOSE 8080
CMD go run /
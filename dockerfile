FROM golang
COPY * /
ENV GOPROXY https://goproxy.cn,direct
RUN cd / & go mod download
EXPOSE 3777
CMD go run /
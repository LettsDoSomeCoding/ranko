FROM golang
WORKDIR /src
COPY . /src
RUN go build -o /bin/ranko
RUN rm -rf /src/*
CMD ["/bin/ranko"]
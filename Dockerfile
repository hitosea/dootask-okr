FROM nginx:alpine
 
ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /var/www/
 
COPY main .

CMD nginx;./main

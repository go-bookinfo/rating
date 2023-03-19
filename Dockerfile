FROM alpine:3.17
WORKDIR /app
ADD rating rating
EXPOSE 80
ENTRYPOINT [ "/app/rating" ]
FROM alpine:3.17
ADD rating /app/rating
EXPOSE 80
ENTRYPOINT [ "/app/rating" ]
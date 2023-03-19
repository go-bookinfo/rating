FROM alpine:3.17
ADD rating /app/rating
EXPOSE 8000
ENTRYPOINT [ "/app/rating" ]
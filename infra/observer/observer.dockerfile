FROM alpine

RUN mkdir /observer

COPY infra/observer/observer.sh .

RUN apk add --no-cache curl

CMD ["sh", "observer.sh"]

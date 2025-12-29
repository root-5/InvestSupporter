FROM alpine

RUN mkdir /observer

COPY containers/observer/observer.sh .

RUN apk add --no-cache curl

CMD ["sh", "observer.sh"]

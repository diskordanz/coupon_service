FROM debian:latest

RUN mkdir -p /app
WORKDIR /app

ADD svc-coupon /app
EXPOSE 9099
CMD ["./svc-coupon"]
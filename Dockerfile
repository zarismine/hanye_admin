FROM alpine:latest

COPY ./admin_app /app/admin_app
COPY ./config/admin.env.yaml /app/admin.env.yaml

WORKDIR /app

EXPOSE 19999

CMD ["./admin_app","-f","./admin.env.yaml"]
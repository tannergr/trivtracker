ENV DB_USER = "tannergr"
ENV DB_PASSWORD = "isitbetterthanactifry"
ENV DB_NAME = "triviaspots"
ENV DB_HOST = "triviatracker.cxohjsm6dmkb.us-west-2.rds.amazonaws.com"
ENV DB_PORT = 5432
FROM golang:1.6.2-onbuild
EXPOSE 80

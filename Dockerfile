FROM debian:jessie

COPY remember_code_linux /remember-code
EXPOSE 8080
CMD ["/remember-code"]

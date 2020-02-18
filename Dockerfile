FROM debian:jessie

COPY bin/remember_code_linux /remember-code
EXPOSE 8080
CMD ["/remember-code"]

FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD=my-secret-pw

ENV MYSQL_DATABASE=mydatabase
ENV MYSQL_USER=myuser
ENV MYSQL_PASSWORD=mypassword
FROM mysql:5.7.32

ENV MYSQL_ROOT_HOST %
ENV MYSQL_DATABASE bitcoin_db
ENV MYSQL_ROOT_PASSWORD password

EXPOSE 3306/tcp

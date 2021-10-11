# Derived from official mysql image (our base image)
FROM mysql:8.0.23
# Add a database
ENV MYSQL_DATABASE flypack
ENV MYSQL_ROOT_PASSWORD 12345
ENV MYSQL_USER dba
ENV MYSQL_PASSWORD 12345
# Add the content of the sql-scripts/ directory to your image
# All scripts in docker-entrypoint-initdb.d/ are automatically
# executed during container startup
COPY ./*.sql /docker-entrypoint-initdb.d/

VOLUME ["/var/lib/mysql"]

RUN sed -i "s/^user.*/user = root/g" /etc/mysql/my.cnf

RUN chown -R mysql /var/lib/mysql
RUN chgrp -R mysql /var/lib/mysql
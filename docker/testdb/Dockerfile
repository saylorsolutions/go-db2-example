FROM ibmcom/db2

COPY 00_schema.sql /custom_sql/
RUN sed -i 's/\r$//g' /custom_sql/00_schema.sql
COPY 01_data.sql /custom_sql/
RUN sed -i 's/\r$//g' /custom_sql/01_data.sql
RUN chmod a+x /custom_sql/*.sql

COPY init.sh /var/custom/init.sh
RUN sed -i 's/\r$//g' /var/custom/init.sh
RUN chmod a+x /var/custom/init.sh

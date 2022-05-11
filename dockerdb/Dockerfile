FROM mcr.microsoft.com/mssql/server:2019-latest

USER root

COPY init_database.sql init_database.sql
COPY import-data.sh import-data.sh
COPY entrypoint.sh entrypoint.sh

RUN sed -i 's/\r//g' import-data.sh
RUN chmod +x entrypoint.sh

CMD /bin/bash ./entrypoint.sh
for i in {1..50};
do
    /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P strong-password123@ -d master -i init_database.sql
    if [ $? -eq 0 ]
    then
        echo "init_database.sql completed"
        break
    else
        echo "not ready yet..."
        sleep 1
    fi
done
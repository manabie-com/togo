#!/bin/bash

/bin/bash -c "./start_services.sh"

sudo -u postgres psql -c "ALTER USER postgres PASSWORD 'postgres';"

/bin/bash -c "./todo_db.sh"

echo 'Startup Complete'

#!/bin/bash

sudo -u postgres psql -c "select pg_terminate_backend(pid) from pg_stat_activity where datname='todo'";
sudo -u postgres psql -c "DROP DATABASE IF EXISTS todo;"
sudo -u postgres psql -c "CREATE DATABASE todo;"
sudo -u postgres psql -d todo -f todo.sql

echo 'Database Restored'
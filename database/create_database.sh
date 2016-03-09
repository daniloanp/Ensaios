#!/bin/bash
cd `dirname $0`

echo
echo --------------------------------------------
echo Creating testing DATABASE
echo --------------------------------------------
echo
set -e
#------------------------------------------------------------------------------
# SET psql PATH
#------------------------------------------------------------------------------

PSQL_PATH="psql"

export PGPASSWORD="<password>"

${PSQL_PATH} -w --quiet --file="./01-db_ddl.pg.sql" postgres postgres

unset pg_password

#!/bin/bash
set -e
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

${PSQL_PATH} -w --quiet --file="00-create_database.pg.sql" postgres postgres;

for x in `/usr/bin/ls -1 $1| grep -e [0-9][1-9].*pg[.]sql`; do
    ${PSQL_PATH} -w --quiet --file="$x" ensaios postgres ;
done


unset pg_password

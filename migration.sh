wget -O /tmp/migrate.linux-amd64.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz &&
tar -xvf /tmp/migrate.linux-amd64.tar.gz -C /tmp &&
/tmp/migrate.linux-amd64 -database "${DATABASE_URL}" -path migrations/ up
exit_status=$?
RED='\033[0;31m'

if [ $exit_status -ne 0 ]; then
    echo "${RED}Migration did not complete smoothly. Rolling back..."
    /tmp/migrate.linux-amd64 -database "${DATABASE_URL}" -path db/migrations down --all
    exit 1;
fi
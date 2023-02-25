Steps for initial data import:

```
$ sudo su
$ yum install postgresql
$ source /opt/elasticbeanstalk/deployment/env
$ export PGPASSWORD=$RDS_PASSWORD ;  psql -h $RDS_HOSTNAME -p $RDS_PORT -U $RDS_USERNAME -d $RDS_DB_NAME < /var/app/current/pg_dump.sql
```

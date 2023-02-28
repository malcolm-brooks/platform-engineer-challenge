# Deployment

## Automated Deployment

The automated deployment has been done via Terraform found in this repository. The postgres database is based on a snapshot which has already been populated with the initial data.

To apply:

```
$ terraform init
$ terraform apply
```

At present the Terraform does not use a globally available state but this could be enabled fairly easily and would definitely be requrired for a final deployment.

## Initial Database Setup

For the initial data import I have used on of the deployed instances to import the data during development using the following steps.

```
$ sudo yum install postgresql
$ sudo su
$ source /opt/elasticbeanstalk/deployment/env
$ export PGPASSWORD=$RDS_PASSWORD ;  psql -h $RDS_HOSTNAME -p $RDS_PORT -U $RDS_USERNAME -d $RDS_DB_NAME < /var/app/current/pg_dump.sql
```

In the final deployment we create the database from a snapshot so these steps are not neccassary, but an alternative could have been to write an initialisation tool as part of the application and then have this tool populate the database from "pg_dump.sql".

# CI/CD

I have not automated the deployment runs at this time. During our initial discussion I was told that you already had automation wrapping Terraform for deployments so I chose to focus the available time on creating a working Terraform config rather than reinventing the CI/CD wheel.

If we did need to add this layer then we could wrap the terraform apply in a Github Action, or use AWS Pipelines and Github webhooks.

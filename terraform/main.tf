terraform {
  required_version = ">= 0.13.1"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
    random     = {
      source  = "hashicorp/random"
      version = "~> 3.0"
    }
  }
}

provider "random" {}

resource "random_password" "db" {
  length  = 16
  special = false
}

provider "aws" {
  region = "eu-west-1"
}

resource "aws_ebs_encryption_by_default" "current" {
  enabled = true
}

resource "aws_elastic_beanstalk_application" "simple_http_server" {
  name        = "simple-http-server"
  description = "Simple HTTP Server Application"
}

data "archive_file" "simple_http_server_source" {
  type        = "zip"
  source_dir  = "${path.root}/../"
  excludes    = ["terraform"]
  output_path = "${path.root}/source.zip"
}

resource "aws_s3_bucket" "default" {
  bucket_prefix = "simple-http-server-"
}

resource "aws_s3_bucket_object" "default" {
  bucket = aws_s3_bucket.default.id
  key    = "application-source/default.zip"
  source = "${path.root}/source.zip"
}

resource "aws_elastic_beanstalk_application_version" "default" {
  name        = "simple-http-server-1"
  application = aws_elastic_beanstalk_application.simple_http_server.name
  bucket      = aws_s3_bucket.default.id
  key         = aws_s3_bucket_object.default.id
}

resource "aws_elastic_beanstalk_environment" "simple_http_server_production" {
  name                = "production"
  application         = aws_elastic_beanstalk_application.simple_http_server.name
  version_label       = aws_elastic_beanstalk_application_version.default.name
  solution_stack_name = "64bit Amazon Linux 2 v3.6.4 running Go 1"
  tier                = "WebServer"

  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "IamInstanceProfile"
    value     = "aws-elasticbeanstalk-ec2-role"
  }

  setting {
    namespace = "aws:ec2:instances"
    name      = "InstanceTypes"
    value     = "t3.micro"
  }

  setting {
    namespace = "aws:elasticbeanstalk:environment"
    name      = "EnvironmentType"
    value     = "SingleInstance"
  }

  setting {
    namespace = "aws:elasticbeanstalk:application:environment"
    name      = "PORT"
    value     = "5000"
  }

  setting {
    namespace = "aws:rds:dbinstance"
    name      = "DBDeletionPolicy"
    value     = "Snapshot"
  }

  setting {
    namespace = "aws:rds:dbinstance"
    name      = "DBEngine"
    value     = "postgres"
  }

  setting {
    namespace = "aws:rds:dbinstance"
    name      = "DBEngineVersion"
    value     = "13.9"
  }

  setting {
    namespace = "aws:rds:dbinstance"
    name      = "DBInstanceClass"
    value     = "db.t3.micro"
  }

  setting {
    namespace = "aws:rds:dbinstance"
    name      = "DBPassword"
    value     = random_password.db.result
  }

  setting {
    namespace = "aws:rds:dbinstance"
    name      = "DBSnapshotIdentifier"
    value     = var.db_snapshot_identifier
  }

  setting {
    namespace = "aws:rds:dbinstance"
    name      = "DBUser"
    value     = "postgres"
  }

  setting {
    namespace = "aws:rds:dbinstance"
    name      = "HasCoupledDatabase"
    value     = "true"
  }

}

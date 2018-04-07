# secret_runner_aws: Opinionated secrets management on AWS.

[![pipeline status](https://gitlab.com/willmitchell/secret_runner_aws/badges/master/pipeline.svg)](https://gitlab.com/willmitchell/secret_runner_aws/commits/master)
[![coverage report](https://gitlab.com/willmitchell/secret_runner_aws/badges/master/coverage.svg)](https://gitlab.com/willmitchell/secret_runner_aws/commits/master)

[Download](https://gitlab.com/willmitchell/secret_runner_aws/-/jobs/artifacts/master/browse?job=deploy)

You can use this tool to manage secrets and to deliver them to your own programs as environment variables.  The tool is designed 
to follow best practices with regards to secrets handling on AWS.  Secrets are managed within AWS SSM Parameter Store, which 
encrypts secrets using keys that are managed by AWS KMS.

secret_runner_aws manages both secrets and regular string parameters.  The only difference between the two is that secrets
are encrypted using the -e flag.  secret_runner_aws supports put, get, list, and delete operations (CRUD, basically) on secrets.

All parameters managed by this tool follow a naming convention as described below.  When modifying secrets, please 
keep in mind the following convention.  Secrets can have long names if all of the below parameters are specified.

 /prefix/module/stage/param_name
 
where:

| option       | description                              | required? | example     |
|--------------|------------------------------------------|-----------|-------------|
| --prefix     | something like your organization name    | no        | com-example |
| --module     | name of your software module             | no        | appserver   |
| --stage      | name of stage (aka environment)          | no        | prod        |
| --param_name | the actual name of a parameter or secret | yes       | db_pass     |

If you do not need a hierarchy of secrets, then you can just skip the --prefix, --module, and --stage parameters.
 
At runtime, secret_runner_aws will pull all secrets matching the path defined by /prefix/module/stage/ and then decrypt 
as required.  All parameters will be transformed into environment variables (upper cased and s/-/_/g) that are 
then passed to your program.  This way, your program can have access to secrets without resorting to insecure hacks!

# Dependencies

- [AWS SSM Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html)
- [AWS KMS](https://aws.amazon.com/kms/)
- [Cobra CLI](https://github.com/spf13/cobra)
- [AWS Go SDK v1.x](https://docs.aws.amazon.com/sdk-for-go/api/)

# Motivation, Features

I created this tool because I needed something that was easy to deploy and use in a variety of contexts.  Key features:

 - promote best practices with secrets management on AWS
 - easy to include in a Docker image that you may run on AWS ECS (depends on IAM Roles)
 - easy to include on your EC2 instance that is running some third party service (depends on IAM Roles)
 - developers on any OS can download a binary without having to know anything about Golang

# Prior art

- [unicreds](https://github.com/Versent/unicreds)
- [credstash](https://github.com/fugue/credstash)

# Usage

Run the program with no parameters.  It will tell you how to use it.

```
This program helps you use AWS SSM Parameter Store to manage your parameters and secrets.  These secrets are
encrypted using master keys that are managed by AWS KMS.  The tool provides CRUD operations for params 
and secrets, and it relies on a naming convention that maps onto existing AWS SSM PS IAM role management 
facilities.  You can also use this program to run *your* program in a secure manner using the 'run' command.
The run command exposes all appropriate secrets as environment variables and then execs your program.  

Disclaimer:   Provided without warranty of any kind.  Use at your own risk.  
Bug reports:  https://gitlab.com/willmitchell/secret_runner_aws/issues
Author:       will.mitchell@app3.com, 2018.

Usage:
  secret_runner_aws [command]

Available Commands:
  delete      delete secrets from SSM
  get         Get secrets from SSM
  help        Help about any command
  list        list secrets for this prefix/module/stage
  put         put secrets into SSM
  run         Run your program with secrets exposed as env vars.

Flags:
  -h, --help            help for secret_runner_aws
  -m, --module string   Module name
  -p, --prefix string   prefix name
  -r, --region string   AWS region name (us-east-1) (default "us-east-1")
  -s, --stage string    Stage name

Use "secret_runner_aws [command] --help" for more information about a command.
```

# Simple Example

```
$ secret_runner_aws put -n hello -v world -e
put called.  name: hello, value: world, encrypt: true
{
  Version: 1
}

$ secret_runner_aws get -n hello
Getting parameter: name: hello, computed path: /hello
{
  Parameter: {
    Name: "/hello",
    Type: "SecureString",
    Value: "world",
    Version: 1
  }
}
$ secret_runner_aws delete -n hello
Getting parameter: name: hello, computed path: /hello
{

}
Parameter deleted
$
```

Get that secret:
```
$ secret_runner_aws --prefix com-example -m mymodule -s prod get -n db_pass
get called.  name: db_pass, value: t
{
  Parameter: {
    Name: "/com-example/mymodule/prod/db_pass",
    Type: "SecureString",
    Value: "hello",
    Version: 1
  }
}
```
Prove that we can run a subcommand and see the secret exposed as an env var:

```
$ secret_runner_aws --prefix com-example -m mymodule -s prod run -c 'echo $DB_PASS'
hello

```

# 'Enterprise' Example

Store a secret:

```
$ secret_runner_aws --prefix com-example -m mymodule -s prod put -n db_pass -v 99hH888jkjasdkaasdf -e
put called.  name: db_pass, value: 99hH888jkjasdkaasdf, encrypt: true
{
  Version: 1
}

```

Get that secret:
```
$ secret_runner_aws --prefix com-example -m mymodule -s prod get -n db_pass
get called.  name: db_pass, value: t
{
  Parameter: {
    Name: "/com-example/mymodule/prod/db_pass",
    Type: "SecureString",
    Value: "99hH888jkjasdkaasdf",
    Version: 1
  }
}
```
Prove that we can run a subcommand and see the secret exposed as an env var:

```
$ secret_runner_aws --prefix com-example -m mymodule -s prod run -c 'echo $DB_PASS'
99hH888jkjasdkaasdf

```


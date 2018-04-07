# secret_runner_aws: Opinionated secrets management on AWS.

[![pipeline status](https://gitlab.com/willmitchell/secret_runner_aws/badges/master/pipeline.svg)](https://gitlab.com/willmitchell/secret_runner_aws/commits/master)
[![coverage report](https://gitlab.com/willmitchell/secret_runner_aws/badges/master/coverage.svg)](https://gitlab.com/willmitchell/secret_runner_aws/commits/master)

You can use this tool to manage secrets and to run your own programs in a secure manner on AWS.  This program manages secrets within
AWS SSM Parameter Store, which encrypts secrets using KMS.

Parameter naming convention:

 /prefix/module/stage/param_name
 
where:

- prefix: your company name. Example: com-example
- module: name of your module. Example: appserver
- stage: name of your runtime environment or stage.  Example: prod
-  param_name: name of a specific parameter.  Example: db_pass
 
secret_runner_aws will pull all secrets matching the path defined by /prefix/module/stage/* and then decrypt as required.
All resulting params will be transformed into environment variables that are then passed to your program.  This way,
your program can have access to secrets without resorting to insecure hacks!

# Dependencies

- [AWS SSM Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html)
- [AWS KMS](https://aws.amazon.com/kms/)
- [Cobra CLI](https://github.com/spf13/cobra)
- [AWS Go SDK v1.x](https://docs.aws.amazon.com/sdk-for-go/api/)

# Prior art

- [unicreds](https://github.com/Versent/unicreds)
- [credstash](https://github.com/fugue/credstash)

# Usage

Run the program with no parameters.  It will tell you how to use it.

You generally use the tool with a long command line.  This was done because I wanted a tool with zero config files
that works well in a Docker environment.  As a result, you will need to provide the prefix (-p), 
module (-m), and stage (-s) parameters for commands like run, get, list, and put.  These parameters provide
a context that is used to form a 'path' that is used for searching AWS SSM PS.  The path is basically just a 
part of a parameter name in AWS SSM PS.

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
  run        Run your program with secrets exposed as env vars.
  get         Get secrets from SSM
  help        Help about any command
  list        list secrets for this prefix/module/stage
  put         put secrets into SSM

Flags:
  -h, --help            help for secret_runner_aws
  -m, --module string   Module name
  -p, --prefix string   prefix name
  -s, --stage string    Stage name

Use "secret_runner_aws [command] --help" for more information about a command.

```

# Examples

Store a secret:

```
$ secret_runner_aws --prefix com-example -m mymodule -s prod put -n db_pass -v hello -e
put called.  name: db_pass, value: hello, encrypt: true
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
    Value: "hello",
    Version: 1
  }
}
```
Prove that we can run a subcommand and see the secret exposed as an env var:

```
$ secret_runner_aws --prefix com-example -m mymodule -s prod run -c 'env' | grep DB_PASS
DB_PASS=hello

```


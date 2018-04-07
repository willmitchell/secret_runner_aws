# ssm_executor: Secure secret delivery to arbitrary programs using AWS SSM Parameter Store and KMS

[![pipeline status](https://gitlab.com/willmitchell/ssm_executor/badges/master/pipeline.svg)](https://gitlab.com/willmitchell/ssm_executor/commits/master)
[![coverage report](https://gitlab.com/willmitchell/ssm_executor/badges/master/coverage.svg)](https://gitlab.com/willmitchell/ssm_executor/commits/master)


You can use this tool to manage and load secrets for other programs that you run.
This program pulls secrets from SSM, which can encrypt secrets using the AWS KMS.
Secrets are pulled from KMS based on a parameter naming convention:

Parameter naming convention:

 /prefix/module/stage/param_name
 
where:

 prefix: your company name. Example: com-example
 module: name of your module. Example: appserver
 stage: name of your runtime environment or stage.  Example: prod
 param_name: name of a specific parameter.  Example: db_pass
 
ssm_executor will pull all secrets matching the path defined by 
/prefix/module/stage/* and then decrypt as required.  All resulting 
params will be transformed into environment variables that are then passed
to other programs using the Go Exec (fork) facility.  This way, 
your program can have access to secrets without resorting to insecure hacks.

# Dependencies

- [AWS SSM Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html)
- [AWS KMS](https://aws.amazon.com/kms/)
- [Cobra CLI](https://github.com/spf13/cobra)
- [AWS Go SDK v1.x](https://docs.aws.amazon.com/sdk-for-go/api/)

# prior art

- [unicreds](https://github.com/Versent/unicreds)
- [credstash](https://github.com/fugue/credstash)

# Usage

Run the program with no parameters.  It will tell you how to use it.

You generally use the tool with a long command line.  This was done because I wanted a tool with zero config files
that works well in a Docker environment.  As a result, you will need to provide the prefix (-p), 
module (-m), and stage (-s) parameters for commands like exec, get, list, and put.  These parameters provide
a context that is used to form a 'path' that is used for searching AWS SSM PS.  The path is basically just a 
part of a parameter name in AWS SSM PS.

```
This program helps you use AWS SSM Parameter Store to manage your parameters and secrets.  These secrets are
encrypted using master keys that are managed by AWS KMS.  The tool provides CRUD operations for params 
and secrets, and it relies on a naming convention that maps onto existing AWS SSM PS IAM role management 
facilities.  You can also use this program to run *your* program in a secure manner using the 'exec' command.
The exec command exposes all appropriate secrets as environment variables and then execs your program.  

Disclaimer:   Provided without warranty of any kind.  Use at your own risk.  
Bug reports:  https://gitlab.com/willmitchell/ssm_executor/issues
Author:       will.mitchell@app3.com, 2018.

Usage:
  ssm_executor [command]

Available Commands:
  exec        Run your program with secrets exposed as env vars.
  get         Get secrets from SSM
  help        Help about any command
  list        list secrets for this prefix/module/stage
  put         put secrets into SSM

Flags:
  -h, --help            help for ssm_executor
  -m, --module string   Module name
  -p, --prefix string   prefix name
  -s, --stage string    Stage name

Use "ssm_executor [command] --help" for more information about a command.

```

# Examples

Store a secret:

```
$ ssm_executor --prefix com-example -m mymodule -s prod put -n db_pass -v hello -e
put called.  name: db_pass, value: hello, encrypt: true
{
  Version: 1
}

```

Get that secret:
```
$ ssm_executor --prefix com-example -m mymodule -s prod get -n db_pass
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
Prove that we can exec a subcommand and see the secret exposed as an env var:

```
$ ssm_executor --prefix com-example -m mymodule -s prod exec -c 'env' | grep DB_PASS
DB_PASS=hello

```


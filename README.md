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


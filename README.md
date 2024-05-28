# AWS Report

A slowly growing utility for reporting on various AWS things for bean counters. It currently has basic coverage for:-

* AWS Identity center (formerly AWS SSO)
* AWD Cognito user pools, groups and users

## Requirements

* A configured and working AWS CLI with at least one working profile configured.
* A working version of GoLang

## Installing AWS Report

```bash
go install github.com/trickyearlobe/awsreport@latest
```

The source will be downloaded, compiled and installed to `~/go/bin/awsreport` by default so make sure it gets added to your path, ideally into a shell startup script like `.bash_profile` or `.zshrc`

```bash
export PATH=$PATH:~/go/bin
```

## Using

```bash
awsreport --profile <an aws profile>
```

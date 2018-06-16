# AppSync Resolvers

[![GitHub release](https://img.shields.io/github/release/sbstjn/appsync-router.svg?maxAge=600)](https://github.com/sbstjn/appsync-router/releases)
[![MIT License](https://img.shields.io/github/license/sbstjn/appsync-router.svg?maxAge=3600)](https://github.com/sbstjn/appsync-router/blob/master/LICENSE.md)
[![GoDoc](https://godoc.org/github.com/sbstjn/appsync-router?status.svg)](https://godoc.org/github.com/sbstjn/appsync-router)
[![Go Report Card](https://goreportcard.com/badge/github.com/sbstjn/appsync-router)](https://goreportcard.com/report/github.com/sbstjn/appsync-router)
[![Build Status](https://img.shields.io/circleci/project/sbstjn/appsync-router.svg?maxAge=600)](https://circleci.com/gh/sbstjn/appsync-router)
[![Test Coverage](https://api.codeclimate.com/v1/badges/ae56f89b122d14b9749e/test_coverage)](https://codeclimate.com/github/sbstjn/appsync-router/test_coverage)

> Wrapper for handling AWS AppSync resolvers with AWS Lambda in Go. See [appsync-router-example] for an example project and how to set up a complete GraphQL API using the [Serverless Application Model].

## Usage

### Installation

```
$ > go get github.com/sbstjn/appsync-router
```

### Routing

```go
import (
  "github.com/sbstjn/appsync-router"
)

func resolverA() (interface{}, error) {
	return nil, fmt.Errorf("Nothing here in resolver A: %s", args.Foo)
}

func resolverB(args struct {
	Bar string `json:"bar"`
}) (interface{}, error) {
	return nil, fmt.Errorf("Nothing here in resolver B: %s", args.Bar)
}

var (
	r = router.New()
)

func init() {
  r = router.New()

	r.Add("fieldA", resolverA)
	r.Add("fieldB", resolverB)
}

func main() {
	lambda.Start(r.Handle)
}
```

### AppSync Configuration

Resolver lookup is based on a `field` property in your `RequestMappingTemplate`, which can be configured using the AWS Console or CloudFormation as well. This approach works fine with the recommended [AWS setup] for multiple custom resolvers and AWS Lambda data sources:

```yaml
  AppSyncDataSource:
    Type: AWS::AppSync::DataSource
    Properties:
      ApiId: !GetAtt [ AppSyncAPI, ApiId ]
      Name: resolver
      Type: AWS_LAMBDA
      LambdaConfig:
        LambdaFunctionArn: !GetAtt [ Lambda, Arn ]
      ServiceRoleArn: !GetAtt [ Role, Arn ]

  AppSyncResolverA:
    Type: AWS::AppSync::Resolver
    Properties:
      ApiId: !GetAtt [ AppSyncAPI, ApiId ]
      TypeName: Query
      FieldName: fieldA
      DataSourceName: !GetAtt [ AppSyncDataSource, Name ]
      RequestMappingTemplate: '{ "version" : "2017-02-28", "operation": "Invoke", "payload": { "field": "fieldA", "arguments": $utils.toJson($context.arguments) } }'
      ResponseMappingTemplate: $util.toJson($context.result)

  AppSyncResolverB:
    Type: AWS::AppSync::Resolver
    Properties:
      ApiId: !GetAtt [ AppSyncAPI, ApiId ]
      TypeName: Query
      FieldName: fieldB
      DataSourceName: !GetAtt [ AppSyncDataSource, Name ]
      RequestMappingTemplate: '{ "version" : "2017-02-28", "operation": "Invoke", "payload": { "field": "fieldB", "arguments": $utils.toJson($context.arguments) } }'
      ResponseMappingTemplate: $util.toJson($context.result)
```

See [appsync-router-example] for an example project and how to set up a serverless GraphQL API with AWS AppSync using the [Serverless Application Model].

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘 

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.

[appsync-router-example]: https://github.com/sbstjn/appsync-router-example
[Serverless Application Model]: https://github.com/awslabs/serverless-application-model
[AWS setup]: https://docs.aws.amazon.com/appsync/latest/devguide/tutorial-lambda-resolvers.html
# AppSync Resolvers

[![GitHub release](https://img.shields.io/github/release/sbstjn/appsync-resolvers.svg?maxAge=600)](https://github.com/sbstjn/appsync-resolvers/releases)
[![MIT License](https://img.shields.io/github/license/sbstjn/appsync-resolvers.svg?maxAge=3600)](https://github.com/sbstjn/appsync-resolvers/blob/master/LICENSE.md)
[![GoDoc](https://godoc.org/github.com/sbstjn/appsync-resolvers?status.svg)](https://godoc.org/github.com/sbstjn/appsync-resolvers)
[![Go Report Card](https://goreportcard.com/badge/github.com/sbstjn/appsync-resolvers)](https://goreportcard.com/report/github.com/sbstjn/appsync-resolvers)
[![Build Status](https://img.shields.io/circleci/project/sbstjn/appsync-resolvers.svg?maxAge=600)](https://circleci.com/gh/sbstjn/appsync-resolvers)
[![Test Coverage](https://api.codeclimate.com/v1/badges/ae56f89b122d14b9749e/test_coverage)](https://codeclimate.com/github/sbstjn/appsync-resolvers/test_coverage)

> Easily create AWS AppSync resolvers with AWS Lambda using Go. See [appsync-resolvers-example] for an example project with custon `Field` and `Query` resolvers and how to set up, maintain, and deploy a working GraphQL API using the [Serverless Application Model] and without any third-party frameworks.

## Usage

### Installation

```
$ > go get github.com/sbstjn/appsync-resolvers
```

### Routing

```go
import (
    "github.com/sbstjn/appsync-resolvers"
)

type personArguments struct {
    ID int `json:"id"`
}

func resolvePeople() (interface{}, error) {
    return dataPeople, nil
}

func resolvePerson(p personArguments) (interface{}, error) {
    return dataPeople.byID(p.ID)
}

func resolveFriends(p person) (interface{}, error) {
    return p.getFriends()
}

var (
    r = resolvers.New()
)

func init() {
    r.Add("query.people", resolvePeople)
    r.Add("query.person", resolvePerson)
    r.Add("field.person.friends", resolveFriends)
}

func main() {
    lambda.Start(r.Handle)
}
```

### AppSync Configuration

Resolver lookup is based on a `resolve` property in your `RequestMappingTemplate`, which can be configured using the AWS Console or CloudFormation as well. This approach works fine with the recommended [AWS setup] for multiple custom resolvers and AWS Lambda data sources:

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

AppSyncResolverPeople:
  Type: AWS::AppSync::Resolver
  Properties:
    ApiId: !GetAtt [ AppSyncAPI, ApiId ]
    TypeName: Query
    FieldName: people
    DataSourceName: !GetAtt [ AppSyncDataSource, Name ]
    RequestMappingTemplate: '{ "version" : "2017-02-28", "operation": "Invoke", "payload": { "resolve": "query.people", "context": $utils.toJson($context) } }'
    ResponseMappingTemplate: $util.toJson($context.result)

AppSyncResolverPerson:
  Type: AWS::AppSync::Resolver
  Properties:
    ApiId: !GetAtt [ AppSyncAPI, ApiId ]
    TypeName: Query
    FieldName: person
    DataSourceName: !GetAtt [ AppSyncDataSource, Name ]
    RequestMappingTemplate: '{ "version" : "2017-02-28", "operation": "Invoke", "payload": { "resolve": "query.person", "context": $utils.toJson($context) } }'
    ResponseMappingTemplate: $util.toJson($context.result)

AppSyncResolverFriends:
  Type: AWS::AppSync::Resolver
  Properties:
    ApiId: !GetAtt [ AppSyncAPI, ApiId ]
    TypeName: Person
    FieldName: friends
    DataSourceName: !GetAtt [ AppSyncDataSource, Name ]
    RequestMappingTemplate: '{ "version" : "2017-02-28", "operation": "Invoke", "payload": { "resolve": "field.person.friends", "context": $utils.toJson($context) } }'
    ResponseMappingTemplate: $util.toJson($context.result)
```

Head over to [appsync-resolvers-example] for an example project and how simple it can be to set up, maintain, and deploy a serverless GraphQL API with AWS AppSync using the [Serverless Application Model].

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘 

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.

[appsync-resolvers-example]: https://github.com/sbstjn/appsync-resolvers-example
[Serverless Application Model]: https://github.com/awslabs/serverless-application-model
[AWS setup]: https://docs.aws.amazon.com/appsync/latest/devguide/tutorial-lambda-resolvers.html

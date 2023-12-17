# Gin Example

[Gin](https://github.com/gin-gonic/gin) is a HTTP web framework written in Go (Golang).

# Getting Started

## Init go module

Init go module:

```
go mod init github.com/IkeMurami-Examples/go-gin-example
```

Install basic dependencies:

```
go get -u github.com/spf13/cobra@latest
go get -u github.com/spf13/viper@latest
go get -u github.com/gin-gonic/gin
go get -u golang.org/x/sync/errgroup
```

Install Cobra CLI (if it isn't installed):

```
go install github.com/spf13/cobra-cli@latest
```

Initialize Cobra Interface and add a Viper support:

```
cobra-cli init --author "Ike Murami murami.ike@gmail.com" --license GPLv3 --viper
cobra-cli add start --author "Ike Murami murami.ike@gmail.com" 
```

So, we can start our program:

```
go run main.go start
```

## Deploy setting

```
mkdir build && cd build
mkdir config
touch config/gin-example.yaml
touch Dockerfile docker-compose.yml
```

## Base Gin HTTP server

```
mkdir pkg
mkdir pkg/cmd
mkdir pkg/server
touch pkg/server/server.go pkg/server/handler.go pkg/cmd/gin-example.go
```

Create that files and modify `cmd/root.go` and `cmd/start.go`.  
Add support of yaml config (`build/config/gin-example.yaml`).

## Add a logger middleware to the Gin HTTP server

Install the Zap logger:

```
go get -u go.uber.org/zap
mkdir pkg/utils
touch pkg/utils/log.go
mkdir pkg/server/middleware
touch pkg/server/middleware/logger.go
```

## Add GraphQL Endpoint

### Get Started

How to: https://www.apollographql.com/blog/graphql/golang/using-graphql-with-golang/

Install gqlgen:

```
go get -u github.com/99designs/gqlgen
```

Next, add gqlgen to your project’s `tools/tools.go`:

```go
//go:build tools
package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
```

Add `configs/gqlgen.yml` config:

```yml
# Where are all the schema files located? globs are supported eg  src/**/*.graphqls
schema:
  - ../graph/*.graphqls

# Where should the generated server code go?
exec:
  filename: ../graph/generated.go
  package: graph

# Uncomment to enable federation
# federation:
#   filename: graph/federation.go
#   package: graph

# Where should any generated models go?
model:
  filename: ../graph/model/models_gen.go
  package: model

# Where should the resolver implementations go?
resolver:
  layout: follow-schema
  dir: ../graph
  package: graph

# Optional: turn on use ` + "`" + `gqlgen:"fieldName"` + "`" + ` tags in your models
# struct_tag: json

# Optional: turn on to use []Thing instead of []*Thing
# omit_slice_element_pointers: false

# Optional: turn off to make struct-type struct fields not use pointers
# e.g. type Thing struct { FieldA OtherThing } instead of { FieldA *OtherThing }
# struct_fields_always_pointers: true

# Optional: turn off to make resolvers return values instead of pointers for structs
# resolvers_always_return_pointers: true

# Optional: set to speed up generation time by not performing a final validation pass.
# skip_validation: true

# gqlgen will search for any type names in the schema in these go packages
# if they match it will use them, otherwise it will generate them.
autobind:
#  - "test/graph/model"

# This section declares type mapping between the GraphQL and go type systems
#
# The first line in each type will be used as defaults for resolver arguments and
# modelgen, the others will be allowed when binding to fields. Configure them to
# your liking
models:
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.ID
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
  Int:
    model:
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32

```

Add GraphQL Scheme — `pkg/graph/schema.graphqls`:

```graphql
```

Add cmd to generate Go files for GraphQL API (`tools/generate.go`):

```go
package tools

//go:generate go run github.com/99designs/gqlgen generate --config ../configs/gqlgen.yml
```

And add to `.gitignore` the generated Go-files:

```
# We don't add a generated code
pkg/graph/**/*.go
```


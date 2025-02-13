# `tempo-api`

The `tempo-api` package provides core interfaces and types for extending the **tempo CLI** functionality.

## 💻 Installation

```bash
go get github.com/indaco/tempo-api@latest
```

## 📦 Packages Overview

### `tempo-api/templatefuncs`

This package defines the `TemplateFuncProvider` interface, which external function providers must implement to register custom template functions.

#### Implementing a Custom Template Function Provider

To register custom template functions, your Go package must implement the `TemplateFuncProvider` interface **and define the provider in a file named `provider.go`**.

#### Example

```go
// provider.go
package myprovider

import (
    "text/template"

    "github.com/indaco/tempo-api/templatefuncs"
)

// MyProvider implements the TemplateFuncProvider interface
type MyProvider struct{}

// GetFunctions returns a map of function names to implementations
func (p *MyProvider) GetFunctions() template.FuncMap {
    return template.FuncMap{
        "myFunc": func() string { return "Hello from myFunc!" },
    }
}

// Provider instance (must be declared in provider.go)
var Provider templatefuncs.TemplateFuncProvider = &MyProvider{}
```

> [!IMPORTANT]
> The `file` must be named **provider.go** for tempo to detect it.
>
> The Provider variable **must** be:
>
> - Exported (`Provider`, not `provider`).
> - Declared at the **package level** (not inside a function).
> - Implementing `templatefuncs.TemplateFuncProvider`.

## 🔗 Registering the Provider in Tempo

Once your provider is implemented, you can register it using the tempo register functions command:

### Register from a Git Repository

```bash
tempo register functions --url https://github.com/user/myprovider.git
```

### Register from a Local Path

```bash
tempo register functions --path /path/to/myprovider_module
```

## Existing Providers

- [tempo-provider-sprig](https://github.com/indaco/tempo-provider-sprig)
- [tempo-provider-goname](https://github.com/indaco/tempo-provider-goname)

## 📖 Related Documentation

- [Tempo CLI](https://github.com/indaco/tempo)
- [Extending Tempo](https://github.com/indaco/tempo/blob/main/README.md#-extending-tempo)

## 🆓 License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

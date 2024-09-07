<!-- markdownlint-disable -->

> Tiny launcher for long running processes that terminates on SIGTERM or SIGINT

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# neolauncher

```go
import "github.com/byteartis/neolauncher"
```

## Index

- [Variables](<#variables>)
- [func Launch\(ctx context.Context, svc Launcher\)](<#Launch>)
- [type Launcher](<#Launcher>)


## Variables

<a name="ErrLauncherIsNil"></a>ErrLauncherIsNil for when nil is passed to the launcher

```go
var ErrLauncherIsNil = errors.New("launcher implementation is nil")
```

<a name="Launch"></a>
## func [Launch](<https://github.com/byteartis/neolauncher/blob/main/launch.go#L21>)

```go
func Launch(ctx context.Context, svc Launcher)
```

Launch takes in a context and a Launcher implementation. The launcher implementation is expected to be a long running process that will be terminated once either SIGTERM or SIGINT is received.

<a name="Launcher"></a>
## type [Launcher](<https://github.com/byteartis/neolauncher/blob/main/launch.go#L14-L16>)



```go
type Launcher interface {
    Launch(context.Context) error
}
```

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

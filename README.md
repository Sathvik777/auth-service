
# go-api-skeleton


# Installation
Add the following lines to ~/.bach_profile or ~/.zshrc (if using zsh)

    export GOPATH=/Users/username/go

    export PATH=$GOPATH/bin:$PATH

Where username is the username of your profile.

Then install dep:

```
brew install dep
```

Run this command in the correct folder:

```
dep ensure
```

# Testing

Install mockery

```
go get github.com/vektra/mockery/.../
```

Generate files
```
go generate ./....
```

Run tests
```
go test ./...
```

# Code Guide

## Api
The api is found in the `goservice` folder and has the service it uses injected to it's structure.

To have multiple endpoints, simply add endpoints in `InitAPIRoute()`.

Middleware specific to endpoint can be added on line 25 in `api.go`:

```
Handler(negroni.New(
    potentialMiddleware,
    negroni.HandlerFunc(a.GoEndpoint()),
))
```

## Database
The database will automatically apply migrations if the variable in `.env` called `ENV` is either `dev` or `test`.

I am using https://github.com/mattes/migrate for database migrations.

Run a local database in docker with `docker-compose up`.

## Bootstrap
I am using https://github.com/facebookgo/inject for dependency injection.

By injecting implementations of interfaces in the `bootstrapApp` you can easily inject them in structs such as:

```
type TestStruct struct {
    Variable pkg.InterfaceType `inject:""`
}
```

## Vendor folder
I have the vendor folder checked into the repo, for reproducibility.

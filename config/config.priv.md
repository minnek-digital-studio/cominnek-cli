## config.priv.go Example

This example shows how to use the `config.priv.go` file to set the `config.priv.go` file.

```go
package config

type IConfigPrivate struct {
	GithubClient	string
	EncryptKey 		string
}

var Private = IConfigPrivate{
	GithubClient: "GITHUB_CLIENT_ID",
	EncryptKey: "PRIVATE_ENCRYPT_KEY",
}
```
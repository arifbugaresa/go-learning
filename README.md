# go-learning
#### long-term project to learn the go programming language.

Example Golang project for reading env files.

## External Lib

```sh
go go get github.com/spf13/viper
```
## What is Viper?

Viper is a complete configuration solution for Go applications including [12-Factor apps](https://12factor.net/#the_twelve_factors).
It is designed to work within an application, and can handle all types of configuration needs
and formats. It supports:

* setting defaults
* reading from JSON, TOML, YAML, HCL, envfile and Java properties config files
* live watching and re-reading of config files (optional)
* reading from environment variables
* reading from remote config systems (etcd or Consul), and watching changes
* reading from command line flags
* reading from buffer
* setting explicit values

Viper can be thought of as a registry for all of your applications configuration needs.


## Why Viper?

When building a modern application, you don’t want to worry about
configuration file formats; you want to focus on building awesome software.
Viper is here to help with that.

Viper does the following for you:

1. Find, load, and unmarshal a configuration file in JSON, TOML, YAML, HCL, INI, envfile or Java properties formats.
2. Provide a mechanism to set default values for your different configuration options.
3. Provide a mechanism to set override values for options specified through command line flags.
4. Provide an alias system to easily rename parameters without breaking existing code.
5. Make it easy to tell the difference between when a user has provided a command line or config file which is the same as the default.




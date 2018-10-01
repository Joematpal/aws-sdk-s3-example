package config

// RPCList of arguments
var RPCList = []string{
	"AWS_ACCESS_KEY_ID",
	"AWS_SECRET_ACCESS_KEY",
}

// RPCmap ...
// TODO: Maybe this could work
var RPCmap = map[string]*Arg{
	"AWS_ACCESS_KEY_ID":     &Arg{},
	"AWS_SECRET_ACCESS_KEY": &Arg{},
}

// Arg ...
type Arg struct {
	Value, Default, Description, Shorthand string
}

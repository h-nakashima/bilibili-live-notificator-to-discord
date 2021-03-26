module bilibili-live-notificator

go 1.15

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dghubble/go-twitter v0.0.0-20201011215211-4b180d0cc78d
	github.com/dghubble/oauth1 v0.7.0
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/dghubble/go-twitter => github.com/janisz/go-twitter v0.0.0-20201206102041-3fe237ed29f3

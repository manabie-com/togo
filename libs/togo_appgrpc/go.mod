module github.com/phathdt/libs/togo_appgrpc

go 1.19

replace (
	github.com/phathdt/libs/go-sdk => ../go-sdk
	github.com/phathdt/libs/togo_proto => ../togo_proto
)

require (
	github.com/phathdt/libs/go-sdk v0.0.0
	github.com/phathdt/libs/togo_proto v0.0.0
	google.golang.org/grpc v1.49.0
)

require (
	github.com/facebookgo/flagenv v0.0.0-20160425205200-fcd59fca7456 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/lib/pq v1.10.6 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/x-cray/logrus-prefixed-formatter v0.5.2 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220812174116-3211cb980234 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

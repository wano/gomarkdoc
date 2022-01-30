module github.com/princjef/gomarkdoc

go 1.16

replace (
	github.com/mgutz/ansi => github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d
	github.com/onsi/ginkgo => github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega => github.com/onsi/gomega v1.18.1 // indirect
	github.com/princjef/gomarkdoc => ./
)

require (
	github.com/go-git/go-git/v5 v5.3.0
	github.com/matryer/is v1.4.0
	github.com/mgutz/ansi v0.0.0-00010101000000-000000000000 // indirect
	github.com/onsi/ginkgo v0.0.0-00010101000000-000000000000 // indirect
	github.com/princjef/mageutil v0.1.0
	github.com/russross/blackfriday/v2 v2.1.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	mvdan.cc/xurls/v2 v2.2.0
)

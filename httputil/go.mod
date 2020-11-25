module github.com/puppetlabs/leg/httputil

go 1.14

replace github.com/puppetlabs/leg/lifecycle => ../lifecycle

replace github.com/puppetlabs/leg/logging => ../logging

replace github.com/puppetlabs/leg/request => ../request

require (
	github.com/gorilla/websocket v1.4.2
	github.com/puppetlabs/errawr-gen v1.0.1
	github.com/puppetlabs/errawr-go/v2 v2.2.0
	github.com/puppetlabs/leg/instrumentation v0.0.0-00010101000000-000000000000
	github.com/puppetlabs/leg/lifecycle v0.0.0-00010101000000-000000000000
	github.com/puppetlabs/leg/logging v0.0.0-00010101000000-000000000000
	github.com/puppetlabs/leg/request v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.6.1
	golang.org/x/text v0.3.4
)

replace github.com/puppetlabs/leg/instrumentation => ../instrumentation

replace github.com/puppetlabs/leg/netutil => ../netutil

replace github.com/puppetlabs/leg/scheduler => ../scheduler
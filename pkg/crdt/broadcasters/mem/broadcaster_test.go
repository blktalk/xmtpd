package membroadcaster_test

import (
	"testing"

	membroadcaster "github.com/xmtp/xmtpd/pkg/crdt/broadcasters/mem"
	crdttest "github.com/xmtp/xmtpd/pkg/crdt/testing"
	test "github.com/xmtp/xmtpd/pkg/testing"
)

func TestMemoryBroadcaster(t *testing.T) {
	crdttest.RunBroadcasterTests(t, func(t *testing.T) crdttest.TestBroadcaster {
		log := test.NewLogger(t)
		return membroadcaster.New(log)
	})
}

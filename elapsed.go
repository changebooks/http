package http

import (
	"fmt"
	"time"
)

type Elapsed struct {
	Total             time.Duration
	DNSLookup         time.Duration // Zero, when Reused or (with IP, no DNS)
	Connect           time.Duration // Zero, when Reused
	TLSHandshake      time.Duration // Zero, when Reused or no TLS
	Server            time.Duration
	Transfer          time.Duration
	Start             time.Time
	Done              time.Time
	DNSStart          time.Time // Zero, when Reused or (with IP, no DNS)
	DNSDone           time.Time // Zero, when Reused or (with IP, no DNS)
	ConnectStart      time.Time // Zero, when Reused
	ConnectDone       time.Time // Zero, when Reused
	TLSHandshakeStart time.Time // Zero, when Reused or no TLS
	TLSHandshakeDone  time.Time // Zero, when Reused or no TLS
	ServerStart       time.Time
	ServerDone        time.Time
	TransferStart     time.Time
	TransferDone      time.Time
}

func (x *Elapsed) ToString() string {
	return fmt.Sprintf("Total:        %v, Start: %v, Done: %v\n"+
		"DNSLookup:    %v, Start: %v, Done: %v\n"+
		"Connect:      %v, Start: %v, Done: %v\n"+
		"TLSHandshake: %v, Start: %v, Done: %v\n"+
		"Server:       %v, Start: %v, Done: %v\n"+
		"Transfer:     %v, Start: %v, Done: %v",
		x.Total, x.Start, x.Done,
		x.DNSLookup, x.DNSStart, x.DNSDone,
		x.Connect, x.ConnectStart, x.ConnectDone,
		x.TLSHandshake, x.TLSHandshakeStart, x.TLSHandshakeDone,
		x.Server, x.ServerStart, x.ServerDone,
		x.Transfer, x.TransferStart, x.TransferDone,
	)
}

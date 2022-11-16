package utils

import "github.com/shravan9912/testmpquicback/internal/protocol"

// PacketInterval is an interval from one PacketNumber to the other
// +gen linkedlist
type PacketInterval struct {
	Start protocol.PacketNumber
	End   protocol.PacketNumber
}

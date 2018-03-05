package il

import "net"

const (
	ProtoID = 40
)

type State byte

const (
	Closed State = iota
	Syncer
	Syncee
	Established
	Closing
)

func Dial(net, addr string) (net.Conn, error) {
	panic("not implemented")
}

func Listen(net, addr string) (net.Listener, error) {
	panic("not implemented")
}

type IPHdr struct {
	VIHL  byte   // Version and header length
	TOS   byte   // Type of service
	Len   uint16 // packet length
	ID    uint16 // Identification
	Frag  uint16 // Fragment information
	TTL   uint8  // Time to live
	Proto byte   // Protocol
	Sum   uint16 // Header checksum
	Src   uint32 // Ip source
	Dst   uint32 //Ip destination
}

type Kind byte

const (
	Ksync Kind = iota
	Kdata
	Kdataquery
	Kack
	Kquery
	Kstat
	Kclose
)

type ILHdr struct {
	Sum  uint16
	Len  uint16 // size of IL header (18 bytes) + data.
	Type byte
	Spec byte
	Src  uint16
	Dst  uint16
	ID   uint32
	Ack  uint32
}

type IPIL struct {
	IPHdr
	ILHdr
}

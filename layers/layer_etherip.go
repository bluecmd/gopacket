// Copyright 2012 Google, Inc. All rights reserved.

package gopacket

import (
"github.com/gconnell/gopacket"
	"encoding/binary"
)

// EtherIP is the struct for storing RFC 3378 EtherIP packet headers.
type EtherIP struct {
	Version  uint8
	Reserved uint16
}

// LayerType returns gopacket.LayerTypeEtherIP.
func (e *EtherIP) LayerType() gopacket.LayerType { return gopacket.LayerTypeEtherIP }

func decodeEtherIP(data []byte) (out gopacket.DecodeResult, _ error) {
	out.DecodedLayer = &EtherIP{
		Version:  data[0] >> 4,
		Reserved: binary.BigEndian.Uint16(data[:2]) & 0x0fff,
	}
	out.NextDecoder = decoderFunc(decodeEthernet)
	out.RemainingBytes = data[2:]
	return
}
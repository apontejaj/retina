// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package packetparser

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type packetparserCtEntry struct {
	EvictionTime     uint32
	LastReportTxDir  uint32
	LastReportRxDir  uint32
	TrafficDirection uint8
	FlagsSeenTxDir   uint8
	FlagsSeenRxDir   uint8
	IsClosing        bool
}

type packetparserCtV4Key struct {
	SrcIp   uint32
	DstIp   uint32
	SrcPort uint16
	DstPort uint16
	Proto   uint8
	_       [3]byte
}

type packetparserMapKey struct {
	Prefixlen uint32
	Data      uint32
}

type packetparserPacket struct {
	T_nsec      uint64
	Bytes       uint32
	SrcIp       uint32
	DstIp       uint32
	SrcPort     uint16
	DstPort     uint16
	TcpMetadata struct {
		Seq    uint32
		AckNum uint32
		Tsval  uint32
		Tsecr  uint32
	}
	ObservationPoint uint8
	TrafficDirection uint8
	Proto            uint8
	Flags            uint8
	IsReply          bool
	_                [3]byte
}

// loadPacketparser returns the embedded CollectionSpec for packetparser.
func loadPacketparser() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_PacketparserBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load packetparser: %w", err)
	}

	return spec, err
}

// loadPacketparserObjects loads packetparser and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*packetparserObjects
//	*packetparserPrograms
//	*packetparserMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadPacketparserObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadPacketparser()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// packetparserSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type packetparserSpecs struct {
	packetparserProgramSpecs
	packetparserMapSpecs
}

// packetparserSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type packetparserProgramSpecs struct {
	EndpointEgressFilter  *ebpf.ProgramSpec `ebpf:"endpoint_egress_filter"`
	EndpointIngressFilter *ebpf.ProgramSpec `ebpf:"endpoint_ingress_filter"`
	HostEgressFilter      *ebpf.ProgramSpec `ebpf:"host_egress_filter"`
	HostIngressFilter     *ebpf.ProgramSpec `ebpf:"host_ingress_filter"`
}

// packetparserMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type packetparserMapSpecs struct {
	PacketparserEvents *ebpf.MapSpec `ebpf:"packetparser_events"`
	RetinaConntrackMap *ebpf.MapSpec `ebpf:"retina_conntrack_map"`
	RetinaFilterMap    *ebpf.MapSpec `ebpf:"retina_filter_map"`
}

// packetparserObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadPacketparserObjects or ebpf.CollectionSpec.LoadAndAssign.
type packetparserObjects struct {
	packetparserPrograms
	packetparserMaps
}

func (o *packetparserObjects) Close() error {
	return _PacketparserClose(
		&o.packetparserPrograms,
		&o.packetparserMaps,
	)
}

// packetparserMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadPacketparserObjects or ebpf.CollectionSpec.LoadAndAssign.
type packetparserMaps struct {
	PacketparserEvents *ebpf.Map `ebpf:"packetparser_events"`
	RetinaConntrackMap *ebpf.Map `ebpf:"retina_conntrack_map"`
	RetinaFilterMap    *ebpf.Map `ebpf:"retina_filter_map"`
}

func (m *packetparserMaps) Close() error {
	return _PacketparserClose(
		m.PacketparserEvents,
		m.RetinaConntrackMap,
		m.RetinaFilterMap,
	)
}

// packetparserPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadPacketparserObjects or ebpf.CollectionSpec.LoadAndAssign.
type packetparserPrograms struct {
	EndpointEgressFilter  *ebpf.Program `ebpf:"endpoint_egress_filter"`
	EndpointIngressFilter *ebpf.Program `ebpf:"endpoint_ingress_filter"`
	HostEgressFilter      *ebpf.Program `ebpf:"host_egress_filter"`
	HostIngressFilter     *ebpf.Program `ebpf:"host_ingress_filter"`
}

func (p *packetparserPrograms) Close() error {
	return _PacketparserClose(
		p.EndpointEgressFilter,
		p.EndpointIngressFilter,
		p.HostEgressFilter,
		p.HostIngressFilter,
	)
}

func _PacketparserClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed packetparser_bpfel_arm64.o
var _PacketparserBytes []byte

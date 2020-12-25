package osgi

import (
	"encoding/json"
	"log"
)

type OSGIPacketStandard struct {
	ProxySequence int
	Data          string
}

func ToPacket(packetString string) *OSGIPacketStandard {
	packetStandard := OSGIPacketStandard{}
	json.Unmarshal([]byte(packetString), &packetStandard)

	return &packetStandard
}

func ToString(packetStandard OSGIPacketStandard) string {
	e, err := json.Marshal(packetStandard)

	if err != nil {
		log.Println(err)
		//TODO: 에러처리
		return ""
	}

	return string(e)
}

func CreatePacket(proxySequence int, json string) *OSGIPacketStandard {
	return &OSGIPacketStandard{
		ProxySequence: proxySequence,
		Data:          json,
	}
}

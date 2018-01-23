package xmpp // import "fluux.io/xmpp"

import "encoding/xml"

// ============================================================================
// Presence Packet

type Presence struct {
	XMLName xml.Name `xml:"presence"`
	PacketAttrs
	Show     string `xml:"show,attr,omitempty"` // away, chat, dnd, xa
	Status   string `xml:"status,attr,omitempty"`
	Priority string `xml:"priority,attr,omitempty"`
	Error    Err    `xml:"error,omitempty"`
}

func (Presence) Name() string {
	return "presence"
}

type presenceDecoder struct{}

var presence presenceDecoder

func (presenceDecoder) decode(p *xml.Decoder, se xml.StartElement) (Presence, error) {
	var packet Presence
	err := p.DecodeElement(&packet, &se)
	// TODO Add default presence type (when omitted)
	return packet, err
}

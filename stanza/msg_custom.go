package stanza

import (
	"encoding/xml"
)

// ============================================================================
// Custom Message extensions

const (
	GIMCommandStatusReceived  = "received"
	GIMCommandStatusCompleted = "completed"
	GIMCommandStatusExecuting = "executing"
	GIMCommandStatusFailed    = "failed"
)

const NSMsgRead = "urn:xmpp:read"
const NSMsgUnread = "urn:xmpp:unread"
const NSMsgGIMS = "urn:xmpp:gims"

type ReadReceipt struct {
	MsgExtension
	XMLName xml.Name `xml:"urn:xmpp:read read"`
	ID      string   `xml:"id,attr"`
}

// Unread implements an extenstion for receiving unread message notifications
type Unread struct {
	MsgExtension
	XMLName xml.Name `xml:"urn:xmpp:unread unread"`
	Data    string   `xml:"data,attr"`
}

// PublishUnread implements an extenstion for sending publish unread request
type PublishUnread struct {
	MsgExtension
	XMLName xml.Name `xml:"urn:xmpp:unread publish"`
	Ref     string   `xml:"ref,attr,omitempty"`
	ID      string   `xml:"id,attr,omitempty"`
}

// UnpublishUnread implements an extenstion for sending unpublish unread request
type UnpublishUnread struct {
	MsgExtension
	XMLName xml.Name `xml:"urn:xmpp:unread unpublish"`
	Ref     string   `xml:"ref,attr,omitempty"`
	ID      string   `xml:"id,attr,omitempty"`
}

// GIMSCommand implements an extenstion for sending commands to/from GIMS client
type GIMSCommand struct {
	MsgExtension
	XMLName xml.Name `xml:"urn:xmpp:gims command"`
	Action  string   `xml:"action,attr"`
}

//GIMSQuery implements an extenstion for sending query to/from GIMS client
type GIMSQuery struct {
	MsgExtension
	XMLName xml.Name `xml:"urn:xmpp:gims query"`
	Context string   `xml:"context,attr"`
	Subject string   `xml:"subject,attr"`
}

//GIMSReply implements an extenstion for sending reply to a query
type GIMSReply struct {
	MsgExtension
	XMLName xml.Name `xml:"urn:xmpp:gims reply"`
	Context string   `xml:"context,attr"`
	ID      string   `xml:"id,attr"`
}

//GIMSAcknowledgement implements an extenstion for sending acknowledgement for command/query
type GIMSAcknowledgement struct {
	MsgExtension
	XMLName xml.Name `xml:"urn:xmpp:gims ack"`
	Status  string   `xml:"status,attr"`
	ID      string   `xml:"id,attr"`
}

func init() {
	TypeRegistry.MapExtension(PKTMessage, xml.Name{Space: NSMsgRead, Local: "read"}, ReadReceipt{})
	TypeRegistry.MapExtension(PKTMessage, xml.Name{Space: NSMsgUnread, Local: "unread"}, Unread{})
	TypeRegistry.MapExtension(PKTMessage, xml.Name{Space: NSMsgUnread, Local: "publish"}, PublishUnread{})
	TypeRegistry.MapExtension(PKTMessage, xml.Name{Space: NSMsgUnread, Local: "unpublish"}, UnpublishUnread{})
	TypeRegistry.MapExtension(PKTMessage, xml.Name{Space: NSMsgGIMS, Local: "command"}, GIMSCommand{})
	TypeRegistry.MapExtension(PKTMessage, xml.Name{Space: NSMsgGIMS, Local: "ack"}, GIMSCommand{})

}

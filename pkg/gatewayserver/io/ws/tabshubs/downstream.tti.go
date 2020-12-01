// Copyright © 2020 The Things Industries B.V.

package tabshubs

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io/ws"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/scheduling"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var errNoClockSync = errors.DefineUnavailable("no_clock_sync", "no clock sync")

// DownlinkMessage is the LoRaWAN downlink message sent to the Tabs Hubs device.
type DownlinkMessage struct {
	DevEUI  string  `json:"DevEui"`
	SeqNo   int64   `json:"seqno"`
	Pdu     string  `json:"pdu"`
	DR      int     `json:"DR"`
	Freq    int     `json:"Freq"`
	XTime   int64   `json:"xtime"`
	RCtx    int64   `json:"rctx"`
	MuxTime float64 `json:"MuxTime"`
}

// MarshalJSON marshals dnmsg to a JSON byte array.
func (dnmsg DownlinkMessage) MarshalJSON() ([]byte, error) {
	type Alias DownlinkMessage
	return json.Marshal(struct {
		Type string `json:"msgtype"`
		Alias
	}{
		Type:  TypeDownstreamDownlinkMessage,
		Alias: Alias(dnmsg),
	})
}

// unmarshalJSON unmarshals dnmsg from a JSON byte array.
func (dnmsg *DownlinkMessage) unmarshalJSON(data []byte) error {
	return json.Unmarshal(data, dnmsg)
}

// FromDownlink implements Formatter.
func (f *tabsHubs) FromDownlink(ctx context.Context, uid string, down ttnpb.DownlinkMessage, concentratorTime scheduling.ConcentratorTime, dlTime time.Time) ([]byte, error) {
	var dnmsg DownlinkMessage
	settings := down.GetScheduled()
	dnmsg.Pdu = hex.EncodeToString(down.GetRawPayload())
	dnmsg.SeqNo = int64(f.tokens.Next(down.CorrelationIDs, dlTime))

	// The first 16 bits of XTime gets the session ID from the upstream latestXTime and the other 48 bits are concentrator timestamp accounted for rollover.
	var (
		state State
		ok    bool
	)
	session := ws.SessionFromContext(ctx)
	session.DataMu.Lock()
	defer session.DataMu.Unlock()
	if state, ok = session.Data.(State); !ok {
		return nil, errSessionStateNotFound
	}

	dnmsg.XTime = int64(state.ID)<<48 | (int64(concentratorTime) / int64(time.Microsecond) & 0xFFFFFFFFFF)

	dnmsg.DevEUI = "00-00-00-00-00-00-00-00"

	// Fix the Tx Parameters since we don't use the gateway scheduler.
	dnmsg.DR = int(settings.DataRateIndex)
	dnmsg.Freq = int(settings.Frequency)

	// Add the MuxTime for RTT measurement
	dnmsg.MuxTime = float64(dlTime.UnixNano()) / float64(time.Second)

	return dnmsg.MarshalJSON()
}

// ToDownlinkMessage translates the LNS DownlinkMessage "dnmsg" to ttnpb.DownlinkMessage.
func (dnmsg *DownlinkMessage) ToDownlinkMessage() ttnpb.DownlinkMessage {
	return ttnpb.DownlinkMessage{
		RawPayload: []byte(dnmsg.Pdu),
		Settings: &ttnpb.DownlinkMessage_Scheduled{
			Scheduled: &ttnpb.TxSettings{
				DataRateIndex: ttnpb.DataRateIndex(dnmsg.DR),
				Frequency:     uint64(dnmsg.Freq),
				Timestamp:     uint32(dnmsg.XTime),
			},
		},
	}
}
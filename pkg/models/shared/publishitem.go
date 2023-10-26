// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PublishItemFormats - Transport-specific message payload representations to be used for delivery. At least one format (`http-response`, `http-stream`, and/or `ws-message`) must be specified. Messages are only delivered to subscribers interested in the provided formats. For example, the `ws-message` format will only be sent to WebSocket clients.
type PublishItemFormats struct {
	// Payload format for delivering to subscribers of whole HTTP responses (`response` hold mode). One of `body` or `body-bin` must be specified.
	HTTPResponse *HTTPResponseFormat `json:"http-response,omitempty"`
	// Payload format for delivering to subscribers of HTTP streaming response bodies (`stream` hold mode). One of `content` or `content-bin` must be specified.
	HTTPStream *HTTPStreamFormat `json:"http-stream,omitempty"`
	// Payload format for delivering to subscribers of WebSocket messages. One of `content` or `content-bin` must be specified.
	WsMessage *WsMessageFormat `json:"ws-message,omitempty"`
}

func (o *PublishItemFormats) GetHTTPResponse() *HTTPResponseFormat {
	if o == nil {
		return nil
	}
	return o.HTTPResponse
}

func (o *PublishItemFormats) GetHTTPStream() *HTTPStreamFormat {
	if o == nil {
		return nil
	}
	return o.HTTPStream
}

func (o *PublishItemFormats) GetWsMessage() *WsMessageFormat {
	if o == nil {
		return nil
	}
	return o.WsMessage
}

// PublishItem - An individual message.
type PublishItem struct {
	// The channel to publish the message on.
	Channel string `json:"channel"`
	// Transport-specific message payload representations to be used for delivery. At least one format (`http-response`, `http-stream`, and/or `ws-message`) must be specified. Messages are only delivered to subscribers interested in the provided formats. For example, the `ws-message` format will only be sent to WebSocket clients.
	Formats PublishItemFormats `json:"formats"`
	// The ID of the message.
	ID *string `json:"id,omitempty"`
	// The ID of the previous message published on the same channel.
	PrevID *string `json:"prev-id,omitempty"`
}

func (o *PublishItem) GetChannel() string {
	if o == nil {
		return ""
	}
	return o.Channel
}

func (o *PublishItem) GetFormats() PublishItemFormats {
	if o == nil {
		return PublishItemFormats{}
	}
	return o.Formats
}

func (o *PublishItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PublishItem) GetPrevID() *string {
	if o == nil {
		return nil
	}
	return o.PrevID
}

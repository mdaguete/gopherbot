package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewConversationsGetConversationMembersParams creates a new ConversationsGetConversationMembersParams object
// with the default values initialized.
func NewConversationsGetConversationMembersParams() *ConversationsGetConversationMembersParams {
	var ()
	return &ConversationsGetConversationMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConversationsGetConversationMembersParamsWithTimeout creates a new ConversationsGetConversationMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConversationsGetConversationMembersParamsWithTimeout(timeout time.Duration) *ConversationsGetConversationMembersParams {
	var ()
	return &ConversationsGetConversationMembersParams{

		timeout: timeout,
	}
}

// NewConversationsGetConversationMembersParamsWithContext creates a new ConversationsGetConversationMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewConversationsGetConversationMembersParamsWithContext(ctx context.Context) *ConversationsGetConversationMembersParams {
	var ()
	return &ConversationsGetConversationMembersParams{

		Context: ctx,
	}
}

/*ConversationsGetConversationMembersParams contains all the parameters to send to the API endpoint
for the conversations get conversation members operation typically these are written to a http.Request
*/
type ConversationsGetConversationMembersParams struct {

	/*ConversationID
	  Conversation ID

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the conversations get conversation members params
func (o *ConversationsGetConversationMembersParams) WithTimeout(timeout time.Duration) *ConversationsGetConversationMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the conversations get conversation members params
func (o *ConversationsGetConversationMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the conversations get conversation members params
func (o *ConversationsGetConversationMembersParams) WithContext(ctx context.Context) *ConversationsGetConversationMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the conversations get conversation members params
func (o *ConversationsGetConversationMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithConversationID adds the conversationID to the conversations get conversation members params
func (o *ConversationsGetConversationMembersParams) WithConversationID(conversationID string) *ConversationsGetConversationMembersParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the conversations get conversation members params
func (o *ConversationsGetConversationMembersParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *ConversationsGetConversationMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

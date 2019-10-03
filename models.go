package main

type Hangouts struct {
	Conversations []Conversations `json:"conversations"`
}
type ConversationID struct {
	ID string `json:"id"`
}
type ID struct {
	ID string `json:"id"`
}
type ParticipantID struct {
	GaiaID string `json:"gaia_id"`
	ChatID string `json:"chat_id"`
}
type SelfReadState struct {
	ParticipantID       ParticipantID `json:"participant_id"`
	LatestReadTimestamp string        `json:"latest_read_timestamp"`
}
type InviterID struct {
	GaiaID string `json:"gaia_id"`
	ChatID string `json:"chat_id"`
}
type DeliveryMedium struct {
	MediumType string `json:"medium_type"`
}
type DeliveryMediumOption struct {
	DeliveryMedium DeliveryMedium `json:"delivery_medium"`
	CurrentDefault bool           `json:"current_default"`
}
type SelfConversationState struct {
	SelfReadState        SelfReadState          `json:"self_read_state"`
	Status               string                 `json:"status"`
	NotificationLevel    string                 `json:"notification_level"`
	View                 []string               `json:"view"`
	InviterID            InviterID              `json:"inviter_id"`
	InviteTimestamp      string                 `json:"invite_timestamp"`
	SortTimestamp        string                 `json:"sort_timestamp"`
	ActiveTimestamp      string                 `json:"active_timestamp"`
	DeliveryMediumOption []DeliveryMediumOption `json:"delivery_medium_option"`
	IsGuest              bool                   `json:"is_guest"`
}
type ReadState struct {
	ParticipantID       ParticipantID `json:"participant_id"`
	LatestReadTimestamp string        `json:"latest_read_timestamp"`
}
type CurrentParticipant struct {
	GaiaID string `json:"gaia_id"`
	ChatID string `json:"chat_id"`
}
type ParticipantData struct {
	ID                             ID     `json:"id"`
	FallbackName                   string `json:"fallback_name"`
	InvitationStatus               string `json:"invitation_status"`
	ParticipantType                string `json:"participant_type"`
	NewInvitationStatus            string `json:"new_invitation_status"`
	InDifferentCustomerAsRequester bool   `json:"in_different_customer_as_requester"`
	DomainID                       string `json:"domain_id"`
}
type Conversation struct {
	ID                     ID                    `json:"id"`
	Type                   string                `json:"type"`
	SelfConversationState  SelfConversationState `json:"self_conversation_state"`
	ReadState              []ReadState           `json:"read_state"`
	HasActiveHangout       bool                  `json:"has_active_hangout"`
	OtrStatus              string                `json:"otr_status"`
	OtrToggle              string                `json:"otr_toggle"`
	CurrentParticipant     []CurrentParticipant  `json:"current_participant"`
	ParticipantData        []ParticipantData     `json:"participant_data"`
	ForkOnExternalInvite   bool                  `json:"fork_on_external_invite"`
	NetworkType            []string              `json:"network_type"`
	ForceHistoryState      string                `json:"force_history_state"`
	GroupLinkSharingStatus string                `json:"group_link_sharing_status"`
}
type ConversationWrapper struct {
	ConversationID ConversationID `json:"conversation_id"`
	Conversation   Conversation   `json:"conversation"`
}
type SenderID struct {
	GaiaID string `json:"gaia_id"`
	ChatID string `json:"chat_id"`
}
type UserID struct {
	GaiaID string `json:"gaia_id"`
	ChatID string `json:"chat_id"`
}

type Segment struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type MessageContent struct {
	Segment []Segment `json:"segment"`
}
type ChatMessage struct {
	MessageContent MessageContent `json:"message_content"`
}
type SelfEventState struct {
	UserID            UserID `json:"user_id"`
	ClientGeneratedID string `json:"client_generated_id"`
	NotificationLevel string `json:"notification_level"`
}
type Event struct {
	ConversationID        ConversationID `json:"conversation_id"`
	SenderID              SenderID       `json:"sender_id"`
	Timestamp             string         `json:"timestamp"`
	SelfEventState        SelfEventState `json:"self_event_state,omitempty"`
	ChatMessage           ChatMessage    `json:"chat_message"`
	EventID               string         `json:"event_id"`
	AdvancesSortTimestamp bool           `json:"advances_sort_timestamp"`
	EventOtr              string         `json:"event_otr"`
	DeliveryMedium        DeliveryMedium `json:"delivery_medium"`
	EventType             string         `json:"event_type"`
	EventVersion          string         `json:"event_version"`
}
type Conversations struct {
	Conversation ConversationWrapper `json:"conversation"`
	Events       []Event             `json:"events"`
}

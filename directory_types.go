package fsxml

type DirectoryParam string

const (
	VertoContext                 DirectoryParam = "verto-context"
	VertoDialplan                DirectoryParam = "verto-dialplan"
	JsonRPCAllowedMethods        DirectoryParam = "jsonrpc-allowed-methods"
	JsonRPCAllowedEventsChannels DirectoryParam = "jsonrpc-allowed-event-channels"
)

type UserParam string

const (
	Password        UserParam = "password"
	ReverseAuthUser UserParam = "reverse-auth-user"
	ReverseAuthPass UserParam = "reverse-auth-pass"
	MWIAccount      UserParam = "MWI-Account"
	DialString      UserParam = "dial-string"
	AOneHash        UserParam = "a1-hash"
)

type UserVariable string

const (
	UserContext               UserVariable = "user_context"
	EffectiveCallerIDName     UserVariable = "effective_caller_id_name"
	EffectiveCallerIDNumber   UserVariable = "effective_caller_id_number"
	OutboundCallerIDNumber    UserVariable = "outbound_caller_id_number"
	OutboundCallerIDName      UserVariable = "outbound_caller_id_name"
	TollAllow                 UserVariable = "toll_allow"
	Mailbox                   UserVariable = "mailbox"
	TakeVM                    UserVariable = "take_vm"
	DomainName                UserVariable = "domain_name"
	HoldMusic                 UserVariable = "hold_music"
	MaxCalls                  UserVariable = "max_calls"
	PresenceID                UserVariable = "presence_id"
	DefaultAreaCode           UserVariable = "default_area_code"
	UserLanguage              UserVariable = "user_language"
	DirectoryFullName         UserVariable = "directory_full_name"
	DirectoryVisible          UserVariable = "directory-visible"
	RecordStereo              UserVariable = "record_stereo"
	TransferFallbackExtension UserVariable = "transfer_fallback_extension"
	CallGroup                 UserVariable = "callgroup"
	RecordCall                UserVariable = "record_call"
	DnD                       UserVariable = "dnd"
	SipForceExpires           UserVariable = "sip-force-expires"
	UserGateway               UserVariable = "user_gateway"
	RTPSecureMedia            UserVariable = "rtp_secure_media"
	TenantUUID                UserVariable = "tenant_uuid"
	UserUUID                  UserVariable = "user_uuid"
)

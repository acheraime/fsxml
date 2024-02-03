package fsxml

type DialplanApplication string

const (
	SetApplication        DialplanApplication = "set"
	PlaybackApplication   DialplanApplication = "playback"
	CallCenterApplication DialplanApplication = "callcenter"
	VoicemailApplication  DialplanApplication = "voicemail"
	TransferApplication   DialplanApplication = "transfer"
	HangupApplication     DialplanApplication = "hangup"
	BridgeApplication     DialplanApplication = "bridge"
	SleepApplication      DialplanApplication = "sleep"
	AnswerApplication     DialplanApplication = "answer"
	ConferenceApplication DialplanApplication = "conference"
	ExportApplication     DialplanApplication = "export"
	LogApplication        DialplanApplication = "log"
	BindMetaApplication   DialplanApplication = "bind_meta_app"
)

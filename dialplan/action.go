package dialplan

import (
	"encoding/xml"
	"strconv"

	"github.com/acheraime/fsxml"
)

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

type Action struct {
	XMLName     xml.Name            `xml:"action"`
	Application DialplanApplication `xml:"application,attr"`
	Data        *string             `xml:"data,attr,omitempty"`
	Inline      string              `xml:"inline,attr,omitempty"`
}

func NewAction(application DialplanApplication, data string) Action {
	return Action{Application: application, Data: fsxml.String(data)}
}

func (a *Action) SetInline(flag bool) {
	a.Inline = strconv.FormatBool(flag)
}

func Set(data string) Action {
	return NewAction(SetApplication, data)
}

func Playback(data string) Action {
	return NewAction(PlaybackApplication, data)
}

func Hangup() Action {
	return NewAction(HangupApplication, "")
}

func CallCenter(data string) Action {
	return NewAction(CallCenterApplication, data)
}

func Bridge(data string) Action {
	return NewAction(BridgeApplication, data)
}

func Transfer(data string) Action {
	return NewAction(TransferApplication, data)
}

func Voicemail(data string) Action {
	return NewAction(VoicemailApplication, data)
}

func Sleep(data int) Action {
	return NewAction(SleepApplication, strconv.Itoa(data))
}

func Answer() Action {
	return NewAction(AnswerApplication, "")
}

func Conference(data string) Action {
	return NewAction(ConferenceApplication, data)
}

func Export(data string) Action {
	return NewAction(ExportApplication, data)
}

func Log(data string) Action {
	return NewAction(LogApplication, data)
}

func BindMetaApp(data string) Action {
	return NewAction(BindMetaApplication, data)
}

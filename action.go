package fsxml

import (
	"encoding/xml"
	"strconv"
)

type Action struct {
	XMLName     xml.Name            `xml:"action"`
	Application DialplanApplication `xml:"application,attr"`
	Data        *string             `xml:"data,attr,omitempty"`
	Inline      string              `xml:"inline,attr,omitempty"`
}

func NewAction(application DialplanApplication, data string) Action {
	return Action{Application: application, Data: String(data)}
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

func WithCallCenter(data string) Action {
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

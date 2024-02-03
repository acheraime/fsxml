package directory

import (
	"encoding/xml"
	"strconv"

	"github.com/acheraime/fsxml"
)

type User struct {
	XMLName   xml.Name        `xml:"user"`
	ID        string          `xml:"id,attr"`
	Mailbox   string          `xml:"mailbox,attr,omitempty"`
	Params    fsxml.Params    `xml:"params>param,omitempty"`
	Variables fsxml.Variables `xml:"variables>variable,omitempty"`
}

func (u *User) addParam(name UserParam, value string) {
	param := fsxml.Param{
		Name:  string(name),
		Value: value,
	}

	u.Params.Add(param)
}

func (u *User) addVariable(name UserVariable, value string) {
	variable := fsxml.Variable{
		Name:  string(name),
		Value: value,
	}

	u.Variables.Add(variable)
}

func (u *User) SetPasswordParam(password string) {
	u.addParam(Password, password)
}

func (u *User) SetReverseAuthUserParam(user string) {
	u.addParam(ReverseAuthUser, user)
}

func (u *User) SetReverseAuthPassParam(password string) {
	u.addParam(ReverseAuthPass, password)
}

func (u *User) SetMWIAccountParam(mwi string) {
	u.addParam(MWIAccount, mwi)
}

func (u *User) SetDialStringParam(dialString string) {
	u.addParam(DialString, dialString)
}

func (u *User) SetOneHashParam(hash string) {
	u.addParam(AOneHash, hash)
}

func (u *User) SetUserContext(context string) {
	u.addVariable(UserContext, context)
}

func (u *User) SetEffectiveCallerIDName(name string) {
	u.addVariable(EffectiveCallerIDName, name)
}

func (u *User) SetEffectiveCallerIDNumber(number string) {
	u.addVariable(EffectiveCallerIDNumber, number)
}

func (u *User) SetOutboundCallerIDName(name string) {
	u.addVariable(OutboundCallerIDName, name)
}

func (u *User) SetOutboundCallerIDNumber(number string) {
	u.addVariable(OutboundCallerIDNumber, number)
}

func (u *User) SetTollAllow(allow string) {
	u.addVariable(TollAllow, allow)
}

func (u *User) SetMailbox(mailbox string) {
	u.addVariable(Mailbox, mailbox)
}

func (u *User) SetTakeVm(flag bool) {
	u.addVariable(TakeVM, strconv.FormatBool(flag))
}

func (u *User) SetDomainName(domain string) {
	u.addVariable(DomainName, domain)
}

func (u *User) SetHoldMusic(moh string) {
	u.addVariable(HoldMusic, moh)
}

func (u *User) SetMaxCall(value int) {
	u.addVariable(MaxCalls, strconv.Itoa(value))
}

func (u *User) SetPresenceID(presenceID string) {
	u.addVariable(PresenceID, presenceID)
}

func (u *User) SetDefaultAreaCode(code string) {
	u.addVariable(DefaultAreaCode, code)
}

func (u *User) SetUserLanguage(lang string) {
	u.addVariable(UserLanguage, lang)
}

func (u *User) SetDirectoryFullName(name string) {
	u.addVariable(DirectoryFullName, name)
}

func (u *User) SetDirectoryVisible(flag bool) {
	u.addVariable(DirectoryVisible, strconv.FormatBool(flag))
}

func (u *User) SetRecordStereo(flag bool) {
	u.addVariable(RecordStereo, strconv.FormatBool(flag))
}

func (u *User) SetTrFallbackExten(extension string) {
	u.addVariable(TransferFallbackExtension, extension)
}

func (u *User) SetCallGroup(group string) {
	u.addVariable(CallGroup, group)
}

func (u *User) SetRecordCall(flag bool) {
	u.addVariable(RecordCall, strconv.FormatBool(flag))
}

func (u *User) SetDnD(flag bool) {
	u.addVariable(DnD, strconv.FormatBool(flag))
}

func (u *User) SetSIPForceExpire(flag bool) {
	u.addVariable(SipForceExpires, strconv.FormatBool(flag))
}

func (u *User) SetUserGateway(gateway string) {
	u.addVariable(UserGateway, gateway)
}

func (u *User) SetRTPSecureMedia(flag bool) {
	u.addVariable(RTPSecureMedia, strconv.FormatBool(flag))
}

func (u *User) SetTenantUUID(uuid string) {
	u.addVariable(TenantUUID, uuid)
}

func (u *User) SetUserUUID(uuid string) {
	u.addVariable(UserUUID, uuid)
}

func NewUser(userId string) User {
	defaultPass := "1878adb"
	u := User{ID: userId, Mailbox: userId}
	// Set defaults
	u.SetPasswordParam(defaultPass)
	u.SetReverseAuthUserParam(userId)
	u.SetReverseAuthPassParam(defaultPass)
	u.SetUserContext("default")
	u.SetMailbox(userId)

	return u
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// SkypeForBusinessParticipantActivityUserCounts undocumented
type SkypeForBusinessParticipantActivityUserCounts struct {
	// Entity is the base model of SkypeForBusinessParticipantActivityUserCounts
	Entity
	// Im undocumented
	Im *int `json:"im,omitempty"`
	// AudioVideo undocumented
	AudioVideo *int `json:"audioVideo,omitempty"`
	// AppSharing undocumented
	AppSharing *int `json:"appSharing,omitempty"`
	// Web undocumented
	Web *int `json:"web,omitempty"`
	// DialInOut3rdParty undocumented
	DialInOut3rdParty *int `json:"dialInOut3rdParty,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *time.Time `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
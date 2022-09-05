// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Member struct {
	ID          string      `json:"id"`
	TwitchID    string      `json:"twitch_id"`
	Login       string      `json:"login"`
	DisplayName string      `json:"display_name"`
	Color       string      `json:"color"`
	Pfp         string      `json:"pfp"`
	Links       *MemberLink `json:"links"`
}

type MemberConnection struct {
	Members  []*Member `json:"members"`
	PageInfo *PageInfo `json:"page_info"`
}

type MemberLink struct {
	Twitch        string `json:"twitch"`
	Reddit        string `json:"reddit"`
	Instagram     string `json:"instagram"`
	Twitter       string `json:"twitter"`
	Discord       string `json:"discord"`
	Youtube       string `json:"youtube"`
	Tiktok        string `json:"tiktok"`
	VrchatLegends string `json:"vrchat_legends"`
}

type PageInfo struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	PerPage   int64 `json:"perPage"`
	Prev      int64 `json:"prev"`
	Next      int64 `json:"next"`
	TotalPage int64 `json:"totalPage"`
}

type StatEntry struct {
	TwitchID        string `json:"twitch_id"`
	Rank            int64  `json:"rank"`
	MinutesStreamed int64  `json:"minutes_streamed"`
	AvgViewers      int64  `json:"avg_viewers"`
	MaxViewers      int64  `json:"max_viewers"`
	HoursWatched    int64  `json:"hours_watched"`
	Followers       int64  `json:"followers"`
	Views           int64  `json:"views"`
	FollowersTotal  int64  `json:"followers_total"`
	ViewsTotal      int64  `json:"views_total"`
}

type Stream struct {
	ID        string       `json:"id"`
	TwitchID  string       `json:"twitch_id"`
	Member    *Member      `json:"member"`
	Status    StreamStatus `json:"status"`
	Title     string       `json:"title"`
	GameID    string       `json:"game_id"`
	Game      string       `json:"game"`
	Viewers   int          `json:"viewers"`
	Thumbnail string       `json:"thumbnail"`
	StartedAt string       `json:"started_at"`
}

type StreamConnection struct {
	Streams  []*Stream `json:"streams"`
	PageInfo *PageInfo `json:"page_info"`
}

type MemberSort string

const (
	MemberSortAz MemberSort = "AZ"
	MemberSortZa MemberSort = "ZA"
)

var AllMemberSort = []MemberSort{
	MemberSortAz,
	MemberSortZa,
}

func (e MemberSort) IsValid() bool {
	switch e {
	case MemberSortAz, MemberSortZa:
		return true
	}
	return false
}

func (e MemberSort) String() string {
	return string(e)
}

func (e *MemberSort) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MemberSort(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MemberSort", str)
	}
	return nil
}

func (e MemberSort) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StreamStatus string

const (
	StreamStatusOnline  StreamStatus = "ONLINE"
	StreamStatusOffline StreamStatus = "OFFLINE"
	StreamStatusAll     StreamStatus = "ALL"
)

var AllStreamStatus = []StreamStatus{
	StreamStatusOnline,
	StreamStatusOffline,
	StreamStatusAll,
}

func (e StreamStatus) IsValid() bool {
	switch e {
	case StreamStatusOnline, StreamStatusOffline, StreamStatusAll:
		return true
	}
	return false
}

func (e StreamStatus) String() string {
	return string(e)
}

func (e *StreamStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StreamStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StreamStatus", str)
	}
	return nil
}

func (e StreamStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

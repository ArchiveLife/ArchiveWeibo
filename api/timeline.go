package api

import (
	"fmt"

	"encoding/json"

	"github.com/imroc/req"
)

// GetTimeLine for current user, need the 'SUB' part of cookie
func (api *WeiboAPI) GetTimeLine(cookieSub string, recentBlogId string) (*WeiboTimeLine, error) {
	res, err := req.Get(
		"https://m.weibo.cn/feed/friends",
		req.QueryParam{
			"max_id": recentBlogId,
		},
		req.Header{
			"Referer":    "https://m.weibo.cn/",
			"MWeibo-Pwa": "1",
			"Cookie":     fmt.Sprintf("SUB=%s", cookieSub),
		},
	)
	if err != nil {
		return nil, err
	}
	body := &WeiboTimeLine{}
	if err = res.ToJSON(body); err != nil {
		return nil, err
	}
	return body, nil
}

func UnmarshalWeiboTimeLine(data []byte) (WeiboTimeLine, error) {
	var r WeiboTimeLine
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WeiboTimeLine) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WeiboTimeLine struct {
	Data     TimeLineData `json:"data"`
	Ok       int64        `json:"ok"`
	HTTPCode int64        `json:"http_code"`
}

type TimeLineData struct {
	Statuses          []Status      `json:"statuses"`
	Advertises        []interface{} `json:"advertises"`
	Ad                []interface{} `json:"ad"`
	Hasvisible        bool          `json:"hasvisible"`
	PreviousCursor    int64         `json:"previous_cursor"`
	NextCursor        int64         `json:"next_cursor"`
	PreviousCursorStr string        `json:"previous_cursor_str"`
	NextCursorStr     string        `json:"next_cursor_str"`
	TotalNumber       int64         `json:"total_number"`
	Interval          int64         `json:"interval"`
	UveBlank          int64         `json:"uve_blank"`
	SinceID           int64         `json:"since_id"`
	SinceIDStr        string        `json:"since_id_str"`
	MaxID             int64         `json:"max_id"`
	MaxIDStr          string        `json:"max_id_str"`
	HasUnread         int64         `json:"has_unread"`
}

type Status struct {
	Visible                  Visible                `json:"visible"`
	CreatedAt                string                 `json:"created_at"`
	ID                       string                 `json:"id"`
	Mid                      string                 `json:"mid"`
	CanEdit                  bool                   `json:"can_edit"`
	Version                  *int64                 `json:"version,omitempty"`
	ShowAdditionalIndication int64                  `json:"show_additional_indication"`
	Text                     string                 `json:"text"`
	TextLength               *int64                 `json:"textLength,omitempty"`
	Source                   string                 `json:"source"`
	Favorited                bool                   `json:"favorited"`
	PicIDS                   []string               `json:"pic_ids"`
	PicTypes                 string                 `json:"pic_types"`
	ThumbnailPic             *string                `json:"thumbnail_pic,omitempty"`
	BmiddlePic               *string                `json:"bmiddle_pic,omitempty"`
	OriginalPic              *string                `json:"original_pic,omitempty"`
	IsPaid                   bool                   `json:"is_paid"`
	MblogVipType             int64                  `json:"mblog_vip_type"`
	User                     StatusUser             `json:"user"`
	PicStatus                *string                `json:"picStatus,omitempty"`
	RepostsCount             int64                  `json:"reposts_count"`
	CommentsCount            int64                  `json:"comments_count"`
	AttitudesCount           int64                  `json:"attitudes_count"`
	PendingApprovalCount     int64                  `json:"pending_approval_count"`
	IsLongText               bool                   `json:"isLongText"`
	RewardExhibitionType     int64                  `json:"reward_exhibition_type"`
	HideFlag                 int64                  `json:"hide_flag"`
	Mlevel                   int64                  `json:"mlevel"`
	DarwinTags               []DarwinTag            `json:"darwin_tags"`
	Mblogtype                int64                  `json:"mblogtype"`
	Rid                      string                 `json:"rid"`
	MoreInfoType             int64                  `json:"more_info_type"`
	EnableCommentGuide       *bool                  `json:"enable_comment_guide,omitempty"`
	ContentAuth              int64                  `json:"content_auth"`
	PicNum                   int64                  `json:"pic_num"`
	AlchemyParams            AlchemyParams          `json:"alchemy_params"`
	PageInfo                 *StatusPageInfo        `json:"page_info,omitempty"`
	Pics                     []StatusPic            `json:"pics,omitempty"`
	Bid                      string                 `json:"bid"`
	EditCount                *int64                 `json:"edit_count,omitempty"`
	EditAt                   *string                `json:"edit_at,omitempty"`
	PicFocusPoint            []PicFocusPoint        `json:"pic_focus_point,omitempty"`
	PicRectangleObject       []interface{}          `json:"pic_rectangle_object,omitempty"`
	PicFlag                  *int64                 `json:"pic_flag,omitempty"`
	TopicID                  *string                `json:"topic_id,omitempty"`
	SyncMblog                *bool                  `json:"sync_mblog,omitempty"`
	IsImportedTopic          *bool                  `json:"is_imported_topic,omitempty"`
	Cardid                   *string                `json:"cardid,omitempty"`
	NumberDisplayStrategy    *NumberDisplayStrategy `json:"number_display_strategy,omitempty"`
	RewardScheme             *string                `json:"reward_scheme,omitempty"`
	PID                      *int64                 `json:"pid,omitempty"`
	Pidstr                   *string                `json:"pidstr,omitempty"`
	RetweetedStatus          *RetweetedStatus       `json:"retweeted_status,omitempty"`
	RepostType               *int64                 `json:"repost_type,omitempty"`
	RawText                  *string                `json:"raw_text,omitempty"`
	SafeTags                 *int64                 `json:"safe_tags,omitempty"`
	Fid                      *int64                 `json:"fid,omitempty"`
}

type DarwinTag struct {
	ObjectType    string      `json:"object_type"`
	ObjectID      string      `json:"object_id"`
	DisplayName   string      `json:"display_name"`
	EnterpriseUid interface{} `json:"enterprise_uid"`
	PCURL         string      `json:"pc_url"`
	MAPIURL       string      `json:"mapi_url"`
	BdObjectType  string      `json:"bd_object_type"`
}

type NumberDisplayStrategy struct {
	ApplyScenarioFlag    int64       `json:"apply_scenario_flag"`
	DisplayTextMinNumber int64       `json:"display_text_min_number"`
	DisplayText          DisplayText `json:"display_text"`
}

type StatusPageInfo struct {
	Type             string        `json:"type"`
	ObjectType       int64         `json:"object_type"`
	PagePic          PurplePagePic `json:"page_pic"`
	PageURL          string        `json:"page_url"`
	PageTitle        string        `json:"page_title"`
	Content1         string        `json:"content1"`
	URLOri           *string       `json:"url_ori,omitempty"`
	ObjectID         *string       `json:"object_id,omitempty"`
	Title            *string       `json:"title,omitempty"`
	Content2         *string       `json:"content2,omitempty"`
	VideoOrientation *string       `json:"video_orientation,omitempty"`
	PlayCount        *string       `json:"play_count,omitempty"`
	MediaInfo        *MediaInfo    `json:"media_info,omitempty"`
	Urls             *Urls         `json:"urls,omitempty"`
}

type PurplePagePic struct {
	URL         string  `json:"url"`
	Width       *int64  `json:"width,omitempty"`
	Height      *int64  `json:"height,omitempty"`
	PID         *string `json:"pid,omitempty"`
	Source      *int64  `json:"source,omitempty"`
	IsSelfCover *int64  `json:"is_self_cover,omitempty"`
	Type        *int64  `json:"type,omitempty"`
}

type PicFocusPoint struct {
	FocusPoint FocusPoint `json:"focus_point"`
	PicID      string     `json:"pic_id"`
}

type FocusPoint struct {
	Left   float64 `json:"left"`
	Top    float64 `json:"top"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type StatusPic struct {
	PID   string     `json:"pid"`
	URL   string     `json:"url"`
	Size  string     `json:"size"`
	Geo   PurpleGeo  `json:"geo"`
	Large LargeClass `json:"large"`
}

type PurpleGeo struct {
	Croped bool `json:"croped"`
}

type LargeClass struct {
	Size string   `json:"size"`
	URL  string   `json:"url"`
	Geo  LargeGeo `json:"geo"`
}

type RetweetedStatusPageInfo struct {
	Type       string        `json:"type"`
	ObjectType int64         `json:"object_type"`
	PagePic    FluffyPagePic `json:"page_pic"`
	PageURL    string        `json:"page_url"`
	PageTitle  string        `json:"page_title"`
	Content1   string        `json:"content1"`
	Content2   string        `json:"content2"`
}

type FluffyPagePic struct {
	URL string `json:"url"`
}

type RetweetedStatusPic struct {
	PID   string     `json:"pid"`
	URL   string     `json:"url"`
	Size  string     `json:"size"`
	Geo   FluffyGeo  `json:"geo"`
	Large LargeClass `json:"large"`
}

type FluffyGeo struct {
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
	Croped bool  `json:"croped"`
}

type RetweetedStatusUser struct {
	ID              int64            `json:"id"`
	ScreenName      string           `json:"screen_name"`
	ProfileImageURL string           `json:"profile_image_url"`
	ProfileURL      string           `json:"profile_url"`
	StatusesCount   int64            `json:"statuses_count"`
	Verified        bool             `json:"verified"`
	VerifiedType    int64            `json:"verified_type"`
	VerifiedTypeEXT int64            `json:"verified_type_ext"`
	VerifiedReason  string           `json:"verified_reason"`
	CloseBlueV      bool             `json:"close_blue_v"`
	Description     string           `json:"description"`
	Gender          Gender           `json:"gender"`
	Mbtype          int64            `json:"mbtype"`
	Urank           int64            `json:"urank"`
	Mbrank          int64            `json:"mbrank"`
	FollowMe        bool             `json:"follow_me"`
	Following       bool             `json:"following"`
	FollowersCount  int64            `json:"followers_count"`
	FollowCount     int64            `json:"follow_count"`
	CoverImagePhone string           `json:"cover_image_phone"`
	AvatarHD        string           `json:"avatar_hd"`
	Like            bool             `json:"like"`
	LikeMe          bool             `json:"like_me"`
	Badge           map[string]int64 `json:"badge"`
}

type StatusUser struct {
	ID              int64            `json:"id"`
	ScreenName      string           `json:"screen_name"`
	ProfileImageURL string           `json:"profile_image_url"`
	ProfileURL      string           `json:"profile_url"`
	StatusesCount   int64            `json:"statuses_count"`
	Verified        bool             `json:"verified"`
	VerifiedType    int64            `json:"verified_type"`
	VerifiedTypeEXT *int64           `json:"verified_type_ext,omitempty"`
	VerifiedReason  *string          `json:"verified_reason,omitempty"`
	CloseBlueV      bool             `json:"close_blue_v"`
	Description     string           `json:"description"`
	Gender          Gender           `json:"gender"`
	Mbtype          int64            `json:"mbtype"`
	Urank           int64            `json:"urank"`
	Mbrank          int64            `json:"mbrank"`
	FollowMe        bool             `json:"follow_me"`
	Following       bool             `json:"following"`
	FollowersCount  int64            `json:"followers_count"`
	FollowCount     int64            `json:"follow_count"`
	CoverImagePhone string           `json:"cover_image_phone"`
	AvatarHD        string           `json:"avatar_hd"`
	Like            bool             `json:"like"`
	LikeMe          bool             `json:"like_me"`
	Badge           map[string]int64 `json:"badge"`
	Remark          *string          `json:"remark,omitempty"`
}

type DisplayText string

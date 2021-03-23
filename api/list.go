package api

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/imroc/req"
)

func (api *WeiboAPI) GetUserPagesIndex(uid string, page int) (*WeiboUserListPageIndex, error) {
	containerId, err := api.GetContainerId(uid)
	if err != nil {
		return nil, err
	}
	res, err := req.Get(
		"https://m.weibo.cn/api/container/getIndex",
		req.QueryParam{
			"type":        "uid",
			"value":       uid,
			"containerid": containerId,
		},
		req.Header{
			"Referer":    "https://m.weibo.cn/",
			"MWeibo-Pwa": "1",
		},
	)
	if err != nil {
		return nil, err
	}

	body := &WeiboUserListPageIndex{}

	if err = res.ToJSON(body); err != nil {
		return nil, err
	}

	return body, nil
}

func UnmarshalWeiboUserListPageIndex(data []byte) (WeiboUserListPageIndex, error) {
	var r WeiboUserListPageIndex
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WeiboUserListPageIndex) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WeiboUserListPageIndex struct {
	Ok   int64       `json:"ok"`
	Data ListPageDat `json:"data"`
}

type ListPageDat struct {
	CardlistInfo CardlistInfo `json:"cardlistInfo"`
	Cards        []Card       `json:"cards"`
	Scheme       string       `json:"scheme"`
	ShowAppTips  int64        `json:"showAppTips"`
}

type CardlistInfo struct {
	Containerid string `json:"containerid"`
	VP          int64  `json:"v_p"`
	ShowStyle   int64  `json:"show_style"`
	Total       int64  `json:"total"`
	SinceID     int64  `json:"since_id"`
}

type Card struct {
	CardType       int64       `json:"card_type"`
	CardStyle      *int64      `json:"card_style,omitempty"`
	DisplayArrow   *int64      `json:"display_arrow,omitempty"`
	SkipGroupTitle *bool       `json:"skip_group_title,omitempty"`
	CardGroup      []CardGroup `json:"card_group,omitempty"`
	Itemid         *string     `json:"itemid,omitempty"`
	Scheme         *string     `json:"scheme,omitempty"`
	Mblog          *Mblog      `json:"mblog,omitempty"`
}

type CardGroup struct {
	CardType        int64               `json:"card_type"`
	Col             *int64              `json:"col,omitempty"`
	Group           []Group             `json:"group,omitempty"`
	Desc            *string             `json:"desc,omitempty"`
	Scheme          *string             `json:"scheme,omitempty"`
	Actionlog       *CardGroupActionlog `json:"actionlog,omitempty"`
	DisplayArrow    *int64              `json:"display_arrow,omitempty"`
	TitleExtraText  *string             `json:"title_extra_text,omitempty"`
	Itemid          *string             `json:"itemid,omitempty"`
	BackgroundColor *int64              `json:"background_color,omitempty"`
	Openurl         *string             `json:"openurl,omitempty"`
	RecomRemark     *string             `json:"recom_remark,omitempty"`
	Recommend       *string             `json:"recommend,omitempty"`
	Desc1           *string             `json:"desc1,omitempty"`
	Desc2           *string             `json:"desc2,omitempty"`
	User            *User               `json:"user,omitempty"`
	Buttons         []Button            `json:"buttons,omitempty"`
}

type CardGroupActionlog struct {
	ActCode     *ActCode `json:"act_code"`
	Cardid      string   `json:"cardid"`
	OID         *string  `json:"oid,omitempty"`
	Featurecode *string  `json:"featurecode,omitempty"`
	Mark        *string  `json:"mark,omitempty"`
	EXT         *string  `json:"ext,omitempty"`
	Uicode      *string  `json:"uicode,omitempty"`
	Luicode     *string  `json:"luicode,omitempty"`
	Fid         *string  `json:"fid,omitempty"`
	Lfid        *string  `json:"lfid,omitempty"`
	Lcardid     *string  `json:"lcardid,omitempty"`
}

type Button struct {
	Type       string          `json:"type"`
	SubType    int64           `json:"sub_type"`
	Name       string          `json:"name"`
	SkipFormat int64           `json:"skip_format"`
	Params     Params          `json:"params"`
	Actionlog  ButtonActionlog `json:"actionlog"`
	Scheme     string          `json:"scheme"`
}

type ButtonActionlog struct {
	ActCode     string `json:"act_code"`
	Cardid      string `json:"cardid"`
	OID         string `json:"oid"`
	Featurecode string `json:"featurecode"`
	Mark        string `json:"mark"`
	EXT         string `json:"ext"`
	Luicode     string `json:"luicode"`
	Fid         string `json:"fid"`
	Lfid        string `json:"lfid"`
	Lcardid     string `json:"lcardid"`
}

type Params struct {
	Uid            int64  `json:"uid"`
	NeedFollow     int64  `json:"need_follow"`
	TrendEXT       string `json:"trend_ext"`
	TrendType      int64  `json:"trend_type"`
	Itemid         int64  `json:"itemid"`
	AllowReplenish int64  `json:"allow_replenish"`
	APIType        string `json:"api_type"`
}

type Group struct {
	TitleSub   string    `json:"title_sub"`
	WordScheme string    `json:"word_scheme"`
	Scheme     string    `json:"scheme"`
	ActionLog  ActionLog `json:"action_log"`
	Icon       *string   `json:"icon,omitempty"`
}

type ActionLog struct {
	ActCode int64  `json:"act_code"`
	OID     string `json:"oid"`
	EXT     string `json:"ext"`
	Fid     string `json:"fid"`
}

type User struct {
	ID              int64            `json:"id"`
	ScreenName      string           `json:"screen_name"`
	ProfileImageURL string           `json:"profile_image_url"`
	ProfileURL      string           `json:"profile_url"`
	StatusesCount   *int64           `json:"statuses_count"`
	Verified        bool             `json:"verified"`
	VerifiedType    int64            `json:"verified_type"`
	VerifiedTypeEXT int64            `json:"verified_type_ext"`
	VerifiedReason  string           `json:"verified_reason"`
	CloseBlueV      bool             `json:"close_blue_v"`
	Description     string           `json:"description"`
	Gender          Gender           `json:"gender"`
	Mbtype          *int64           `json:"mbtype"`
	Urank           *int64           `json:"urank"`
	Mbrank          *int64           `json:"mbrank"`
	FollowMe        *bool            `json:"follow_me"`
	Following       *bool            `json:"following"`
	FollowersCount  int64            `json:"followers_count"`
	FollowCount     int64            `json:"follow_count"`
	CoverImagePhone string           `json:"cover_image_phone"`
	AvatarHD        string           `json:"avatar_hd"`
	Desc1           *string          `json:"desc1,omitempty"`
	Desc2           interface{}      `json:"desc2"`
	Like            *bool            `json:"like,omitempty"`
	LikeMe          *bool            `json:"like_me,omitempty"`
	Badge           map[string]int64 `json:"badge,omitempty"`
}

type Mblog struct {
	Visible                  *Visible         `json:"visible,omitempty"`
	CreatedAt                *string          `json:"created_at,omitempty"`
	ID                       *string          `json:"id,omitempty"`
	Mid                      *string          `json:"mid,omitempty"`
	EditCount                *int64           `json:"edit_count,omitempty"`
	CanEdit                  *bool            `json:"can_edit,omitempty"`
	EditAt                   *string          `json:"edit_at,omitempty"`
	Version                  *int64           `json:"version,omitempty"`
	ShowAdditionalIndication *int64           `json:"show_additional_indication,omitempty"`
	Text                     *string          `json:"text,omitempty"`
	TextLength               *int64           `json:"textLength,omitempty"`
	Source                   *Source          `json:"source,omitempty"`
	Favorited                *bool            `json:"favorited,omitempty"`
	PicIDS                   []string         `json:"pic_ids,omitempty"`
	PicTypes                 *string          `json:"pic_types,omitempty"`
	ThumbnailPic             *string          `json:"thumbnail_pic,omitempty"`
	BmiddlePic               *string          `json:"bmiddle_pic,omitempty"`
	OriginalPic              *string          `json:"original_pic,omitempty"`
	IsPaid                   *bool            `json:"is_paid,omitempty"`
	MblogVipType             *int64           `json:"mblog_vip_type,omitempty"`
	User                     *User            `json:"user,omitempty"`
	PicStatus                *string          `json:"picStatus,omitempty"`
	RepostsCount             *int64           `json:"reposts_count,omitempty"`
	CommentsCount            *int64           `json:"comments_count,omitempty"`
	AttitudesCount           *int64           `json:"attitudes_count,omitempty"`
	PendingApprovalCount     *int64           `json:"pending_approval_count,omitempty"`
	IsLongText               *bool            `json:"isLongText,omitempty"`
	RewardExhibitionType     *int64           `json:"reward_exhibition_type,omitempty"`
	HideFlag                 *int64           `json:"hide_flag,omitempty"`
	Mlevel                   *int64           `json:"mlevel,omitempty"`
	DarwinTags               []interface{}    `json:"darwin_tags,omitempty"`
	Mblogtype                *int64           `json:"mblogtype,omitempty"`
	Rid                      *string          `json:"rid,omitempty"`
	MoreInfoType             *int64           `json:"more_info_type,omitempty"`
	ExternSafe               *int64           `json:"extern_safe,omitempty"`
	ContentAuth              *int64           `json:"content_auth,omitempty"`
	SafeTags                 *int64           `json:"safe_tags,omitempty"`
	PicNum                   *int64           `json:"pic_num,omitempty"`
	AlchemyParams            *AlchemyParams   `json:"alchemy_params,omitempty"`
	MblogMenuNewStyle        *int64           `json:"mblog_menu_new_style,omitempty"`
	EditConfig               *MblogEditConfig `json:"edit_config,omitempty"`
	IsTop                    *int64           `json:"isTop,omitempty"`
	PageInfo                 *PageInfo        `json:"page_info,omitempty"`
	Pics                     []Pic            `json:"pics,omitempty"`
	Title                    *Title           `json:"title,omitempty"`
	Bid                      *string          `json:"bid,omitempty"`
	RetweetedStatus          *RetweetedStatus `json:"retweeted_status,omitempty"`
	RepostType               *int64           `json:"repost_type,omitempty"`
	RawText                  *string          `json:"raw_text,omitempty"`
	Fid                      *int64           `json:"fid,omitempty"`
}

type AlchemyParams struct {
	UgRedEnvelope bool `json:"ug_red_envelope"`
}

type MblogEditConfig struct {
	Edited          bool             `json:"edited"`
	MenuEditHistory *MenuEditHistory `json:"menu_edit_history,omitempty"`
}

type MenuEditHistory struct {
	Scheme string `json:"scheme"`
	Title  string `json:"title"`
}

type PageInfo struct {
	Type             Type       `json:"type"`
	ObjectType       int64      `json:"object_type"`
	PagePic          PagePic    `json:"page_pic"`
	PageURL          string     `json:"page_url"`
	PageTitle        string     `json:"page_title"`
	Content1         string     `json:"content1"`
	URLOri           *string    `json:"url_ori,omitempty"`
	ObjectID         *string    `json:"object_id,omitempty"`
	Title            *string    `json:"title,omitempty"`
	Content2         *string    `json:"content2,omitempty"`
	VideoOrientation *string    `json:"video_orientation,omitempty"`
	PlayCount        *string    `json:"play_count,omitempty"`
	MediaInfo        *MediaInfo `json:"media_info,omitempty"`
	Urls             *Urls      `json:"urls,omitempty"`
}

type MediaInfo struct {
	StreamURL   string  `json:"stream_url"`
	StreamURLHD string  `json:"stream_url_hd"`
	Duration    float64 `json:"duration"`
}

type PagePic struct {
	URL         string  `json:"url"`
	Width       *int64  `json:"width,omitempty"`
	PID         *string `json:"pid,omitempty"`
	Source      *int64  `json:"source,omitempty"`
	IsSelfCover *int64  `json:"is_self_cover,omitempty"`
	Type        *int64  `json:"type,omitempty"`
	Height      *int64  `json:"height,omitempty"`
}

type Urls struct {
	Mp4720PMp4 string  `json:"mp4_720p_mp4"`
	Mp4HDMp4   string  `json:"mp4_hd_mp4"`
	Mp4LdMp4   string  `json:"mp4_ld_mp4"`
	HevcMp4HD  *string `json:"hevc_mp4_hd,omitempty"`
}

type Pic struct {
	PID   string `json:"pid"`
	URL   string `json:"url"`
	Size  string `json:"size"`
	Geo   PicGeo `json:"geo"`
	Large Large  `json:"large"`
}

type PicGeo struct {
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
	Croped bool  `json:"croped"`
}

type Large struct {
	Size string   `json:"size"`
	URL  string   `json:"url"`
	Geo  LargeGeo `json:"geo"`
}

type LargeGeo struct {
	Width  string `json:"width"`
	Height string `json:"height"`
	Croped bool   `json:"croped"`
}

type RetweetedStatus struct {
	Visible                  Visible                   `json:"visible"`
	CreatedAt                string                    `json:"created_at"`
	ID                       string                    `json:"id"`
	Mid                      string                    `json:"mid"`
	CanEdit                  bool                      `json:"can_edit"`
	ShowAdditionalIndication int64                     `json:"show_additional_indication"`
	Text                     string                    `json:"text"`
	TextLength               int64                     `json:"textLength"`
	Source                   string                    `json:"source"`
	Favorited                bool                      `json:"favorited"`
	PicIDS                   []interface{}             `json:"pic_ids"`
	PicTypes                 string                    `json:"pic_types"`
	IsPaid                   bool                      `json:"is_paid"`
	MblogVipType             int64                     `json:"mblog_vip_type"`
	User                     User                      `json:"user"`
	RepostsCount             int64                     `json:"reposts_count"`
	CommentsCount            int64                     `json:"comments_count"`
	AttitudesCount           int64                     `json:"attitudes_count"`
	PendingApprovalCount     int64                     `json:"pending_approval_count"`
	IsLongText               bool                      `json:"isLongText"`
	RewardExhibitionType     int64                     `json:"reward_exhibition_type"`
	HideFlag                 int64                     `json:"hide_flag"`
	Mlevel                   int64                     `json:"mlevel"`
	DarwinTags               []interface{}             `json:"darwin_tags"`
	Mblogtype                int64                     `json:"mblogtype"`
	Rid                      string                    `json:"rid"`
	MoreInfoType             int64                     `json:"more_info_type"`
	ContentAuth              int64                     `json:"content_auth"`
	PicNum                   int64                     `json:"pic_num"`
	EditConfig               RetweetedStatusEditConfig `json:"edit_config"`
	PageInfo                 PageInfo                  `json:"page_info"`
	Bid                      string                    `json:"bid"`
}

type RetweetedStatusEditConfig struct {
	Edited bool `json:"edited"`
}

type Visible struct {
	Type   int64 `json:"type"`
	ListID int64 `json:"list_id"`
}

type Title struct {
	Text      string `json:"text"`
	BaseColor int64  `json:"base_color"`
}

type Gender string

const (
	F Gender = "f"
	M Gender = "m"
)

type Type string

const (
	SearchTopic Type = "search_topic"
	Video       Type = "video"
)

type Source string

const (
	微博WeiboCOM Source = "微博 weibo.com"
	微博云剪       Source = "微博云剪"
)

type ActCode struct {
	Integer *int64
	String  *string
}

func (x *ActCode) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *ActCode) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}

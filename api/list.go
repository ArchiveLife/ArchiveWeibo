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
			"uid":         "uid",
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
	Ok   int64        `json:"ok"`
	Data ListPageData `json:"data"`
}

type ListPageData struct {
	Cards        []Card       `json:"cards"`
	CardlistInfo CardlistInfo `json:"cardlistInfo"`
	Scheme       string       `json:"scheme"`
	ShowAppTips  int64        `json:"showAppTips"`
}

type CardlistInfo struct {
	ShowStyle     int64          `json:"show_style"`
	CanShared     int64          `json:"can_shared"`
	CardlistMenus []CardlistMenu `json:"cardlist_menus"`
	CardlistTitle string         `json:"cardlist_title"`
	VP            string         `json:"v_p"`
	Desc          string         `json:"desc"`
	Containerid   string         `json:"containerid"`
	Page          interface{}    `json:"page"`
}

type CardlistMenu struct {
	Name   string              `json:"name"`
	Type   string              `json:"type"`
	Params *CardlistMenuParams `json:"params,omitempty"`
}

type CardlistMenuParams struct {
	Scheme string `json:"scheme"`
}

type Card struct {
	CardType     int64          `json:"card_type"`
	CardGroup    []CardGroup    `json:"card_group,omitempty"`
	Itemid       *string        `json:"itemid,omitempty"`
	IsAsyn       *int64         `json:"is_asyn,omitempty"`
	AsyncAPI     *string        `json:"async_api,omitempty"`
	Pic          *string        `json:"pic,omitempty"`
	Actionlog    *CardActionlog `json:"actionlog,omitempty"`
	Buttons      []Button       `json:"buttons,omitempty"`
	Desc2        *string        `json:"desc2,omitempty"`
	TitleSub     *string        `json:"title_sub,omitempty"`
	CardID       *string        `json:"card_id,omitempty"`
	Desc1        *string        `json:"desc1,omitempty"`
	ShowType     *int64         `json:"show_type,omitempty"`
	Title        *string        `json:"title,omitempty"`
	Scheme       *string        `json:"scheme,omitempty"`
	Desc         *string        `json:"desc,omitempty"`
	CardTypeName *string        `json:"card_type_name,omitempty"`
}

type CardActionlog struct {
	ActCode string  `json:"act_code"`
	EXT     string  `json:"ext"`
	Fid     string  `json:"fid"`
	Cardid  string  `json:"cardid"`
	Uicode  *string `json:"uicode,omitempty"`
	Unicode *string `json:"unicode,omitempty"`
}

type Button struct {
	Pic        string          `json:"pic"`
	SkipFormat int64           `json:"skip_format"`
	Params     ButtonParams    `json:"params"`
	Actionlog  ButtonActionlog `json:"actionlog"`
	Name       string          `json:"name"`
	Type       string          `json:"type"`
	SubType    int64           `json:"sub_type"`
	Scheme     string          `json:"scheme"`
}

type ButtonActionlog struct {
	ActCode string  `json:"act_code"`
	EXT     string  `json:"ext"`
	Fid     string  `json:"fid"`
	Cardid  string  `json:"cardid"`
	Uicode  *string `json:"uicode,omitempty"`
}

type ButtonParams struct {
	Action string `json:"action"`
}

type CardGroup struct {
	Title          *string       `json:"title,omitempty"`
	LeftTagImg     *string       `json:"left_tag_img,omitempty"`
	TitleIsBold    *int64        `json:"title_is_bold,omitempty"`
	CardType       int64         `json:"card_type"`
	Scheme         *string       `json:"scheme,omitempty"`
	Actionlog      CardActionlog `json:"actionlog"`
	SubTitle       *string       `json:"sub_title,omitempty"`
	Itemid         *string       `json:"itemid,omitempty"`
	Users          []User        `json:"users,omitempty"`
	ShowType       *int64        `json:"show_type,omitempty"`
	Desc           *string       `json:"desc,omitempty"`
	DisplayArrow   *int64        `json:"display_arrow,omitempty"`
	TitleExtraText *string       `json:"title_extra_text,omitempty"`
	Col            *int64        `json:"col,omitempty"`
	Mode           *string       `json:"mode,omitempty"`
	Group          []Group       `json:"group,omitempty"`
}

type Group struct {
	Actionlog ButtonActionlog `json:"actionlog"`
	ItemDesc  string          `json:"item_desc"`
	Scheme    string          `json:"scheme"`
	ItemTitle *ItemTitle      `json:"item_title"`
}

type User struct {
	ID              int64       `json:"id"`
	ScreenName      string      `json:"screen_name"`
	ProfileImageURL string      `json:"profile_image_url"`
	ProfileURL      string      `json:"profile_url"`
	StatusesCount   interface{} `json:"statuses_count"`
	Verified        bool        `json:"verified"`
	VerifiedType    int64       `json:"verified_type"`
	VerifiedTypeEXT int64       `json:"verified_type_ext"`
	VerifiedReason  string      `json:"verified_reason"`
	CloseBlueV      bool        `json:"close_blue_v"`
	Description     string      `json:"description"`
	Gender          string      `json:"gender"`
	Mbtype          interface{} `json:"mbtype"`
	Urank           interface{} `json:"urank"`
	Mbrank          interface{} `json:"mbrank"`
	FollowMe        bool        `json:"follow_me"`
	Following       bool        `json:"following"`
	FollowersCount  int64       `json:"followers_count"`
	FollowCount     int64       `json:"follow_count"`
	CoverImagePhone string      `json:"cover_image_phone"`
	AvatarHD        string      `json:"avatar_hd"`
}

type ItemTitle struct {
	Integer *int64
	String  *string
}

func (x *ItemTitle) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *ItemTitle) MarshalJSON() ([]byte, error) {
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

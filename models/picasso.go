package models

import "encoding/json"

type Message struct {
	Msg  string          `json:"msg"`
	Res  json.RawMessage `json:"res"`
	Code int             `json:"code"`
}

type Categories struct {
	Data []*Category `json:"category"`
}

type Category struct {
	Count        int      `json:"count"`
	Ename        string   `json:"ename"`
	Rname        string   `json:"rname"`
	CoverTemp    string   `json:"cover_temp"`
	Name         string   `json:"name"`
	Cover        string   `json:"cover"`
	Rank         int      `json:"rank"`
	Filter       []string `json:"filter"` // 猜测是 string
	SN           int      `json:"sn"`
	Icover       string   `json:"icover"`
	Atime        float64  `json:"atime"`
	Type         int      `json:"type"`
	ID           string   `json:"id"`
	PicassoCover string   `json:"picasso_cover"`
}

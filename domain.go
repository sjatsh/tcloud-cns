package cns

import (
	"net/url"
	"strconv"
)

//“域名”的数据结构
type Domain struct {
	Id               int    `json:"id"`
	Status           string `json:"status"`
	GroupId          string `json:"group_id"`
	SearchenginePush string `json:"searchengine_push"`
	IsMark           string `json:"is_mark"`
	Ttl              string `json:"ttl"`
	CnameSpeedup     string `json:"cname_speedup"`
	Remark           string `json:"remark"`
	CreatedOn        string `json:"created_on"`
	UpdatedOn        string `json:"updated_on"`
	QProjectId       int    `json:"q_project_id"`
	Punycode         string `json:"punycode"`
	ExtStatus        string `json:"ext_status"`
	SrcFlag          string `json:"src_flag"`
	Name             string `json:"name"`
	Grade            string `json:"DP_Free"`
	GradeTitle       string `json:"grade_title"`
	IsVip            string `json:"is_vip"`
	Owner            string `json:"owner"`
	Records          string `json:"records"`
	MinTtl           int    `json:"min_ttl"`
}

type DomainListResponse struct {
	BaseResponse
	Data struct {
		Info struct {
			DomainTotal int `json:"domain_total"`
		} `json:"info"`
		Domains []Domain `json:"domains"`
	} `json:"data"`
}

type DomainCreateResponse struct {
	BaseResponse
	Data struct {
		Domain struct {
			Id       int    `json:",string"`
			Punycode string `json:"punycode"`
			Domain   string `json:"domain"`
		} `json:"domain"`
	} `json:"data"`
}

//获取域名列表
func (cli *Client) DomainList(param url.Values) (*DomainListResponse, error) {
	respInfo := &DomainListResponse{}
	err := cli.requestGET("DomainList", param, respInfo)
	if err != nil {
		return nil, err
	}
	return respInfo, nil
}

//添加域名，如果成功，返回创建的域名ID
func (cli *Client) DomainCreate(domain string, projectId ...int) (*DomainCreateResponse, error) {
	param := url.Values{"domain": {domain}}
	if len(projectId) > 0 {
		param.Set("projectId", strconv.Itoa(projectId[0]))
	}
	respInfo := &DomainCreateResponse{}
	err := cli.requestGET("DomainCreate", param, respInfo)
	if err != nil {
		return nil, err
	}
	return respInfo, nil
}

//设置域名解析状态
func (cli *Client) SetDomainStatus(domain string, enable bool) error {
	param := url.Values{"domain": {domain}}

	if enable {
		param.Set("status", "enable")
	} else {
		param.Set("status", "disable")
	}

	var respInfo BaseResponse
	err := cli.requestGET("SetDomainStatus", param, &respInfo)
	if err != nil {
		return err
	}

	return nil
}

//删除域名
func (cli *Client) DomainDelete(domain string) error {
	param := url.Values{"domain": {domain}}

	var respInfo BaseResponse
	err := cli.requestGET("DomainDelete", param, &respInfo)
	if err != nil {
		return err
	}
	return nil
}

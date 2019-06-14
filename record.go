package tcloud_cns

import (
	"net/url"
	"strconv"
)

//“域名”“解析记录”的数据结构
type Record struct {
	Id         int
	Ttl        int
	Value      string
	Enabled    int
	Status     string
	UpdatedOn  string `json:"updated_on"`
	QProjectId int    `json:"q_project_id"`
	Name       string
	Line       string
	LineId     string `json:"line_id"`
	Type       string
	Remark     string
	Mx         int
	Hold       string
}

type RecordResponse struct {
	BaseResponse
	Data struct {
		Record struct {
			Id     int         `json:",string"`
			Name   string      `json:"name"`
			Status string      `json:"status"`
			Weight interface{} `json:"weight"`
		} `json:"record"`
	} `json:"data"`
}

//“域名”“解析记录”列表中的“域名”数据结构，和“域名列表”中的略有不同
type DomainInRecordList struct {
	Domain
	Ttl      int
	Id       int      `json:",string"`
	DnspodNs []string `json:"dnspod_ns"`
}

type RecordModifyResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

//获取指定“域名”的“解析记录”列表
func (cli *Client) RecordList(domain string) ([]Record, error) {
	var respInfo struct {
		BaseResponse
		Data struct {
			Domain  DomainInRecordList
			Records []Record
			Info    struct {
				SubDomains  int `json:"sub_domains,string"`
				RecordTotal int `json:"record_total,string"`
			}
		}
	}

	param := url.Values{
		"domain": {domain},
	}

	err := cli.requestGET("RecordList", param, &respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Data.Records, nil
}

//添加指定“域名”的“解析记录”
func (cli *Client) RecordCreate(domain string, record Record) (*RecordResponse, error) {
	if record.Line == "" {
		record.Line = "默认"
	}

	//必选参数
	param := url.Values{
		"domain":     {domain},
		"subDomain":  {record.Name},
		"recordType": {record.Type},
		"recordLine": {record.Line},
		"value":      {record.Value},
	}

	//可选TTL参数，缺省为600
	if record.Ttl > 0 {
		param.Set("ttl", strconv.Itoa(record.Ttl))
	}

	//MX记录必须的额外参数
	if record.Type == "MX" {
		param.Set("mx", strconv.Itoa(record.Mx))
	}

	respInfo := &RecordResponse{}
	err := cli.requestGET("RecordCreate", param, respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo, nil
}

//设置指定“域名”的“解析记录”状态
func (cli *Client) RecordStatus(domain string, recordId int, enable bool) error {
	param := url.Values{
		"domain":   {domain},
		"recordId": {strconv.Itoa(recordId)},
	}

	if enable {
		param.Set("status", "enable")
	} else {
		param.Set("status", "disable")
	}

	var respInfo BaseResponse
	err := cli.requestGET("RecordStatus", param, &respInfo)
	if err != nil {
		return err
	}

	return nil
}

//添加指定“域名”的“解析记录”
func (cli *Client) RecordModify(domain string, record Record) (*RecordModifyResponse, error) {
	if record.Line == "" {
		record.Line = "默认"
	}

	//必选参数
	param := url.Values{
		"domain":     {domain},
		"recordId":   {strconv.Itoa(record.Id)},
		"subDomain":  {record.Name},
		"recordType": {record.Type},
		"recordLine": {record.Line},
		"value":      {record.Value},
	}

	//可选TTL参数，缺省为600
	if record.Ttl > 0 {
		param.Set("ttl", strconv.Itoa(record.Ttl))
	}

	//MX记录必须的额外参数
	if record.Type == "MX" {
		param.Set("mx", strconv.Itoa(record.Mx))
	}

	respInfo := &RecordModifyResponse{}
	err := cli.requestGET("RecordModify", param, respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo, nil
}

//删除指定“域名”的“解析记录”
func (cli *Client) RecordDelete(domain string, recordId int) (*BaseResponse, error) {
	param := url.Values{
		"domain":   {domain},
		"recordId": {strconv.Itoa(recordId)},
	}

	respInfo := &BaseResponse{}
	err := cli.requestGET("RecordDelete", param, respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo, nil
}

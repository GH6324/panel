package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"github.com/spf13/cast"
)

type CertUpdate struct {
	ID        uint     `form:"id" json:"id"`
	Type      string   `form:"type" json:"type"`
	Domains   []string `form:"domains" json:"domains"`
	AutoRenew bool     `form:"auto_renew" json:"auto_renew"`
	UserID    uint     `form:"user_id" json:"user_id" filter:"uint"`
	DNSID     *uint    `form:"dns_id" json:"dns_id"`
	WebsiteID *uint    `form:"website_id" json:"website_id"`
}

func (r *CertUpdate) Authorize(ctx http.Context) error {
	return nil
}

func (r *CertUpdate) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"id":         "required|uint|min:1|exists:certs,id",
		"type":       "required|in:P256,P384,2048,4096",
		"domains":    "required|array",
		"auto_renew": "required|bool",
		"user_id":    "required|uint|exists:cert_users,id",
		"dns_id":     "uint",
		"website_id": "uint",
	}
}

func (r *CertUpdate) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CertUpdate) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CertUpdate) PrepareForValidation(ctx http.Context, data validation.Data) error {
	// TODO 由于验证器 filter 标签的问题，暂时这里这样处理
	dnsID, exist := data.Get("dns_id")
	if exist {
		err := data.Set("dns_id", cast.ToUint(dnsID))
		if err != nil {
			return err
		}

	}
	websiteID, exist := data.Get("website_id")
	if exist {
		err := data.Set("website_id", cast.ToUint(websiteID))
		if err != nil {
			return err
		}
	}

	return nil
}

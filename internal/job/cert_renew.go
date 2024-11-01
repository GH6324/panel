package job

import (
	"log/slog"
	"time"

	"github.com/TheTNB/panel/internal/app"
	"github.com/TheTNB/panel/internal/biz"
	"github.com/TheTNB/panel/internal/data"
	pkgcert "github.com/TheTNB/panel/pkg/cert"
)

// CertRenew 证书续签
type CertRenew struct {
	certRepo biz.CertRepo
}

func NewCertRenew() *CertRenew {
	return &CertRenew{
		certRepo: data.NewCertRepo(),
	}
}

func (r *CertRenew) Run() {
	if app.Status != app.StatusNormal {
		return
	}

	var certs []biz.Cert
	if err := app.Orm.Preload("Website").Preload("Account").Preload("DNS").Find(&certs).Error; err != nil {
		app.Logger.Warn("获取证书失败", slog.Any("err", err))
		return
	}

	for _, cert := range certs {
		if cert.Type == "upload" || !cert.AutoRenew {
			continue
		}

		decode, err := pkgcert.ParseCert(cert.Cert)
		if err != nil {
			continue
		}

		// 结束时间大于 7 天的证书不续签
		now := time.Now()
		if decode.NotAfter.Sub(now).Hours() > 24*7 {
			continue
		}

		_, err = r.certRepo.Renew(cert.ID)
		if err != nil {
			app.Logger.Warn("续签证书失败", slog.Any("err", err))
		}
	}
}

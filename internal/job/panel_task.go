package job

import (
	"log/slog"
	"runtime"
	"runtime/debug"

	"github.com/TheTNB/panel/internal/app"
	"github.com/TheTNB/panel/internal/biz"
	"github.com/TheTNB/panel/internal/data"
)

// PanelTask 面板每日任务
type PanelTask struct {
	appRepo     biz.AppRepo
	backupRepo  biz.BackupRepo
	settingRepo biz.SettingRepo
}

func NewPanelTask() *PanelTask {
	return &PanelTask{
		appRepo:     data.NewAppRepo(),
		backupRepo:  data.NewBackupRepo(),
		settingRepo: data.NewSettingRepo(),
	}
}

func (r *PanelTask) Run() {
	app.Status = app.StatusMaintain

	// 优化数据库
	if err := app.Orm.Exec("VACUUM").Error; err != nil {
		app.Status = app.StatusFailed
		app.Logger.Warn("优化面板数据库失败", slog.Any("err", err))
	}
	if err := app.Orm.Exec("PRAGMA wal_checkpoint(TRUNCATE);").Error; err != nil {
		app.Status = app.StatusFailed
		app.Logger.Warn("优化面板数据库失败", slog.Any("err", err))
	}

	// 备份面板
	if err := r.backupRepo.Create(biz.BackupTypePanel, ""); err != nil {
		app.Logger.Warn("备份面板失败", slog.Any("err", err))
	}

	// 清理备份
	path, err := r.backupRepo.GetPath("panel")
	if err == nil {
		if err = r.backupRepo.ClearExpired(path, "panel_", 10); err != nil {
			app.Logger.Warn("清理面板备份失败", slog.Any("err", err))
		}
	}

	// 更新商店缓存
	if offline, err := r.settingRepo.GetBool(biz.SettingKeyOfflineMode); err == nil && !offline {
		if err = r.appRepo.UpdateCache(); err != nil {
			app.Logger.Warn("更新商店缓存失败", slog.Any("err", err))
		}
	}

	// 回收内存
	runtime.GC()
	debug.FreeOSMemory()

	app.Status = app.StatusNormal
}

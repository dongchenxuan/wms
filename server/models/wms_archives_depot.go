package models

import (
	"sync"
	"time"
)

var (
	wmsArchivesDepotDao     *WmsArchivesDepotDao
	wmsArchivesDepotDaoOnce sync.Once
)

// WmsArchivesDepotInstance todo: Instance
func WmsArchivesDepotInstance() *WmsArchivesDepotDao {
	wmsArchivesDepotDaoOnce.Do(func() {
		wmsArchivesDepotDao = &WmsArchivesDepotDao{}
	})
	return wmsArchivesDepotDao
}

type WmsArchivesDepotDao struct {
	Id           string    `xorm:"pk not null comment('仓库ID') varchar(64) 'id'" json:"id"`
	OrgId        int64     `xorm:"not null comment('单位ID') bigint(20) 'orgId'" json:"orgId"`
	IsDel        int8      `xorm:"not null default 2 comment('是否删除：1-是 2-否') tinyint(1) 'isDel'" json:"isDel"`
	IsLock       int8      `xorm:"not null default 2 comment('是否锁定：1-是 2-否') tinyint(1) 'isLock'" json:"isLock"`
	IsVerify     int8      `xorm:"not null default 4 comment('审核状态：1-通过 2-拒绝 3-反审 4-未审') tinyint(1) 'isVerify'" json:"isVerify"`
	Name         string    `xorm:"not null comment('仓库名称') varchar(50) 'name'" json:"name"`
	Location     string    `xorm:"not null comment('仓库地址') varchar(100) 'location'" json:"location"`
	Area         int64     `xorm:"not null comment('面积（x100）') bigint(20) 'area'" json:"area"`
	LeaderId     string    `xorm:"default '0' comment('负责人ID') varchar(64) 'leaderId'" json:"leaderId"`
	LeaderName   string    `xorm:"not null comment('负责人姓名') varchar(50) 'leaderName'" json:"leaderName"`
	LeaderMobile string    `xorm:"not null comment('负责人联系方式') varchar(100) 'leaderMobile'" json:"leaderMobile"`
	Capacity     int64     `xorm:"default 0 comment('仓库容量') bigint(20) 'capacity'" json:"capacity"`
	UsedCapacity int64     `xorm:"default 0 comment('已用容量') bigint(20) 'usedCapacity'" json:"usedCapacity"`
	Remark       string    `xorm:"default 'NULL' comment('备注') varchar(100) 'remark'" json:"remark"`
	ReviewerId   string    `xorm:"default 'NULL' comment('审核人ID') varchar(64) 'reviewerId'" json:"reviewerId"`
	ReviewerName string    `xorm:"default 'NULL' comment('审核人姓名') varchar(50) 'reviewerName'" json:"reviewerName"`
	Comment      string    `xorm:"default 'NULL' comment('审核批注') varchar(100) 'comment'" json:"comment"`
	CreatorId    string    `xorm:"not null comment('创建人ID') varchar(64) 'creatorId'" json:"creatorId"`
	CreatorName  string    `xorm:"not null comment('创建人姓名') varchar(50) 'creatorName'" json:"creatorName"`
	UpdaterId    string    `xorm:"not null comment('更新人ID') varchar(64) 'updaterId'" json:"updaterId"`
	UpdaterName  string    `xorm:"not null comment('更新人姓名') varchar(50) 'updaterName'" json:"updaterName"`
	Ctime        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') datetime 'ctime'" json:"ctime"`
	Utime        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') datetime 'utime'" json:"utime"`
}

func (dao *WmsArchivesDepotDao) TableName() string {
	return "wms_archives_depot"
}

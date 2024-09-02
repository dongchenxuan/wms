package models

import (
	"sync"
	"time"
)

var (
	wmsArchivesPartDao     *WmsArchivesPartDao
	wmsArchivesPartDaoOnce sync.Once
)

// WmsArchivesPartInstance todo: Instance
func WmsArchivesPartInstance() *WmsArchivesPartDao {
	wmsArchivesPartDaoOnce.Do(func() {
		wmsArchivesPartDao = &WmsArchivesPartDao{}
	})

	return wmsArchivesPartDao
}

type WmsArchivesPartDao struct {
	Id                string    `xorm:"pk not null varchar(64) 'id'" json:"id"`
	OrgId             int64     `xorm:"not null bigint(20) 'orgId'" json:"orgId"`
	IsDel             int8      `xorm:"not null default 2 comment('是否删除：1-是 2-否') tinyint(1) 'isDel'" json:"isDel"`
	IsLock            int8      `xorm:"not null default 2 comment('是否锁定：1-是 2-否') tinyint(1) 'isLock'" json:"isLock"`
	IsVerify          int8      `xorm:"not null default 4 comment('审核状态：1-通过 2-拒绝 3-反审 4-未审') tinyint(1) 'isVerify'" json:"isVerify"`
	DepotId           string    `xorm:"not null comment('仓库ID') varchar(64) 'depotId'" json:"depotId"`
	DepotName         string    `xorm:"not null comment('仓库名称') varchar(50) 'depotName'" json:"depotName"`
	Name              string    `xorm:"not null comment('配件名称') varchar(50) 'name'" json:"name"`
	FullName          string    `xorm:"default 'NULL' comment('全称') varchar(100) 'fullName'" json:"fullName"`
	Model             string    `xorm:"not null comment('配件编号') varchar(100) 'model'" json:"model"`
	SiteType          string    `xorm:"not null comment('维修部件类型') varchar(20) 'siteType'" json:"siteType"`
	ServiceLife       int32     `xorm:"not null default 0 comment('使用年限') int(11) 'serviceLife'" json:"serviceLife"`
	Unit              string    `xorm:"not null comment('配件单位') varchar(5) 'unit'" json:"unit"`
	UnitPrice         int64     `xorm:"not null comment('单价（x100）') bigint(20) 'unitPrice'" json:"unitPrice"`
	StockQuantity     int64     `xorm:"default NULL comment('当前库存数量') bigint(20) 'stockQuantity'" json:"stockQuantity"`
	AvailableQuantity int64     `xorm:"default NULL comment('可用库存数量') bigint(20) 'availableQuantity'" json:"availableQuantity"`
	Brand             string    `xorm:"not null comment('品牌') varchar(100) 'brand'" json:"brand"`
	Contact           string    `xorm:"not null comment('联系人') varchar(50) 'contact'" json:"contact"`
	Mobile            string    `xorm:"not null comment('联系电话') varchar(100) 'mobile'" json:"mobile"`
	SupplierName      string    `xorm:"default 'NULL' comment('供应商名称') varchar(100) 'supplierName'" json:"supplierName"`
	Address           string    `xorm:"not null comment('地址') varchar(100) 'address'" json:"address"`
	Remark            string    `xorm:"default 'NULL' comment('备注') varchar(100) 'remark'" json:"remark"`
	CreatorId         string    `xorm:"not null comment('创建人ID') varchar(64) 'creatorId'" json:"creatorId"`
	CreatorName       string    `xorm:"default 'NULL' comment('创建人姓名') varchar(50) 'creatorName'" json:"creatorName"`
	UpdaterId         string    `xorm:"not null comment('更新人ID') varchar(64) 'updaterId'" json:"updaterId"`
	UpdaterName       string    `xorm:"default 'NULL' comment('更新人姓名') varchar(50) 'updaterName'" json:"updaterName"`
	ReviewerId        string    `xorm:"default 'NULL' comment('审核人ID') varchar(64) 'reviewerId'" json:"reviewerId"`
	ReviewerName      string    `xorm:"default 'NULL' comment('审核人姓名') varchar(50) 'reviewerName'" json:"reviewerName"`
	Comment           string    `xorm:"default 'NULL' comment('审核皮质') varchar(100) 'comment'" json:"comment"`
	Ctime             time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' datetime 'ctime'" json:"ctime"`
	Utime             time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' datetime 'utime'" json:"utime"`
}

func (dao *WmsArchivesPartDao) TableName() string {
	return "wms_archives_part"
}

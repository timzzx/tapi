// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePermissionResource = "permission_resource"

// PermissionResource mapped from table <permission_resource>
type PermissionResource struct {
	ID     int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 权限资源表主键
	Name   string `gorm:"column:name;not null" json:"name"`                  // 资源名称
	URL    string `gorm:"column:url;not null" json:"url"`                    // 资源url
	Status int32  `gorm:"column:status;not null;default:1" json:"status"`    // 1.有效，2.无效
	Ctime  int32  `gorm:"column:ctime;not null" json:"ctime"`                // 创建时间
	Utime  int32  `gorm:"column:utime;not null" json:"utime"`                // 更新时间
}

// TableName PermissionResource's table name
func (*PermissionResource) TableName() string {
	return TableNamePermissionResource
}

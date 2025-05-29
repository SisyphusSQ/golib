package ob_dto

type Tag map[string]string

type BaseReq struct {
	ClientToken string `json:"clientToken"`

	// cluster
	InstanceId   string `json:"instanceId,omitempty"`
	InstanceName string `json:"instanceName,omitempty"`

	// tenant
	TenantId   string `json:"tenantId,omitempty"`
	TenantName string `json:"tenantName,omitempty"`

	// page
	PageNo   int `json:"pageNo,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

// CreateClusterReq https://cloud.baidu.com/doc/DDC/s/Umaophg3b
type CreateClusterReq struct {
	BaseReq
	Number       int      `json:"number"`
	Tags         []Tag    `json:"tags"`
	Version      string   `json:"version"`
	Cpu          int      `json:"cpu"`
	Memory       int      `json:"memory"`
	InstanceName string   `json:"instanceName"`
	Zones        []string `json:"zones"`
	DeployType   string   `json:"deployType"`
	ReplicaMode  string   `json:"replicaMode"`
	InstanceType string   `json:"instanceType"`
}

// CstListReq https://cloud.baidu.com/doc/DDC/s/1map0t682
type CstListReq struct {
	BaseReq
}

// CstDetailReq https://cloud.baidu.com/doc/DDC/s/5map12blc
type CstDetailReq struct {
	BaseReq
}

// ModCstProxyReq https://cloud.baidu.com/doc/DDC/s/bmap2ewqd
type ModCstProxyReq struct {
	BaseReq
	ProxyClusterId string      `json:"proxyClusterId"`
	UnitNum        int         `json:"unitNum"`
	Mode           ObProxyMode `json:"mode"`
}

// ReleaseCstReq https://cloud.baidu.com/doc/DDC/s/Umap1kapf
type ReleaseCstReq struct {
	BaseReq
}

// DelCstReq https://cloud.baidu.com/doc/DDC/s/umap1qfqp
type DelCstReq struct {
	BaseReq
}

// ChangeCstNameReq https://cloud.baidu.com/doc/DDC/s/Smap1bzm5
type ChangeCstNameReq struct {
	BaseReq
}

// WhiteGrpReq
// 设置OB集群IP白名单: https://cloud.baidu.com/doc/DDC/s/Fmap1ztet
// 创建租户安全白名单: https://cloud.baidu.com/doc/DDC/s/kmapa3xay
// 修改租户安全白名单: https://cloud.baidu.com/doc/DDC/s/Cmapaeehr
// 删除租户安全白名单: https://cloud.baidu.com/doc/DDC/s/gmapapgdc
type WhiteGrpReq struct {
	BaseReq
	SecurityIPs     string `json:"securityIps,omitempty"`
	SecurityIPGroup string `json:"securityIpGroupName"`
}

// GetWhiteGrpReq
// 查询集群IP白名单: https://cloud.baidu.com/doc/DDC/s/Xmap26h2u
// 查询租户安全白名单列表: https://cloud.baidu.com/doc/DDC/s/cmapaiwqd
type GetWhiteGrpReq struct {
	BaseReq
}

// GetPriceReq https://cloud.baidu.com/doc/DDC/s/1maqdrapl
type GetPriceReq struct {
	Number      int    `json:"number"`
	Duration    int    `json:"duration"`
	ProductType string `json:"productType"`
	Cpu         int    `json:"cpu"`
	Memory      int    `json:"memory"`
}

// CreateTenantReq https://cloud.baidu.com/doc/DDC/s/gmap2n6oi
type CreateTenantReq struct {
	InstanceId             string          `json:"instanceId"`
	TenantName             string          `json:"tenantName"`
	Description            string          `json:"description"`
	Cpu                    int             `json:"cpu"`
	Memory                 int             `json:"memory"`
	UnitNum                int             `json:"unitNum"`
	PrimaryZone            string          `json:"primaryZone"`
	TenantMode             string          `json:"tenantMode"`
	TimeZone               string          `json:"timeZone"`
	Charset                ObEncodingType  `json:"charset"`
	Collation              ObCollationType `json:"collation"`
	ProxyAddressMasterZone string          `json:"proxyAddressMasterZone"`
	CreateParams           CreateParams    `json:"createParams"`
}

// TenListReq https://cloud.baidu.com/doc/DDC/s/Pmap3er3z
type TenListReq struct {
	BaseReq
}

// TenDetailReq https://cloud.baidu.com/doc/DDC/s/1map3n7e2
type TenDetailReq struct {
	BaseReq
}

// DelTenReq https://cloud.baidu.com/doc/DDC/s/smapa0fvy
type DelTenReq struct {
	BaseReq
}

// ChangeTenNameReq https://cloud.baidu.com/doc/DDC/s/bmap9wgb7
type ChangeTenNameReq struct {
	BaseReq
}

// ChangePrimaryZoneReq https://cloud.baidu.com/doc/DDC/s/mmapbr440
type ChangePrimaryZoneReq struct {
	BaseReq
	PrimaryZone         string `json:"primaryZone"`
	ProxyAddrMasterZone string `json:"proxyAddressMasterZone"`
}

// CreateTenLinkReq https://cloud.baidu.com/doc/DDC/s/Smapat72a
type CreateTenLinkReq struct {
	BaseReq
	VpcId            string           `json:"vpcId"`
	SubnetId         string           `json:"subnetId"`
	VipServiceType   ObServiceType    `json:"vipServiceType"`
	ZoneIds          []string         `json:"zoneIdList,omitempty"`
	TrafficStrategy  TrafficStrategy  `json:"trafficStrategy,omitempty"`
	ReplParams       []ReplParam      `json:"replParams,omitempty"` // use type ReplParam
	BalStrategy      string           `json:"balStrategy,omitempty"`
	RecoveryStrategy RecoveryStrategy `json:"disasterRecoveryStrategy,omitempty"`
}

// TenLinkReq
// 查询：https://cloud.baidu.com/doc/DDC/s/Mmapbfcvf
// 删除：https://cloud.baidu.com/doc/DDC/s/Ymapbnrbu
type TenLinkReq struct {
	BaseReq
	AddrId string `json:"addressId"`
}

// EditTenUserReq
// 创建租户用户: https://cloud.baidu.com/doc/DDC/s/Lmapbwm5e
// 修改用户密码: https://cloud.baidu.com/doc/DDC/s/Gmaq670n9
type EditTenUserReq struct {
	BaseReq
	UserType ObUserType `json:"userType,omitempty"`
	UserName string     `json:"userName"`
	Pwd      string     `json:"password"`
	Desc     string     `json:"description,omitempty"`
	Role     []ObDBRole `json:"roles,omitempty"`
}

// GetTenUserReq https://cloud.baidu.com/doc/DDC/s/Wmaq56dc4
type GetTenUserReq struct {
	Database string `json:"databaseName,omitempty"`
}

// DelTenUserReq https://cloud.baidu.com/doc/DDC/s/amaq7rl31
type DelTenUserReq struct {
	BaseReq
	Users []string `json:"users"`
}

// TenDBReq
// 创建租户数据库: https://cloud.baidu.com/doc/DDC/s/hmaq80fi8
// 查询租户数据库列表: https://cloud.baidu.com/doc/DDC/s/Wmaq8ggnx
type TenDBReq struct {
	BaseReq
	DB        string          `json:"databaseName,omitempty"`
	Encoding  ObEncodingType  `json:"encoding,omitempty"`
	Collation ObCollationType `json:"collation,omitempty"`
	Desc      string          `json:"description,omitempty"`
}

// DelTenDBReq https://cloud.baidu.com/doc/DDC/s/0maq8ypc9
type DelTenDBReq struct {
	BaseReq
	DBs []string `json:"databaseNames"`
}

// GetTopSQLReq https://cloud.baidu.com/doc/DDC/s/bmaqcrtt6
type GetTopSQLReq struct {
	BaseReq
	StartTime  string `json:"startTime"` // 2023-04-12T04:38:38Z
	EndTime    string `json:"endTime"`   // 2023-04-12T04:38:38Z 开始时间与结束时间间隔不能大于1天
	DBName     string `json:"dbName,omitempty"`
	SearchKey  string `json:"searchKeyWord,omitempty"`
	NodeIP     string `json:"nodeIp,omitempty"`
	SQLTextLen int64  `json:"sqlTextLength,omitempty"`
	FilCond    string `json:"filterCondition,omitempty"`
}

// GetSlowSQLReq https://cloud.baidu.com/doc/DDC/s/bmaqcrtt6
type GetSlowSQLReq struct {
	BaseReq
	StartTime  string `json:"startTime"` // 2023-04-12T04:38:38Z
	EndTime    string `json:"endTime"`   // 2023-04-12T04:38:38Z 开始时间与结束时间间隔不能大于1天
	DBName     string `json:"dbName,omitempty"`
	SearchKey  string `json:"searchKeyWord,omitempty"`
	NodeIP     string `json:"nodeIp,omitempty"`
	SQLTextLen int64  `json:"sqlTextLength,omitempty"`
	FilCond    string `json:"filterCondition,omitempty"`
}

package ob_dto

type ObVersion string

const (
	V4_2_1 ObVersion = "4.2.1"
	V4_2_5 ObVersion = "4.2.5"
	v4_3_3 ObVersion = "v4.3.3"
	v4_3_5 ObVersion = "v4.3.5"
)

type ObTenantModeType string

const (
	TenantModeMySQL ObTenantModeType = "MySQL"
)

type ObAddressType string

const (
	AddressMaster    ObAddressType = "Master"
	AddressReadOnly  ObAddressType = "READONLY"
	AddressReadWrite ObAddressType = "READWRITE"
)

type ObProxyMode string

const (
	ProxyExclusive ObProxyMode = "k8s_exclusive"
	ProxyShared    ObProxyMode = "k8s_shared"
)

type ObServiceType string

const (
	ObProxyIntranet  ObServiceType = "OBPROXY_INTRANET"
	ObProxyReadOnly  ObServiceType = "OBPROXY_READONLY"
	ObProxyReadWrite ObServiceType = "OBPROXY_READWRITE"
)

type ObDeployType string

const (
	DeployMulti  ObDeployType = "multiple"
	DeployDual   ObDeployType = "dual"
	DeploySingle ObDeployType = "single"
)

type ObDataBaseRoleType string

const (
	RoleRW    ObDataBaseRoleType = "readwrite"
	RoleRO    ObDataBaseRoleType = "readonly"
	RoleDDL   ObDataBaseRoleType = "DDL"
	RoleDML   ObDataBaseRoleType = "DML"
	RoleOther ObDataBaseRoleType = "other"
)

type ObDataBasePrivilegeType string

const (
	PrivAlter      ObDataBasePrivilegeType = "ALTER"
	PrivCreate     ObDataBasePrivilegeType = "CREATE"
	PrivDelete     ObDataBasePrivilegeType = "DELETE"
	PrivDrop       ObDataBasePrivilegeType = "DROP"
	PrivInsert     ObDataBasePrivilegeType = "INSERT"
	PrivSelect     ObDataBasePrivilegeType = "SELECT"
	PrivUpdate     ObDataBasePrivilegeType = "UPDATE"
	PrivIndex      ObDataBasePrivilegeType = "INDEX"
	PrivCreateView ObDataBasePrivilegeType = "CREATE VIEW"
	PrivDeleteView ObDataBasePrivilegeType = "SHOW VIEW"
)

type ObUserType string

const (
	UserNormal ObUserType = "normal"
	UserAdmin  ObUserType = "admin"
	UserRo     ObUserType = "ReadonlyAccount"
)

type ObEncodingType string

const (
	UTF8MB4     ObEncodingType = "utf8mb4"
	UTF16       ObEncodingType = "utf16"
	GBK         ObEncodingType = "gbk"
	GB18030     ObEncodingType = "gb18030"
	EncodBinary ObEncodingType = "binary"
)

type ObCollationType string

const (
	UTF8MB4GeneralCi ObCollationType = "utf8mb4_general_ci"
	UTF8MB4Bin       ObCollationType = "utf8mb4_bin"
	UTF8MB4UnicodeCi ObCollationType = "utf8mb4_unicode_ci"
	UTF16GeneralCi   ObCollationType = "utf16_general_ci"
	UTF16Bin         ObCollationType = "utf16_bin"
	UTF16UnicodeCi   ObCollationType = "utf16_unicode_ci"
	GBKChineseCi     ObCollationType = "gbk_chinese_ci"
	GBKBin           ObCollationType = "gbk_bin"
	GB18030ChineseCi ObCollationType = "gb18030_chinese_ci"
	GB18030Bin       ObCollationType = "gb18030_bin"
	CollationBinary  ObCollationType = "binary"
)

type ObInstanceType string

const (
	InstCluster ObInstanceType = "CLUSTER"
)

type ObInstanceStatus string

const (
	InstPendCreate      ObInstanceStatus = "PENDING_CREATE"
	InstOnline          ObInstanceStatus = "ONLINE"
	InstPendDelete      ObInstanceStatus = "PENDING_DELETE"
	InstDeleted         ObInstanceStatus = "DELETED"
	InstClosing         ObInstanceStatus = "CLOSING"
	InstClosed          ObInstanceStatus = "CLOSED"
	InstRestoring       ObInstanceStatus = "RESTORING"
	InstExpanding       ObInstanceStatus = "EXPANDING"
	InstReducing        ObInstanceStatus = "REDUCING"
	InstSpecUp          ObInstanceStatus = "SPEC_UPGRADING"
	InstSpecDown        ObInstanceStatus = "SPEC_DOWNGRADING"
	InstDiskUp          ObInstanceStatus = "DISK_UPGRADING"
	InstDiskDown        ObInstanceStatus = "DISK_DOWNGRADING"
	InstUpgrade         ObInstanceStatus = "UPGRADING"
	InstMigrate         ObInstanceStatus = "MIGRATING"
	InstSplit           ObInstanceStatus = "SPLITING"
	InstNetInit         ObInstanceStatus = "NETWORK_INITIALIZING"
	InstWhiteModify     ObInstanceStatus = "WHITE_LIST_MODIFYING"
	InstParamModify     ObInstanceStatus = "PARAMETER_MODIFYING"
	InstTenCreate       ObInstanceStatus = "TENANT_CREATING"
	InstTenSpecModify   ObInstanceStatus = "TENANT_SPEC_MODIFYING"
	InstTenExpand       ObInstanceStatus = "TENANT_EXPANDING"
	InstPreExpClosed    ObInstanceStatus = "PREPAID_EXPIRE_CLOSED"
	InstArrClosed       ObInstanceStatus = "ARREARS_CLOSED"
	InstStandbyCreate   ObInstanceStatus = "STANDBY_CREATING"
	InstSSLModify       ObInstanceStatus = "SSL_MODIFYING"
	InstSpecModify      ObInstanceStatus = "SPEC_MODIFYING"
	InstRestarting      ObInstanceStatus = "RESTARTING"
	InstAbnormal        ObInstanceStatus = "ABNORMAL"
	InstStandbyDelete   ObInstanceStatus = "STANDBY_DELETING"
	InstSwitch          ObInstanceStatus = "SWITCHOVER_SWITCHING"
	InstStandbyDis      ObInstanceStatus = "STANDBY_DISCONNECTING"
	InstLogDiskUp       ObInstanceStatus = "LOG_DISK_UPGRADING"
	InstDiskTypeModify  ObInstanceStatus = "DISKTYPE_MODIFYING"
	InstProxyCreate     ObInstanceStatus = "PROXY_SERVICE_CREATING"
	InstProxyDelete     ObInstanceStatus = "PROXY_SERVICE_DELETING"
	InstProxyModify     ObInstanceStatus = "PROXY_SERVICE_SPEC_MODIFYING"
	InstModifyPriZone   ObInstanceStatus = "MODIFYING_PRIMARY_ZONE"
	InstDeployModify    ObInstanceStatus = "DEPLOY_MODE_MODIFYING"
	InstCreateTenRORepl ObInstanceStatus = "CREATING_TENANT_READONLY_REPLICA"
	InstDeleteTenRORepl ObInstanceStatus = "DELETING_TENANT_READONLY_REPLICA"
)

var InstanceMp = map[ObInstanceStatus]string{
	InstPendCreate:      "创建中",
	InstOnline:          "运行中",
	InstPendDelete:      "删除中",
	InstDeleted:         "已删除",
	InstClosing:         "关闭中",
	InstClosed:          "已关闭",
	InstRestoring:       "备份恢复中",
	InstExpanding:       "租户规格修改中",
	InstReducing:        "节点扩容中",
	InstSpecUp:          "套餐规格扩容中",
	InstSpecDown:        "变配中",
	InstDiskUp:          "日志盘扩容中",
	InstDiskDown:        "变配中",
	InstUpgrade:         "版本升级中",
	InstMigrate:         "迁移中",
	InstSplit:           "拆分中",
	InstNetInit:         "网络初始化中",
	InstWhiteModify:     "白名单修改中",
	InstParamModify:     "参数修改中",
	InstTenCreate:       "租户创建中",
	InstTenSpecModify:   "租户规格修改中",
	InstTenExpand:       "租户扩容中",
	InstPreExpClosed:    "已隔离",
	InstArrClosed:       "已隔离",
	InstStandbyCreate:   "备集群创建中",
	InstSSLModify:       "ssl配置变更中",
	InstSpecModify:      "套餐规格修改中",
	InstRestarting:      "集群滚动重启",
	InstAbnormal:        "异常",
	InstStandbyDelete:   "备集群删除中",
	InstSwitch:          "主备切换中",
	InstStandbyDis:      "备库解耦中",
	InstLogDiskUp:       "日志盘扩容中",
	InstDiskTypeModify:  "磁盘类型修改中",
	InstProxyCreate:     "代理服务开通中",
	InstProxyDelete:     "代理服务关闭中",
	InstProxyModify:     "代理服务规格修改中",
	InstModifyPriZone:   "修改主可用区",
	InstDeployModify:    "部署模式切换中",
	InstCreateTenRORepl: "租户只读副本创建中",
	InstDeleteTenRORepl: "租户只读副本删除中",
}

func (o ObInstanceStatus) String() string {
	if s, ok := InstanceMp[o]; ok {
		return s
	}
	return string(o)
}

type ObTenantStatus string

const (
	TenPendCreate       ObTenantStatus = "PENDING_CREATE"
	TenRestore          ObTenantStatus = "RESTORE"
	TenOnline           ObTenantStatus = "ONLINE"
	TenSpecModify       ObTenantStatus = "SPEC_MODIFYING"
	TenAllocInterAddr   ObTenantStatus = "ALLOCATING_INTERNET_ADDRESS"
	TenPendOffInterAddr ObTenantStatus = "PENDING_OFFLINE_INTERNET_ADDRESS"
	TenPriZoneModify    ObTenantStatus = "PRIMARY_ZONE_MODIFYING"
	TenParamModify      ObTenantStatus = "PARAMETER_MODIFYING"
	TenWhiteListModify  ObTenantStatus = "WHITE_LIST_MODIFYING"
)

var TenantMp = map[ObTenantStatus]string{
	TenPendCreate:       "创建中",
	TenRestore:          "恢复中",
	TenOnline:           "运行中",
	TenSpecModify:       "规格修改中",
	TenAllocInterAddr:   "公网地址分配中",
	TenPendOffInterAddr: "公网地址关闭中",
	TenPriZoneModify:    "切主可用区中",
	TenParamModify:      "参数修改中",
	TenWhiteListModify:  "白名单修改中",
}

func (t ObTenantStatus) String() string {
	if s, ok := TenantMp[t]; ok {
		return s
	}
	return string(t)
}

type ObAddressStatus string

const (
	AddPendCreate ObAddressStatus = "PENDING_CREATE"
	AddOnline     ObAddressStatus = "ONLINE"
	AddPendDelete ObAddressStatus = "PENDING_DELETE"
	AddDeleted    ObAddressStatus = "DELETED"
)

var AddressMp = map[ObAddressStatus]string{
	AddPendCreate: "创建中",
	AddOnline:     "运行中",
	AddPendDelete: "删除中",
	AddDeleted:    "运行中",
}

func (o ObAddressStatus) String() string {
	if s, ok := AddressMp[o]; ok {
		return s
	}
	return string(o)
}

type LowCase string

const (
	Off LowCase = "0"
	On  LowCase = "1"
)

type TrafficStrategy string

const (
	TraRepl TrafficStrategy = "REPLICA_TYPE"
)

type ReplParam struct {
	Value string `json:"value"`
}

var (
	ReplFull       = []ReplParam{{Value: "FUll"}}
	ReplMaster     = []ReplParam{{Value: "MASTER"}}
	ReplFullMaster = []ReplParam{ReplFull[0], ReplMaster[0]}
)

type BalanceStrategy string

const (
	BalAuto BalanceStrategy = "AUTO_BALANCE"
)

type RecoveryStrategy string

const (
	Yes RecoveryStrategy = "YES"
	No  RecoveryStrategy = "NO"
)

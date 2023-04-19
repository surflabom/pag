package grafana

import "time"

const (
	DS_GRAPHITE       = "graphite"
	DS_INFLUXDB       = "influxdb"
	DS_INFLUXDB_08    = "influxdb_08"
	DS_ES             = "elasticsearch"
	DS_PROMETHEUS     = "prometheus"
	DS_ALERTMANAGER   = "alertmanager"
	DS_JAEGER         = "jaeger"
	DS_LOKI           = "loki"
	DS_OPENTSDB       = "opentsdb"
	DS_TEMPO          = "tempo"
	DS_ZIPKIN         = "zipkin"
	DS_MYSQL          = "mysql"
	DS_POSTGRES       = "postgres"
	DS_MSSQL          = "mssql"
	DS_ACCESS_DIRECT  = "direct"
	DS_ACCESS_PROXY   = "proxy"
	DS_ES_OPEN_DISTRO = "grafana-es-open-distro-datasource"
	DS_ES_OPENSEARCH  = "grafana-opensearch-datasource"
	DS_AZURE_MONITOR  = "grafana-azure-monitor-datasource"
)

type Config struct {
	BaseURL string //grafana 服务地址 例如: "http://127.0.0.1:9090"
	ApiKey  string //身份验证 API Key
}

type Json struct {
	Data interface{}
}

type DashboardFullWithMeta struct {
	Meta      DashboardMeta          `json:"meta"`
	Dashboard map[string]interface{} `json:"dashboard" binding:"Required"`
}

type DashboardMeta struct {
	IsStarred                  bool                  `json:"isStarred,omitempty"`
	IsSnapshot                 bool                  `json:"isSnapshot,omitempty"`
	Type                       string                `json:"type,omitempty"`
	CanSave                    bool                  `json:"canSave"`
	CanEdit                    bool                  `json:"canEdit"`
	CanAdmin                   bool                  `json:"canAdmin"`
	CanStar                    bool                  `json:"canStar"`
	CanDelete                  bool                  `json:"canDelete"`
	Slug                       string                `json:"slug"`
	Url                        string                `json:"url"`
	Expires                    time.Time             `json:"expires"`
	Created                    time.Time             `json:"created"`
	Updated                    time.Time             `json:"updated"`
	UpdatedBy                  string                `json:"updatedBy"`
	CreatedBy                  string                `json:"createdBy"`
	Version                    int                   `json:"version"`
	HasACL                     bool                  `json:"hasAcl" xorm:"has_acl"`
	IsFolder                   bool                  `json:"isFolder"`
	FolderId                   int64                 `json:"folderId"`
	FolderUid                  string                `json:"folderUid"`
	FolderTitle                string                `json:"folderTitle"`
	FolderUrl                  string                `json:"folderUrl"`
	Provisioned                bool                  `json:"provisioned"`
	ProvisionedExternalId      string                `json:"provisionedExternalId"`
	AnnotationsPermissions     *AnnotationPermission `json:"annotationsPermissions"`
	PublicDashboardAccessToken string                `json:"publicDashboardAccessToken"`
	PublicDashboardUID         string                `json:"publicDashboardUid"`
	PublicDashboardEnabled     bool                  `json:"publicDashboardEnabled"`
}

type AnnotationPermission struct {
	Dashboard    AnnotationActions `json:"dashboard"`
	Organization AnnotationActions `json:"organization"`
}

type AnnotationActions struct {
	CanAdd    bool `json:"canAdd"`
	CanEdit   bool `json:"canEdit"`
	CanDelete bool `json:"canDelete"`
}

type SaveDashboardCommand struct {
	Dashboard    map[string]interface{} `json:"dashboard" binding:"Required"`
	UserID       int64                  `json:"userId" xorm:"user_id"`
	Overwrite    bool                   `json:"overwrite"`
	Message      string                 `json:"message"`
	OrgID        int64                  `json:"-" xorm:"org_id"`
	RestoredFrom int                    `json:"-"`
	PluginID     string                 `json:"-" xorm:"plugin_id"`
	FolderID     int64                  `json:"folderId" xorm:"folder_id"`
	FolderUID    string                 `json:"folderUid" xorm:"folder_uid"`
	IsFolder     bool                   `json:"isFolder"`
	UpdatedAt    time.Time
}

type Panels struct {
	Datasource  Datasource  `json:"datasource"`
	FieldConfig FieldConfig `json:"fieldConfig"`
	GridPos     GridPos     `json:"gridPos"`
	ID          int         `json:"id"`
	Options     Options     `json:"options"`
	Targets     []Targets   `json:"targets"`
	Title       string      `json:"title"`
	Type        string      `json:"type"`
}

type Datasource struct {
	Type string `json:"type"`
	UID  string `json:"uid"`
}
type Color struct {
	Mode string `json:"mode"`
}
type HideFrom struct {
	Legend  bool `json:"legend"`
	Tooltip bool `json:"tooltip"`
	Viz     bool `json:"viz"`
}
type ScaleDistribution struct {
	Type string `json:"type"`
}
type Stacking struct {
	Group string `json:"group"`
	Mode  string `json:"mode"`
}
type ThresholdsStyle struct {
	Mode string `json:"mode"`
}
type Custom struct {
	AxisCenteredZero  bool              `json:"axisCenteredZero"`
	AxisColorMode     string            `json:"axisColorMode"`
	AxisLabel         string            `json:"axisLabel"`
	AxisPlacement     string            `json:"axisPlacement"`
	BarAlignment      int               `json:"barAlignment"`
	DrawStyle         string            `json:"drawStyle"`
	FillOpacity       int               `json:"fillOpacity"`
	GradientMode      string            `json:"gradientMode"`
	HideFrom          HideFrom          `json:"hideFrom"`
	LineInterpolation string            `json:"lineInterpolation"`
	LineWidth         int               `json:"lineWidth"`
	PointSize         int               `json:"pointSize"`
	ScaleDistribution ScaleDistribution `json:"scaleDistribution"`
	ShowPoints        string            `json:"showPoints"`
	SpanNulls         bool              `json:"spanNulls"`
	Stacking          Stacking          `json:"stacking"`
	ThresholdsStyle   ThresholdsStyle   `json:"thresholdsStyle"`
}
type Steps struct {
	Color string      `json:"color"`
	Value interface{} `json:"value"`
}
type Thresholds struct {
	Mode  string  `json:"mode"`
	Steps []Steps `json:"steps"`
}
type Defaults struct {
	Color      Color         `json:"color"`
	Custom     Custom        `json:"custom"`
	Mappings   []interface{} `json:"mappings"`
	Thresholds Thresholds    `json:"thresholds"`
}
type FieldConfig struct {
	Defaults  Defaults      `json:"defaults"`
	Overrides []interface{} `json:"overrides"`
}
type GridPos struct {
	H int `json:"h"`
	W int `json:"w"`
	X int `json:"x"`
	Y int `json:"y"`
}
type Legend struct {
	Calcs       []interface{} `json:"calcs"`
	DisplayMode string        `json:"displayMode"`
	Placement   string        `json:"placement"`
	ShowLegend  bool          `json:"showLegend"`
}
type Tooltip struct {
	Mode string `json:"mode"`
	Sort string `json:"sort"`
}
type Options struct {
	Legend  Legend  `json:"legend"`
	Tooltip Tooltip `json:"tooltip"`
}
type Targets struct {
	Datasource   Datasource `json:"datasource"`
	EditorMode   string     `json:"editorMode"`
	Expr         string     `json:"expr"`
	LegendFormat string     `json:"legendFormat"`
	Range        bool       `json:"range"`
	RefID        string     `json:"refId"`
}

type DynMap map[string]interface{}

type DsAccess string

//type Datasource struct {
//	ID          int      `json:"id"`
//	UID         string   `json:"uid"`
//	OrgID       int      `json:"orgId"`
//	Name        string   `json:"name"`
//	Type        string   `json:"type"`
//	TypeName    string   `json:"typeName"`
//	TypeLogoURL string   `json:"typeLogoUrl"`
//	Access      string   `json:"access"`
//	URL         string   `json:"url"`
//	User        string   `json:"user"`
//	Database    string   `json:"database"`
//	BasicAuth   bool     `json:"basicAuth"`
//	IsDefault   bool     `json:"isDefault"`
//	JSONData    JSONData `json:"jsonData,omitempty"`
//	ReadOnly    bool     `json:"readOnly"`
//}

type ResponseError struct {
	Message string `json:"message"`
	TraceID string `json:"traceID"`
}

func (r ResponseError) Error() string {
	return r.TraceID + r.Message
}

type JSONData struct {
	HandleGrafanaManagedAlerts bool   `json:"handleGrafanaManagedAlerts"`
	Implementation             string `json:"implementation"`
	HTTPMethod                 string `json:"httpMethod"`
}

type DtoDataSource struct {
	Id               int64           `json:"id"`
	UID              string          `json:"uid"`
	OrgId            int64           `json:"orgId"`
	Name             string          `json:"name"`
	Type             string          `json:"type"`
	TypeLogoUrl      string          `json:"typeLogoUrl"`
	Access           DsAccess        `json:"access"`
	Url              string          `json:"url"`
	User             string          `json:"user"`
	Database         string          `json:"database"`
	BasicAuth        bool            `json:"basicAuth"`
	BasicAuthUser    string          `json:"basicAuthUser"`
	WithCredentials  bool            `json:"withCredentials"`
	IsDefault        bool            `json:"isDefault"`
	JsonData         *Json           `json:"jsonData,omitempty"`
	SecureJsonFields map[string]bool `json:"secureJsonFields"`
	Version          int             `json:"version"`
	ReadOnly         bool            `json:"readOnly"`
	AccessControl    map[string]bool `json:"accessControl,omitempty"`
}

type DataSource struct {
	ID      int64 `json:"id,omitempty" xorm:"pk autoincr 'id'"`
	OrgID   int64 `json:"orgId,omitempty" xorm:"org_id"`
	Version int   `json:"version,omitempty"`

	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Access DsAccess `json:"access"`
	URL    string   `json:"url" xorm:"url"`
	// swagger:ignore
	Password      string `json:"-"`
	User          string `json:"user"`
	Database      string `json:"database"`
	BasicAuth     bool   `json:"basicAuth"`
	BasicAuthUser string `json:"basicAuthUser"`
	// swagger:ignore
	BasicAuthPassword string            `json:"-"`
	WithCredentials   bool              `json:"withCredentials"`
	IsDefault         bool              `json:"isDefault"`
	JsonData          *Json             `json:"jsonData"`
	SecureJsonData    map[string][]byte `json:"secureJsonData"`
	ReadOnly          bool              `json:"readOnly"`
	UID               string            `json:"uid" xorm:"uid"`

	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
}

// AllowedCookies parses the jsondata.keepCookies and returns a list of
// allowed cookies, otherwise an empty list.
//func (ds DataSource) AllowedCookies() []string {
//	if ds.JsonData != nil {
//		if keepCookies := ds.JsonData.Get("keepCookies"); keepCookies != nil {
//			return keepCookies.MustStringArray()
//		}
//	}
//
//	return []string{}
//}

// Specific error type for grpc secrets management so that we can show more detailed plugin errors to users
type ErrDatasourceSecretsPluginUserFriendly struct {
	Err string
}

func (e ErrDatasourceSecretsPluginUserFriendly) Error() string {
	return e.Err
}

// ----------------------
// COMMANDS

// Also acts as api DTO
type AddDataSourceCommand struct {
	Name            string            `json:"name" binding:"Required"`
	Type            string            `json:"type" binding:"Required"`
	Access          DsAccess          `json:"access" binding:"Required"`
	URL             string            `json:"url"`
	Database        string            `json:"database"`
	User            string            `json:"user"`
	BasicAuth       bool              `json:"basicAuth"`
	BasicAuthUser   string            `json:"basicAuthUser"`
	WithCredentials bool              `json:"withCredentials"`
	IsDefault       bool              `json:"isDefault"`
	JsonData        *Json             `json:"jsonData"`
	SecureJsonData  map[string]string `json:"secureJsonData"`
	UID             string            `json:"uid"`

	OrgID                   int64             `json:"-"`
	UserID                  int64             `json:"-"`
	ReadOnly                bool              `json:"-"`
	EncryptedSecureJsonData map[string][]byte `json:"-"`
	UpdateSecretFn          UpdateSecretFn    `json:"-"`
}

// Also acts as api DTO
type UpdateDataSourceCommand struct {
	Name            string            `json:"name" binding:"Required"`
	Type            string            `json:"type" binding:"Required"`
	Access          DsAccess          `json:"access" binding:"Required"`
	URL             string            `json:"url"`
	User            string            `json:"user"`
	Database        string            `json:"database"`
	BasicAuth       bool              `json:"basicAuth"`
	BasicAuthUser   string            `json:"basicAuthUser"`
	WithCredentials bool              `json:"withCredentials"`
	IsDefault       bool              `json:"isDefault"`
	JsonData        *Json             `json:"jsonData"`
	SecureJsonData  map[string]string `json:"secureJsonData"`
	Version         int               `json:"version"`
	UID             string            `json:"uid"`

	OrgID                   int64             `json:"-"`
	ID                      int64             `json:"-"`
	ReadOnly                bool              `json:"-"`
	EncryptedSecureJsonData map[string][]byte `json:"-"`
	UpdateSecretFn          UpdateSecretFn    `json:"-"`
}

// DeleteDataSourceCommand will delete a DataSource based on OrgID as well as the UID (preferred), ID, or Name.
// At least one of the UID, ID, or Name properties must be set in addition to OrgID.
type DeleteDataSourceCommand struct {
	ID   int64
	UID  string
	Name string

	OrgID int64

	DeletedDatasourcesCount int64

	UpdateSecretFn UpdateSecretFn
}

// Function for updating secrets along with datasources, to ensure atomicity
type UpdateSecretFn func() error

// ---------------------
// QUERIES

type SignedInUser struct {
	UserID             int64 `xorm:"user_id"`
	OrgID              int64 `xorm:"org_id"`
	OrgName            string
	OrgRole            string
	ExternalAuthModule string
	ExternalAuthID     string `xorm:"external_auth_id"`
	Login              string
	Name               string
	Email              string
	ApiKeyID           int64 `xorm:"api_key_id"`
	IsServiceAccount   bool  `xorm:"is_service_account"`
	OrgCount           int
	IsGrafanaAdmin     bool
	IsAnonymous        bool
	IsDisabled         bool
	HelpFlags1         uint64
	LastSeenAt         time.Time
	Teams              []int64
	Analytics          AnalyticsSettings
	// Permissions grouped by orgID and actions
	Permissions map[int64]map[string][]string `json:"-"`
}

type AnalyticsSettings struct {
	Identifier         string
	IntercomIdentifier string
}

type GetDataSourcesQuery struct {
	OrgID           int64
	DataSourceLimit int
	User            *SignedInUser
}

type GetAllDataSourcesQuery struct {
	Result []*DataSource
}

type GetDataSourcesByTypeQuery struct {
	OrgID int64 // optional: filter by org_id
	Type  string
}

type GetDefaultDataSourceQuery struct {
	OrgID int64
	User  *SignedInUser
}

// GetDataSourceQuery will get a DataSource based on OrgID as well as the UID (preferred), ID, or Name.
// At least one of the UID, ID, or Name properties must be set in addition to OrgID.
type GetDataSourceQuery struct {
	ID   int64
	UID  string
	Name string

	OrgID int64
}

type DatasourcesPermissionFilterQuery struct {
	User        *SignedInUser
	Datasources []*DataSource
}

//const (
//	QuotaTargetSrv quota.TargetSrv = "data_source"
//	QuotaTarget    quota.Target    = "data_source"
//)

package models

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Client struct {
	IdClient          sql.NullInt16  `json:"idclient"`
	ClientName        string         `json:"clientname"`
	ClientMail        string         `json:"clientmail"`
	ClientPass        string         `json:"clientpass"`
	ClientPhone       string         `json:"clientphone"`
	ClientTelegram    string         `json:"clienttelegram"`
	ClientDateReg     sql.NullString `json:"clientdatereg"`
	ClientTimeZone    sql.NullInt16  `json:"clienttimezone"`
	ClientTimeInfo    sql.NullInt16  `json:"clienttimeinfo"`
	ClientStatus      sql.NullInt16  `json:"clientstatus"`
	IdcCient_IdClient sql.NullInt16  `json:"idclient_idclient"`
}

type Convoy struct {
	IdConvoy          int    `json:"idconvoy"`
	IdConvoy_IdClient int    `json:"idconvoy_idclient"`
	ConvoyName        string `json:"convoyname"`
	ConvoyStatus      int    `json:"convoystatus"`
}

type Machine struct {
	IdMachine         int            `json:"idmachine"`
	IdMachineIdConvoy int            `json:"idmachine_idconvoy"`
	IdMachineIdModel  int            `json:"idmachine_idmodel"`
	IdMachineIdType   int            `json:"idmachine_idtype"`
	MachineYear       string         `json:"machineyear"`
	MachineGosNumber  string         `json:"machinegosnumber"`
	MachineOption     string         `json:"machineoption"`
	MachineDateCome   string         `json:"machinedatecome"`
	MachineDateOut    sql.NullString `json:"machinedateout"`
	MachineSeason     int            `json:"machineseason"`
	MachineKilometr   int            `json:"machinekilometr"`
	MachineMotoHour   int            `json:"machinemotohour"`
	MachineMiles      int            `json:"machinemiles"`
	MachineStatus     int            `json:"machinestatus"`
}

type Manager struct {
	IdManager         int    `json:"idmanager"`
	IdManagerIdClient int    `json:"idmanager_idclient"`
	ManagerName       string `json:"managername"`
	ManagerMail       string `json:"managermail"`
	ManagerPass       string `json:"managerpass"`
	ManagerPhone      string `json:"managerphone"`
	ManagerTelegram   string `json:"managertelegram"`
	ManagerStatus     int    `json:"managerstatus"`
}

type MhKm struct {
	IdMHKM          int    `json:"idmhkm"`
	IdMHKMIdMachine int    `json:"idmhkm_idmachine"`
	MotoHour        int    `json:"motohour"`
	Kilometr        int    `json:"kilometr"`
	Miles           int    `json:"miles"`
	MHKMDate        string `json:"mhkmdate"`
	MHKMName        string `json:"mhkmname"`
}

type Service struct {
	IdService         int    `json:"idservice"`
	IdServiceIdClient int    `json:"idservice_idclient"`
	ServiceName       string `json:"servicename"`
	MotoHourStandart  int    `json:"motohourstandart"`
	KilometrStandart  int    `json:"kilometrstandart"`
	MilesStandart     int    `json:"milesstandart"`
	ServiceStatus     int    `json:"servicestatus"`
}

type Repair struct {
	IdRepair         int    `json:"idrepair"`
	IdRepairIdClient int    `json:"idrepair_idclient"`
	RepairName       string `json:"repairname"`
	RepairCaption    string `json:"repaircaption"`
	RepairStatus     int    `json:"repairstatus"`
}

type Balance struct {
	IdBalance         int64  `json:"idbalance" gorm:"column:idbalance;primaryKey"`
	IdBalanceIdClient int64  `json:"idbalance_idclient" gorm:"column:idbalance_idclient"`
	BalanceDateStart  string `json:"balancedatestart" gorm:"column:balancedatestart"`
	BalanceDateStop   string `json:"balancedatestop" gorm:"column:balancedatestop"`
	Balance           int64  `json:"balance" gorm:"column:balance"`
	BalanceCaption    string `json:"balancecaption" gorm:"column:balancecaption"`
	BalanceStatus     int64  `json:"balancestatus" gorm:"column:balancestatus"`
}

type ServiceDone struct {
	ID          int           `json:"idservicedone"`
	MachineID   int           `json:"idservicedone_idmachine"`
	ServiceID   int           `json:"idservicedone_idservice"`
	MotoHour    sql.NullInt64 `json:"servicedonemotohour,omitempty"`
	Kilometr    sql.NullInt64 `json:"servicedonekilometr,omitempty"`
	Miles       sql.NullInt64 `json:"servicedonemiles,omitempty"`
	ServiceDate string        `json:"servicedonedate"`
	ServiceName string        `json:"servicedonename"`
}

type RepairDone struct {
	IdRepairDone          int           `json:"idrepairdone"`
	IdRepairDoneIdMachine int           `json:"idrepairdone_idmachine"`
	IdRepairDoneIdRepair  int           `json:"idrepairdone_idrepair"`
	RepairDoneMotoHour    sql.NullInt64 `json:"repairdonemotohour,omitempty"`
	RepairDoneKilometr    sql.NullInt64 `json:"repairdonekilometr,omitempty"`
	RepairDoneMiles       sql.NullInt64 `json:"repairdonemiles,omitempty"`
	RepairDoneCaption     string        `json:"repairdonecaption"`
	RepairDoneDate        string        `json:"repairdonedate"`
	RepairDoneName        string        `json:"repairdonename"`
}

type MachineInfo struct {
	IdMachine         int    `json:"id_machine"`
	TypeRus           string `json:"type_rus"`
	Marka             string `json:"marka"`
	Model             string `json:"model"`
	MachineYear       int    `json:"machine_year"`
	MachineGosNumber  string `json:"machine_gos_number"`
	MachineStatus     int    `json:"machine_status"`
	MachineDateCome   string `json:"machine_date_come"`
	MachineOption     string `json:"machine_option"`
	MachineSeason     int    `json:"machine_season"`
	MachineKilometr   int    `json:"machine_kilometr"`
	MachineMotoHour   int    `json:"machine_moto_hour"`
	MachineMiles      int    `json:"machine_miles"`
	MachineStatusWord string `json:"machine_status_word"`
	MachineSeasonWord string `json:"machine_season_word"`
}
type ConvoyInfo struct {
	IdConvoy   string `json:"idconvoy"`
	ConvoyName string `json:"convoyname"`
}

type GuideInfo struct {
}

type InfoPhoto struct {
	IdInfoPhoto       int    `json:"idinfophoto"`
	IdInfoPhotoIdInfo int    `json:"idinfophoto_idinfo"`
	InfoPhoto         string `json:"infophoto"`
}

type Type struct {
	IdType  int    `json:"idtype"`
	TypeRus string `json:"type_rus"`
}

type Marka struct {
	IdMarka  string `json:"idmarka"`
	MarkaRus string `json:"marka_rus"`
}

type Model struct {
	IdModel        string `json:"idmodel"`
	IdModelIdMarka string `json:"idmodel_idmarka"`
	Model          string `json:"model"`
	ModelStatus    int    `json:"modelstatus"`
}

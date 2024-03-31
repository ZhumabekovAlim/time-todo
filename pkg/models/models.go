package models

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Client struct {
	IdClient          string `json:"idclient"`
	ClientName        string `json:"clientname"`
	ClientMail        string `json:"clientmail"`
	ClientPass        string `json:"clientpass"`
	ClientPhone       string `json:"clientphone"`
	ClientTelegram    string `json:"clienttelegram"`
	ClientDateReg     string `json:"clientdatereg"`
	ClientTimeZone    uint16 `json:"clienttimezone"`
	ClientTimeInfo    uint16 `json:"clienttimeinfo"`
	ClientStatus      uint16 `json:"clientstatus"`
	IdcCient_IdClient uint16 `json:"idclient_idclient"`
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
	MachineGosNumber  sql.NullString `json:"machinegosnumber"`
	MachineOption     sql.NullString `json:"machineoption"`
	MachineDateCome   sql.NullString `json:"machinedatecome"`
	MachineDateOut    sql.NullString `json:"machinedateout"`
	MachineSeason     int            `json:"machineseason"`
	MachineMotoHour   int            `json:"machinemotohour"`
	MachineKilometr   int            `json:"machinekilometr"`
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
	IdMHKM          int           `json:"idmhkm"`
	IdMHKMIdMachine int           `json:"idmhkm_idmachine"`
	MotoHour        sql.NullInt64 `json:"motohour"`
	Kilometr        sql.NullInt64 `json:"kilometr"`
	Miles           sql.NullInt64 `json:"miles"`
	MHKMDate        time.Time     `json:"mhkmdate"`
	MHKMName        string        `json:"mhkmname"`
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
	IdBalance         int            `json:"idbalance"`
	IdBalanceIdClient int            `json:"idbalance_idclient"`
	BalanceDateStart  string         `json:"balancedatestart"`
	BalanceDateStop   sql.NullString `json:"balancedatestop"`
	Balance           int            `json:"balance"`
	BalanceCaption    string         `json:"balancecaption"`
	BalanceStatus     int            `json:"balancestatus"`
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

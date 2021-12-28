package sap_api_output_formatter

type EquipmentReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	Equipment     string `json:"equipment_code"`
	Deleted       bool   `json:"deleted"`
}

type Equipment struct {
	Equipment                     string `json:"Equipment"`
	ValidityEndDate               string `json:"ValidityEndDate"`
	ValidityEndTime               string `json:"ValidityEndTime"`
	ValidityStartDate             string `json:"ValidityStartDate"`
	EquipmentName                 string `json:"EquipmentName"`
	EquipmentCategory             string `json:"EquipmentCategory"`
	TechnicalObjectType           string `json:"TechnicalObjectType"`
	GrossWeight                   string `json:"GrossWeight"`
	GrossWeightUnit               string `json:"GrossWeightUnit"`
	SizeOrDimensionText           string `json:"SizeOrDimensionText"`
	InventoryNumber               string `json:"InventoryNumber"`
	OperationStartDate            string `json:"OperationStartDate"`
	AcquisitionValue              string `json:"AcquisitionValue"`
	Currency                      string `json:"Currency"`
	AcquisitionDate               string `json:"AcquisitionDate"`
	AssetManufacturerName         string `json:"AssetManufacturerName"`
	ManufacturerPartTypeName      string `json:"ManufacturerPartTypeName"`
	ManufacturerCountry           string `json:"ManufacturerCountry"`
	ConstructionYear              string `json:"ConstructionYear"`
	ConstructionMonth             string `json:"ConstructionMonth"`
	ManufacturerPartNmbr          string `json:"ManufacturerPartNmbr"`
	ManufacturerSerialNumber      string `json:"ManufacturerSerialNumber"`
	MaintenancePlant              string `json:"MaintenancePlant"`
	AssetLocation                 string `json:"AssetLocation"`
	AssetRoom                     string `json:"AssetRoom"`
	PlantSection                  string `json:"PlantSection"`
	WorkCenter                    string `json:"WorkCenter"`
	WorkCenterPlant               string `json:"WorkCenterPlant"`
	CompanyCode                   string `json:"CompanyCode"`
	BusinessArea                  string `json:"BusinessArea"`
	MasterFixedAsset              string `json:"MasterFixedAsset"`
	FixedAsset                    string `json:"FixedAsset"`
	CostCenter                    string `json:"CostCenter"`
	WBSElementExternalID          string `json:"WBSElementExternalID"`
	SettlementOrder               string `json:"SettlementOrder"`
	MaintenancePlanningPlant      string `json:"MaintenancePlanningPlant"`
	MaintenancePlannerGroup       string `json:"MaintenancePlannerGroup"`
	MainWorkCenter                string `json:"MainWorkCenter"`
	MainWorkCenterPlant           string `json:"MainWorkCenterPlant"`
	CatalogProfile                string `json:"CatalogProfile"`
	FunctionalLocation            string `json:"FunctionalLocation"`
	SuperordinateEquipment        string `json:"SuperordinateEquipment"`
	EquipInstallationPositionNmbr string `json:"EquipInstallationPositionNmbr"`
	TechnicalObjectSortCode       string `json:"TechnicalObjectSortCode"`
	ConstructionMaterial          string `json:"ConstructionMaterial"`
	Material                      string `json:"Material"`
	EquipmentIsAvailable          bool   `json:"EquipmentIsAvailable"`
	EquipmentIsInstalled          bool   `json:"EquipmentIsInstalled"`
	EquipIsAllocToSuperiorEquip   bool   `json:"EquipIsAllocToSuperiorEquip"`
	EquipHasSubOrdinateEquipment  string `json:"EquipHasSubOrdinateEquipment"`
	CreationDate                  string `json:"CreationDate"`
	LastChangeDateTime            string `json:"LastChangeDateTime"`
	EquipmentIsMarkedForDeletion  bool   `json:"EquipmentIsMarkedForDeletion"`
	ToPartner                     string `json:"to_Partner"`
}

type ToPartner struct {
	Equipment                  string `json:"Equipment"`
	BusinessPartner            string `json:"BusinessPartner"`
	PartnerFunction            string `json:"PartnerFunction"`
	EquipmentPartnerObjectNmbr string `json:"EquipmentPartnerObjectNmbr"`
	Partner                    string `json:"Partner"`
	CreationDate               string `json:"CreationDate"`
}

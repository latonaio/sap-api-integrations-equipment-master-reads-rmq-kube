package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Equipment     struct {
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
		BusinessPartner               struct {
			BusinessPartner            string `json:"BusinessPartner"`
			PartnerFunction            string `json:"PartnerFunction"`
			EquipmentPartnerObjectNmbr string `json:"EquipmentPartnerObjectNmbr"`
			Partner                    string `json:"Partner"`
			CreationDate               string `json:"CreationDate"`
		} `json:"BusinessPartner"`
	} `json:"Equipment"`
	APISchema     string `json:"api_schema"`
	Accepter    []string `json:"accepter"`
	EquipmentCode string `json:"equipment_code"`
	Deleted       bool   `json:"deleted"`
}

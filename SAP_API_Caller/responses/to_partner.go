package responses

type ToPartner struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Equipment                  string `json:"Equipment"`
			BusinessPartner            string `json:"BusinessPartner"`
			PartnerFunction            string `json:"PartnerFunction"`
			EquipmentPartnerObjectNmbr string `json:"EquipmentPartnerObjectNmbr"`
			Partner                    string `json:"Partner"`
			CreationDate               string `json:"CreationDate"`
		} `json:"results"`
	} `json:"d"`
}

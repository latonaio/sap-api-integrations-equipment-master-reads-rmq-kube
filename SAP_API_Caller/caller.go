package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-equipment-master-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetEquipment(equipment, equipmentName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Equipment":
			func() {
				c.Equipment(equipment)
				wg.Done()
			}()
		case "EquipmentName":
			func() {
				c.EquipmentName(equipmentName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Equipment(equipment string) {
	equipmentData, err := c.callEquipmentSrvAPIRequirementEquipment("Equipment", equipment)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": equipmentData, "function": "Equipment"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(equipmentData)

	partnerData, err := c.callToPartner(equipmentData[0].ToPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": partnerData, "function": "EquipmentPartner"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(partnerData)
}

func (c *SAPAPICaller) callEquipmentSrvAPIRequirementEquipment(api, equipment string) ([]sap_api_output_formatter.Equipment, error) {
	url := strings.Join([]string{c.baseURL, "API_EQUIPMENT", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithEquipment(req, equipment)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEquipment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPartner(url string) (*sap_api_output_formatter.ToPartner, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPartner(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) EquipmentName(equipmentName string) {
	data, err := c.callEquipmentSrvAPIRequirementEquipmentName("Equipment", equipmentName)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "Equipment"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callEquipmentSrvAPIRequirementEquipmentName(api, equipmentName string) ([]sap_api_output_formatter.Equipment, error) {
	url := strings.Join([]string{c.baseURL, "API_EQUIPMENT", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithEquipmentName(req, equipmentName)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEquipment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithEquipment(req *http.Request, equipment string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Equipment eq '%s'", equipment))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithEquipmentName(req *http.Request, equipmentName string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("substringof('%s', EquipmentName)", equipmentName))
	req.URL.RawQuery = params.Encode()
}

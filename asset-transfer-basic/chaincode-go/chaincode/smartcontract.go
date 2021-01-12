package chaincode
import (
	"encoding/json"
	"fmt"
	_"github.com/cd1/utils-golang"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}
// Asset describes basic details of what makes up a simple asset
type Asset struct {
	ID             string `json:"ID"`
	Color          string `json:"color"`
	Size           int    `json:"size"`
	Owner          string `json:"owner"`
	AppraisedValue int    `json:"appraisedValue"`
}
//metaDAta STore
//type MetaDataStore struct {
//	Doctype string
//	User    string
//	Metadata string
//	Key string
//}
// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{ID: "asset1", Color: "blue", Size: 5, Owner: "Tomoko", AppraisedValue: 300},
		{ID: "asset2", Color: "red", Size: 5, Owner: "Brad", AppraisedValue: 400},
		{ID: "asset3", Color: "green", Size: 10, Owner: "Jin Soo", AppraisedValue: 500},
		{ID: "asset4", Color: "yellow", Size: 10, Owner: "Max", AppraisedValue: 600},
		{ID: "asset5", Color: "black", Size: 15, Owner: "Adriana", AppraisedValue: 700},
		{ID: "asset6", Color: "white", Size: 15, Owner: "Michel", AppraisedValue: 800},
	}
	metaDatas := [] MetaDataStore{
		{ Doctype: "MetaData Store", User: "www.idp1.org", Metadata: "1234rf", Key: "1231"},
		{ Doctype: "MetaData Store", User: "www.idp2.org", Metadata: "1234rf", Key: "1232"},
		{ Doctype: "MetaData Store", User: "www.idp3.org", Metadata: "1234rf", Key: "1233"},
		{ Doctype: "MetaData Store", User: "www.idp4.org", Metadata: "1234rf", Key: "1234"},
	}
	for _, metaData := range metaDatas {
		metaDataJSON, err := json.Marshal(metaData)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(metaData.Key, metaDataJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}
	asset := Asset{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}
func (s *SmartContract) CreateCar(ctx contractapi.TransactionContextInterface, carNumber string, make string, model string, colour string, owner string) error {
	car := Car{
		Make:   make,
		Model:  model,
		Colour: colour,
		Owner:  owner,
	}

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

// store metadata
//func (s *SmartContract) StoreMetadata(ctx contractapi.TransactionContextInterface, user string, metaData string) error {
//	exists, err := s.MetaDataExists(ctx, user)
//	if err != nil {
//		return err
//	}
//	if exists {
//		return fmt.Errorf("the metaData %s already exists", user)
//	}
//	key := utils.RandomString()
//	metaDataStore := MetaDataStore {
//		Doctype: "MetaData Store",
//		User  : user,
//		Metadata : metaData,
//		Key : key,
//	}
//	metaDataStoreJSON, err := json.Marshal(metaDataStore)
//	if err != nil {
//		return err
//	}
//
//	return ctx.GetStub().PutState(key, metaDataStoreJSON)
//}
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}
// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}
	println(asset.Color)
	return &asset, nil
}
//read metadata
func (s *SmartContract) ReadMetaData(ctx contractapi.TransactionContextInterface, user string) (*MetaDataStore, error) {
	metaDataStoreJSON, err := ctx.GetStub().GetState(user)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if metaDataStoreJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", user)
	}

	var metaDataStore MetaDataStore
	err = json.Unmarshal(metaDataStoreJSON, &metaDataStore)
	if err != nil {
		return nil, err
	}
	//println(metaDataStore)
	println(metaDataStore.User)
	return &metaDataStore, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	// overwriting original asset with new asset
	asset := Asset{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ID exists in world state


//metaDAta exists
func (s *SmartContract) MetaDataExists(ctx contractapi.TransactionContextInterface, User string) (bool, error) {
	metaDataJSON, err := ctx.GetStub().GetState(User)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return metaDataJSON != nil, nil
}

// TransferAsset updates the owner field of asset with given id in world state.
func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) error {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.Owner = newOwner
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

/// GET all metadata
//func (s *SmartContract) GetAllMetaData(ctx contractapi.TransactionContextInterface) ([]*MetaDataStore, error) {
//	// range query with empty string for startKey and endKey does an
//	// open-ended query of all assets in the chaincode namespace.
////	queryString := newCouchQueryBuilder().addSelector("Doctype", "MetaData Store").getQueryString()
////	resultsIterator, err := ctx.GetStub().GetPrivateData("Doctype", "MetaData Store")
//
//	//resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
//	if err != nil {
//		return nil, err
//	}
//	defer resultsIterator.Close()
//
//	var metaDatas []*MetaDataStore
//	for resultsIterator.HasNext() {
//		queryResponse, err := resultsIterator.Next()
//		if err != nil {
//			return nil, err
//		}
//
//		var metaData MetaDataStore
//		err = json.Unmarshal(queryResponse.Value, &metaData)
//		if err != nil {
//			return nil, err
//		}
//		metaDatas = append(metaDatas, &metaData)
//	}
//
//	return metaDatas, nil
//}
//


//func main() {
//	// Create a new Smart Contract
//}

func (s *SmartContract) ReadMeta(ctx contractapi.TransactionContextInterface, User string) (*MetaDataStore, error) {
	assetJSON, err := ctx.GetStub().GetState(User)
	if err != nil {
		return nil, err
	} else if assetJSON == nil {
		return nil, fmt.Errorf("%s does not exist",User)
	}

	asset := new(MetaDataStore)
	err = json.Unmarshal(assetJSON, asset)
	if err != nil {
		return nil, err
	}
   println(asset.User)
	return asset, nil
}
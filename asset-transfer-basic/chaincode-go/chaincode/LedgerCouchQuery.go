package chaincode

//import (
//	"encoding/json"
//	"fmt"
//	"github.com/hyperledger/fabric-chaincode-go/shim"
//	"github.com/hyperledger/fabric-contract-api-go/contractapi"
//	"strings"
//	"time"
//)

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/hyperledger/fabric/core/chaincode/shim"
//	"github.com/jmoiron/jsonq"
//	"strings"
//)
//
//type QueryResponse struct {
//	Key string
//	Record []byte
//	Query *jsonq.JsonQuery
//}
//func lastQueryValueForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
//	resultsIterator, err := stub.GetQueryResult(queryString)
//	if err != nil {
//		return nil, err
//	}
//	defer closeIterator(resultsIterator)
//
//	x,_:=resultsIterator.Next()
//	for resultsIterator.HasNext() {
//		x,_=resultsIterator.Next()
//
//
//	}
//
//	if !resultsIterator.HasNext() {
//		fmt.Println(x)
//		return x.Value, nil
//	}
//
//	return []byte(""), nil
//}
//func firstQueryValueForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
//	resultsIterator, err := stub.GetQueryResult(queryString)
//	if err != nil {
//		return nil, err
//	}
//	// for i:=0; resultsIterator.HasNext(); i++ {
//	// 	data := resultsIterator.Next()
//	// }
//	data, err := resultsIterator.Next()
//	if err != nil {
//		return nil, err
//	}
//	value := data.Value
//	return value, nil
//}
//
//func decodeSingleResponse(jsonResponse []byte) *QueryResponse {
//	data := map[string]interface{}{}
//	dec := json.NewDecoder(strings.NewReader(string(jsonResponse)))
//	err := dec.Decode(&data)
//	if err!=nil {
//		fmt.Println(err.Error())
//	}
//	jq := jsonq.NewQuery(data)
//
//	key, err := jq.String("Key")
//	if err!=nil {
//		fmt.Println(err.Error())
//	}
//	record, err := jq.String("Record")
//	if err!=nil {
//		fmt.Println(err.Error())
//	}
//	recordByteArray := []byte(record)
//
//	return &QueryResponse{Key: key , Record: recordByteArray, Query: jq }
//}
//
//func getJSONQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
//	start := "{\"values\": "
//	end := "}"
//	data, err := getQueryResultForQueryString(stub, queryString)
//	if err!=nil {
//		return nil, err
//	}
//	return []byte(start+string(data)+end), nil
//}
//
//// =========================================================================================
//// getQueryResultForQueryString executes the passed in query string.
//// Result set is built and returned as a byte array containing the JSON results.
//// =========================================================================================
//func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
//
//	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)
//
//	resultsIterator, err := stub.GetQueryResult(queryString)
//	if err != nil {
//		return nil, err
//	}
//	defer closeIterator(resultsIterator)
//
//	buffer, err := constructQueryResponseFromIterator(resultsIterator)
//	if err != nil {
//		return nil, err
//	}
//
//	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
//
//	return buffer.Bytes(), nil
//}
//
//func allUserQueryVAlueForQueryString(stub shim.ChaincodeStubInterface, queryString string) []byte {
//	resultIterator ,err := stub.GetQueryResult(queryString)
//
//	if err != nil{
//		return nil
//	}
//	defer closeIterator(resultIterator)
//
//	metaDataArray := []MetaDataStore{}
//
//	x, _:= resultIterator.Next()
//	for resultIterator.HasNext(){
//
//		var metaData MetaDataStore
//		_=json.Unmarshal(x.Value,&metaData)
//		metaDataArray = append(metaDataArray,metaData)
//		x,_=resultIterator.Next()
//
//	}
//	var userMetaData MetaDataStore
//	_=json.Unmarshal(x.Value,&userMetaData)
//	metaDataArray = append (metaDataArray,userMetaData)
//
//	resultArrayJjson, _:=json.Marshal(metaDataArray)
//
//	return (resultArrayJjson)
//}
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//// ===========================================================================================
//// constructQueryResponseFromIterator constructs a JSON array containing query results from
//// a given result iterator
//// ===========================================================================================
//func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
//	// buffer is a JSON array containing QueryResults
//	var buffer bytes.Buffer
//	buffer.WriteString("[")
//
//	bArrayMemberAlreadyWritten := false
//	for resultsIterator.HasNext() {
//		queryResponse, err := resultsIterator.Next()
//		if err != nil {
//			return nil, err
//		}
//		// Add a comma before array members, suppress it for the first array member
//		if bArrayMemberAlreadyWritten == true {
//			buffer.WriteString(",")
//		}
//		buffer.WriteString("{\"Key\":")
//		buffer.WriteString("\"")
//		buffer.WriteString(queryResponse.Key)
//		buffer.WriteString("\"")
//
//		buffer.WriteString(", \"Record\":")
//		// Record is a JSON object, so we write as-is
//		buffer.WriteString(string(queryResponse.Value))
//		buffer.WriteString("}")
//		bArrayMemberAlreadyWritten = true
//	}
//	buffer.WriteString("]")
//
//	return &buffer, nil
//}
//
//func firstQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
//
//	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)
//
//	resultsIterator, err := stub.GetQueryResult(queryString)
//	if err != nil {
//		return nil, err
//	}
//	defer closeIterator(resultsIterator)
//
//	buffer, err := firstQueryResponseFromIterator(resultsIterator)
//	if err != nil {
//		return nil, err
//	}
//
//	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
//
//	return buffer.Bytes(), nil
//}
//
//func closeIterator(resultsIterator shim.StateQueryIteratorInterface) {
//	err := resultsIterator.Close()
//	if err!=nil {
//		fmt.Println(err.Error())
//	}
//}
//
//// ===========================================================================================
//// firstQueryResponseFromIterator returns query results from
//// a given result iterator
//// ===========================================================================================
//func firstQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
//	// buffer is a JSON array containing QueryResults
//	var buffer bytes.Buffer
//
//	bArrayMemberAlreadyWritten := false
//	for resultsIterator.HasNext() {
//		queryResponse, err := resultsIterator.Next()
//		if err != nil {
//			return nil, err
//		}
//		// Add a comma before array members, suppress it for the first array member
//		if bArrayMemberAlreadyWritten == true {
//			buffer.WriteString(",")
//		}
//		buffer.WriteString("{\"Key\":")
//		buffer.WriteString("\"")
//		buffer.WriteString(queryResponse.Key)
//		buffer.WriteString("\"")
//
//		buffer.WriteString(", \"Record\":")
//		// Record is a JSON object, so we write as-is
//		buffer.WriteString(string(queryResponse.Value))
//		buffer.WriteString("}")
//		bArrayMemberAlreadyWritten = true
//
//		break
//	}
//
//	return &buffer, nil
//}
//
//type QueryResult struct {
//	Record              *MetaDataStore
//	TxId                string    `json:"txId"`
//	Timestamp           time.Time `json:"timestamp"`
//	FetchedRecordsCount int       `json:"fetchedRecordsCount"`
//	Bookmark            string    `json:"bookmark"`
//}
//// constructQueryResponseFromIterator constructs a JSON array containing query results from
//// a given result iterator
//func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*QueryResult, error) {
//	resp := []*QueryResult{}
//
//	for resultsIterator.HasNext() {
//		queryResponse, err := resultsIterator.Next()
//		if err != nil {
//			return nil, err
//		}
//
//		newRecord := new(QueryResult)
//		err = json.Unmarshal(queryResponse.Value, newRecord)
//		if err != nil {
//			return nil, err
//		}
//
//		resp = append(resp, newRecord)
//	}
//	return resp, nil
//}
//// GetAssetsByRange performs a range query based on the start and end keys provided.
//// Read-only function results are not typically submitted to ordering. If the read-only
//// results are submitted to ordering, or if the query is used in an update transaction
//// and submitted to ordering, then the committing peers will re-execute to guarantee that
//// result sets are stable between endorsement time and commit time. The transaction is
//// invalidated by the committing peers if the result set has changed between endorsement
//// time and commit time.
//// Therefore, range queries are a safe option for performing update transactions based on query results.
//func (t *SmartContract) GetAssetsByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*QueryResult, error) {
//	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
//	if err != nil {
//		return nil, err
//	}
//	defer resultsIterator.Close()
//
//	return constructQueryResponseFromIterator(resultsIterator)
//}
//
//// TransferAssetBasedOnColor will transfer assets of a given color to a certain new owner.
//// Uses a GetStateByPartialCompositeKey (range query) against color~name 'index'.
//// Committing peers will re-execute range queries to guarantee that result sets are stable
//// between endorsement time and commit time. The transaction is invalidated by the
//// committing peers if the result set has changed between endorsement time and commit time.
//// Therefore, range queries are a safe option for performing update transactions based on query results.
//// Example: GetStateByPartialCompositeKey/RangeQuery
//func (t *SmartContract) TransferAssetBasedOnColor(ctx contractapi.TransactionContextInterface, color, newOwner string) error {
//	newOwner = strings.ToLower(newOwner)
//
//	// Query the color~name index by color
//	// This will execute a key range query on all keys starting with 'color'
//	coloredAssetResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey("color~name", []string{color})
//	if err != nil {
//		return fmt.Errorf(err.Error())
//	}
//	defer coloredAssetResultsIterator.Close()
//
//	// Iterate through result set and for each asset found, transfer to newOwner
//	var i int
//	for i = 0; coloredAssetResultsIterator.HasNext(); i++ {
//		// Note that we don't get the value (2nd return variable), we'll just get the asset name from the composite key
//		responseRange, err := coloredAssetResultsIterator.Next()
//		if err != nil {
//			return fmt.Errorf(err.Error())
//		}
//
//		// get the color and name from color~name composite key
//		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
//		if err != nil {
//			return fmt.Errorf(err.Error())
//		}
//
//		if len(compositeKeyParts) > 2 {
//			returnedAssetID := compositeKeyParts[1]
//
//			// Now call the transfer function for the found asset.
//			// Re-use the same function that is used to transfer individual assets
//			err = t.TransferAsset(ctx, returnedAssetID, newOwner)
//			// if the transfer failed break out of loop and return error
//			if err != nil {
//				return fmt.Errorf("Transfer failed: %v", err)
//			}
//		}
//	}
//
//	return nil
//}
//
//// QueryAssetsByOwner queries for assets based on a passed in owner.
//// This is an example of a parameterized query where the query logic is baked into the chaincode,
//// and accepting a single query parameter (owner).
//// Only available on state databases that support rich query (e.g. CouchDB)
//// Example: Parameterized rich query
//func (t *SmartContract) QueryAssetsByOwner(ctx contractapi.TransactionContextInterface, owner string) ([]*QueryResult, error) {
//	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"asset\",\"owner\":\"%s\"}}", owner)
//
//	return getQueryResultForQueryString(ctx, queryString)
//}
//
//// QueryAssets uses a query string to perform a query for assets.
//// Query string matching state database syntax is passed in and executed as is.
//// Supports ad hoc queries that can be defined at runtime by the client.
//// If this is not desired, follow the QueryAssetsForOwner example for parameterized queries.
//// Only available on state databases that support rich query (e.g. CouchDB)
//// Example: Ad hoc rich query
//func (t *SmartContract) QueryAssets(ctx contractapi.TransactionContextInterface, queryString string) ([]*QueryResult, error) {
//	return getQueryResultForQueryString(ctx, queryString)
//}
//
//// getQueryResultForQueryString executes the passed in query string.
//// Result set is built and returned as a byte array containing the JSON results.
//func getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*QueryResult, error) {
//
//	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
//	if err != nil {
//		return nil, err
//	}
//	defer resultsIterator.Close()
//
//	return constructQueryResponseFromIterator(resultsIterator)
//}
//
//// GetAssetsByRangeWithPagination performs a range query based on the start & end key,
//// page size and a bookmark.
//// The number of fetched records will be equal to or lesser than the page size.
//// Paginated range queries are only valid for read only transactions.
//// Example: Pagination with Range Query
//func (t *SmartContract) GetAssetsByRangeWithPagination(ctx contractapi.TransactionContextInterface, startKey,
//	endKey, bookmark string, pageSize int) ([]*QueryResult, error) {
//
//	resultsIterator, _, err := ctx.GetStub().GetStateByRangeWithPagination(startKey, endKey, int32(pageSize), bookmark)
//	if err != nil {
//		return nil, err
//	}
//	defer resultsIterator.Close()
//
//	return constructQueryResponseFromIterator(resultsIterator)
//}
//
//// QueryAssetsWithPagination uses a query string, page size and a bookmark to perform a query
//// for assets. Query string matching state database syntax is passed in and executed as is.
//// The number of fetched records would be equal to or lesser than the specified page size.
//// Supports ad hoc queries that can be defined at runtime by the client.
//// If this is not desired, follow the QueryAssetsForOwner example for parameterized queries.
//// Only available on state databases that support rich query (e.g. CouchDB)
//// Paginated queries are only valid for read only transactions.
//// Example: Pagination with Ad hoc Rich Query
//func (t *SmartContract) QueryAssetsWithPagination(ctx contractapi.TransactionContextInterface, queryString,
//	bookmark string, pageSize int) ([]*QueryResult, error) {
//	return getQueryResultForQueryStringWithPagination(ctx, queryString, int32(pageSize), bookmark)
//}
//
//// getQueryResultForQueryStringWithPagination executes the passed in query string with
//// pagination info. Result set is built and returned as a byte array containing the JSON results.
//func getQueryResultForQueryStringWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int32, bookmark string) ([]*QueryResult, error) {
//
//	resultsIterator, _, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
//	if err != nil {
//		return nil, err
//	}
//	defer resultsIterator.Close()
//
//	return constructQueryResponseFromIterator(resultsIterator)
//}
//
//// GetAssetHistory returns the chain of custody for an asset since issuance.
//func (t *SmartContract) GetAssetHistory(ctx contractapi.TransactionContextInterface, assetID string) ([]QueryResult, error) {
//
//	resultsIterator, err := ctx.GetStub().GetHistoryForKey(assetID)
//	if err != nil {
//		return nil, err
//	}
//	defer resultsIterator.Close()
//
//	records := []QueryResult{}
//
//	for resultsIterator.HasNext() {
//		response, err := resultsIterator.Next()
//		if err != nil {
//			return nil, err
//		}
//
//		asset := new(Asset)
//		err = json.Unmarshal(response.Value, asset)
//		if err != nil {
//			return nil, err
//		}
//
//		record := QueryResult{
//			TxId:      response.TxId,
//			Timestamp: time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)),
//			Record:    asset,
//		}
//		records = append(records, record)
//	}
//
//	return records, nil
//}
//
//// AssetExists returns true when asset with given ID exists in world state
////func (f *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
////	assetJSON, err := ctx.GetStub().GetState(id)
////	if err != nil {
////		return false, fmt.Errorf("Failed to read from world state. %s", err.Error())
////	}
////
////	return assetJSON != nil, nil
////}
//
//// InitLedger creates sample assets in the ledger
////func (f *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
////	assets := []asset{
////		asset{ID: "asset1", Color: "blue", Size: 5, Owner: "Tomoko", AppraisedValue: 300},
////		asset{ID: "asset2", Color: "red", Size: 5, Owner: "Brad", AppraisedValue: 400},
////		asset{ID: "asset3", Color: "green", Size: 10, Owner: "Jin Soo", AppraisedValue: 500},
////		asset{ID: "asset4", Color: "yellow", Size: 10, Owner: "Max", AppraisedValue: 600},
////		asset{ID: "asset5", Color: "black", Size: 15, Owner: "Adriana", AppraisedValue: 700},
////		asset{ID: "asset6", Color: "white", Size: 15, Owner: "Michel", AppraisedValue: 800},
////	}
////
////	for _, asset := range assets {
////		assetJSON, err := json.Marshal(asset)
////		if err != nil {
////			return err
////		}
////
////		err = ctx.GetStub().PutState(asset.ID, assetJSON)
////		if err != nil {
////			return fmt.Errorf("Failed to put to world state. %s", err.Error())
////		}
////	}
////
////	return nil
////}
//
//func main() {
//
//	chaincode, err := contractapi.NewChaincode(new(SmartContract))
//
//	if err != nil {
//		fmt.Printf("Error creating asset chaincode: %s", err.Error())
//		return
//	}
//
//	if err := chaincode.Start(); err != nil {
//		fmt.Printf("Error starting asset chaincode: %s", err.Error())
//		return
//	}
//}

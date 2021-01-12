package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	_ "github.com/hyperledger/fabric-contract-api-go/contractapi/utils"
)

//func (s *SmartContract) QueryMetaDataStore(ctx contractapi.TransactionContextInterface, queryString string) ([]MetaDataStore, error) {
//
//	queryResults, err := s.getQueryResultForQueryString(ctx, queryString)
//	if err != nil {
//		return nil, err
//	}
//	return queryResults, nil
//}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func (s *SmartContract) QueryCodeByIdp(ctx contractapi.TransactionContextInterface, idp string) ([]QueryResultCode, error) {

//	userString  := strings.ToLower(user)

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Code Store\",\"whichIdp\":\"%s\"}}", idp)
  fmt.Println(queryString)
	queryResults, err := s.getQueryResultForQueryStringForCode(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}


func (s *SmartContract) getQueryResultForQueryStringForCode(ctx contractapi.TransactionContextInterface, queryString string) ([]QueryResultCode, error) {

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResultCode{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		newMarble := new(CodeStore)

		err = json.Unmarshal(response.Value, newMarble)
		if err != nil {
			return nil, err
		}
	queryResult  := QueryResultCode{response.Key, newMarble}
//	println(queryResult)
	results = append(results, queryResult)
	}
	return results, nil
}
 func (s *SmartContract) GetQueryCode (ctx contractapi.TransactionContextInterface, idp string)([]QueryResultCode, error) {
	 queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Code Store\",\"whichIdp\":\"%s\"}}", idp)
 	 resultIterator, err := ctx.GetStub().GetQueryResult(queryString)
 	  if err != nil{
 	  	return nil, err
	  }
	  defer resultIterator.Close()
 	 var results []QueryResultCode
 	 for resultIterator.HasNext() {
 	 	queryResponse, err :=  resultIterator.Next()
 	 	if err != nil {
 	 		return nil, err
		}
		codeData := new(CodeStore)
		_= json.Unmarshal(queryResponse.Value,codeData)
		queryResults := QueryResultCode{queryResponse.Key,codeData}
		println(queryResponse)
		results = append(results, queryResults)

	 }
 	return results, nil
 }

//
//func (s *SmartContract) getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]MetaDataStore, error) {
//
//	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("MetaData Store", queryString)
//	if err != nil {
//		return nil, err
//	}
//	defer resultsIterator.Close()
//
//	results := []MetaDataStore{}
//
//	for resultsIterator.HasNext() {
//		response, err := resultsIterator.Next()
//		if err != nil {
//			return nil, err
//		}
//
//		newMarble := new(MetaDataStore)
//
//		err = json.Unmarshal(response.Value, newMarble)
//		if err != nil {
//			return nil, err
//		}
//
//		results = append(results, *newMarble)
//	}
//	return results, nil
//}

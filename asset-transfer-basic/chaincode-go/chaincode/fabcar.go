package chaincode
import (
	"encoding/json"
	"fmt"
	"crypto/rand"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	_ "github.com/hyperledger/fabric-contract-api-go/contractapi/utils"
	//"golang.org/x/text/message"
	_ "math/rand"
)
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}
type MetaDataStore struct {
	Doctype string //`json:"docType"`
	User    string //`json:"user"`
	Metadata string //`json:"metaData"`
	Key string
}
type CodeStore struct {
	Doctype string //`json:"docType"`
	ForWhichSP    string ///`json:"forWhichSp"`
	WhichIDP  string //`json:"whichIdp"`
	Code string //`json:"code"`
	//Key string
}
func (s *SmartContract) StoreMetaData(ctx contractapi.TransactionContextInterface, user string, metaData string) error {

	b := make([]byte, 4)
	rand.Read(b)
	key := fmt.Sprintf("%x", b)
	metaDataStore := MetaDataStore {
		Doctype: "MetaData Store",
		User  : user,
		Metadata : metaData,
		Key : key,
	}
	metaDataBytes, _ := json.Marshal(metaDataStore)

	return ctx.GetStub().PutState(metaDataStore.Key, metaDataBytes)
}

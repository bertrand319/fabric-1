package scc

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/common/errors"
)

type SCCContract struct {
}

func (t *SCCContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.ERROR("Incorrect arguments, Expectiong a key and a value")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.ERROR(fmt.Springtf("Failed to create asset: %s", args[0]))
	}
	return shim.Success(nil)
}


func (t *SCCContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	var result string
	var err error
	if fn == "set" {
		result, err = set(stub, args)
	} else {
		result, err = get(stub, args)
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(result))
}

func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments, Expecting a key and a value")
	}

	err := stub.PutState(args[0], []byte(args[1]))

	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return args[1], nil
}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}

func main() {
	if err := shim.Start(new(SCCContract)); err != nil {
		fmt.Printf("Error starting SCCContract chaincode: %s", err)
	}
}

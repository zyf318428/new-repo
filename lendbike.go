/*
	zyf
*/

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	err := stub.PutState("person", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("time", []byte(args[1]))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("position", []byte(args[2]))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("state", []byte(args[3]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation")
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "readPerson" { //read a variable
		return t.readPerson(stub, args)
	}
	// Handle different functions
	if function == "readTime" { //read a variable
		return t.readTime(stub, args)
	}
	// Handle different functions
	if function == "readPosition" { //read a variable
		return t.readPosition(stub, args)
	}
	// Handle different functions
	if function == "readState" { //read a variable
		return t.readState(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query")
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	fmt.Println("running write()")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4. name of the key and value to set")
	}

	err = stub.PutState("person", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("time", []byte(args[1]))
	if err != nil {
		return nil, err
	}

	err = stub.PutState("position", []byte(args[2]))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("state", []byte(args[3]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) readPerson(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) readTime(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) readPosition(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) readState(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

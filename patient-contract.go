package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type contract struct {
	contractapi.Contract
}

//checks the world state to see if a student with a given ID exists
func (c *contract) StudentExists(ctx contractapi.TransactionContextInterface, StudentID string) (bool, error) {
	data, err := ctx.GetStub().GetState(StudentID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// Student registers a university with Student ID and Uni ID
func (c *contract) StuRegUni(ctx contractapi.TransactionContextInterface, StudentID string, UniID string) error {
	asset := StudentRegUni{
		StudentID: StudentID,
		UniID:     UniID,
	}

	dataJSON, err := json.Marshal(asset)

	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(StudentID, dataJSON)
}

//Read the information on a given student via Student ID (also checks if student exists first)
func (c *contract) ReadStudentInfo(ctx contractapi.TransactionContextInterface, StudentID string) (*StudentRegUni, error) {
	exists, err := c.StudentExists(ctx, StudentID)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("Student %s does not exist", StudentID)
	}

	data, _ := ctx.GetStub().GetState(StudentID)

	StudentInfo := new(StudentRegUni)

	err = json.Unmarshal(data, StudentInfo)

	if err != nil {
		return nil, fmt.Errorf("Coule not unmarshal world state data to type StudentRegUni")
	}
	return StudentInfo, nil
}

// Deletes a student via student ID (also checks if student exists first)
func (c *contract) Delete(ctx contractapi.TransactionContextInterface, StudentID string) error {
	exists, err := c.StudentExists(ctx, StudentID)
	if err != nil {
		return fmt.Errorf("Failed to read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("Student %s does not exist", StudentID)
	}
	return ctx.GetStub().DelState(StudentID)
}

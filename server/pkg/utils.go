package faceit_cc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

// GetSetting checks for environment variables in system
func GetSetting(setting string) (string, error) {
	value, ok := os.LookupEnv(setting)
	if !ok {
		err := fmt.Errorf("init error, %s ENV var not found", setting)
		return "", err
	}
	return value, nil
}

// ValidateUuid validates uuid format
func ValidateUuid(input string) bool {
	_, err := uuid.Parse(input)
	return err == nil
}

// PrettyPrint is a helper function to print structs
func PrettyPrint(input interface{}) {
	empJSON, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("PrettyPrint output \n %s\n", string(empJSON))
}

func (p *Pagination) SetPagination() {
	intResults := 10
	if p.ResultsPage == 0 {
		p.ResultsPage = intResults
	}

	if p.SearchBy == "" || p.SearchValue == "" {
		p.SearchBy = "1"
		p.SearchValue = "1"
	}
}

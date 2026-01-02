package api

import (
	"encoding/json"
	"os"
)

type PluginInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type ScanReport struct {
	PluginID string        `json:"plugin_id"`
	Checks   []CheckResult `json:"checks"`
}

type CheckResult struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Passed bool `json:"passed"`

	Score    int `json:"score"`
	MaxScore int `json:"max_score"`

	Remediation string `json:"remediation"`
	DocsURL     string `json:"docs_url"`

	Error string `json:"error,omitempty"`
}

type GrandReport struct {
	Timestamp     string        `json:"timestamp"`
	TotalScore    int           `json:"total_score"`
	MaxScore      int           `json:"max_score"`
	Rank          string        `json:"rank"`
	PluginReports []*ScanReport `json:"details"`
}

func PrintReport(report ScanReport) error {
	encoder := json.NewEncoder(os.Stdout)
	return encoder.Encode(report)
}

package voiceflow

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/analytics"
)

func FetchAnalytics(agentID string, startTime time.Time, endTime time.Time, limit int, analyticsToFetch []string) (string, error) {
	if global.VoiceflowSubdomain != "" {
		global.VoiceflowSubdomain = "." + global.VoiceflowSubdomain
	}
	url := fmt.Sprintf("https://analytics-api%s.voiceflow.com/v1/query/usage", global.VoiceflowSubdomain)
	analyticsRequest := analytics.Query{}
	for _, analytic := range analyticsToFetch {
		analyticsRequest.Query = append(analyticsRequest.Query, analytics.QueryItem{
			Name: analytic,
			Filter: analytics.Filter{
				ProjectID: agentID,
				StartTime: startTime,
				EndTime:   endTime,
				Limit:     limit,
			},
		})
	}

	byts, err := json.Marshal(analyticsRequest)
	if err != nil {
		return "", fmt.Errorf("error marshalling request: %v", err)
	}

	payload := strings.NewReader(string(byts))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", global.VoiceflowAPIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func SaveAnalytics(analytics string, outputFile string) error {

	// Write agent to file
	err := os.WriteFile(outputFile, []byte(analytics), 0644)
	if err != nil {
		return fmt.Errorf("failed to write analytcs file: %w", err)
	}

	return nil
}

func ParseFilters(startTime, endTime string, limit int) (time.Time, time.Time, int, error) {
	if startTime == "" {
		startTime = time.Now().AddDate(0, 0, -30).Format(global.FilterDateFormat)
	}
	if endTime == "" {
		endTime = time.Now().Format(global.FilterDateFormat)
	}

	startTimeDate, err := time.Parse(global.FilterDateFormat, startTime)

	if err != nil {
		return time.Time{}, time.Time{}, 0, fmt.Errorf("invalid start time: %w", err)
	}

	endTimeDate, err := time.Parse(global.FilterDateFormat, endTime)

	if err != nil {
		return time.Time{}, time.Time{}, 0, fmt.Errorf("invalid end time: %w", err)
	}

	if limit <= 0 {
		return time.Time{}, time.Time{}, 0, fmt.Errorf("invalid limit: %w", err)
	}

	return startTimeDate, endTimeDate, limit, nil

}
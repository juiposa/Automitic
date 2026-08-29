package fflogs

import (
	"context"

	"github.com/gqlgo/gqlgenc/clientv2"
	"github.com/juiposa/Automitic/internal/config"
	"github.com/juiposa/Automitic/internal/fflogs/query"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type GQL interface {
	GetReportMetadata(reportId string) (*ReportMasterData, error)
	GetReportFights(reportId string) ([]*ReportFight, error)
	GetFightPlayers(reportId string, fightNumber int) (any, error) //TODO player object
	GetFightEvents(reportId string, fightNumber int, startTime float64, endTime float64) ([]Event, error)
}

type fflogsClient struct {
	v2 *clientv2.Client
}

func NewClient(cli clientv2.HttpClient, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) GQL {
	return &fflogsClient{
		v2: clientv2.NewClient(cli, baseURL, options, interceptors...),
	}
}

func NewFFLogsClient(cfg *config.Config) GQL {
	ctx := context.Background()
	clientCfg := clientcredentials.Config{
		ClientID:     cfg.FFLogsClientID,
		ClientSecret: cfg.FFLogsClientSecret,
		TokenURL:     cfg.FFLogsAccessTokenURL,
	}
	client := oauth2.NewClient(ctx, clientCfg.TokenSource(ctx))
	return NewClient(client, cfg.FFLogsApiEndpoint, &clientv2.Options{})
}

type ResponseBody struct {
	ReportData ReportData `json:"reportData,omitempty"`
}

// TODO context
func (f *fflogsClient) sendQuery(opName string, queryString string, variables map[string]any) (*ReportData, error) {
	var responseBody ResponseBody
	err := f.v2.Post(context.Background(), opName, queryString, &responseBody, variables)
	if err != nil {
		return nil, err
	}
	return &responseBody.ReportData, nil
}

func (f *fflogsClient) GetReportMetadata(reportId string) (*ReportMasterData, error) {
	vars := map[string]any{
		"reportId": reportId,
	}
	response, err := f.sendQuery("", query.GetMetadataQuery, vars)
	if err != nil {
		return nil, err
	}
	return response.Report.MasterData, nil
}

func (f *fflogsClient) GetReportFights(reportId string) ([]*ReportFight, error) {
	vars := map[string]any{
		"reportId": reportId,
	}
	response, err := f.sendQuery("", query.GetFightsQuery, vars)
	if err != nil {
		return nil, err
	}
	return response.Report.Fights, nil
}

func (f *fflogsClient) GetFightPlayers(reportId string, fightNumber int) (any, error) { //TODO player object
	vars := map[string]any{
		"reportId":    reportId,
		"fightNumber": fightNumber,
	}
	response, err := f.sendQuery("", query.GetFightPlayersQuery, vars)
	if err != nil {
		return nil, err
	}
	return response.Report.PlayerDetails, nil
}

func (f *fflogsClient) GetFightEvents(reportId string, fightNumber int, startTime float64, endTime float64) ([]Event, error) {
	vars := map[string]any{
		"reportId":    reportId,
		"fightNumber": fightNumber,
		"startTime":   startTime,
		"endTime":     endTime,
	}
	var fullEventsList []Event
	for {
		response, err := f.sendQuery("", query.GetPaginatedFightEvents, vars)
		if err != nil {
			return nil, err
		}
		paginator := response.Report.Events
		fullEventsList = append(fullEventsList, paginator.Data...)
		if paginator.NextPageTimestamp == nil {
			break
		}
		vars["startTime"] = paginator.NextPageTimestamp
	}
	return fullEventsList, nil
}

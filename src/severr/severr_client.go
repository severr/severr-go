package severr

import (
	"severr_client"
	"os"
	"runtime"
	"time"
	"fmt"
)

type SeverrClient struct {
	apiKey                  string
	url                     string
	contextAppVersion       string
	contextEnvName          string
	contextEnvVersion       string
	contextEnvHostname      string
	contextAppOS            string
	contextAppOSVersion     string
	contextDataCenter       string
	contextDataCenterRegion string
	eventsApi               severr_client.EventsApi
	eventTraceBuilder       EventTraceBuilder
}

// Create a new SeverrClient and return it with the data.
//
// Most parameters are optional i.e. empty (pass "" to use defaults) with the exception of apiKey which is required.
// url is the location of the serverr service, if "" is passed it defaults to https://severr.io/api/v1
func NewSeverrClientWithDefaults(
	apiKey string,
	url string,
	contextAppVersion string,
	contextEnvName string) *SeverrClient {
	return NewSeverrClient(apiKey, url, contextAppVersion, contextEnvName, "", "", "", "", "", "")
}

// Create a new SeverrClient and return it with the data.
//
// Most parameters are optional i.e. empty (pass "" to use defaults) with the exception of apiKey which is required.
// url is the location of the serverr service, if "" is passed it defaults to https://severr.io/api/v1
func NewSeverrClient(
	apiKey string,
	url string,
	contextAppVersion string,
	contextEnvName string,
	contextEnvVersion string,
	contextEnvHostname string,
	contextAppOS string,
	contextAppOSVersion string,
	contextDataCenter string,
	contextDataCenterRegion string) *SeverrClient {

	if contextEnvName == "" { contextEnvName = "development" }
	if contextAppVersion == "" { contextAppVersion = "1.0" }
	if contextEnvHostname == "" { contextEnvHostname, _ = os.Hostname() }

	if contextAppOS == "" {
		contextAppOS = runtime.GOOS
		contextAppOSVersion = "N/A (arch:" + runtime.GOARCH + ")"
	}
	var eventsApi severr_client.EventsApi

	if url != "" {
		eventsApi = *severr_client.NewEventsApiWithBasePath(url);
	} else {
		eventsApi = *severr_client.NewEventsApi()
	}

	return &SeverrClient{
		apiKey: apiKey,
		url: url,
		contextAppVersion: contextAppVersion,
		contextEnvName: contextEnvName,
		contextEnvVersion: contextEnvVersion,
		contextEnvHostname: contextEnvHostname,
		contextAppOS: contextAppOS,
		contextAppOSVersion: contextAppOSVersion,
		contextDataCenter: contextDataCenter,
		contextDataCenterRegion: contextDataCenterRegion,
		eventsApi: eventsApi,
		eventTraceBuilder: EventTraceBuilder{} }
}

func (severrClient *SeverrClient) NewAppEvent(classification string, eventType string, eventMessage string) *severr_client.AppEvent {
	if classification == "" { classification = "Error" }
	if eventType == "" { eventType = "unknown" }
	if eventMessage == "" { eventMessage = "unknown "}
	return severrClient.FillDefaults(&severr_client.AppEvent{Classification: classification, EventType:eventType, EventMessage: eventMessage })
}

func (severrClient *SeverrClient) SendEvent(appEvent *severr_client.AppEvent) (*severr_client.APIResponse, error) {
	return severrClient.eventsApi.EventsPost(*severrClient.FillDefaults(appEvent))
}

func (severrClient *SeverrClient) SendError(err interface{}) (*severr_client.APIResponse, error) {
	appEvent := severrClient.CreateAppEventFromError(err)

	return severrClient.eventsApi.EventsPost(*appEvent)
}

func (severrClient *SeverrClient) CreateAppEventFromError(err interface{}) *severr_client.AppEvent {
	stacktrace := severrClient.eventTraceBuilder.GetEventTraces(err, 4)
	event := severr_client.AppEvent{}
	event.EventType = fmt.Sprintf("%T", err)
	event.EventMessage = fmt.Sprint(err)
	event.Classification = "Error"

	result := severrClient.FillDefaults(&event)
	event.EventStacktrace = stacktrace
	return result
}

func (severrClient *SeverrClient) FillDefaults(appEvent *severr_client.AppEvent) *severr_client.AppEvent {
	if appEvent.ApiKey == "" {
		appEvent.ApiKey = severrClient.apiKey
	}

	if (appEvent.ContextAppVersion == "") {
		appEvent.ContextAppVersion = severrClient.contextAppVersion
	}

	if (appEvent.ContextEnvName == "") {
		appEvent.ContextEnvName = severrClient.contextEnvName
	}
	if (appEvent.ContextEnvVersion == "") {
		appEvent.ContextEnvVersion = severrClient.contextEnvVersion
	}
	if (appEvent.ContextEnvHostname == "") {
		appEvent.ContextEnvHostname = severrClient.contextEnvHostname
	}

	if (appEvent.ContextAppOS == "") {
		appEvent.ContextAppOS = severrClient.contextAppOS
		appEvent.ContextAppOSVersion = severrClient.contextAppOSVersion
	}

	if (appEvent.ContextDataCenter == "") {
		appEvent.ContextDataCenter = severrClient.contextDataCenter
	}
	if (appEvent.ContextDataCenterRegion == "") {
		appEvent.ContextDataCenterRegion = severrClient.contextDataCenterRegion
	}

	if (appEvent.EventTime <= 0) {
		appEvent.EventTime = makeTimestamp()
	}
	return appEvent
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}




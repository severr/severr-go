# Go API client for severr_client

Get your application events and errors to Severr via the *Severr API*.

## Overview
- API version: 1.0.0
- Package version: 1.0.0

## Installation
Put the packages under your project folder and add the following in import:
```
    "severr"
```

## Getting Started

There are a few options (illustrated below with comment Option-#) to send events to Severr. The easiest of 
which is to send only errors to Severr (Option-1).

```$golang
package main

import (
	"severr"
	"errors"
)

func main() {
	client := severr.NewSeverrClientWithDefaults(
		"ceba200baf79b1b5e9dc73d4054d6c9618388477122",
		"http://192.168.0.117:3000/api/v1",
		"1.0",
		"development")
	err := errors.New("Something bad happened here")

	// Option-1: send error
	client.SendError(err)

	// Option-2: send error with custom properties
	appEventWithErr := client.CreateAppEventFromError(err)

	// set any custom data on appEvent
	appEventWithErr.CustomProperties.StringData.CustomData1 = "foo"
	appEventWithErr.CustomProperties.StringData.CustomData2 = "bar"

	client.SendEvent(appEventWithErr)

	// Option-3: send event manually
	appEvent := client.NewAppEvent("Info", "SomeType", "SomeMessage")

	// set any custom data on appEvent
	appEvent.CustomProperties.StringData.CustomData1 = "foo"
	appEvent.CustomProperties.StringData.CustomData2 = "bar"

	client.SendEvent(appEvent)
}
```

## Documentation For Models

 - [AppEvent](src/severr_client/docs/AppEvent.md)




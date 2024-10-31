// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Defines values for ProcessInstanceState.
const (
	Active     ProcessInstanceState = "active"
	Completed  ProcessInstanceState = "completed"
	Terminated ProcessInstanceState = "terminated"
)

// Activity defines model for Activity.
type Activity struct {
	BpmnElementType      *string    `json:"bpmnElementType,omitempty"`
	CreatedAt            *time.Time `json:"createdAt,omitempty"`
	ElementId            *string    `json:"elementId,omitempty"`
	Key                  *string    `json:"key,omitempty"`
	ProcessDefinitionKey *string    `json:"processDefinitionKey,omitempty"`
	ProcessInstanceKey   *string    `json:"processInstanceKey,omitempty"`
	State                *string    `json:"state,omitempty"`
}

// ActivityPage defines model for ActivityPage.
type ActivityPage struct {
	Count  *int        `json:"count,omitempty"`
	Items  *[]Activity `json:"items,omitempty"`
	Offset *int        `json:"offset,omitempty"`
	Size   *int        `json:"size,omitempty"`
}

// ClusterInfo defines model for ClusterInfo.
type ClusterInfo struct {
	Partitions *[]ClusterPartition `json:"partitions,omitempty"`
}

// ClusterPartition defines model for ClusterPartition.
type ClusterPartition struct {
	Id      *string   `json:"id,omitempty"`
	Leader  *string   `json:"leader,omitempty"`
	Members *[]string `json:"members,omitempty"`
}

// Job defines model for Job.
type Job struct {
	CreatedAt          *time.Time `json:"createdAt,omitempty"`
	ElementId          *string    `json:"elementId,omitempty"`
	ElementInstanceKey *string    `json:"elementInstanceKey,omitempty"`
	Key                *string    `json:"key,omitempty"`
	ProcessInstanceKey *string    `json:"processInstanceKey,omitempty"`
	State              *string    `json:"state,omitempty"`
}

// JobPage defines model for JobPage.
type JobPage struct {
	Count  *int   `json:"count,omitempty"`
	Items  *[]Job `json:"items,omitempty"`
	Offset *int   `json:"offset,omitempty"`
	Size   *int   `json:"size,omitempty"`
}

// PageMetadata defines model for PageMetadata.
type PageMetadata struct {
	Count  *int `json:"count,omitempty"`
	Offset *int `json:"offset,omitempty"`
	Size   *int `json:"size,omitempty"`
}

// ProcessDefinitionDetail defines model for ProcessDefinitionDetail.
type ProcessDefinitionDetail struct {
	BpmnData      *string `json:"bpmnData,omitempty"`
	BpmnProcessId *string `json:"bpmnProcessId,omitempty"`
	Key           *string `json:"key,omitempty"`
	Version       *int    `json:"version,omitempty"`
}

// ProcessDefinitionSimple defines model for ProcessDefinitionSimple.
type ProcessDefinitionSimple struct {
	BpmnProcessId *string `json:"bpmnProcessId,omitempty"`
	Key           *string `json:"key,omitempty"`
	Version       *int    `json:"version,omitempty"`
}

// ProcessDefinitionsPage defines model for ProcessDefinitionsPage.
type ProcessDefinitionsPage struct {
	Count  *int                       `json:"count,omitempty"`
	Items  *[]ProcessDefinitionSimple `json:"items,omitempty"`
	Offset *int                       `json:"offset,omitempty"`
	Size   *int                       `json:"size,omitempty"`
}

// ProcessInstance defines model for ProcessInstance.
type ProcessInstance struct {
	Activities           *string               `json:"activities,omitempty"`
	CaughtEvents         *string               `json:"caughtEvents,omitempty"`
	CreatedAt            *time.Time            `json:"createdAt,omitempty"`
	Key                  *string               `json:"key,omitempty"`
	ProcessDefinitionKey *string               `json:"processDefinitionKey,omitempty"`
	State                *ProcessInstanceState `json:"state,omitempty"`
	VariableHolder       *string               `json:"variableHolder,omitempty"`
}

// ProcessInstanceState defines model for ProcessInstance.State.
type ProcessInstanceState string

// ProcessInstancePage defines model for ProcessInstancePage.
type ProcessInstancePage struct {
	Count  *int               `json:"count,omitempty"`
	Items  *[]ProcessInstance `json:"items,omitempty"`
	Offset *int               `json:"offset,omitempty"`
	Size   *int               `json:"size,omitempty"`
}

// CompleteJobJSONBody defines parameters for CompleteJob.
type CompleteJobJSONBody struct {
	JobKey string `json:"jobKey"`
}

// CreateProcessInstanceJSONBody defines parameters for CreateProcessInstance.
type CreateProcessInstanceJSONBody struct {
	ProcessDefinitionKey string                  `json:"processDefinitionKey"`
	Variables            *map[string]interface{} `json:"variables,omitempty"`
}

// GetProcessInstancesParams defines parameters for GetProcessInstances.
type GetProcessInstancesParams struct {
	ProcessDefinitionKey *int64 `form:"processDefinitionKey,omitempty" json:"processDefinitionKey,omitempty"`
	Offset               *int   `form:"offset,omitempty" json:"offset,omitempty"`
	Size                 *int   `form:"size,omitempty" json:"size,omitempty"`
}

// CompleteJobJSONRequestBody defines body for CompleteJob for application/json ContentType.
type CompleteJobJSONRequestBody CompleteJobJSONBody

// CreateProcessInstanceJSONRequestBody defines body for CreateProcessInstance for application/json ContentType.
type CreateProcessInstanceJSONRequestBody CreateProcessInstanceJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get cluster information
	// (GET /cluster)
	GetClusterInfo(ctx echo.Context) error
	// Rebalance cluster
	// (POST /cluster/rebalance)
	Rebalance(ctx echo.Context) error
	// Complete a job
	// (POST /jobs)
	CompleteJob(ctx echo.Context) error
	// Get list of process definitions
	// (GET /process-definitions)
	GetProcessDefinitions(ctx echo.Context) error
	// Deploy a new process definition
	// (POST /process-definitions)
	CreateProcessDefinition(ctx echo.Context) error
	// Get process definition
	// (GET /process-definitions/{processDefinitionKey})
	GetProcessDefinition(ctx echo.Context, processDefinitionKey int64) error
	// Create a new process instance
	// (POST /process-instances)
	CreateProcessInstance(ctx echo.Context) error
	// Get list of running process instances
	// (GET /process-instances/)
	GetProcessInstances(ctx echo.Context, params GetProcessInstancesParams) error
	// Get state of a process instance selected by processInstanceId
	// (GET /process-instances/{processInstanceKey})
	GetProcessInstance(ctx echo.Context, processInstanceKey int64) error
	// Get list of activities for a process instance
	// (GET /process-instances/{processInstanceKey}/activities)
	GetActivities(ctx echo.Context, processInstanceKey int64) error
	// Get list of jobs for a process instance
	// (GET /process-instances/{processInstanceKey}/jobs)
	GetJobs(ctx echo.Context, processInstanceKey int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetClusterInfo converts echo context to params.
func (w *ServerInterfaceWrapper) GetClusterInfo(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetClusterInfo(ctx)
	return err
}

// Rebalance converts echo context to params.
func (w *ServerInterfaceWrapper) Rebalance(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Rebalance(ctx)
	return err
}

// CompleteJob converts echo context to params.
func (w *ServerInterfaceWrapper) CompleteJob(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CompleteJob(ctx)
	return err
}

// GetProcessDefinitions converts echo context to params.
func (w *ServerInterfaceWrapper) GetProcessDefinitions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProcessDefinitions(ctx)
	return err
}

// CreateProcessDefinition converts echo context to params.
func (w *ServerInterfaceWrapper) CreateProcessDefinition(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateProcessDefinition(ctx)
	return err
}

// GetProcessDefinition converts echo context to params.
func (w *ServerInterfaceWrapper) GetProcessDefinition(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "processDefinitionKey" -------------
	var processDefinitionKey int64

	err = runtime.BindStyledParameterWithOptions("simple", "processDefinitionKey", ctx.Param("processDefinitionKey"), &processDefinitionKey, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter processDefinitionKey: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProcessDefinition(ctx, processDefinitionKey)
	return err
}

// CreateProcessInstance converts echo context to params.
func (w *ServerInterfaceWrapper) CreateProcessInstance(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateProcessInstance(ctx)
	return err
}

// GetProcessInstances converts echo context to params.
func (w *ServerInterfaceWrapper) GetProcessInstances(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProcessInstancesParams
	// ------------- Optional query parameter "processDefinitionKey" -------------

	err = runtime.BindQueryParameter("form", true, false, "processDefinitionKey", ctx.QueryParams(), &params.ProcessDefinitionKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter processDefinitionKey: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, false, "size", ctx.QueryParams(), &params.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter size: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProcessInstances(ctx, params)
	return err
}

// GetProcessInstance converts echo context to params.
func (w *ServerInterfaceWrapper) GetProcessInstance(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "processInstanceKey" -------------
	var processInstanceKey int64

	err = runtime.BindStyledParameterWithOptions("simple", "processInstanceKey", ctx.Param("processInstanceKey"), &processInstanceKey, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter processInstanceKey: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProcessInstance(ctx, processInstanceKey)
	return err
}

// GetActivities converts echo context to params.
func (w *ServerInterfaceWrapper) GetActivities(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "processInstanceKey" -------------
	var processInstanceKey int64

	err = runtime.BindStyledParameterWithOptions("simple", "processInstanceKey", ctx.Param("processInstanceKey"), &processInstanceKey, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter processInstanceKey: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetActivities(ctx, processInstanceKey)
	return err
}

// GetJobs converts echo context to params.
func (w *ServerInterfaceWrapper) GetJobs(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "processInstanceKey" -------------
	var processInstanceKey int64

	err = runtime.BindStyledParameterWithOptions("simple", "processInstanceKey", ctx.Param("processInstanceKey"), &processInstanceKey, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter processInstanceKey: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetJobs(ctx, processInstanceKey)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/cluster", wrapper.GetClusterInfo)
	router.POST(baseURL+"/cluster/rebalance", wrapper.Rebalance)
	router.POST(baseURL+"/jobs", wrapper.CompleteJob)
	router.GET(baseURL+"/process-definitions", wrapper.GetProcessDefinitions)
	router.POST(baseURL+"/process-definitions", wrapper.CreateProcessDefinition)
	router.GET(baseURL+"/process-definitions/:processDefinitionKey", wrapper.GetProcessDefinition)
	router.POST(baseURL+"/process-instances", wrapper.CreateProcessInstance)
	router.GET(baseURL+"/process-instances/", wrapper.GetProcessInstances)
	router.GET(baseURL+"/process-instances/:processInstanceKey", wrapper.GetProcessInstance)
	router.GET(baseURL+"/process-instances/:processInstanceKey/activities", wrapper.GetActivities)
	router.GET(baseURL+"/process-instances/:processInstanceKey/jobs", wrapper.GetJobs)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xYTXPbNhD9Kxy0R9mU206n1c2JPa3cptEkuWV0AMmVDAcEGGDpVtXwv3cAfhOgRNmy",
	"p5OTZQJY7r739gPck1immRQgUJPFnuj4HlJqf17HyB4Z7szvTMkMFDKwK1GWilsOKQj8tMvAPEL7l2hU",
	"TGxJMSOxAoqQXKNZ3UiVUiQLklCEC2QpkJl7BEqTy8Rr8AvsvM8zJWPQ+gY2TDBkUvxxeONSaKQihrFt",
	"Gin6Qioaj2X0ADGavTVEK7q1Ryjn7zdk8XkIGENI+z++V7AhC/Jd2MIfVtiHDfDtK6lSdOf14bAp49g7",
	"QJpQpKRYe0J4y3ONoJZiI12iM6rQYjrd+creqj45KQjinHJcYX5VcKAJKO9SCmkEqu+5s+m4a3cycr05",
	"u7rr1SPi/PJ62r6T0XllbYB8DUX3drjMyVxgBwImELagzEG52WgYWdPsX/Ct+JBbDSvSDSBlvI/kwRiH",
	"Bj6yNONgsXEr8U0V51FO11N8rV7lrfnV5hML9CMoXaX0k9DT55XhOLivIM1+mroo07L0V/+5TZXm23u8",
	"fay79Tm67rP7alNUQOQpWXwugzDvMhhxQEjMe0GlTBjnyNrjxSNVjEYcfpfcX8+L43C+iE4arl5eH8Yi",
	"q/pwAjpWLCt7Iflw+/FTcL1aBhupAs6iC5ONFyC2TFhOGZqUJW9W7/4Kbu3T4H0G4nq1JJ38I/PLq8u5",
	"rXQZCJoxsiA/Xs4v52RGMor3FpIwLlux+b0tq6HBkBpPTOKT3wC7M8OMKNCZFLrE+If5vKyyAqGsszTL",
	"OIvt+fBBl3WgRGTiKGFfY8Hpg1ItBwYyI3Q7bBg95mlK1a50NYg9u2YE6VYbqdbBrs3BOvRQQUR5k6BS",
	"e1D40GzxA+B3tTGcDDxtzNX+jvr4ICM97tbbKuNMrzWOfc1B4xuZ7E4ipZ82DzLyp35RvoEpSIyf1T6/",
	"rtuNqHIoHNCuXNDuZBS0FaSPVx1nQIMHG2oNloWnRKoqYBdJ20kOidrtOy+p7ZEu55H5n0xjIDdBFU7Q",
	"DceVOz+wu0XJB83a1Hy/qGxTcTyeLLB/Ut4HpmlMZsVpBlPkMn+Gnic2tsKr4z43KwflIIGMy50j2Bv7",
	"OKCBgL895BzlZkTQ4d4XTXGSzm31VzQFtLelz3vCTGymI5AZETQ1GHhBG/I083HMBP78U8tyO/utXzO9",
	"qhF8EoeetHomY6waIg5V7m6SNUPHuWr45GGunsO642Wj/36999p8WvU/O/Ht1DZOeE1KUA3NwxZjnw4y",
	"lrXEDNlvOR7hPpyQlcvGiD8pv+agdkez8pQsnPktV3fhrq0ENjTnSBbz6Wbstdlr5Gr+AiXBK/plN/me",
	"MvqXnXk4/s8ISqR82rV2vK2rXAgmto7CDjX38TNPkeXe/XhUnCDVU9pH9/vU/755HKohH82d11BB3dLg",
	"8qbHdwcaOMQISRDtggFKy+RsfIb9Dwtj1F63u745Vnvf6g8kZAep8QxsN9krOT1Pg/AyV1/4xji7M+vf",
	"HFv11+cDRFlgxikyy08kx9gE9VhjmStOFuQeMVuEIZcx5fdS4+KX+a9X5Wec0tZ+AHRvHjQNcrDcqfTt",
	"Yn3ZL9bFfwEAAP//fY+RLZsbAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

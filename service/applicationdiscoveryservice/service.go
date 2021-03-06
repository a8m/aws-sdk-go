// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// This is the AWS Discovery Service API Reference. The AWS Discovery Service
// streamlines the process of migrating to Amazon Web Services by helping you
// identify assets in your data center, including servers, virtual machines,
// applications, application dependencies, and network infrastructure. You can
// use this information to find the workloads that make up an application, analyze
// dependencies, and build migration strategies. The service also collects performance
// data about your workloads which you can use to assess migration outcomes.
//
// This API reference provides descriptions, syntax, and usage examples for
// each of the actions and data types for the Discovery Service. The topic for
// each action shows the API request parameters and the response. Alternatively,
// you can use one of the AWS SDKs to access an API that's tailored to the programming
// language or platform that you're using. For more information, see AWS SDKs
// (http://aws.amazon.com/tools/#SDKs).
//
// This guide is intended for use with the AWS Discovery Service user guide
// (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/remote-commands-prereq.html).
//
// The following are short descriptions of each API action, organized by function.
//
//  Managing AWS Agents
//
// The AWS agent is an Amazon application that you install on servers and virtual
// machines in your data center or on Amazon EC2 instances. The agent captures
// server configuration and activity information (including hardware profile,
// network, file system, and process activity) and sends this data to the AWS
// Application Discovery Service. The Discovery Service processes this data
// and maps the application dependencies for your workloads.
//
//    StartDataCollectionByAgentIds: Instructs the specified agents to start
// collecting data. Agents can reside on host servers or virtual machines in
// your data center or on AWS EC2 instances.
//
//    StopDataCollectionByAgentIds: Instructs the specified agents to stop
// collecting data.
//
//    DescribeAgents: Lists AWS agents by ID or lists all agents associated
// with your user account if you did not specify an agent ID. The output includes
// agent IDs, IP addresses, MAC addresses, agent health, host name where the
// agent resides, and the version number of each agent.
//
//    Querying Configuration Items
//
// A configuration item is an IT asset that was discovered in your data center
// by an AWS agent. With the Discovery Service, you can specify filters and
// query specific configuration items. For example, using this API, you could
// create a filter to query for a process configuration item named apache and
// an operating system configuration item named Ubuntu.
//
//    GetConfigurationAttributes: Retrieves a list of attributes for a specific
// configuration ID. For example, the output for a server configuration item
// includes a list of attributes about the server, including host name, operating
// system, number of network cards, etc.
//
//    ListConfigurations: Retrieves a list of configurations items according
// to the criteria you specify in a filter. The filter criteria identify relationship
// requirements. For example, the following filter specifies criteria of process.name
// and values of nginx and apache.
//
//  ConfigurationType = Process Filters = [WebServerCriteria] WebServerCriteria
// = { ‘key’ : process.name, ‘values’ : [ ‘nginx’, ‘apache’ ], ‘condition’ :
// ‘contains’ }
//
//    Tagging Discovered Configuration Items
//
// You can tag discovered configuration items. Tags are metadata that help
// you categorize IT assets in your data center. Tags use a key,value format.
// For example, {"key": "serverType", "value": "webServer"}.
//
//    CreateTags: Creates one or more tags for a configuration item. Tags are
// metadata that help you categorize IT assets.
//
//    DescribeTags: Retrieve a list of configuration items that are tagged
// with a specific tag. Or retrieve a list all tags assigned to a specific configuration
// item.
//
//    DeleteTags: Deletes one or more tags associated with a configuration
// item.
//
//    Exporting Data
//
// You can export discovered data to an Amazon S3 bucket in the form of CSV
// files.
//
//    ExportConfigurations: Exports all discovered configuration data to an
// Amazon S3 bucket. Data includes processes, connections, servers, and system
// performance.
//
//    GetExportStatus: Gets the status of the data export. When the export
// is complete, the service returns an Amazon S3 URL where you can download
// CSV files that include the data.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type ApplicationDiscoveryService struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "discovery"

// New creates a new instance of the ApplicationDiscoveryService client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ApplicationDiscoveryService client from just a session.
//     svc := applicationdiscoveryservice.New(mySession)
//
//     // Create a ApplicationDiscoveryService client with additional configuration
//     svc := applicationdiscoveryservice.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ApplicationDiscoveryService {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *ApplicationDiscoveryService {
	svc := &ApplicationDiscoveryService{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-11-01",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSPoseidonService_V2015_11_01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ApplicationDiscoveryService operation and runs any
// custom request initialization.
func (c *ApplicationDiscoveryService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}

package service

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/gofrs/uuid"
	"github.com/redis/go-redis/v9"
	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.temporal.io/sdk/client"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/instill-ai/pipeline-backend/pkg/acl"
	"github.com/instill-ai/pipeline-backend/pkg/logger"
	"github.com/instill-ai/pipeline-backend/pkg/repository"
	"github.com/instill-ai/pipeline-backend/pkg/resource"

	componentstore "github.com/instill-ai/component/store"
	pb "github.com/instill-ai/protogen-go/vdp/pipeline/v1beta"
)

// Service interface
type Service interface {
	GetHubStats(ctx context.Context) (*pb.GetHubStatsResponse, error)
	ListPipelines(ctx context.Context, pageSize int32, pageToken string, view pb.Pipeline_View, visibility *pb.Pipeline_Visibility, filter filtering.Filter, showDeleted bool, order ordering.OrderBy) ([]*pb.Pipeline, int32, string, error)
	GetPipelineByUID(ctx context.Context, uid uuid.UUID, view pb.Pipeline_View) (*pb.Pipeline, error)
	CreateNamespacePipeline(ctx context.Context, ns resource.Namespace, pipeline *pb.Pipeline) (*pb.Pipeline, error)
	ListNamespacePipelines(ctx context.Context, ns resource.Namespace, pageSize int32, pageToken string, view pb.Pipeline_View, visibility *pb.Pipeline_Visibility, filter filtering.Filter, showDeleted bool, order ordering.OrderBy) ([]*pb.Pipeline, int32, string, error)
	GetNamespacePipelineByID(ctx context.Context, ns resource.Namespace, id string, view pb.Pipeline_View) (*pb.Pipeline, error)
	UpdateNamespacePipelineByID(ctx context.Context, ns resource.Namespace, id string, updatedPipeline *pb.Pipeline) (*pb.Pipeline, error)
	UpdateNamespacePipelineIDByID(ctx context.Context, ns resource.Namespace, id string, newID string) (*pb.Pipeline, error)
	DeleteNamespacePipelineByID(ctx context.Context, ns resource.Namespace, id string) error
	ValidateNamespacePipelineByID(ctx context.Context, ns resource.Namespace, id string) ([]*pb.PipelineValidationError, error)
	GetNamespacePipelineLatestReleaseUID(ctx context.Context, ns resource.Namespace, id string) (uuid.UUID, error)
	CloneNamespacePipeline(ctx context.Context, ns resource.Namespace, id string, targetNS resource.Namespace, targetID string) (*pb.Pipeline, error)

	ListPipelinesAdmin(ctx context.Context, pageSize int32, pageToken string, view pb.Pipeline_View, filter filtering.Filter, showDeleted bool) ([]*pb.Pipeline, int32, string, error)
	GetPipelineByUIDAdmin(ctx context.Context, uid uuid.UUID, view pb.Pipeline_View) (*pb.Pipeline, error)

	CreateNamespacePipelineRelease(ctx context.Context, ns resource.Namespace, pipelineUID uuid.UUID, pipelineRelease *pb.PipelineRelease) (*pb.PipelineRelease, error)
	ListNamespacePipelineReleases(ctx context.Context, ns resource.Namespace, pipelineUID uuid.UUID, pageSize int32, pageToken string, view pb.Pipeline_View, filter filtering.Filter, showDeleted bool) ([]*pb.PipelineRelease, int32, string, error)
	GetNamespacePipelineReleaseByID(ctx context.Context, ns resource.Namespace, pipelineUID uuid.UUID, id string, view pb.Pipeline_View) (*pb.PipelineRelease, error)
	UpdateNamespacePipelineReleaseByID(ctx context.Context, ns resource.Namespace, pipelineUID uuid.UUID, id string, updatedPipelineRelease *pb.PipelineRelease) (*pb.PipelineRelease, error)
	DeleteNamespacePipelineReleaseByID(ctx context.Context, ns resource.Namespace, pipelineUID uuid.UUID, id string) error
	RestoreNamespacePipelineReleaseByID(ctx context.Context, ns resource.Namespace, pipelineUID uuid.UUID, id string) error
	UpdateNamespacePipelineReleaseIDByID(ctx context.Context, ns resource.Namespace, pipelineUID uuid.UUID, id string, newID string) (*pb.PipelineRelease, error)

	CreateNamespaceSecret(ctx context.Context, ns resource.Namespace, secret *pb.Secret) (*pb.Secret, error)
	ListNamespaceSecrets(ctx context.Context, ns resource.Namespace, pageSize int32, pageToken string, filter filtering.Filter) ([]*pb.Secret, int32, string, error)
	GetNamespaceSecretByID(ctx context.Context, ns resource.Namespace, id string) (*pb.Secret, error)
	UpdateNamespaceSecretByID(ctx context.Context, ns resource.Namespace, id string, updatedSecret *pb.Secret) (*pb.Secret, error)
	DeleteNamespaceSecretByID(ctx context.Context, ns resource.Namespace, id string) error

	TriggerNamespacePipelineByID(ctx context.Context, ns resource.Namespace, id string, data []*pb.TriggerData, pipelineTriggerID string, returnTraces bool) ([]*structpb.Struct, *pb.TriggerMetadata, error)
	TriggerAsyncNamespacePipelineByID(ctx context.Context, ns resource.Namespace, id string, data []*pb.TriggerData, pipelineTriggerID string, returnTraces bool) (*longrunningpb.Operation, error)

	TriggerNamespacePipelineReleaseByID(ctx context.Context, ns resource.Namespace, pipelineUID uuid.UUID, id string, data []*pb.TriggerData, pipelineTriggerID string, returnTraces bool) ([]*structpb.Struct, *pb.TriggerMetadata, error)
	TriggerAsyncNamespacePipelineReleaseByID(ctx context.Context, ns resource.Namespace, pipelineUID uuid.UUID, id string, data []*pb.TriggerData, pipelineTriggerID string, returnTraces bool) (*longrunningpb.Operation, error)
	GetOperation(ctx context.Context, workflowID string) (*longrunningpb.Operation, error)

	GetCtxUserNamespace(ctx context.Context) (resource.Namespace, error)
	GetRscNamespaceAndNameID(ctx context.Context, path string) (resource.Namespace, string, error)
	GetRscNamespaceAndPermalinkUID(ctx context.Context, path string) (resource.Namespace, uuid.UUID, error)
	GetRscNamespaceAndNameIDAndReleaseID(ctx context.Context, path string) (resource.Namespace, string, string, error)

	ListComponentDefinitions(context.Context, *pb.ListComponentDefinitionsRequest) (*pb.ListComponentDefinitionsResponse, error)
	ListOperatorDefinitions(context.Context, *pb.ListOperatorDefinitionsRequest) (*pb.ListOperatorDefinitionsResponse, error)
	GetOperatorDefinitionByID(ctx context.Context, defID string) (*pb.OperatorDefinition, error)
	ListConnectorDefinitions(context.Context, *pb.ListConnectorDefinitionsRequest) (*pb.ListConnectorDefinitionsResponse, error)
	GetConnectorDefinitionByID(ctx context.Context, id string) (*pb.ConnectorDefinition, error)
}

type service struct {
	repository     repository.Repository
	redisClient    *redis.Client
	temporalClient client.Client
	component      *componentstore.Store
	aclClient      *acl.ACLClient
	converter      Converter
}

// NewService initiates a service instance
func NewService(
	r repository.Repository,
	rc *redis.Client,
	t client.Client,
	acl *acl.ACLClient,
	c Converter,
) Service {
	logger, _ := logger.GetZapLogger(context.Background())

	return &service{
		repository:     r,
		redisClient:    rc,
		temporalClient: t,
		component:      componentstore.Init(logger, nil, nil),
		aclClient:      acl,
		converter:      c,
	}
}

module github.com/leonidweinberg/clouddumper/dumper

go 1.18

require (
	github.com/aws/aws-sdk-go-v2 v1.17.0
	github.com/aws/aws-sdk-go-v2/config v1.17.9
	github.com/aws/aws-sdk-go-v2/service/acm v1.15.1
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.18.2
	github.com/aws/aws-sdk-go-v2/service/amp v1.15.6
	github.com/aws/aws-sdk-go-v2/service/appflow v1.18.1
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.15.19
	github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice v1.14.10
	github.com/aws/aws-sdk-go-v2/service/applicationinsights v1.16.13
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.16.3
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.13.1
	github.com/aws/aws-sdk-go-v2/service/appstream v1.17.12
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.23.17
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.12.18
	github.com/aws/aws-sdk-go-v2/service/backup v1.17.10
	github.com/aws/aws-sdk-go-v2/service/batch v1.18.17
	github.com/aws/aws-sdk-go-v2/service/budgets v1.13.12
	github.com/aws/aws-sdk-go-v2/service/chime v1.21.10
	github.com/aws/aws-sdk-go-v2/service/chimesdkidentity v1.9.19
	github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging v1.11.7
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.16.18
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.22.11
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.20.6
	github.com/aws/aws-sdk-go-v2/service/cloudhsm v1.12.19
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.13.20
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.13.18
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.19.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.21.7
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.14.20
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.15.21
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.13.10
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.19.18
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.13.18
	github.com/aws/aws-sdk-go-v2/service/codeguruprofiler v1.12.18
	github.com/aws/aws-sdk-go-v2/service/codegurureviewer v1.16.16
	github.com/aws/aws-sdk-go-v2/service/codestar v1.12.2
	github.com/aws/aws-sdk-go-v2/service/codestarnotifications v1.13.3
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.14.2
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.21.0
	github.com/aws/aws-sdk-go-v2/service/cognitosync v1.11.18
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.19.1
	github.com/aws/aws-sdk-go-v2/service/comprehendmedical v1.13.18
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.17.20
	github.com/aws/aws-sdk-go-v2/service/configservice v1.27.1
	github.com/aws/aws-sdk-go-v2/service/connect v1.33.1
	github.com/aws/aws-sdk-go-v2/service/connectcampaigns v1.1.7
	github.com/aws/aws-sdk-go-v2/service/costandusagereportservice v1.14.1
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.21.1
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.21.13
	github.com/aws/aws-sdk-go-v2/service/databrew v1.20.15
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.13.18
	github.com/aws/aws-sdk-go-v2/service/datasync v1.18.13
	github.com/aws/aws-sdk-go-v2/service/dax v1.11.18
	github.com/aws/aws-sdk-go-v2/service/detective v1.16.10
	github.com/aws/aws-sdk-go-v2/service/devopsguru v1.20.1
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.17.19
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.15.1
	github.com/aws/aws-sdk-go-v2/service/docdb v1.19.12
	github.com/aws/aws-sdk-go-v2/service/drs v1.8.3
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.17.2
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.13.21
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.63.2
	github.com/aws/aws-sdk-go-v2/service/ecr v1.17.19
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.13.18
	github.com/aws/aws-sdk-go-v2/service/ecs v1.18.25
	github.com/aws/aws-sdk-go-v2/service/efs v1.17.17
	github.com/aws/aws-sdk-go-v2/service/eks v1.22.2
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.22.11
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.14.19
	github.com/aws/aws-sdk-go-v2/service/elasticinference v1.11.18
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.14.19
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.18.21
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.16.11
	github.com/aws/aws-sdk-go-v2/service/emr v1.20.12
	github.com/aws/aws-sdk-go-v2/service/emrcontainers v1.14.2
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.16.16
	github.com/aws/aws-sdk-go-v2/service/firehose v1.14.20
	github.com/aws/aws-sdk-go-v2/service/forecast v1.23.6
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.20.11
	github.com/aws/aws-sdk-go-v2/service/fsx v1.25.2
	github.com/aws/aws-sdk-go-v2/service/gamelift v1.15.6
	github.com/aws/aws-sdk-go-v2/service/glacier v1.13.18
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.15.1
	github.com/aws/aws-sdk-go-v2/service/grafana v1.9.14
	github.com/aws/aws-sdk-go-v2/service/greengrassv2 v1.19.1
	github.com/aws/aws-sdk-go-v2/service/groundstation v1.13.19
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.16.1
	github.com/aws/aws-sdk-go-v2/service/health v1.15.21
	github.com/aws/aws-sdk-go-v2/service/healthlake v1.14.18
	github.com/aws/aws-sdk-go-v2/service/honeycode v1.12.18
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.15.6
	github.com/aws/aws-sdk-go-v2/service/inspector v1.12.18
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.8.2
	github.com/aws/aws-sdk-go-v2/service/iot v1.29.2
	github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice v1.10.18
	github.com/aws/aws-sdk-go-v2/service/iot1clickprojects v1.11.18
	github.com/aws/aws-sdk-go-v2/service/iotanalytics v1.13.2
	github.com/aws/aws-sdk-go-v2/service/iotevents v1.14.20
	github.com/aws/aws-sdk-go-v2/service/ioteventsdata v1.12.13
	github.com/aws/aws-sdk-go-v2/service/iotfleethub v1.12.18
	github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane v1.11.18
	github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling v1.14.2
	github.com/aws/aws-sdk-go-v2/service/iotsitewise v1.25.2
	github.com/aws/aws-sdk-go-v2/service/iotthingsgraph v1.13.4
	github.com/aws/aws-sdk-go-v2/service/kafka v1.17.20
	github.com/aws/aws-sdk-go-v2/service/kafkaconnect v1.8.18
	github.com/aws/aws-sdk-go-v2/service/kendra v1.35.1
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.15.20
	github.com/aws/aws-sdk-go-v2/service/kinesisanalytics v1.13.19
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.14.19
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.12.15
	github.com/aws/aws-sdk-go-v2/service/kms v1.18.14
	github.com/aws/aws-sdk-go-v2/service/lakeformation v1.17.6
	github.com/aws/aws-sdk-go-v2/service/lexmodelsv2 v1.25.1
	github.com/aws/aws-sdk-go-v2/service/location v1.19.1
	github.com/aws/aws-sdk-go-v2/service/lookoutequipment v1.15.5
	github.com/aws/aws-sdk-go-v2/service/lookoutmetrics v1.18.2
	github.com/aws/aws-sdk-go-v2/service/lookoutvision v1.14.10
	github.com/aws/aws-sdk-go-v2/service/machinelearning v1.14.18
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.23.5
	github.com/aws/aws-sdk-go-v2/service/marketplacecatalog v1.13.10
	github.com/aws/aws-sdk-go-v2/service/mediaconnect v1.17.2
	github.com/aws/aws-sdk-go-v2/service/mediaconvert v1.26.1
	github.com/aws/aws-sdk-go-v2/service/medialive v1.24.1
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.19.3
	github.com/aws/aws-sdk-go-v2/service/mediapackagevod v1.18.1
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.12.18
	github.com/aws/aws-sdk-go-v2/service/mediastoredata v1.12.18
	github.com/aws/aws-sdk-go-v2/service/mediatailor v1.17.15
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.9.19
	github.com/aws/aws-sdk-go-v2/service/mgn v1.15.13
	github.com/aws/aws-sdk-go-v2/service/migrationhub v1.12.18
	github.com/aws/aws-sdk-go-v2/service/migrationhubconfig v1.12.18
	github.com/aws/aws-sdk-go-v2/service/mobile v1.11.18
	github.com/aws/aws-sdk-go-v2/service/mq v1.13.14
	github.com/aws/aws-sdk-go-v2/service/neptune v1.17.13
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.20.2
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.15.6
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.10.11
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.13.18
	github.com/aws/aws-sdk-go-v2/service/opsworkscm v1.14.18
	github.com/aws/aws-sdk-go-v2/service/organizations v1.16.14
	github.com/aws/aws-sdk-go-v2/service/panorama v1.9.1
	github.com/aws/aws-sdk-go-v2/service/personalize v1.21.9
	github.com/aws/aws-sdk-go-v2/service/pi v1.15.2
	github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2 v1.0.16
	github.com/aws/aws-sdk-go-v2/service/polly v1.18.1
	github.com/aws/aws-sdk-go-v2/service/pricing v1.17.2
	github.com/aws/aws-sdk-go-v2/service/qldb v1.14.19
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.26.1
	github.com/aws/aws-sdk-go-v2/service/rds v1.26.2
	github.com/aws/aws-sdk-go-v2/service/redshift v1.26.11
	github.com/aws/aws-sdk-go-v2/service/redshiftdata v1.16.12
	github.com/aws/aws-sdk-go-v2/service/rekognition v1.20.6
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.7.1
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.13.20
	github.com/aws/aws-sdk-go-v2/service/robomaker v1.17.19
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.9.18
	github.com/aws/aws-sdk-go-v2/service/s3control v1.24.1
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.49.0
	github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime v1.14.2
	github.com/aws/aws-sdk-go-v2/service/savingsplans v1.11.18
	github.com/aws/aws-sdk-go-v2/service/schemas v1.14.18
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.16.3
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.23.6
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.14.18
	github.com/aws/aws-sdk-go-v2/service/ses v1.14.19
	github.com/aws/aws-sdk-go-v2/service/sfn v1.14.2
	github.com/aws/aws-sdk-go-v2/service/shield v1.17.10
	github.com/aws/aws-sdk-go-v2/service/signer v1.13.18
	github.com/aws/aws-sdk-go-v2/service/snowball v1.16.1
	github.com/aws/aws-sdk-go-v2/service/snowdevicemanagement v1.8.18
	github.com/aws/aws-sdk-go-v2/service/ssm v1.31.2
	github.com/aws/aws-sdk-go-v2/service/ssmcontacts v1.13.19
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.15.12
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.17.15
	github.com/aws/aws-sdk-go-v2/service/sts v1.17.0
	github.com/aws/aws-sdk-go-v2/service/support v1.13.18
	github.com/aws/aws-sdk-go-v2/service/swf v1.13.19
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.16.12
	github.com/aws/aws-sdk-go-v2/service/timestreamquery v1.14.18
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.14.3
	github.com/aws/aws-sdk-go-v2/service/transcribe v1.21.10
	github.com/aws/aws-sdk-go-v2/service/transfer v1.23.1
	github.com/aws/aws-sdk-go-v2/service/translate v1.15.2
	github.com/aws/aws-sdk-go-v2/service/voiceid v1.11.5
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.22.10
	github.com/aws/aws-sdk-go-v2/service/workdocs v1.11.18
	github.com/aws/aws-sdk-go-v2/service/worklink v1.12.15
	github.com/aws/aws-sdk-go-v2/service/workmail v1.17.1
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.23.1
	github.com/lib/pq v1.10.7
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.8 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.12.22 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.18 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.24 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.18 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.13.7 // indirect
	github.com/aws/smithy-go v1.13.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

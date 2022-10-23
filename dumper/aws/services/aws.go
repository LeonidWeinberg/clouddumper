package services

import (
	
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/apprunner"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cloudwatch"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/detective"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/directoryservice"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/securityhub"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/storagegateway"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/identitystore"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/kinesisvideo"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/lexmodelsv2"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/snowdevicemanagement"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/ssmcontacts"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/applicationautoscaling"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/ioteventsdata"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/panorama"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/support"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/workmail"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/kafka"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/mediapackagevod"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/neptune"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/signer"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/swf"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/translate"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cloudhsmv2"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cloudtrail"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/devopsguru"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/emrcontainers"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/lakeformation"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/pinpointsmsvoicev2"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/dynamodb"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/elasticinference"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/fsx"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/iot1clickdevicesservice"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/polly"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/quicksight"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/machinelearning"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/autoscalingplans"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/codegurureviewer"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/dax"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/ecrpublic"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/elasticache"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/elasticloadbalancing"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/rds"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/chime"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/ecr"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/elasticbeanstalk"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/inspector"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/kinesisanalytics"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/networkfirewall"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/ecs"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/iotthingsgraph"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/mediapackage"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/mediastoredata"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/savingsplans"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/wafv2"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/autoscaling"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cloudwatchevents"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cognitoidentityprovider"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/lookoutequipment"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/s3control"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/eventbridge"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/pricing"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/timestreamquery"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/connect"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/drs"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/honeycode"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/iotfleethub"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/kinesisanalyticsv2"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/redshiftdata"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/opsworks"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/snowball"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cloud9"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/comprehend"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/configservice"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/emr"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/glacier"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/iotsecuretunneling"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/resourcegroupstaggingapi"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/voiceid"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cloudfront"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/codecommit"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/codestarnotifications"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/greengrassv2"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/mobile"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/organizations"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/batch"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/chimesdkmessaging"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/qldb"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/workdocs"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/costexplorer"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/mediaconvert"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/robomaker"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/schemas"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/costandusagereportservice"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/datapipeline"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/iotsitewise"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/networkmanager"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/sagemakera2iruntime"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/macie2"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/mediaconnect"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/redshift"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/ssm"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/worklink"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/transfer"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/acm"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/codeguruprofiler"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/eks"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/migrationhub"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/secretsmanager"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/synthetics"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cognitosync"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/pi"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/amp"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/applicationdiscoveryservice"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/codebuild"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/connectcampaigns"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/iot"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/applicationinsights"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/databrew"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/kendra"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/resiliencehub"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/ses"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/budgets"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/chimesdkidentity"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/mgn"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/opsworkscm"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/appstream"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/databasemigrationservice"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/location"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/shield"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/transcribe"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/mq"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/ssoadmin"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/codeartifact"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/datasync"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/docdb"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/gamelift"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/grafana"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/groundstation"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/workspaces"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/memorydb"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/sagemaker"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/appmesh"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cloudwatchlogs"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/globalaccelerator"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/inspector2"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/lookoutmetrics"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/mediatailor"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/sfn"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/computeoptimizer"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/marketplacecatalog"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/servicecatalog"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/mediastore"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/acmpca"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/comprehendmedical"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/elasticsearchservice"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/iotanalytics"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/iotevents"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/kms"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/frauddetector"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/guardduty"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/kafkaconnect"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/medialive"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/personalize"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/timestreamwrite"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/lookoutvision"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/route53recoverycontrolconfig"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/codestar"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cognitoidentity"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/firehose"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/forecast"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/health"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/healthlake"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cloudhsm"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/ec2"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/efs"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/migrationhubconfig"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/opensearch"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/elasticloadbalancingv2"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/iot1clickprojects"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/appflow"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/backup"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cloudformation"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/cloudsearch"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/directconnect"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/dynamodbstreams"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/iotjobsdataplane"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/kinesis"
    
    "github.com/leonidweinberg/clouddumper/dumper/aws/services/rekognition"
    
    "github.com/aws/aws-sdk-go-v2/aws"
)

type Service interface {
	Dump(aws.Config, func(map[string]interface{}, string, string))
}

func GetServices() []Service {
    return []Service{
        
        &apprunner.Service{},
        
        &cloudwatch.Service{},
        
        &detective.Service{},
        
        &directoryservice.Service{},
        
        &securityhub.Service{},
        
        &storagegateway.Service{},
        
        &identitystore.Service{},
        
        &kinesisvideo.Service{},
        
        &lexmodelsv2.Service{},
        
        &snowdevicemanagement.Service{},
        
        &ssmcontacts.Service{},
        
        &applicationautoscaling.Service{},
        
        &ioteventsdata.Service{},
        
        &panorama.Service{},
        
        &support.Service{},
        
        &workmail.Service{},
        
        &kafka.Service{},
        
        &mediapackagevod.Service{},
        
        &neptune.Service{},
        
        &signer.Service{},
        
        &swf.Service{},
        
        &translate.Service{},
        
        &cloudhsmv2.Service{},
        
        &cloudtrail.Service{},
        
        &devopsguru.Service{},
        
        &emrcontainers.Service{},
        
        &lakeformation.Service{},
        
        &pinpointsmsvoicev2.Service{},
        
        &dynamodb.Service{},
        
        &elasticinference.Service{},
        
        &fsx.Service{},
        
        &iot1clickdevicesservice.Service{},
        
        &polly.Service{},
        
        &quicksight.Service{},
        
        &machinelearning.Service{},
        
        &autoscalingplans.Service{},
        
        &codegurureviewer.Service{},
        
        &dax.Service{},
        
        &ecrpublic.Service{},
        
        &elasticache.Service{},
        
        &elasticloadbalancing.Service{},
        
        &rds.Service{},
        
        &chime.Service{},
        
        &ecr.Service{},
        
        &elasticbeanstalk.Service{},
        
        &inspector.Service{},
        
        &kinesisanalytics.Service{},
        
        &networkfirewall.Service{},
        
        &ecs.Service{},
        
        &iotthingsgraph.Service{},
        
        &mediapackage.Service{},
        
        &mediastoredata.Service{},
        
        &savingsplans.Service{},
        
        &wafv2.Service{},
        
        &autoscaling.Service{},
        
        &cloudwatchevents.Service{},
        
        &cognitoidentityprovider.Service{},
        
        &lookoutequipment.Service{},
        
        &s3control.Service{},
        
        &eventbridge.Service{},
        
        &pricing.Service{},
        
        &timestreamquery.Service{},
        
        &connect.Service{},
        
        &drs.Service{},
        
        &honeycode.Service{},
        
        &iotfleethub.Service{},
        
        &kinesisanalyticsv2.Service{},
        
        &redshiftdata.Service{},
        
        &opsworks.Service{},
        
        &snowball.Service{},
        
        &cloud9.Service{},
        
        &comprehend.Service{},
        
        &configservice.Service{},
        
        &emr.Service{},
        
        &glacier.Service{},
        
        &iotsecuretunneling.Service{},
        
        &resourcegroupstaggingapi.Service{},
        
        &voiceid.Service{},
        
        &cloudfront.Service{},
        
        &codecommit.Service{},
        
        &codestarnotifications.Service{},
        
        &greengrassv2.Service{},
        
        &mobile.Service{},
        
        &organizations.Service{},
        
        &batch.Service{},
        
        &chimesdkmessaging.Service{},
        
        &qldb.Service{},
        
        &workdocs.Service{},
        
        &costexplorer.Service{},
        
        &mediaconvert.Service{},
        
        &robomaker.Service{},
        
        &schemas.Service{},
        
        &costandusagereportservice.Service{},
        
        &datapipeline.Service{},
        
        &iotsitewise.Service{},
        
        &networkmanager.Service{},
        
        &sagemakera2iruntime.Service{},
        
        &macie2.Service{},
        
        &mediaconnect.Service{},
        
        &redshift.Service{},
        
        &ssm.Service{},
        
        &worklink.Service{},
        
        &transfer.Service{},
        
        &acm.Service{},
        
        &codeguruprofiler.Service{},
        
        &eks.Service{},
        
        &migrationhub.Service{},
        
        &secretsmanager.Service{},
        
        &synthetics.Service{},
        
        &cognitosync.Service{},
        
        &pi.Service{},
        
        &amp.Service{},
        
        &applicationdiscoveryservice.Service{},
        
        &codebuild.Service{},
        
        &connectcampaigns.Service{},
        
        &iot.Service{},
        
        &applicationinsights.Service{},
        
        &databrew.Service{},
        
        &kendra.Service{},
        
        &resiliencehub.Service{},
        
        &ses.Service{},
        
        &budgets.Service{},
        
        &chimesdkidentity.Service{},
        
        &mgn.Service{},
        
        &opsworkscm.Service{},
        
        &appstream.Service{},
        
        &databasemigrationservice.Service{},
        
        &location.Service{},
        
        &shield.Service{},
        
        &transcribe.Service{},
        
        &mq.Service{},
        
        &ssoadmin.Service{},
        
        &codeartifact.Service{},
        
        &datasync.Service{},
        
        &docdb.Service{},
        
        &gamelift.Service{},
        
        &grafana.Service{},
        
        &groundstation.Service{},
        
        &workspaces.Service{},
        
        &memorydb.Service{},
        
        &sagemaker.Service{},
        
        &appmesh.Service{},
        
        &cloudwatchlogs.Service{},
        
        &globalaccelerator.Service{},
        
        &inspector2.Service{},
        
        &lookoutmetrics.Service{},
        
        &mediatailor.Service{},
        
        &sfn.Service{},
        
        &computeoptimizer.Service{},
        
        &marketplacecatalog.Service{},
        
        &servicecatalog.Service{},
        
        &mediastore.Service{},
        
        &acmpca.Service{},
        
        &comprehendmedical.Service{},
        
        &elasticsearchservice.Service{},
        
        &iotanalytics.Service{},
        
        &iotevents.Service{},
        
        &kms.Service{},
        
        &frauddetector.Service{},
        
        &guardduty.Service{},
        
        &kafkaconnect.Service{},
        
        &medialive.Service{},
        
        &personalize.Service{},
        
        &timestreamwrite.Service{},
        
        &lookoutvision.Service{},
        
        &route53recoverycontrolconfig.Service{},
        
        &codestar.Service{},
        
        &cognitoidentity.Service{},
        
        &firehose.Service{},
        
        &forecast.Service{},
        
        &health.Service{},
        
        &healthlake.Service{},
        
        &cloudhsm.Service{},
        
        &ec2.Service{},
        
        &efs.Service{},
        
        &migrationhubconfig.Service{},
        
        &opensearch.Service{},
        
        &elasticloadbalancingv2.Service{},
        
        &iot1clickprojects.Service{},
        
        &appflow.Service{},
        
        &backup.Service{},
        
        &cloudformation.Service{},
        
        &cloudsearch.Service{},
        
        &directconnect.Service{},
        
        &dynamodbstreams.Service{},
        
        &iotjobsdataplane.Service{},
        
        &kinesis.Service{},
        
        &rekognition.Service{},
        
    }
}

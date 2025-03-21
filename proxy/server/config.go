package server

import (
	"time"

	"github.com/mason-leap-lab/infinicache/proxy/lambdastore"
)

const AWSRegion = "us-west-2"
const LambdaMaxDeployments = 20
const NumLambdaClusters = 20
const LambdaStoreName = "LambdaStore" // replica version (no use)
const LambdaPrefix = "CacheNodeA"
const InstanceWarmTimout = 1 * time.Minute
const InstanceCapacity = 128 * 1000000 // MB
const InstanceOverhead = 100 * 1000000 // MB
const ServerPublicIp = ""              // Leave it empty if using VPC.

func init() {
	lambdastore.WarmTimout = InstanceWarmTimout
}

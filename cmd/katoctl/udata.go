package main

//-----------------------------------------------------------------------------
// 'katoctl udata' command flags definitions:
//-----------------------------------------------------------------------------

var (

	//--------------------------
	// udata: top level command
	//--------------------------

	cmdUdata = app.Command("udata", "Generate CoreOS cloud-config user-data.")

	flUdataMasterCount = cmdUdata.Flag("master-count",
		"Number of master nodes [ 1 | 3 | 5 ]").
		Default("3").OverrideDefaultFromEnvar("KATO_UDATA_MASTER_COUNT").
		HintOptions("1", "3", "5").Int()

	flUdataClusterID = cmdUdata.Flag("cluster-id",
		"Cluster ID.").
		Required().PlaceHolder("KATO_UDATA_CLUSTER_ID").
		OverrideDefaultFromEnvar("KATO_UDATA_CLUSTER_ID").
		String()

	flUdataHostID = cmdUdata.Flag("hostid",
		"Must be a number: hostname = <role>-<hostid>").
		Required().PlaceHolder("KATO_UDATA_HOSTID").
		OverrideDefaultFromEnvar("KATO_UDATA_HOSTID").
		String()

	flUdataDomain = cmdUdata.Flag("domain",
		"Domain name as in (hostname -d)").
		Required().PlaceHolder("KATO_UDATA_DOMAIN").
		OverrideDefaultFromEnvar("KATO_UDATA_DOMAIN").
		String()

	flUdataRole = cmdUdata.Flag("role",
		"Choose one of [ master | worker | edge ]").
		Required().PlaceHolder("KATO_UDATA_ROLE").
		OverrideDefaultFromEnvar("KATO_UDATA_ROLE").
		Enum("master", "worker", "edge")

	flUdataNs1Apikey = cmdUdata.Flag("ns1-api-key",
		"NS1 private API key.").
		Required().PlaceHolder("KATO_UDATA_NS1_API_KEY").
		OverrideDefaultFromEnvar("KATO_UDATA_NS1_API_KEY").
		String()

	flUdataCaCert = cmdUdata.Flag("ca-cert",
		"Path to CA certificate.").
		PlaceHolder("KATO_UDATA_CA_CERT").
		OverrideDefaultFromEnvar("KATO_UDATA_CA_CERT").
		ExistingFile()

	flUdataEtcdToken = cmdUdata.Flag("etcd-token",
		"Provide an etcd discovery token.").
		PlaceHolder("KATO_UDATA_ETCD_TOKEN").
		OverrideDefaultFromEnvar("KATO_UDATA_ETCD_TOKEN").
		String()

	flUdataGzipUdata = cmdUdata.Flag("gzip-udata",
		"Enable udata compression.").
		Default("false").OverrideDefaultFromEnvar("KATO_UDATA_GZIP_UDATA").
		Bool()

	flUdataFlannelNetwork = cmdUdata.Flag("flannel-network",
		"Flannel entire overlay network.").
		Default("10.128.0.0/21").
		OverrideDefaultFromEnvar("KATO_UDATA_FLANNEL_NETWORK").
		String()

	flUdataFlannelSubnetLen = cmdUdata.Flag("flannel-subnet-len",
		"Subnet len to llocate to each host.").
		Default("27").OverrideDefaultFromEnvar("KATO_UDATA_FLANNEL_SUBNET_LEN").
		String()

	flUdataFlannelSubnetMin = cmdUdata.Flag("flannel-subnet-min",
		"Minimum subnet IP addresses.").
		Default("10.128.0.192").
		OverrideDefaultFromEnvar("KATO_UDATA_FLANNEL_SUBNET_MIN").
		String()

	flUdataFlannelSubnetMax = cmdUdata.Flag("flannel-subnet-max",
		"Maximum subnet IP addresses.").
		Default("10.128.7.224").
		OverrideDefaultFromEnvar("KATO_UDATA_FLANNEL_SUBNET_MAX").
		String()

	flUdataFlannelBackend = cmdUdata.Flag("flannel-backend",
		"Flannel backend: [ udp | vxlan | host-gw | gce | aws-vpc | alloc ]").
		Default("vxlan").OverrideDefaultFromEnvar("KATO_UDATA_FLANNEL_BACKEND").
		Enum("udp", "vxlan", "host-gw", "gce", "aws-vpc", "alloc")

	flUdataRexrayStorageDriver = cmdUdata.Flag("rexray-storage-driver",
		"REX-Ray storage driver: [ ec2 | virtualbox ]").
		PlaceHolder("KATO_UDATA_REXRAY_STORAGE_DRIVER").
		OverrideDefaultFromEnvar("KATO_UDATA_REXRAY_STORAGE_DRIVER").
		Enum("virtualbox", "ec2")

	flUdataRexrayEndpointIP = cmdUdata.Flag("rexray-endpoint-ip",
		"REX-Ray endpoint IP address.").
		PlaceHolder("KATO_UDATA_REXRAY_ENDPOINT_IP").
		OverrideDefaultFromEnvar("KATO_UDATA_REXRAY_ENDPOINT_IP").
		String()

	flUdataEc2Region = cmdUdata.Flag("ec2-region",
		"EC2 region.").
		Default("eu-west-1").PlaceHolder("KATO_UDATA_EC2_REGION").
		OverrideDefaultFromEnvar("KATO_UDATA_EC2_REGION").
		Enum(ec2Regions...)
)

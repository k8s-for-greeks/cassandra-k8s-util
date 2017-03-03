package cass

type CassandraConf struct {
	Authenticator              string `yaml:"authenticator" json:"authenticator"`
	Authorizer                 string `yaml:"authorizer" json:"authorizer"`
	AutoSnapshot               bool   `yaml:"auto_snapshot" json:"auto_snapshot"`
	BatchSizeFailThresholdInKb int    `yaml:"batch_size_fail_threshold_in_kb" json:"batch_size_fail_threshold_in_kb"`
	BatchSizeWarnThresholdInKb int    `yaml:"batch_size_warn_threshold_in_kb" json:"batch_size_warn_threshold_in_kb"`
	BatchlogReplayThrottleInKb int    `yaml:"batchlog_replay_throttle_in_kb" json:"batchlog_replay_throttle_in_kb"`
	CasContentionTimeoutInMs   int    `yaml:"cas_contention_timeout_in_ms" json:"cas_contention_timeout_in_ms"`
	ClientEncryptionOptions    struct {
		Enabled          bool   `yaml:"enabled" json:"enabled"`
		Keystore         string `yaml:"keystore" json:"keystore"`
		KeystorePassword string `yaml:"keystore_password" json:"keystore_password"`
		Optional         bool   `yaml:"optional" json:"optional"`
	} `yaml:"client_encryption_options" json:"client_encryption_options"`
	ClusterName                                string      `yaml:"cluster_name" json:"cluster_name"`
	ColumnIndexSizeInKb                        int         `yaml:"column_index_size_in_kb" json:"column_index_size_in_kb"`
	CommitFailurePolicy                        string      `yaml:"commit_failure_policy" json:"commit_failure_policy"`
	CommitlogDirectory                         string      `yaml:"commitlog_directory" json:"commitlog_directory"`
	CommitlogSegmentSizeInMb                   int         `yaml:"commitlog_segment_size_in_mb" json:"commitlog_segment_size_in_mb"`
	CommitlogSync                              string      `yaml:"commitlog_sync" json:"commitlog_sync"`
	CommitlogSyncPeriodInMs                    int         `yaml:"commitlog_sync_period_in_ms" json:"commitlog_sync_period_in_ms"`
	CompactionLargePartitionWarningThresholdMb int         `yaml:"compaction_large_partition_warning_threshold_mb" json:"compaction_large_partition_warning_threshold_mb"`
	CompactionThroughputMbPerSec               int         `yaml:"compaction_throughput_mb_per_sec" json:"compaction_throughput_mb_per_sec"`
	ConcurrentCounterWrites                    int         `yaml:"concurrent_counter_writes" json:"concurrent_counter_writes"`
	ConcurrentMaterializedViewWrites           int         `yaml:"concurrent_materialized_view_writes" json:"concurrent_materialized_view_writes"`
	ConcurrentReads                            int         `yaml:"concurrent_reads" json:"concurrent_reads"`
	ConcurrentWrites                           int         `yaml:"concurrent_writes" json:"concurrent_writes"`
	CounterCacheSavePeriod                     int         `yaml:"counter_cache_save_period" json:"counter_cache_save_period"`
	CounterCacheSizeInMb                       interface{} `yaml:"counter_cache_size_in_mb" json:"counter_cache_size_in_mb"`
	CounterWriteRequestTimeoutInMs             int         `yaml:"counter_write_request_timeout_in_ms" json:"counter_write_request_timeout_in_ms"`
	CredentialsValidityInMs                    int         `yaml:"credentials_validity_in_ms" json:"credentials_validity_in_ms"`
	CrossNodeTimeout                           bool        `yaml:"cross_node_timeout" json:"cross_node_timeout"`
	DataFileDirectories                        []string    `yaml:"data_file_directories" json:"data_file_directories"`
	DiskFailurePolicy                          string      `yaml:"disk_failure_policy" json:"disk_failure_policy"`
	DynamicSnitchBadnessThreshold              float64     `yaml:"dynamic_snitch_badness_threshold" json:"dynamic_snitch_badness_threshold"`
	DynamicSnitchResetIntervalInMs             int         `yaml:"dynamic_snitch_reset_interval_in_ms" json:"dynamic_snitch_reset_interval_in_ms"`
	DynamicSnitchUpdateIntervalInMs            int         `yaml:"dynamic_snitch_update_interval_in_ms" json:"dynamic_snitch_update_interval_in_ms"`
	EnableScriptedUserDefinedFunctions         bool        `yaml:"enable_scripted_user_defined_functions" json:"enable_scripted_user_defined_functions"`
	EnableUserDefinedFunctions                 bool        `yaml:"enable_user_defined_functions" json:"enable_user_defined_functions"`
	EndpointSnitch                             string      `yaml:"endpoint_snitch" json:"endpoint_snitch"`
	GcWarnThresholdInMs                        int         `yaml:"gc_warn_threshold_in_ms" json:"gc_warn_threshold_in_ms"`
	HintedHandoffEnabled                       bool        `yaml:"hinted_handoff_enabled" json:"hinted_handoff_enabled"`
	HintedHandoffThrottleInKb                  int         `yaml:"hinted_handoff_throttle_in_kb" json:"hinted_handoff_throttle_in_kb"`
	HintsDirectory                             string      `yaml:"hints_directory" json:"hints_directory"`
	HintsFlushPeriodInMs                       int         `yaml:"hints_flush_period_in_ms" json:"hints_flush_period_in_ms"`
	IncrementalBackups                         bool        `yaml:"incremental_backups" json:"incremental_backups"`
	IndexSummaryCapacityInMb                   interface{} `yaml:"index_summary_capacity_in_mb" json:"index_summary_capacity_in_mb"`
	IndexSummaryResizeIntervalInMinutes        int         `yaml:"index_summary_resize_interval_in_minutes" json:"index_summary_resize_interval_in_minutes"`
	InterDcTcpNodelay                          bool        `yaml:"inter_dc_tcp_nodelay" json:"inter_dc_tcp_nodelay"`
	InternodeCompression                       string      `yaml:"internode_compression" json:"internode_compression"`
	KeyCacheSavePeriod                         int         `yaml:"key_cache_save_period" json:"key_cache_save_period"`
	KeyCacheSizeInMb                           interface{} `yaml:"key_cache_size_in_mb" json:"key_cache_size_in_mb"`
	ListenAddress                              string      `yaml:"listen_address" json:"listen_address"`
	MaxHintWindowInMs                          int         `yaml:"max_hint_window_in_ms" json:"max_hint_window_in_ms"`
	MaxHintsDeliveryThreads                    int         `yaml:"max_hints_delivery_threads" json:"max_hints_delivery_threads"`
	MaxHintsFileSizeInMb                       int         `yaml:"max_hints_file_size_in_mb" json:"max_hints_file_size_in_mb"`
	MemtableAllocationType                     string      `yaml:"memtable_allocation_type" json:"memtable_allocation_type"`
	NativeTransportPort                        int         `yaml:"native_transport_port" json:"native_transport_port"`
	NumTokens                                  int         `yaml:"num_tokens" json:"num_tokens"`
	Partitioner                                string      `yaml:"partitioner" json:"partitioner"`
	PermissionsValidityInMs                    int         `yaml:"permissions_validity_in_ms" json:"permissions_validity_in_ms"`
	RangeRequestTimeoutInMs                    int         `yaml:"range_request_timeout_in_ms" json:"range_request_timeout_in_ms"`
	ReadRequestTimeoutInMs                     int         `yaml:"read_request_timeout_in_ms" json:"read_request_timeout_in_ms"`
	RequestScheduler                           string      `yaml:"request_scheduler" json:"request_scheduler"`
	RequestTimeoutInMs                         int         `yaml:"request_timeout_in_ms" json:"request_timeout_in_ms"`
	RoleManager                                string      `yaml:"role_manager" json:"role_manager"`
	RolesValidityInMs                          int         `yaml:"roles_validity_in_ms" json:"roles_validity_in_ms"`
	RowCacheSavePeriod                         int         `yaml:"row_cache_save_period" json:"row_cache_save_period"`
	RowCacheSizeInMb                           int         `yaml:"row_cache_size_in_mb" json:"row_cache_size_in_mb"`
	RPCAddress                                 string      `yaml:"rpc_address" json:"rpc_address"`
	RPCKeepalive                               bool        `yaml:"rpc_keepalive" json:"rpc_keepalive"`
	RPCPort                                    int         `yaml:"rpc_port" json:"rpc_port"`
	RPCServerType                              string      `yaml:"rpc_server_type" json:"rpc_server_type"`
	SavedCachesDirectory                       string      `yaml:"saved_caches_directory" json:"saved_caches_directory"`
	SeedProvider                               []struct {
		ClassName  string `yaml:"class_name" json:"class_name"`
		Parameters []struct {
			Seeds string `yaml:"seeds" json:"seeds"`
		} `yaml:"parameters" json:"parameters"`
	} `yaml:"seed_provider" json:"seed_provider"`
	ServerEncryptionOptions struct {
		InternodeEncryption string `yaml:"internode_encryption" json:"internode_encryption"`
		Keystore            string `yaml:"keystore" json:"keystore"`
		KeystorePassword    string `yaml:"keystore_password" json:"keystore_password"`
		Truststore          string `yaml:"truststore" json:"truststore"`
		TruststorePassword  string `yaml:"truststore_password" json:"truststore_password"`
	} `yaml:"server_encryption_options" json:"server_encryption_options"`
	SnapshotBeforeCompaction          bool `yaml:"snapshot_before_compaction" json:"snapshot_before_compaction"`
	SslStoragePort                    int  `yaml:"ssl_storage_port" json:"ssl_storage_port"`
	SstablePreemptiveOpenIntervalInMb int  `yaml:"sstable_preemptive_open_interval_in_mb" json:"sstable_preemptive_open_interval_in_mb"`
	StartNativeTransport              bool `yaml:"start_native_transport" json:"start_native_transport"`
	StartRPC                          bool `yaml:"start_rpc" json:"start_rpc"`
	StoragePort                       int  `yaml:"storage_port" json:"storage_port"`
	ThriftFramedTransportSizeInMb     int  `yaml:"thrift_framed_transport_size_in_mb" json:"thrift_framed_transport_size_in_mb"`
	TombstoneFailureThreshold         int  `yaml:"tombstone_failure_threshold" json:"tombstone_failure_threshold"`
	TombstoneWarnThreshold            int  `yaml:"tombstone_warn_threshold" json:"tombstone_warn_threshold"`
	TracetypeQueryTTL                 int  `yaml:"tracetype_query_ttl" json:"tracetype_query_ttl"`
	TracetypeRepairTTL                int  `yaml:"tracetype_repair_ttl" json:"tracetype_repair_ttl"`
	TransparentDataEncryptionOptions  struct {
		ChunkLengthKb int    `yaml:"chunk_length_kb" json:"chunk_length_kb"`
		Cipher        string `yaml:"cipher" json:"cipher"`
		Enabled       bool   `yaml:"enabled" json:"enabled"`
		KeyAlias      string `yaml:"key_alias" json:"key_alias"`
		KeyProvider   []struct {
			ClassName  string `yaml:"class_name" json:"class_name"`
			Parameters []struct {
				KeyPassword      string `yaml:"key_password" json:"key_password"`
				Keystore         string `yaml:"keystore" json:"keystore"`
				KeystorePassword string `yaml:"keystore_password" json:"keystore_password"`
				StoreType        string `yaml:"store_type" json:"store_type"`
			} `yaml:"parameters" json:"parameters"`
		} `yaml:"key_provider" json:"key_provider"`
	} `yaml:"transparent_data_encryption_options" json:"transparent_data_encryption_options"`
	TrickleFsync               bool `yaml:"trickle_fsync" json:"trickle_fsync"`
	TrickleFsyncIntervalInKb   int  `yaml:"trickle_fsync_interval_in_kb" json:"trickle_fsync_interval_in_kb"`
	TruncateRequestTimeoutInMs int  `yaml:"truncate_request_timeout_in_ms" json:"truncate_request_timeout_in_ms"`
	WindowsTimerInterval       int  `yaml:"windows_timer_interval" json:"windows_timer_interval"`
	WriteRequestTimeoutInMs    int  `yaml:"write_request_timeout_in_ms" json:"write_request_timeout_in_ms"`
}

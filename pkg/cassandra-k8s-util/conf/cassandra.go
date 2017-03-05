//Copyright © 2017 NAME HERE K8S For Greeks / Vorstella
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

////
// Generated code, more details cassandra-k8s-util codegen -h
////

package cass

type CassandraYaml struct {
	Authenticator                                 string               `yaml:"authenticator" json:"authenticator"`
	Authorizer                                    string               `yaml:"authorizer" json:"authorizer"`
	AutoSnapshot                                  bool                 `yaml:"auto_snapshot" json:"auto_snapshot"`
	BackPressureEnabled                           bool                 `yaml:"back_pressure_enabled" json:"back_pressure_enabled"`
	BackPressureStrategy                          []CassandraYaml_sub2 `yaml:"back_pressure_strategy" json:"back_pressure_strategy"`
	BatchSizeFailThresholdInKb                    int                  `yaml:"batch_size_fail_threshold_in_kb" json:"batch_size_fail_threshold_in_kb"`
	BatchSizeWarnThresholdInKb                    int                  `yaml:"batch_size_warn_threshold_in_kb" json:"batch_size_warn_threshold_in_kb"`
	BatchlogReplayThrottleInKb                    int                  `yaml:"batchlog_replay_throttle_in_kb" json:"batchlog_replay_throttle_in_kb"`
	BufferPoolUseHeapIfExhausted                  bool                 `yaml:"buffer_pool_use_heap_if_exhausted" json:"buffer_pool_use_heap_if_exhausted"`
	CasContentionTimeoutInMs                      int                  `yaml:"cas_contention_timeout_in_ms" json:"cas_contention_timeout_in_ms"`
	CdcEnabled                                    bool                 `yaml:"cdc_enabled" json:"cdc_enabled"`
	CdcFreeSpaceCheckIntervalMs                   int                  `yaml:"cdc_free_space_check_interval_ms" json:"cdc_free_space_check_interval_ms"`
	CdcRawDirectory                               string               `yaml:"cdc_raw_directory" json:"cdc_raw_directory"`
	ClientEncryptionOptions                       CassandraYaml_sub3   `yaml:"client_encryption_options" json:"client_encryption_options"`
	ClusterName                                   string               `yaml:"cluster_name" json:"cluster_name"`
	ColumnIndexCacheSizeInKb                      int                  `yaml:"column_index_cache_size_in_kb" json:"column_index_cache_size_in_kb"`
	ColumnIndexSizeInKb                           int                  `yaml:"column_index_size_in_kb" json:"column_index_size_in_kb"`
	CommitFailurePolicy                           string               `yaml:"commit_failure_policy" json:"commit_failure_policy"`
	CommitlogDirectory                            string               `yaml:"commitlog_directory" json:"commitlog_directory"`
	CommitlogSegmentSizeInMb                      int                  `yaml:"commitlog_segment_size_in_mb" json:"commitlog_segment_size_in_mb"`
	CommitlogSync                                 string               `yaml:"commitlog_sync" json:"commitlog_sync"`
	CommitlogSyncBatchWindowInMs                  int                  `yaml:"commitlog_sync_batch_window_in_ms" json:"commitlog_sync_batch_window_in_ms"`
	CommitlogSyncPeriodInMs                       int                  `yaml:"commitlog_sync_period_in_ms" json:"commitlog_sync_period_in_ms"`
	CompactionLargePartitionWarningThresholdMb    int                  `yaml:"compaction_large_partition_warning_threshold_mb" json:"compaction_large_partition_warning_threshold_mb"`
	CompactionThroughputMbPerSec                  int                  `yaml:"compaction_throughput_mb_per_sec" json:"compaction_throughput_mb_per_sec"`
	ConcurrentCompactors                          int                  `yaml:"concurrent_compactors" json:"concurrent_compactors"`
	ConcurrentCounterWrites                       int                  `yaml:"concurrent_counter_writes" json:"concurrent_counter_writes"`
	ConcurrentMaterializedViewWrites              int                  `yaml:"concurrent_materialized_view_writes" json:"concurrent_materialized_view_writes"`
	ConcurrentReads                               int                  `yaml:"concurrent_reads" json:"concurrent_reads"`
	ConcurrentWrites                              int                  `yaml:"concurrent_writes" json:"concurrent_writes"`
	CounterCacheKeysToSave                        int                  `yaml:"counter_cache_keys_to_save" json:"counter_cache_keys_to_save"`
	CounterCacheSavePeriod                        int                  `yaml:"counter_cache_save_period" json:"counter_cache_save_period"`
	CounterWriteRequestTimeoutInMs                int                  `yaml:"counter_write_request_timeout_in_ms" json:"counter_write_request_timeout_in_ms"`
	CredentialsUpdateIntervalInMs                 int                  `yaml:"credentials_update_interval_in_ms" json:"credentials_update_interval_in_ms"`
	CredentialsValidityInMs                       int                  `yaml:"credentials_validity_in_ms" json:"credentials_validity_in_ms"`
	CrossNodeTimeout                              bool                 `yaml:"cross_node_timeout" json:"cross_node_timeout"`
	DataFileDirectories                           []string             `yaml:"data_file_directories" json:"data_file_directories"`
	DiskFailurePolicy                             string               `yaml:"disk_failure_policy" json:"disk_failure_policy"`
	DiskOptimizationStrategy                      string               `yaml:"disk_optimization_strategy" json:"disk_optimization_strategy"`
	DynamicSnitchBadnessThreshold                 float64              `yaml:"dynamic_snitch_badness_threshold" json:"dynamic_snitch_badness_threshold"`
	DynamicSnitchResetIntervalInMs                int                  `yaml:"dynamic_snitch_reset_interval_in_ms" json:"dynamic_snitch_reset_interval_in_ms"`
	DynamicSnitchUpdateIntervalInMs               int                  `yaml:"dynamic_snitch_update_interval_in_ms" json:"dynamic_snitch_update_interval_in_ms"`
	EnableScriptedUserDefinedFunctions            bool                 `yaml:"enable_scripted_user_defined_functions" json:"enable_scripted_user_defined_functions"`
	EnableUserDefinedFunctions                    bool                 `yaml:"enable_user_defined_functions" json:"enable_user_defined_functions"`
	EndpointSnitch                                string               `yaml:"endpoint_snitch" json:"endpoint_snitch"`
	FileCacheSizeInMb                             int                  `yaml:"file_cache_size_in_mb" json:"file_cache_size_in_mb"`
	GcLogThresholdInMs                            int                  `yaml:"gc_log_threshold_in_ms" json:"gc_log_threshold_in_ms"`
	GcWarnThresholdInMs                           int                  `yaml:"gc_warn_threshold_in_ms" json:"gc_warn_threshold_in_ms"`
	HintedHandoffEnabled                          bool                 `yaml:"hinted_handoff_enabled" json:"hinted_handoff_enabled"`
	HintedHandoffThrottleInKb                     int                  `yaml:"hinted_handoff_throttle_in_kb" json:"hinted_handoff_throttle_in_kb"`
	HintsDirectory                                string               `yaml:"hints_directory" json:"hints_directory"`
	HintsFlushPeriodInMs                          int                  `yaml:"hints_flush_period_in_ms" json:"hints_flush_period_in_ms"`
	IncrementalBackups                            bool                 `yaml:"incremental_backups" json:"incremental_backups"`
	IndexSummaryResizeIntervalInMinutes           int                  `yaml:"index_summary_resize_interval_in_minutes" json:"index_summary_resize_interval_in_minutes"`
	InterDcStreamThroughputOutboundMegabitsPerSec int                  `yaml:"inter_dc_stream_throughput_outbound_megabits_per_sec" json:"inter_dc_stream_throughput_outbound_megabits_per_sec"`
	InterDcTcpNodelay                             bool                 `yaml:"inter_dc_tcp_nodelay" json:"inter_dc_tcp_nodelay"`
	InternodeCompression                          string               `yaml:"internode_compression" json:"internode_compression"`
	KeyCacheKeysToSave                            int                  `yaml:"key_cache_keys_to_save" json:"key_cache_keys_to_save"`
	KeyCacheSavePeriod                            int                  `yaml:"key_cache_save_period" json:"key_cache_save_period"`
	ListenInterface                               string               `yaml:"listen_interface" json:"listen_interface"`
	MaxHintWindowInMs                             int                  `yaml:"max_hint_window_in_ms" json:"max_hint_window_in_ms"`
	MaxHintsDeliveryThreads                       int                  `yaml:"max_hints_delivery_threads" json:"max_hints_delivery_threads"`
	MaxHintsFileSizeInMb                          int                  `yaml:"max_hints_file_size_in_mb" json:"max_hints_file_size_in_mb"`
	MaxValueSizeInMb                              int                  `yaml:"max_value_size_in_mb" json:"max_value_size_in_mb"`
	MemtableAllocationType                        string               `yaml:"memtable_allocation_type" json:"memtable_allocation_type"`
	MemtableFlushWriters                          int                  `yaml:"memtable_flush_writers" json:"memtable_flush_writers"`
	MemtableHeapSpaceInMb                         int                  `yaml:"memtable_heap_space_in_mb" json:"memtable_heap_space_in_mb"`
	MemtableOffheapSpaceInMb                      int                  `yaml:"memtable_offheap_space_in_mb" json:"memtable_offheap_space_in_mb"`
	NativeTransportMaxConcurrentConnections       int                  `yaml:"native_transport_max_concurrent_connections" json:"native_transport_max_concurrent_connections"`
	NativeTransportMaxConcurrentConnectionsPerIP  int                  `yaml:"native_transport_max_concurrent_connections_per_ip" json:"native_transport_max_concurrent_connections_per_ip"`
	NativeTransportMaxFrameSizeInMb               int                  `yaml:"native_transport_max_frame_size_in_mb" json:"native_transport_max_frame_size_in_mb"`
	NativeTransportMaxThreads                     int                  `yaml:"native_transport_max_threads" json:"native_transport_max_threads"`
	NativeTransportPort                           int                  `yaml:"native_transport_port" json:"native_transport_port"`
	NumTokens                                     int                  `yaml:"num_tokens" json:"num_tokens"`
	OtcCoalescingEnoughCoalescedMessages          int                  `yaml:"otc_coalescing_enough_coalesced_messages" json:"otc_coalescing_enough_coalesced_messages"`
	OtcCoalescingStrategy                         string               `yaml:"otc_coalescing_strategy" json:"otc_coalescing_strategy"`
	OtcCoalescingWindowUs                         int                  `yaml:"otc_coalescing_window_us" json:"otc_coalescing_window_us"`
	Partitioner                                   string               `yaml:"partitioner" json:"partitioner"`
	PermissionsUpdateIntervalInMs                 int                  `yaml:"permissions_update_interval_in_ms" json:"permissions_update_interval_in_ms"`
	PermissionsValidityInMs                       int                  `yaml:"permissions_validity_in_ms" json:"permissions_validity_in_ms"`
	PhiConvictThreshold                           int                  `yaml:"phi_convict_threshold" json:"phi_convict_threshold"`
	RangeRequestTimeoutInMs                       int                  `yaml:"range_request_timeout_in_ms" json:"range_request_timeout_in_ms"`
	ReadRequestTimeoutInMs                        int                  `yaml:"read_request_timeout_in_ms" json:"read_request_timeout_in_ms"`
	RequestTimeoutInMs                            int                  `yaml:"request_timeout_in_ms" json:"request_timeout_in_ms"`
	RoleManager                                   string               `yaml:"role_manager" json:"role_manager"`
	RolesUpdateIntervalInMs                       int                  `yaml:"roles_update_interval_in_ms" json:"roles_update_interval_in_ms"`
	RolesValidityInMs                             int                  `yaml:"roles_validity_in_ms" json:"roles_validity_in_ms"`
	RowCacheClassName                             string               `yaml:"row_cache_class_name" json:"row_cache_class_name"`
	RowCacheSavePeriod                            int                  `yaml:"row_cache_save_period" json:"row_cache_save_period"`
	RowCacheSizeInMb                              int                  `yaml:"row_cache_size_in_mb" json:"row_cache_size_in_mb"`
	RPCAddress                                    string               `yaml:"rpc_address" json:"rpc_address"`
	RPCKeepalive                                  bool                 `yaml:"rpc_keepalive" json:"rpc_keepalive"`
	SavedCachesDirectory                          string               `yaml:"saved_caches_directory" json:"saved_caches_directory"`
	SeedProvider                                  []CassandraYaml_sub5 `yaml:"seed_provider" json:"seed_provider"`
	ServerEncryptionOptions                       CassandraYaml_sub6   `yaml:"server_encryption_options" json:"server_encryption_options"`
	SlowQueryLogTimeoutInMs                       int                  `yaml:"slow_query_log_timeout_in_ms" json:"slow_query_log_timeout_in_ms"`
	SnapshotBeforeCompaction                      bool                 `yaml:"snapshot_before_compaction" json:"snapshot_before_compaction"`
	SslStoragePort                                int                  `yaml:"ssl_storage_port" json:"ssl_storage_port"`
	SstablePreemptiveOpenIntervalInMb             int                  `yaml:"sstable_preemptive_open_interval_in_mb" json:"sstable_preemptive_open_interval_in_mb"`
	StartNativeTransport                          bool                 `yaml:"start_native_transport" json:"start_native_transport"`
	StoragePort                                   int                  `yaml:"storage_port" json:"storage_port"`
	StreamThroughputOutboundMegabitsPerSec        int                  `yaml:"stream_throughput_outbound_megabits_per_sec" json:"stream_throughput_outbound_megabits_per_sec"`
	StreamingConnectionsPerHost                   int                  `yaml:"streaming_connections_per_host" json:"streaming_connections_per_host"`
	StreamingKeepAlivePeriodInSecs                int                  `yaml:"streaming_keep_alive_period_in_secs" json:"streaming_keep_alive_period_in_secs"`
	TombstoneFailureThreshold                     int                  `yaml:"tombstone_failure_threshold" json:"tombstone_failure_threshold"`
	TombstoneWarnThreshold                        int                  `yaml:"tombstone_warn_threshold" json:"tombstone_warn_threshold"`
	TracetypeQueryTTL                             int                  `yaml:"tracetype_query_ttl" json:"tracetype_query_ttl"`
	TracetypeRepairTTL                            int                  `yaml:"tracetype_repair_ttl" json:"tracetype_repair_ttl"`
	TransparentDataEncryptionOptions              CassandraYaml_sub9   `yaml:"transparent_data_encryption_options" json:"transparent_data_encryption_options"`
	TrickleFsync                                  bool                 `yaml:"trickle_fsync" json:"trickle_fsync"`
	TrickleFsyncIntervalInKb                      int                  `yaml:"trickle_fsync_interval_in_kb" json:"trickle_fsync_interval_in_kb"`
	TruncateRequestTimeoutInMs                    int                  `yaml:"truncate_request_timeout_in_ms" json:"truncate_request_timeout_in_ms"`
	UnloggedBatchAcrossPartitionsWarnThreshold    int                  `yaml:"unlogged_batch_across_partitions_warn_threshold" json:"unlogged_batch_across_partitions_warn_threshold"`
	WindowsTimerInterval                          int                  `yaml:"windows_timer_interval" json:"windows_timer_interval"`
	WriteRequestTimeoutInMs                       int                  `yaml:"write_request_timeout_in_ms" json:"write_request_timeout_in_ms"`
}

type CassandraYaml_sub9 struct {
	ChunkLengthKb int                  `yaml:"chunk_length_kb" json:"chunk_length_kb"`
	Cipher        string               `yaml:"cipher" json:"cipher"`
	Enabled       bool                 `yaml:"enabled" json:"enabled"`
	KeyAlias      string               `yaml:"key_alias" json:"key_alias"`
	KeyProvider   []CassandraYaml_sub8 `yaml:"key_provider" json:"key_provider"`
}

type CassandraYaml_sub2 struct {
	ClassName  string               `yaml:"class_name" json:"class_name"`
	Parameters []CassandraYaml_sub1 `yaml:"parameters" json:"parameters"`
}

type CassandraYaml_sub5 struct {
	ClassName  string               `yaml:"class_name" json:"class_name"`
	Parameters []CassandraYaml_sub4 `yaml:"parameters" json:"parameters"`
}

type CassandraYaml_sub8 struct {
	ClassName  string               `yaml:"class_name" json:"class_name"`
	Parameters []CassandraYaml_sub7 `yaml:"parameters" json:"parameters"`
}

type CassandraYaml_sub3 struct {
	Enabled          bool   `yaml:"enabled" json:"enabled"`
	Keystore         string `yaml:"keystore" json:"keystore"`
	KeystorePassword string `yaml:"keystore_password" json:"keystore_password"`
	Optional         bool   `yaml:"optional" json:"optional"`
}

type CassandraYaml_sub1 struct {
	Factor    int     `yaml:"factor" json:"factor"`
	Flow      string  `yaml:"flow" json:"flow"`
	HighRatio float64 `yaml:"high_ratio" json:"high_ratio"`
}

type CassandraYaml_sub6 struct {
	InternodeEncryption string `yaml:"internode_encryption" json:"internode_encryption"`
	Keystore            string `yaml:"keystore" json:"keystore"`
	KeystorePassword    string `yaml:"keystore_password" json:"keystore_password"`
	Truststore          string `yaml:"truststore" json:"truststore"`
	TruststorePassword  string `yaml:"truststore_password" json:"truststore_password"`
}

type CassandraYaml_sub7 struct {
	KeyPassword      string `yaml:"key_password" json:"key_password"`
	Keystore         string `yaml:"keystore" json:"keystore"`
	KeystorePassword string `yaml:"keystore_password" json:"keystore_password"`
	StoreType        string `yaml:"store_type" json:"store_type"`
}

type CassandraYaml_sub4 struct {
	Seeds string `yaml:"seeds" json:"seeds"`
}

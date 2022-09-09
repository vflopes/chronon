// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: store.proto

package chrononpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StoreConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to EventStore:
	//	*StoreConfiguration_PostgresEvents
	//	*StoreConfiguration_InfluxdbEvents
	//	*StoreConfiguration_RedisEvents
	//	*StoreConfiguration_DynamodbEvents
	EventStore isStoreConfiguration_EventStore `protobuf_oneof:"event_store"`
	// Types that are assignable to SnapshotStore:
	//	*StoreConfiguration_PostgresSnapshots
	//	*StoreConfiguration_InfluxdbSnapshots
	//	*StoreConfiguration_RedisSnapshots
	//	*StoreConfiguration_DynamodbSnapshots
	SnapshotStore isStoreConfiguration_SnapshotStore `protobuf_oneof:"snapshot_store"`
}

func (x *StoreConfiguration) Reset() {
	*x = StoreConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreConfiguration) ProtoMessage() {}

func (x *StoreConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreConfiguration.ProtoReflect.Descriptor instead.
func (*StoreConfiguration) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0}
}

func (x *StoreConfiguration) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *StoreConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *StoreConfiguration) GetEventStore() isStoreConfiguration_EventStore {
	if m != nil {
		return m.EventStore
	}
	return nil
}

func (x *StoreConfiguration) GetPostgresEvents() *StoreConfiguration_Postgres {
	if x, ok := x.GetEventStore().(*StoreConfiguration_PostgresEvents); ok {
		return x.PostgresEvents
	}
	return nil
}

func (x *StoreConfiguration) GetInfluxdbEvents() *StoreConfiguration_InfluxDB {
	if x, ok := x.GetEventStore().(*StoreConfiguration_InfluxdbEvents); ok {
		return x.InfluxdbEvents
	}
	return nil
}

func (x *StoreConfiguration) GetRedisEvents() *StoreConfiguration_Redis {
	if x, ok := x.GetEventStore().(*StoreConfiguration_RedisEvents); ok {
		return x.RedisEvents
	}
	return nil
}

func (x *StoreConfiguration) GetDynamodbEvents() *StoreConfiguration_DynamoDB {
	if x, ok := x.GetEventStore().(*StoreConfiguration_DynamodbEvents); ok {
		return x.DynamodbEvents
	}
	return nil
}

func (m *StoreConfiguration) GetSnapshotStore() isStoreConfiguration_SnapshotStore {
	if m != nil {
		return m.SnapshotStore
	}
	return nil
}

func (x *StoreConfiguration) GetPostgresSnapshots() *StoreConfiguration_Postgres {
	if x, ok := x.GetSnapshotStore().(*StoreConfiguration_PostgresSnapshots); ok {
		return x.PostgresSnapshots
	}
	return nil
}

func (x *StoreConfiguration) GetInfluxdbSnapshots() *StoreConfiguration_InfluxDB {
	if x, ok := x.GetSnapshotStore().(*StoreConfiguration_InfluxdbSnapshots); ok {
		return x.InfluxdbSnapshots
	}
	return nil
}

func (x *StoreConfiguration) GetRedisSnapshots() *StoreConfiguration_Redis {
	if x, ok := x.GetSnapshotStore().(*StoreConfiguration_RedisSnapshots); ok {
		return x.RedisSnapshots
	}
	return nil
}

func (x *StoreConfiguration) GetDynamodbSnapshots() *StoreConfiguration_DynamoDB {
	if x, ok := x.GetSnapshotStore().(*StoreConfiguration_DynamodbSnapshots); ok {
		return x.DynamodbSnapshots
	}
	return nil
}

type isStoreConfiguration_EventStore interface {
	isStoreConfiguration_EventStore()
}

type StoreConfiguration_PostgresEvents struct {
	PostgresEvents *StoreConfiguration_Postgres `protobuf:"bytes,3,opt,name=postgres_events,json=postgresEvents,proto3,oneof"`
}

type StoreConfiguration_InfluxdbEvents struct {
	InfluxdbEvents *StoreConfiguration_InfluxDB `protobuf:"bytes,5,opt,name=influxdb_events,json=influxdbEvents,proto3,oneof"`
}

type StoreConfiguration_RedisEvents struct {
	RedisEvents *StoreConfiguration_Redis `protobuf:"bytes,7,opt,name=redis_events,json=redisEvents,proto3,oneof"`
}

type StoreConfiguration_DynamodbEvents struct {
	DynamodbEvents *StoreConfiguration_DynamoDB `protobuf:"bytes,9,opt,name=dynamodb_events,json=dynamodbEvents,proto3,oneof"`
}

func (*StoreConfiguration_PostgresEvents) isStoreConfiguration_EventStore() {}

func (*StoreConfiguration_InfluxdbEvents) isStoreConfiguration_EventStore() {}

func (*StoreConfiguration_RedisEvents) isStoreConfiguration_EventStore() {}

func (*StoreConfiguration_DynamodbEvents) isStoreConfiguration_EventStore() {}

type isStoreConfiguration_SnapshotStore interface {
	isStoreConfiguration_SnapshotStore()
}

type StoreConfiguration_PostgresSnapshots struct {
	PostgresSnapshots *StoreConfiguration_Postgres `protobuf:"bytes,4,opt,name=postgres_snapshots,json=postgresSnapshots,proto3,oneof"`
}

type StoreConfiguration_InfluxdbSnapshots struct {
	InfluxdbSnapshots *StoreConfiguration_InfluxDB `protobuf:"bytes,6,opt,name=influxdb_snapshots,json=influxdbSnapshots,proto3,oneof"`
}

type StoreConfiguration_RedisSnapshots struct {
	RedisSnapshots *StoreConfiguration_Redis `protobuf:"bytes,8,opt,name=redis_snapshots,json=redisSnapshots,proto3,oneof"`
}

type StoreConfiguration_DynamodbSnapshots struct {
	DynamodbSnapshots *StoreConfiguration_DynamoDB `protobuf:"bytes,10,opt,name=dynamodb_snapshots,json=dynamodbSnapshots,proto3,oneof"`
}

func (*StoreConfiguration_PostgresSnapshots) isStoreConfiguration_SnapshotStore() {}

func (*StoreConfiguration_InfluxdbSnapshots) isStoreConfiguration_SnapshotStore() {}

func (*StoreConfiguration_RedisSnapshots) isStoreConfiguration_SnapshotStore() {}

func (*StoreConfiguration_DynamodbSnapshots) isStoreConfiguration_SnapshotStore() {}

type InfluxDBPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measurement string            `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	Tags        map[string]string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InfluxDBPoint) Reset() {
	*x = InfluxDBPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfluxDBPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfluxDBPoint) ProtoMessage() {}

func (x *InfluxDBPoint) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfluxDBPoint.ProtoReflect.Descriptor instead.
func (*InfluxDBPoint) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{1}
}

func (x *InfluxDBPoint) GetMeasurement() string {
	if x != nil {
		return x.Measurement
	}
	return ""
}

func (x *InfluxDBPoint) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type StoreConfiguration_Postgres struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionString string `protobuf:"bytes,1,opt,name=connection_string,json=connectionString,proto3" json:"connection_string,omitempty"`
	EventTable       string `protobuf:"bytes,2,opt,name=event_table,json=eventTable,proto3" json:"event_table,omitempty"`
	SourceTable      string `protobuf:"bytes,3,opt,name=source_table,json=sourceTable,proto3" json:"source_table,omitempty"`
	SnapshotTable    string `protobuf:"bytes,4,opt,name=snapshot_table,json=snapshotTable,proto3" json:"snapshot_table,omitempty"`
	ScanBatchSize    int64  `protobuf:"varint,5,opt,name=scan_batch_size,json=scanBatchSize,proto3" json:"scan_batch_size,omitempty"`
}

func (x *StoreConfiguration_Postgres) Reset() {
	*x = StoreConfiguration_Postgres{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreConfiguration_Postgres) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreConfiguration_Postgres) ProtoMessage() {}

func (x *StoreConfiguration_Postgres) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreConfiguration_Postgres.ProtoReflect.Descriptor instead.
func (*StoreConfiguration_Postgres) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StoreConfiguration_Postgres) GetConnectionString() string {
	if x != nil {
		return x.ConnectionString
	}
	return ""
}

func (x *StoreConfiguration_Postgres) GetEventTable() string {
	if x != nil {
		return x.EventTable
	}
	return ""
}

func (x *StoreConfiguration_Postgres) GetSourceTable() string {
	if x != nil {
		return x.SourceTable
	}
	return ""
}

func (x *StoreConfiguration_Postgres) GetSnapshotTable() string {
	if x != nil {
		return x.SnapshotTable
	}
	return ""
}

func (x *StoreConfiguration_Postgres) GetScanBatchSize() int64 {
	if x != nil {
		return x.ScanBatchSize
	}
	return 0
}

type StoreConfiguration_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreConfiguration_Redis) Reset() {
	*x = StoreConfiguration_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreConfiguration_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreConfiguration_Redis) ProtoMessage() {}

func (x *StoreConfiguration_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreConfiguration_Redis.ProtoReflect.Descriptor instead.
func (*StoreConfiguration_Redis) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0, 1}
}

type StoreConfiguration_InfluxDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url          string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Token        string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Organization string `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty"`
	EventBucket  string `protobuf:"bytes,4,opt,name=event_bucket,json=eventBucket,proto3" json:"event_bucket,omitempty"`
	SourceBucket string `protobuf:"bytes,5,opt,name=source_bucket,json=sourceBucket,proto3" json:"source_bucket,omitempty"`
}

func (x *StoreConfiguration_InfluxDB) Reset() {
	*x = StoreConfiguration_InfluxDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreConfiguration_InfluxDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreConfiguration_InfluxDB) ProtoMessage() {}

func (x *StoreConfiguration_InfluxDB) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreConfiguration_InfluxDB.ProtoReflect.Descriptor instead.
func (*StoreConfiguration_InfluxDB) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0, 2}
}

func (x *StoreConfiguration_InfluxDB) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *StoreConfiguration_InfluxDB) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *StoreConfiguration_InfluxDB) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *StoreConfiguration_InfluxDB) GetEventBucket() string {
	if x != nil {
		return x.EventBucket
	}
	return ""
}

func (x *StoreConfiguration_InfluxDB) GetSourceBucket() string {
	if x != nil {
		return x.SourceBucket
	}
	return ""
}

type StoreConfiguration_DynamoDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreConfiguration_DynamoDB) Reset() {
	*x = StoreConfiguration_DynamoDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreConfiguration_DynamoDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreConfiguration_DynamoDB) ProtoMessage() {}

func (x *StoreConfiguration_DynamoDB) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreConfiguration_DynamoDB.ProtoReflect.Descriptor instead.
func (*StoreConfiguration_DynamoDB) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0, 3}
}

var File_store_proto protoreflect.FileDescriptor

var file_store_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x6e, 0x22, 0xf4, 0x08, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x70,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x6e, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0e, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x0f,
	0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x62, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x44, 0x42, 0x48, 0x00, 0x52, 0x0e, 0x69,
	0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x46, 0x0a,
	0x0c, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x6e, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x73, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x0f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64,
	0x62, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x79, 0x6e, 0x61,
	0x6d, 0x6f, 0x44, 0x42, 0x48, 0x00, 0x52, 0x0e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x55, 0x0a, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x48, 0x01, 0x52, 0x11, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x55, 0x0a,
	0x12, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x62, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x44, 0x42, 0x48,
	0x01, 0x52, 0x11, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x62, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x48, 0x01, 0x52, 0x0e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x73, 0x12, 0x55, 0x0a, 0x12, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62, 0x5f, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x79, 0x6e, 0x61,
	0x6d, 0x6f, 0x44, 0x42, 0x48, 0x01, 0x52, 0x11, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x1a, 0xca, 0x01, 0x0a, 0x08, 0x50, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x26,
	0x0a, 0x0f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x63, 0x61, 0x6e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x07, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x1a,
	0x9e, 0x01, 0x0a, 0x08, 0x49, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x44, 0x42, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x1a, 0x0a, 0x0a, 0x08, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x44, 0x42, 0x42, 0x0d, 0x0a, 0x0b,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xa0, 0x01,
	0x0a, 0x0d, 0x49, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x44, 0x42, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6c, 0x75, 0x78,
	0x44, 0x42, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x66, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x6e, 0x2f, 0x67,
	0x65, 0x6e, 0x3b, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_proto_rawDescOnce sync.Once
	file_store_proto_rawDescData = file_store_proto_rawDesc
)

func file_store_proto_rawDescGZIP() []byte {
	file_store_proto_rawDescOnce.Do(func() {
		file_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_proto_rawDescData)
	})
	return file_store_proto_rawDescData
}

var file_store_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_store_proto_goTypes = []interface{}{
	(*StoreConfiguration)(nil),          // 0: chronon.StoreConfiguration
	(*InfluxDBPoint)(nil),               // 1: chronon.InfluxDBPoint
	(*StoreConfiguration_Postgres)(nil), // 2: chronon.StoreConfiguration.Postgres
	(*StoreConfiguration_Redis)(nil),    // 3: chronon.StoreConfiguration.Redis
	(*StoreConfiguration_InfluxDB)(nil), // 4: chronon.StoreConfiguration.InfluxDB
	(*StoreConfiguration_DynamoDB)(nil), // 5: chronon.StoreConfiguration.DynamoDB
	nil,                                 // 6: chronon.InfluxDBPoint.TagsEntry
}
var file_store_proto_depIdxs = []int32{
	2, // 0: chronon.StoreConfiguration.postgres_events:type_name -> chronon.StoreConfiguration.Postgres
	4, // 1: chronon.StoreConfiguration.influxdb_events:type_name -> chronon.StoreConfiguration.InfluxDB
	3, // 2: chronon.StoreConfiguration.redis_events:type_name -> chronon.StoreConfiguration.Redis
	5, // 3: chronon.StoreConfiguration.dynamodb_events:type_name -> chronon.StoreConfiguration.DynamoDB
	2, // 4: chronon.StoreConfiguration.postgres_snapshots:type_name -> chronon.StoreConfiguration.Postgres
	4, // 5: chronon.StoreConfiguration.influxdb_snapshots:type_name -> chronon.StoreConfiguration.InfluxDB
	3, // 6: chronon.StoreConfiguration.redis_snapshots:type_name -> chronon.StoreConfiguration.Redis
	5, // 7: chronon.StoreConfiguration.dynamodb_snapshots:type_name -> chronon.StoreConfiguration.DynamoDB
	6, // 8: chronon.InfluxDBPoint.tags:type_name -> chronon.InfluxDBPoint.TagsEntry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_store_proto_init() }
func file_store_proto_init() {
	if File_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfluxDBPoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreConfiguration_Postgres); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreConfiguration_Redis); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreConfiguration_InfluxDB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreConfiguration_DynamoDB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_store_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StoreConfiguration_PostgresEvents)(nil),
		(*StoreConfiguration_InfluxdbEvents)(nil),
		(*StoreConfiguration_RedisEvents)(nil),
		(*StoreConfiguration_DynamodbEvents)(nil),
		(*StoreConfiguration_PostgresSnapshots)(nil),
		(*StoreConfiguration_InfluxdbSnapshots)(nil),
		(*StoreConfiguration_RedisSnapshots)(nil),
		(*StoreConfiguration_DynamodbSnapshots)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_proto_goTypes,
		DependencyIndexes: file_store_proto_depIdxs,
		MessageInfos:      file_store_proto_msgTypes,
	}.Build()
	File_store_proto = out.File
	file_store_proto_rawDesc = nil
	file_store_proto_goTypes = nil
	file_store_proto_depIdxs = nil
}

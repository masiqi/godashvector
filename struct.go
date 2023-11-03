package godashvector

// BaseResponse base response
type BaseResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

// Doc struct
type Doc struct {
	ID     string                 `json:"id"`
	Vector interface{}            `json:"vector"`
	Fields map[string]interface{} `json:"fields"`
	Score  float64                `json:"score"`
}

// CollectionMeta struct
type CollectionMeta struct {
	Name       string            `json:"name"`
	Dimension  int               `json:"dimension"`
	Dtype      string            `json:"dtype"`
	Metric     string            `json:"metric"`
	Status     Status            `json:"status"`
	Fields     map[string]string `json:"fields"`
	Partitions map[string]Status `json:"partitions"`
}

// CollectionStats struct
type CollectionStats struct {
	TotalDocCount     string                    `json:"total_doc_count"`
	IndexCompleteness float64                   `json:"index_completeness"`
	Partitions        map[string]PartitionStats `json:"partitions"`
}

// PartitionStats struct
type PartitionStats struct {
	TotalDocCount string `json:"total_doc_count"`
}

// Status enum
type Status string

// FieldDataType field data类型
type FieldDataType map[string]interface{}

// CollectionCreateRequest 创建Collection请求
type CollectionCreateRequest struct {
	Name         string            `json:"name"`
	Dimension    int               `json:"dimension"`
	Dtype        string            `json:"dtype"`
	Metric       string            `json:"metric"`
	FieldsSchema map[string]string `json:"fields_schema"`
}

// CollectionCreateResponse 创建Collection响应
type CollectionCreateResponse struct {
	BaseResponse
}

// CollectionDetailRequest 描述Collection请求
type CollectionDetailRequest struct {
	Name string `json:"name"`
}

// CollectionDetailResponse 描述Collection响应
type CollectionDetailResponse struct {
	BaseResponse
	Output CollectionMeta `json:"output"`
}

// CollectionListRequest 列Collection请求
type CollectionListRequest struct {
}

// CollectionListResponse 列Collection响应
type CollectionListResponse struct {
	BaseResponse
	Output []string `json:"output"`
}

// CollectionStatsRequest 统计Collection请求
type CollectionStatsRequest struct {
	Name string `json:"name"`
}

// CollectionStatsResponse 统计Collection响应
type CollectionStatsResponse struct {
	BaseResponse
	Output CollectionStats `json:"output"`
}

// CollectionDeleteRequest 删除Collection请求
type CollectionDeleteRequest struct {
	Name string `json:"name"`
}

// CollectionDeleteResponse 删除Collection响应
type CollectionDeleteResponse struct {
	BaseResponse
}

// PartitionCreateRequest 创建Partition请求
type PartitionCreateRequest struct {
	CollectionName string `json:"collection_name"`
	Name           string `json:"name"`
}

// PartitionCreateResponse 创建Partition响应
type PartitionCreateResponse struct {
	BaseResponse
}

// PartitionDetailRequest 描述Partition请求
type PartitionDetailRequest struct {
	CollectionName string `json:"collection_name"`
	Name           string `json:"name"`
}

// PartitionDetailResponse 描述Partition响应
type PartitionDetailResponse struct {
	BaseResponse
	Output Status `json:"output"`
}

// PartitionListRequest 列Partition请求
type PartitionListRequest struct {
	CollectionName string `json:"collection_name"`
}

// PartitionListResponse 列Partition响应
type PartitionListResponse struct {
	BaseResponse
	Output []string `json:"output"`
}

// PartitionStatsRequest 统计Partition请求
type PartitionStatsRequest struct {
	CollectionName string `json:"collection_name"`
	Name           string `json:"name"`
}

// PartitionStatsResponse 统计Partition响应
type PartitionStatsResponse struct {
	BaseResponse
	Output PartitionStats `json:"output"`
}

// PartitionDeleteRequest 删除Partition请求
type PartitionDeleteRequest struct {
	CollectionName string `json:"collection_name"`
	Name           string `json:"name"`
}

// PartitionDeleteResponse 删除Partition响应
type PartitionDeleteResponse struct {
	BaseResponse
}

// DocumentCreateRequest 创建Document请求
type DocumentCreateRequest struct {
	CollectionName string  `json:"collection_name"`
	Partition      *string `json:"partition"`
	Docs           []Doc   `json:"docs"`
}

// DocumentCreateResponse 创建Document响应
type DocumentCreateResponse struct {
	BaseResponse
}

// DocumentQueryRequest 查询Document请求
type DocumentQueryRequest struct {
	CollectionName string     `json:"collection_name"`
	Partition      *string    `json:"partition"`
	Vector         *[]float64 `json:"vector"`
	ID             *string    `json:"id"`
	TopK           *int       `json:"topk"`
	Filter         *string    `json:"filter"`
	IncludeVector  *bool      `json:"include_vector"`
	OutputFields   *[]string  `json:"output_fields"`
}

// DocumentQueryResponse 查询Document响应
type DocumentQueryResponse struct {
	BaseResponse
	Output []Doc `json:"output"`
}

// DocumentUpsertRequest 插入或更新Document请求
type DocumentUpsertRequest struct {
	CollectionName string  `json:"collection_name"`
	Partition      *string `json:"partition"`
	Docs           []Doc   `json:"docs"`
}

// DocumentUpsertResponse 插入或更新Document响应
type DocumentUpsertResponse struct {
	BaseResponse
}

// DocumentUpdateRequest 更新Document请求
type DocumentUpdateRequest struct {
	CollectionName string  `json:"collection_name"`
	Partition      *string `json:"partition"`
	Docs           []Doc   `json:"docs"`
}

// DocumentUpdateResponse 更新Document响应
type DocumentUpdateResponse struct {
	BaseResponse
}

// DocumentGetRequest 获取Document请求
type DocumentGetRequest struct {
	CollectionName string   `json:"collection_name"`
	Partition      *string  `json:"partition"`
	IDs            []string `json:"ids"`
}

// DocumentGetResponse 获取Document响应
type DocumentGetResponse struct {
	BaseResponse
	Output map[string]Doc `json:"output"`
}

// DocumentDeleteRequest 删除Document请求
type DocumentDeleteRequest struct {
	CollectionName string   `json:"collection_name"`
	Partition      *string  `json:"partition"`
	IDs            []string `json:"ids"`
	DeleteAll      *bool    `json:"delete_all"`
}

// DocumentDeleteResponse 删除Document响应
type DocumentDeleteResponse struct {
	BaseResponse
}

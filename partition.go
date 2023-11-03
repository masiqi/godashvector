package godashvector

// CreatePartition 创建Partition
func (c *Client) CreatePartition(req *PartitionCreateRequest) (resp *PartitionCreateResponse, err error) {
	_, err = c.RestyClient.R().
		SetBody(req).
		SetResult(&resp).
		Post("/v1/collections/" + req.CollectionName + "/partitions")
	return
}

// DescribePartition 描述Partition
func (c *Client) DescribePartition(req *PartitionDetailRequest) (resp *PartitionDetailResponse, err error) {
	_, err = c.RestyClient.R().
		SetResult(&resp).
		Get("/v1/collections/" + req.CollectionName + "/partitions/" + req.Name)
	return
}

// ListPartition 列Partition
func (c *Client) ListPartition(req *PartitionListRequest) (resp *PartitionListResponse, err error) {
	_, err = c.RestyClient.R().
		SetResult(&resp).
		Get("/v1/collections/" + req.CollectionName + "/partitions")
	return
}

// StatsPartition 统计Partition
func (c *Client) StatsPartition(req *PartitionStatsRequest) (resp *PartitionStatsResponse, err error) {
	_, err = c.RestyClient.R().
		SetResult(&resp).
		Get("/v1/collections/" + req.CollectionName + "/partitions/" + req.Name + "/stats")
	return
}

// DeletePartition 删除Partition
func (c *Client) DeletePartition(req *PartitionDeleteRequest) (resp *PartitionDeleteResponse, err error) {
	_, err = c.RestyClient.R().
		SetResult(&resp).
		Delete("/v1/collections/" + req.CollectionName + "/partitions/" + req.Name)
	return
}

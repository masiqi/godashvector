package godashvector

// CreateCollection 创建Collection
func (c *Client) CreateCollection(req *CollectionCreateRequest) (resp *CollectionCreateResponse, err error) {
	_, err = c.RestyClient.R().
		SetBody(req).
		SetResult(&resp).
		Post("/v1/collections")

	return
}

// DescribeCollection 描述Collection
func (c *Client) DescribeCollection(req *CollectionDetailRequest) (resp *CollectionDetailResponse, err error) {
	_, err = c.RestyClient.R().
		SetResult(&resp).
		Get("/v1/collections/" + req.Name)

	return
}

// ListCollection 列Collection
func (c *Client) ListCollection(req *CollectionListRequest) (resp *CollectionListResponse, err error) {
	_, err = c.RestyClient.R().
		SetResult(&resp).
		Get("/v1/collections")
	return
}

// StatsCollection 统计Collection
func (c *Client) StatsCollection(req *CollectionStatsRequest) (resp *CollectionStatsResponse, err error) {
	_, err = c.RestyClient.R().
		SetResult(&resp).
		Get("/v1/collections/" + req.Name + "/stats")
	return
}

// DeleteCollection 删除Collection
func (c *Client) DeleteCollection(req *CollectionDeleteRequest) (resp *CollectionDeleteResponse, err error) {
	_, err = c.RestyClient.R().
		SetResult(&resp).
		Delete("/v1/collections/" + req.Name)
	return nil, nil
}

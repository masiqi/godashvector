package godashvector

import (
	"strings"
)

// CreateDocument 创建文档
func (c *Client) CreateDocument(req *DocumentCreateRequest) (resp *DocumentCreateResponse, err error) {
	_, err = c.RestyClient.R().
		SetBody(req).
		SetResult(&resp).
		Post("/v1/collections/" + req.CollectionName + "/docs")
	return
}

// QueryDocument 查询文档
func (c *Client) QueryDocument(req *DocumentQueryRequest) (resp *DocumentQueryResponse, err error) {
	_, err = c.RestyClient.R().
		SetBody(req).
		SetResult(&resp).
		Post("/v1/collections/" + req.CollectionName + "/query")
	return
}

// UpsertDocument 插入或更新文档
func (c *Client) UpsertDocument(req *DocumentUpsertRequest) (resp *DocumentUpsertResponse, err error) {
	_, err = c.RestyClient.R().
		SetBody(req).
		SetResult(&resp).
		Post("/v1/collections/" + req.CollectionName + "/docs/upsert")
	return
}

// UpdateDocument 更新文档
func (c *Client) UpdateDocument(req *DocumentUpdateRequest) (resp *DocumentUpdateResponse, err error) {
	_, err = c.RestyClient.R().
		SetBody(req).
		SetResult(&resp).
		Put("/v1/collections/" + req.CollectionName + "/docs")
	return
}

// GetDocument 获取文档
func (c *Client) GetDocument(req *DocumentGetRequest) (resp *DocumentGetResponse, err error) {
	_, err = c.RestyClient.R().
		SetQueryParam("ids", strings.Join(req.IDs, ",")).
		SetQueryParam("partition", *req.Partition).
		SetResult(&resp).
		Get("/v1/collections/" + req.CollectionName + "/docs")
	return
}

// DeleteDocument 删除文档
func (c *Client) DeleteDocument(req *DocumentDeleteRequest) (resp *DocumentDeleteResponse, err error) {
	_, err = c.RestyClient.R().
		SetBody(req).
		SetResult(&resp).
		Delete("/v1/collections/" + req.CollectionName + "/docs")
	return
}

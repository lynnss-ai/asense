// Package difyknow
// @File    : difyknow.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/9 18:24
// @Desc    :
package difyknow

import "context"

const (
	createDocumentByTextURL   = "/datasets/%s/document/create_by_text"     //通过文本创建文档 [dataset_id]
	createDocumentByFileURL   = "/datasets/%s/document/create_by_file"     //通过文件创建文档 [dataset_id]
	updateDocumentByTextURL   = "/datasets/%s/documents/%s/update_by_text" //通过文本更新文档 [dataset_id / document_id]
	updateDocumentByFileURL   = "/datasets/%s/documents/%s/update_by_file" //通过文件更新文档 [dataset_id / document_id]
	deleteDocumentURL         = "/datasets/%s/documents/%s"                //删除文档 [dataset_id / document_id]
	createEmptyDatasetsURL    = "/datasets"                                //创建空知识库
	deleteDatasetsURL         = "/datasets/%s"                             //删除知识库 [dataset_id]
	listDatasetsURL           = "/datasets"                                //知识库列表
	listDatasetsDocumentsURL  = "/datasets/%s/documents"                   //知识库文档列表 [dataset_id]
	newDocumentSegmentsURL    = "/datasets/%s/documents/%s/segments"       //新增分段 [dataset_id / document_id]
	queryDocumentSegmentsURL  = "/datasets/%s/documents/%s/segments"       //查询文档分段 [dataset_id / document_id]
	delDocumentSegmentsURL    = "/datasets/%s/documents/%s/segments/%s"    //删除文档分段 [dataset_id / document_id / segment_id]
	updateDocumentSegmentsURL = "/datasets/%s/documents/%s/segments/%s"    //更新文档分段 [dataset_id / document_id / segment_id]
)

type (
	DifyKnowModel struct {
		BaseURL string
		ApiKey  string
	}

	// [01|02]  CreateDocumentResp
	CreateDocumentResp struct {
		Document struct {
			Id             string `json:"id"`
			Position       int    `json:"position"`
			DataSourceType string `json:"data_source_type"`
			DataSourceInfo struct {
				UploadFileId string `json:"upload_file_id"`
			} `json:"data_source_info"`
			DatasetProcessRuleId string      `json:"dataset_process_rule_id"`
			Name                 string      `json:"name"`
			CreatedFrom          string      `json:"created_from"`
			CreatedBy            string      `json:"created_by"`
			CreatedAt            int         `json:"created_at"`
			Tokens               int         `json:"tokens"`
			IndexingStatus       string      `json:"indexing_status"`
			Error                interface{} `json:"error"`
			Enabled              bool        `json:"enabled"`
			DisabledAt           interface{} `json:"disabled_at"`
			DisabledBy           interface{} `json:"disabled_by"`
			Archived             bool        `json:"archived"`
			DisplayStatus        string      `json:"display_status"`
			WordCount            int         `json:"word_count"`
			HitCount             int         `json:"hit_count"`
			DocForm              string      `json:"doc_form"`
		} `json:"document"`
		Batch string `json:"batch"`
	}

	//
	CreateDatasets struct {
		Id                     string      `json:"id"`
		Name                   string      `json:"name"`
		Description            interface{} `json:"description"`
		Provider               string      `json:"provider"`
		Permission             string      `json:"permission"`
		DataSourceType         interface{} `json:"data_source_type"`
		IndexingTechnique      interface{} `json:"indexing_technique"`
		AppCount               int         `json:"app_count"`
		DocumentCount          int         `json:"document_count"`
		WordCount              int         `json:"word_count"`
		CreatedBy              string      `json:"created_by"`
		CreatedAt              int         `json:"created_at"`
		UpdatedBy              string      `json:"updated_by"`
		UpdatedAt              int         `json:"updated_at"`
		EmbeddingModel         interface{} `json:"embedding_model"`
		EmbeddingModelProvider interface{} `json:"embedding_model_provider"`
		EmbeddingAvailable     interface{} `json:"embedding_available"`
	}

	//
	DocumentSegmentsResp struct {
		Data []struct {
			Id            string      `json:"id"`
			Position      int         `json:"position"`
			DocumentId    string      `json:"document_id"`
			Content       string      `json:"content"`
			Answer        string      `json:"answer"`
			WordCount     int         `json:"word_count"`
			Tokens        int         `json:"tokens"`
			Keywords      []string    `json:"keywords"`
			IndexNodeId   string      `json:"index_node_id"`
			IndexNodeHash string      `json:"index_node_hash"`
			HitCount      int         `json:"hit_count"`
			Enabled       bool        `json:"enabled"`
			DisabledAt    interface{} `json:"disabled_at"`
			DisabledBy    interface{} `json:"disabled_by"`
			Status        string      `json:"status"`
			CreatedBy     string      `json:"created_by"`
			CreatedAt     int         `json:"created_at"`
			IndexingAt    int         `json:"indexing_at"`
			CompletedAt   int         `json:"completed_at"`
			Error         interface{} `json:"error"`
			StoppedAt     interface{} `json:"stopped_at"`
		} `json:"data"`
		DocForm string `json:"doc_form"`
	}

	DifyKnowInterface interface {
		CreateDocumentByText(ctx context.Context, datasetID string, bodyArgs []byte) (*CreateDocumentResp, error)                                          //通过文本创建文档
		CreateDocumentByFile(ctx context.Context, datasetID string, bodyArgs []byte) (*CreateDocumentResp, error)                                          //通过文件创建文档
		UpdateDocumentByText(ctx context.Context, datasetID string, documentID string, bodyArgs []byte) (*CreateDocumentResp, error)                       //通过文本更新文档
		UpdateDocumentByFile(ctx context.Context, datasetID string, documentID string, bodyArgs []byte) (*CreateDocumentResp, error)                       //通过文件更新文档
		DeleteDocument(ctx context.Context, datasetID string, documentID string) error                                                                     //删除文档
		NewDocumentSegments(ctx context.Context, datasetID string, documentID string, bodyArgs []byte) (*DocumentSegmentsResp, error)                      //新增文档分段
		QueryDocumentSegments(ctx context.Context, datasetID string, documentID string, queryArgs map[string]interface{}) (*DocumentSegmentsResp, error)   //查询文档分段
		DeleteDocumentSegments(ctx context.Context, datasetID string, documentID string, segmentID string) error                                           //删除文档分段
		UpdateDocumentSegments(ctx context.Context, datasetID string, documentID string, segmentID string, bodyArgs []byte) (*DocumentSegmentsResp, error) //更新文档分段
		CreateEmptyDatasetsURL(ctx context.Context, knowName string) (*CreateDatasets, error)                                                              //创建空知识库
		DeleteDatasets(ctx context.Context, datasetID string) error                                                                                        //删除知识库
	}
)

func NewDifyKnowModel(baseUrl, apiKey string) *DifyKnowModel {
	return &DifyKnowModel{
		BaseURL: baseUrl,
		ApiKey:  apiKey,
	}
}

package godashvector_test

import (
	"fmt"
	"github.com/masiqi/godashvector"
	"math/rand"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	APIKey := os.Getenv("DASHVECTOR_API_KEY")
	Endpoint := os.Getenv("DASHVECTOR_ENDPOINT")
	if APIKey == "" {
		t.Fatal("DASHVECTOR_API_KEY is empty")
	}

	if Endpoint == "" {
		t.Fatal("DASHVECTOR_ENDPOINT is empty")
	}

	client := godashvector.NewClient(APIKey, Endpoint)

	dim := 1536

	// 生成随机浮点数slice
	floats := make([]float64, dim)

	// 填充随机浮点数
	for i := range floats {
		floats[i] = rand.Float64()
	}
	t.Run("collection_list", func(t *testing.T) {
		resp, err := client.ListCollection(&godashvector.CollectionListRequest{})
		fmt.Printf("collection_list: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	/*
		t.Run("collection_create", func(t *testing.T) {
			resp, err := client.CreateCollection(&godashvector.CollectionCreateRequest{
				Name:      "test",
				Dimension: 1536,
				Metric:    "cosine",
			})
			fmt.Printf("collection_create: \n%v\n", resp)
			if err != nil {
				t.Error(err)
			}
		})
	*/

	t.Run("collection_describe", func(t *testing.T) {
		resp, err := client.DescribeCollection(&godashvector.CollectionDetailRequest{
			Name: "test",
		})
		fmt.Printf("collection_describe: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("collection_stats", func(t *testing.T) {
		resp, err := client.StatsCollection(&godashvector.CollectionStatsRequest{
			Name: "test",
		})
		fmt.Printf("collection_stats: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("partition_create", func(t *testing.T) {
		resp, err := client.CreatePartition(&godashvector.PartitionCreateRequest{
			CollectionName: "test",
			Name:           "test",
		})
		fmt.Printf("partition_create: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("partition_list", func(t *testing.T) {
		resp, err := client.ListPartition(&godashvector.PartitionListRequest{
			CollectionName: "test",
		})
		fmt.Printf("partition_list: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("document_create", func(t *testing.T) {

		doc := godashvector.Doc{
			ID:     "2",
			Vector: floats,
			Fields: map[string]interface{}{
				"age":  20,
				"name": "张三",
			},
		}
		partition := "test"
		resp, err := client.CreateDocument(&godashvector.DocumentCreateRequest{
			CollectionName: "test",
			Partition:      &partition,
			Docs:           []godashvector.Doc{doc},
		})
		fmt.Printf("document_create: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("document_upsert", func(t *testing.T) {
		partition := "test"
		doc := godashvector.Doc{
			ID:     "2",
			Vector: floats,
			Fields: map[string]interface{}{
				"age":  20,
				"name": "李四",
			},
		}
		resp, err := client.UpsertDocument(&godashvector.DocumentUpsertRequest{
			CollectionName: "test",
			Partition:      &partition,
			Docs:           []godashvector.Doc{doc},
		})
		fmt.Printf("document_upsert: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("document_update", func(t *testing.T) {
		partition := "test"
		doc := godashvector.Doc{
			ID:     "2",
			Vector: floats,
			Fields: map[string]interface{}{
				"age":  20,
				"name": "王五",
			},
		}
		resp, err := client.UpdateDocument(&godashvector.DocumentUpdateRequest{
			CollectionName: "test",
			Partition:      &partition,
			Docs:           []godashvector.Doc{doc},
		})
		fmt.Printf("document_update: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("document_get", func(t *testing.T) {
		partition := "test"
		resp, err := client.GetDocument(&godashvector.DocumentGetRequest{
			CollectionName: "test",
			Partition:      &partition,
			IDs:            []string{"2", "3"},
		})
		fmt.Printf("document_get: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("document_query", func(t *testing.T) {
		partition := "test"
		resp, err := client.QueryDocument(&godashvector.DocumentQueryRequest{
			CollectionName: "test",
			Partition:      &partition,
			Vector:         &floats,
		})
		fmt.Printf("document_query: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("document_delete", func(t *testing.T) {
		partition := "test"
		resp, err := client.DeleteDocument(&godashvector.DocumentDeleteRequest{
			CollectionName: "test",
			Partition:      &partition,
			IDs:            []string{"2", "3"},
		})
		fmt.Printf("document_delete: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("partition_describe", func(t *testing.T) {
		resp, err := client.DescribePartition(&godashvector.PartitionDetailRequest{
			CollectionName: "test",
			Name:           "test",
		})
		fmt.Printf("partition_describe: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("partition_stats", func(t *testing.T) {
		resp, err := client.StatsPartition(&godashvector.PartitionStatsRequest{
			CollectionName: "test",
			Name:           "test",
		})
		fmt.Printf("partition_stats: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("partition_delete", func(t *testing.T) {
		resp, err := client.DeletePartition(&godashvector.PartitionDeleteRequest{
			CollectionName: "test",
			Name:           "test",
		})
		fmt.Printf("partition_delete: \n%v\n", resp)
		if err != nil {
			t.Error(err)
		}
	})

	/*
		t.Run("collection_delete", func(t *testing.T) {
			resp, err := client.DeleteCollection(&godashvector.CollectionDeleteRequest{
				Name: "test",
			})
			fmt.Printf("collection_delete: \n%v\n", resp)
			if err != nil {
				t.Error(err)
			}
		})
	*/
}

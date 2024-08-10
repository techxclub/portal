package azure

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/config"
	"github.com/techx/portal/utils"
)

const (
	containerNameLogos = "logos"
)

type Client interface {
	UploadLogo(ctx context.Context, params UploadImageParams) error
	FetchLogos(ctx context.Context, params FetchImageParams) ([]string, error)
}

type azureClient struct {
	client *azblob.Client
}

type UploadImageParams struct {
	CompanyID   int64  `json:"company_id"`
	CompanyName string `json:"company_name"`
	LogoFile    multipart.File
	LogoHeader  *multipart.FileHeader
}

type FetchImageParams struct {
	CompanyID int64 `json:"company_id"`
}

func NewAzureClient(azureCfg config.AzureStorage) Client {
	client, err := azblob.NewClientFromConnectionString(azureCfg.ConnectionString, nil)
	if err != nil {
		log.Err(err).Msg("failed to create azure client")
		panic(fmt.Sprintf("failed to create azure client: %v", err))
	}

	return &azureClient{
		client: client,
	}
}

func (c *azureClient) UploadLogo(ctx context.Context, params UploadImageParams) error {
	containerClient := c.client.ServiceClient().NewContainerClient(containerNameLogos)

	randomID, _ := utils.GenerateRandomNumber(4)
	fileExtension := filepath.Ext(params.LogoHeader.Filename)
	blobName := fmt.Sprintf("logos/%d/%s_%s%s", params.CompanyID, params.CompanyName, randomID, fileExtension)
	blobClient := containerClient.NewBlockBlobClient(blobName)

	_, err := blobClient.UploadStream(ctx, params.LogoFile, &azblob.UploadStreamOptions{})
	if err != nil {
		return fmt.Errorf("failed to upload image: %v", err)
	}

	return nil
}

func (c *azureClient) FetchLogos(ctx context.Context, params FetchImageParams) ([]string, error) {
	containerClient := c.client.ServiceClient().NewContainerClient(containerNameLogos)
	prefix := fmt.Sprintf("logos/%d/", params.CompanyID)
	pager := containerClient.NewListBlobsFlatPager(&azblob.ListBlobsFlatOptions{
		Prefix: &prefix,
	})

	var urls []string
	for pager.More() {
		resp, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch image URLs: %v", err)
		}
		for _, blob := range resp.Segment.BlobItems {
			blobURL := containerClient.NewBlobClient(*blob.Name).URL()
			urls = append(urls, blobURL)
		}
	}

	return urls, nil
}

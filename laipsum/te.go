import (
	"context"
	"fmt"
	"io"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func indexDetails(w io.Writer, db string) error {
	ctx := context.Background()
	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		return err
	}
	defer adminClient.Close()

	req := &adminpb.GetIndexRequest{
		Index: fmt.Sprintf("%s/indexes/SingersByAlbum", db),
	}
	index, err := adminClient.GetIndex(ctx, req)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return fmt.Errorf("index not found: %v", err)
		}
		return err
	}
	fmt.Fprintf(w, "Index %s:\n", index.Name)
	fmt.Fprintf(w, "\tColumns: %v\n", index.Columns)
	fmt.Fprintf(w, "\tStoring: %v\n", index.Storing)
	fmt.Fprintf(w, "\tIs unique: %v\n", index.Unique)
	return nil
}
  

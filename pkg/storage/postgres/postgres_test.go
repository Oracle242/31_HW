package postgres

import (
	"reflect"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestStorage_Tasks(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		taskID   int
		authorID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.Tasks(tt.args.taskID, tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Tasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.Tasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
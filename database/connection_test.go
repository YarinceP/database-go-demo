package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"testing"
)

func TestNewConnectorRepositoryImplement(t *testing.T) {
	type args struct {
		connectionString string
	}
	tests := []struct {
		name string
		args args
		want ConnectorRepositoryImplement
	}{
		{
			name: "ok",
			args: args{
				connectionString: "test",
			},
			want: ConnectorRepositoryImplement{
				ConnectionString: "test",
				db:               nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConnectorRepositoryImplement(tt.args.connectionString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnectorRepositoryImplement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func NewMockDB() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestConnectorRepositoryImplement_Connect(t *testing.T) {
	//defaultConnectionString := "root:@tcp(localhost:3306)/db_lib_go"
	mockDB, _ := NewMockDB()
	defer mockDB.Close()

	type fields struct {
		ConnectionString string
		db               *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    *sql.DB
		wantErr bool
	}{
		{
			name: "Open Connection Fail",
			fields: fields{
				ConnectionString: "fakeConnectionString",
				db:               nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Pinging database fail",
			fields: fields{
				ConnectionString: "",
				db:               mockDB,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConnectorRepositoryImplement{
				ConnectionString: tt.fields.ConnectionString,
				db:               tt.fields.db,
			}
			got, err := c.Connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect() got = %v, want %v", got, tt.want)
			}

		})
	}

}

func TestConnectorRepositoryImplement_Disconnect(t *testing.T) {
	mockDB, _ := NewMockDB()
	defer mockDB.Close()
	type fields struct {
		ConnectionString string
		db               *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "No active database connection",
			fields: fields{
				ConnectionString: "",
				db:               nil,
			},
			wantErr: true,
		},
		{
			name: "Error closing the database connection",
			fields: fields{
				ConnectionString: "",
				db:               mockDB,
			},
			wantErr: true,
		},
		{
			name: "Successfully disconnected",
			fields: fields{
				ConnectionString: "",
				db:               mockDB,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConnectorRepositoryImplement{
				ConnectionString: tt.fields.ConnectionString,
				db:               tt.fields.db,
			}
			if err := c.Disconnect(); (err != nil) != tt.wantErr {
				t.Errorf("Disconnect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

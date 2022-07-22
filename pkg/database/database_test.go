package database

import "testing"

func TestMySQLConfig_DSN(t *testing.T) {
	type fields struct {
		Host     string
		Database string
		Port     int
		Username string
		Password string
		Options  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "happy case",
			fields: fields{
				Host:     "127.0.0.1",
				Database: "sample",
				Port:     3306,
				Username: "default",
				Password: "secret",
				Options:  "",
			},
			want: "default:secret@tcp(127.0.0.1:3306)/sample",
		}, {
			name: "password with special characters",
			fields: fields{
				Host:     "127.0.0.1",
				Database: "sample",
				Port:     3306,
				Username: "default",
				Password: "secret@!(/:.1234",
				Options:  "",
			},
			want: "default:secret@!(/:.1234@tcp(127.0.0.1:3306)/sample",
		}, {
			name: "with options case having ?",
			fields: fields{
				Host:     "127.0.0.1",
				Database: "sample",
				Port:     3306,
				Username: "default",
				Password: "secret",
				Options:  "?parseTime=true",
			},
			want: "default:secret@tcp(127.0.0.1:3306)/sample?parseTime=true",
		}, {
			name: "with options case not having ?",
			fields: fields{
				Host:     "127.0.0.1",
				Database: "sample",
				Port:     3306,
				Username: "default",
				Password: "secret",
				Options:  "parseTime=true",
			},
			want: "default:secret@tcp(127.0.0.1:3306)/sample?parseTime=true",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := MySQLConfig{
				Host:     tt.fields.Host,
				Database: tt.fields.Database,
				Port:     tt.fields.Port,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				Options:  tt.fields.Options,
			}
			if got := c.DSN(); got != tt.want {
				t.Errorf("Config.DSN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySQLConfig_String(t *testing.T) {
	type fields struct {
		MySQLConfig MySQLConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{{
		name: "happy case",
		fields: fields{MySQLConfig{
			Host:     "127.0.0.1",
			Database: "sample",
			Port:     3306,
			Username: "default",
			Password: "secret",
			Options:  "",
		}},
		want: "mysql://default:secret@tcp(127.0.0.1:3306)/sample",
	}, {
		name: "password with special characters",
		fields: fields{MySQLConfig{
			Host:     "127.0.0.1",
			Database: "sample",
			Port:     3306,
			Username: "default",
			Password: "secret@!(/:.1234",
			Options:  "",
		}},
		want: "mysql://default:secret@!(/:.1234@tcp(127.0.0.1:3306)/sample",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.MySQLConfig
			if got := c.String(); got != tt.want {
				t.Errorf("MySQLConfig.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

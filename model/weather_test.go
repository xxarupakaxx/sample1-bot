package model

import (
	"database/sql"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/xxarupakaxx/sample1-bot/domain"
	"reflect"
	"testing"
)

func TestConvertTelop(t *testing.T) {
	type args struct {
		telop string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertTelop(tt.args.telop); got != tt.want {
				t.Errorf("ConvertTelop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnect(t *testing.T) {
	tests := []struct {
		name string
		want *sql.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DBConnect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBConnect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRestoInfo(t *testing.T) {
	type args struct {
		lat string
		lng string
	}
	tests := []struct {
		name string
		args args
		want []*linebot.CarouselColumn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRestoInfo(tt.args.lat, tt.args.lng); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRestoInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWeather(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want *domain.Weather
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWeather(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWeather() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrefCode(t *testing.T) {
	type args struct {
		cityName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrefCode(tt.args.cityName); got != tt.want {
				t.Errorf("PrefCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendRestoInfo(t *testing.T) {
	type args struct {
		bot *linebot.Client
		e   *linebot.Event
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSendWeather(t *testing.T) {
	type args struct {
		bot      *linebot.Client
		event    *linebot.Event
		cityName string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestUserListGET(t *testing.T) {
	type args struct {
		bot   *linebot.Client
		event *linebot.Event
	}
	tests := []struct {
		name string
		args args
		want []domain.Video
	}{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserListGET(tt.args.bot, tt.args.event); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserListGET() = %v, want %v", got, tt.want)
			}
		})
	}
}

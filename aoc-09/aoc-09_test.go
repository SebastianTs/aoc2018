package main

import "testing"

func Test_play(t *testing.T) {
	type args struct {
		players     int
		last_marble int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"#1", args{10, 1618}, 8317},
		{"#2", args{13, 7999}, 146373},
		{"#3", args{17, 1104}, 2764},
		{"#4", args{21, 6111}, 54718},
		{"#5", args{30, 5807}, 37305},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := play(tt.args.players, tt.args.last_marble); got != tt.want {
				t.Errorf("play() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_playList(t *testing.T) {
	type args struct {
		players     int
		last_marble int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"#1", args{10, 1618}, 8317},
		{"#2", args{13, 7999}, 146373},
		{"#3", args{17, 1104}, 2764},
		{"#4", args{21, 6111}, 54718},
		{"#5", args{30, 5807}, 37305},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playList(tt.args.players, tt.args.last_marble); got != tt.want {
				t.Errorf("playList() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_loadImage(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want <-chan Job
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadImage(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resize(t *testing.T) {
	type args struct {
		input <-chan Job
	}
	tests := []struct {
		name string
		args args
		want <-chan Job
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resize(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToGrayscale(t *testing.T) {
	type args struct {
		input <-chan Job
	}
	tests := []struct {
		name string
		args args
		want <-chan Job
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToGrayscale(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToGrayscale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_saveImage(t *testing.T) {
	type args struct {
		input <-chan Job
	}
	tests := []struct {
		name string
		args args
		want <-chan bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := saveImage(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("saveImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_benchmarkPipeline(t *testing.T) {
	type args struct {
		imagePaths []string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := benchmarkPipeline(tt.args.imagePaths)
			if (err != nil) != tt.wantErr {
				t.Errorf("benchmarkPipeline() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("benchmarkPipeline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

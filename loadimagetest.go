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


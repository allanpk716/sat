package sat

import (
	"testing"
)

func TestDefaultDict_Init(t *testing.T) {
	err := InitDefaultDict(SetPath("./word.txt"))
	if err != nil {
		t.Error(err)
		return
	}
	dicter := DefaultDict()
	t.Log(dicter.ReadReverse("一繁"))
	t.Log(dicter.Read("五"))
}

func BenchmarkDefaultDict_Read(b *testing.B) {
	dicter := DefaultDict()
	for i := 0; i < b.N; i++ {
		_ = dicter.Read("什麼")
	}
}

func BenchmarkDefaultDict_ReadReverse(b *testing.B) {
	dicter := DefaultDict()
	for i := 0; i < b.N; i++ {
		_ = dicter.ReadReverse("什么")
	}
}

func Test_defaultDict_IsChs(t *testing.T) {

	type args struct {
		s         string
		threshold float32
	}
	tests := []struct {
		name   string
		args   args
		want   bool
	}{
		{"cht", args{s: "我會再確認時間地點", threshold: 0.9}, false},
		{"cht", args{s: "我再打給你", threshold: 0.9}, false},
		{"cht", args{s: "他中午會到", threshold: 0.9}, false},
		{"chs", args{s: "太好了", threshold: 0.8}, true},
	}
	d := DefaultDict()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := d.IsChs(tt.args.s, tt.args.threshold); got != tt.want {
				t.Errorf("IsChs() = %v, want %v", got, tt.want)
			}
		})
	}
}
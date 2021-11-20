package main

import (
	"testing"
	"time"
)

func TestBlockchain_ProofOfWork(t *testing.T) {
	type fields struct {
		chain []Block
	}
	type args struct {
		previousProof uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint
	}{
		{
			name: "should return first proof of 533",
			fields: fields{
				NewBlockchain().Chain,
			},
			args: args{
				previousProof: 1,
			},
			want: 4216,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Blockchain{
				Chain: tt.fields.chain,
			}
			if got := b.ProofOfWork(tt.args.previousProof); got != tt.want {
				t.Errorf("Blockchain.ProofOfWork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hash(t *testing.T) {
	type args struct {
		block Block
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return hash for block",
			args: args{
				Block{
					Index:        1,
					Timestamp:    time.Date(2021, time.April, 1, 1, 1, 1, 1, time.Local),
					Proof:        4216,
					PreviousHash: "0",
				},
			},
			want: "a5e020f46f8eb7e3f3b5f1c5cb86fd270ae6cbd864cb45a4f948d8ae5bafb381",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.block); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

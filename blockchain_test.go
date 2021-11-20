package main

import (
	"testing"
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
				NewBlockchain().chain,
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
				chain: tt.fields.chain,
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
				NewBlockchain().GetPreviousBlock(),
			},
			want: "44136fa355b3678a1146ad16f7e8649e94fb4fc21fe77e8310c060f61caaff8a",
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

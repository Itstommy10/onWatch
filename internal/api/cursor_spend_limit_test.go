package api

import (
	"encoding/json"
	"testing"
)

func TestCursorSpendLimitUsage_AcceptsStringsAndNumbers(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want CursorSpendLimitUsage
	}{
		{"numbers", `{"totalSpend":150,"pooledLimit":2000,"pooledUsed":150,"pooledRemaining":1850,"limitType":"user"}`,
			CursorSpendLimitUsage{TotalSpend: 150, PooledLimit: 2000, PooledUsed: 150, PooledRemaining: 1850, LimitType: "user"}},
		{"numeric strings", `{"pooledLimit":"2000","pooledUsed":"0","individualLimit":"500"}`,
			CursorSpendLimitUsage{PooledLimit: 2000, IndividualLimit: 500}},
		{"empty and null", `{"pooledLimit":"","pooledUsed":null,"limitType":null}`,
			CursorSpendLimitUsage{}},
		{"float", `{"totalSpend":12.0}`, CursorSpendLimitUsage{TotalSpend: 12}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got CursorSpendLimitUsage
			if err := json.Unmarshal([]byte(tc.in), &got); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if got != tc.want {
				t.Fatalf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestCursorSpendLimitUsage_NestedInUsageResponse(t *testing.T) {
	var wrapper struct {
		SpendLimitUsage CursorSpendLimitUsage `json:"spendLimitUsage"`
	}
	if err := json.Unmarshal([]byte(`{"spendLimitUsage":{"pooledLimit":"0","totalSpend":0}}`), &wrapper); err != nil {
		t.Fatalf("unmarshal nested: %v", err)
	}
}

package db_test

import (
	"context"
	"testing"

	DB "github.com/PepsiKingIV/KeyValueDB/pkg/db"
)

func BenchmarkHash(b *testing.B) {
	ctxb := context.Background()

	db := DB.NewDB(ctxb, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range TCS {
			db.Hash(tc.Key, 100)
		}
	}
}

func BenchmarkPut(b *testing.B) {
	ctxb := context.Background()

	db := DB.NewDB(ctxb, false)
	ctxb.Done()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range TCS {
			db.Put(tc.Key, tc.Value)
		}
	}
}

func BenchmarkGet(b *testing.B) {
	ctxb := context.Background()
	db := DB.NewDB(ctxb, true)

	// for i := 0; i < b.N; i++ {
	// 	for _, tc := range TCS {
	// 		db.Put(tc.Key, tc.Value)
	// 	}
	// }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range TCS {
			db.Get(tc.Key)
		}
	}
	ctxb.Done()
}

func BenchmarkDelete(b *testing.B) {
	ctxb := context.Background()
	db := DB.NewDB(ctxb, true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range TCS {
			db.Delete(tc.Key)
		}
	}
	b.StopTimer()
	ctxb.Done()
}

type TestCase struct {
	ID    int
	Key   string
	Value string
}

var TCS []TestCase = []TestCase{
	{
		ID:    1,
		Key:   "keyAlphaBravo",
		Value: "The quick brown fox jumps over",
	},
	{
		ID:    2,
		Key:   "keyCharlieDelta",
		Value: "Lazy dog near the red house",
	},
	{
		ID:    3,
		Key:   "keyEchoFoxtrot",
		Value: "Zebras and lions share the savannah",
	},
	{
		ID:    4,
		Key:   "keyGolfHotel",
		Value: "Baking fresh bread early in the morning",
	},
	{
		ID:    5,
		Key:   "keyIndiaJuliet",
		Value: "Exploring the mountains in wintertime",
	},
	{
		ID:    6,
		Key:   "keyKiloLima",
		Value: "Waves crashing onto the sandy beach",
	},
	{
		ID:    7,
		Key:   "keyMikeNovember",
		Value: "A bright sunny day with clear skies",
	},
	{
		ID:    8,
		Key:   "keyOscarPapa",
		Value: "Coding a new program until late night",
	},
	{
		ID:    9,
		Key:   "keyQuebecRomeo",
		Value: "Reading an intriguing mystery novel",
	},
	{
		ID:    10,
		Key:   "keySierraTango",
		Value: "Listening to soothing classical music",
	},
	{
		ID:    11,
		Key:   "keyUniformVictor",
		Value: "Cooking delicious meals from scratch",
	},
	{
		ID:    12,
		Key:   "keyWhiskeyXray",
		Value: "A long journey through the forest path",
	},
	{
		ID:    13,
		Key:   "keyYankeeZulu",
		Value: "Building a cozy cabin in the woods",
	},
	{
		ID:    14,
		Key:   "keyAlphaBeta",
		Value: "Writing a captivating fantasy novel",
	},
	{
		ID:    15,
		Key:   "keyGammaDelta",
		Value: "Playing chess with a challenging opponent",
	},
	{
		ID:    16,
		Key:   "keyEpsilonZeta",
		Value: "Enjoying a peaceful evening by the lake",
	},
	{
		ID:    17,
		Key:   "keyEtaTheta",
		Value: "Discovering hidden treasures in the attic",
	},
	{
		ID:    18,
		Key:   "keyIotaKappa",
		Value: "Learning a new language for travel",
	},
	{
		ID:    19,
		Key:   "keyLambdaMu",
		Value: "Gardening in the spring sunshine",
	},
	{
		ID:    20,
		Key:   "keyNuXi",
		Value: "Crafting handmade gifts for friends",
	},
}

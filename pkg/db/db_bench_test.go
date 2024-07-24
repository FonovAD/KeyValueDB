package db_test

import (
	"testing"
)

func PutBenchTest(b *testing.B) {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range TCS {
			db.Put(tc.Key, tc.Value)
		}
	}
	b.StopTimer()
}

func GetBenchTest(b *testing.B) {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range TCS {
			db.Get(tc.Key)
		}
	}
	b.StopTimer()
}

type TestCase struct {
	ID    int
	Key   string
	Value string
}

var TCS []TestCase = []TestCase{
	TestCase{
		ID:    1,
		Key:   "keyAlphaBravo",
		Value: "The quick brown fox jumps over",
	},
	TestCase{
		ID:    2,
		Key:   "keyCharlieDelta",
		Value: "Lazy dog near the red house",
	},
	TestCase{
		ID:    3,
		Key:   "keyEchoFoxtrot",
		Value: "Zebras and lions share the savannah",
	},
	TestCase{
		ID:    4,
		Key:   "keyGolfHotel",
		Value: "Baking fresh bread early in the morning",
	},
	TestCase{
		ID:    5,
		Key:   "keyIndiaJuliet",
		Value: "Exploring the mountains in wintertime",
	},
	TestCase{
		ID:    6,
		Key:   "keyKiloLima",
		Value: "Waves crashing onto the sandy beach",
	},
	TestCase{
		ID:    7,
		Key:   "keyMikeNovember",
		Value: "A bright sunny day with clear skies",
	},
	TestCase{
		ID:    8,
		Key:   "keyOscarPapa",
		Value: "Coding a new program until late night",
	},
	TestCase{
		ID:    9,
		Key:   "keyQuebecRomeo",
		Value: "Reading an intriguing mystery novel",
	},
	TestCase{
		ID:    10,
		Key:   "keySierraTango",
		Value: "Listening to soothing classical music",
	},
	TestCase{
		ID:    11,
		Key:   "keyUniformVictor",
		Value: "Cooking delicious meals from scratch",
	},
	TestCase{
		ID:    12,
		Key:   "keyWhiskeyXray",
		Value: "A long journey through the forest path",
	},
	TestCase{
		ID:    13,
		Key:   "keyYankeeZulu",
		Value: "Building a cozy cabin in the woods",
	},
	TestCase{
		ID:    14,
		Key:   "keyAlphaBeta",
		Value: "Writing a captivating fantasy novel",
	},
	TestCase{
		ID:    15,
		Key:   "keyGammaDelta",
		Value: "Playing chess with a challenging opponent",
	},
	TestCase{
		ID:    16,
		Key:   "keyEpsilonZeta",
		Value: "Enjoying a peaceful evening by the lake",
	},
	TestCase{
		ID:    17,
		Key:   "keyEtaTheta",
		Value: "Discovering hidden treasures in the attic",
	},
	TestCase{
		ID:    18,
		Key:   "keyIotaKappa",
		Value: "Learning a new language for travel",
	},
	TestCase{
		ID:    19,
		Key:   "keyLambdaMu",
		Value: "Gardening in the spring sunshine",
	},
	TestCase{
		ID:    20,
		Key:   "keyNuXi",
		Value: "Crafting handmade gifts for friends",
	},
}

package model

type ChTexture struct{
	Contrast        float64 `db:"contrast"`
	Entropy		   	float64	`db:"entropy"`
	Idm			   	float64	`db:"idm"`
	MaxProbability 	float64	`db:"maxProbability"`
	Correlation	   	float64	`db:"correlation"`
}
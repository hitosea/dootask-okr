package core

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(5*time.Hour, 10*time.Hour)

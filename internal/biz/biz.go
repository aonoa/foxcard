package biz

import (
	"foxcard/internal/biz/card"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(card.NewCardUsecase)

package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/libmonsoon-dev/LonginusNightmare/logger"
)

var _ logger.Logger = (*logrus.Logger)(nil)

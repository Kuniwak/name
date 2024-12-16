package printer

import (
	"github.com/Kuniwak/name/filter"
)

type Func func(<-chan filter.Target)

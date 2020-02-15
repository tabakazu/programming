package repository

import "github.com/tabakazu/csharp-tdd-ddd-udemy/domain"

type IWhetherRepository interface {
	GetLatest(areaId int) domain.Whether
}

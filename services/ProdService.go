package services

import (
    "context"
)

type ProdService struct {
}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {

    var stocks int32 = 0
    if request.Area == Area_Area_A {
        stocks = 90
    } else if request.Area == Area_Area_B {
        stocks = 78
    } else {
        stocks = 99
    }
    return &ProdResponse{ProdStock: stocks}, nil
}

func (this *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdStockList, error) {

    ProdStocks := []*ProdResponse{
        &ProdResponse{ProdStock: 23},
        &ProdResponse{ProdStock: 323},
        &ProdResponse{ProdStock: 234},
        &ProdResponse{ProdStock: 2333},
    }
    return &ProdStockList{ProStocks: ProdStocks}, nil
}

func (this *ProdService) GetProdInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error) {
    res := ProdModel{
        ProdId:    100,
        ProdName:  "Iphone 12 Pro max",
        ProdPrice: 5999.00,
        ProdKc:    32,
    }
    return &res, nil
}

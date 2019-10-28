package main

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"errors"
)

func main() {
	urlsAll := []string{
		"/vCompreStaByCity/getList",
		"/budgetExpendMonth/exportBudgetExpendYearMonthList",
		"/ElectricityByPro/initElectricityPage",
		"/vCityLaborCostByCom/export",
		"/budgetLaborCostByBuss/export",
		"/budgetFixedNTMonth/export",
		"/vProBusinessHall/export",
		"/vCityLaborCostByCom/getList",
		"/vElecBillByCity/export",
		"/budgetComprehensiveStation/export",
		"/vCityUserSizeByCom/getList",
		"/FixedNetworkByPro/initFixedNetworkPage",
		"/BudgetPurchaseRate/getAllList",
		"/budgetExpendMonth/initBudgetExpendMonth",
		"/LandlineTerminal/getList",
		"/FixedNetworkByPro/fixedNetworkListMap",
		"/FixedNetworkByPro/fixedNetworkHeadMap",
		"/QuanShiElectricBill/getHeadColumn",
		"/vElecBillByCityByUnit/export",
		"/ElectricityByPro/exportElectricityMonthList",
		"/budgetBusinessHall/getList",
		"/budgetLaborCostByProTest/initBudgetLaborCostByProList",
		"/vCityBusinessHall/getList",
		"/budgetComprehensiveStation/getList",
		"/budgetLaborCostByBuss/getList",
		"/vUserSizeByPro/export",
		"/budgetBusinessHall/export",
		"/vCityLaborCostByBuss/getList",
		"/vGridInstallationByProvince/export",
		"/vUserSizeByPro/getAllUrl",
		"/QuanShiElectricBill/export",
		"/budgetExpendByPro/export",
		"/BudgetJobCostByPro/export",
		"/vGridInstallationByUnit/export",
		"/budgetTower/getList",
		"/ElectricityByPro/getElectricityMonthList",
		"/vCityUserSizeByCom/export",
		"/vGridInstallationByProvince/getList",
		"/vCompreStaByProvince/export",
		"/budgetGridInstallation/export",
		"/ElectricityByPro/getHeadColumn",
		"/vProLaborCost/export",
		"/budgetUserSizeByBuss/getList",
		"/budgetLaborCostByProTest/budgetLaborCostByProList",
		"/vProLaborCost/getList",
		"/budgetTower/export",
		"/vGridInstallationByBuss/export",
		"/vElecBillByCityByUnit/getList",
		"/budgetExpendByBuss/getList",
		"/budgetLaborCostByPro/export",
		"/budgetLaborCostByPro/budgetLaborCostByProList",
		"/budgetExpendByPro/getRengongGongzuoliang",
		"/eachCityElectricCharge/getList",
		"/budgetUserSizeByBuss/export",
		"/BudgetPurchaseRate/getLocalList",
		"/vUserSizeByPro/getList",
		"/budgetExpendMonth/getBudgetExpendYearMonthList",
		"/vCompreStaByCity/export",
		"/budgetLaborCostByProTest/budgetJobCostByProList",
		"/FixedNetworkByPro/exportFixedNetworkByPro",
		"/TowerRentByPro/export",
		"/budgetFixedNTMonth/getHeadColumn",
		"/eachCityElectricCharge/export",
		"/budgetLaborCostByPro/initBudgetLaborCostByProList",
		"/vCompreStaByProvince/getList",
		"/budgetFixedNetworkTerminal/export",
		"/budgetByCom/getUnit",
		"/vElecBillByCity/getList",
		"/budgetLaborCostByCom/export",
		"/vCityBusinessHall/export",
		"/BudgetJobCostByPro/budgetJobCostByProList",
		"/vCityUserSizeByBuss/export",
		"/budgetLaborCostByCom/getList",
		"/budgetFixedNTMonth/getList",
		"/budgetFixedNetworkTerminal/getList",
		"/vCityUserSizeByBuss/getList",
		"/vCityLaborCostByBuss/export",
		"/budgetGridInstallation/getList",
		"/budgetExpendByPro/getRengongLaochanlv",
		"/budgetExpendByPro/getList",
		"/vElecBillByProvince/getList",
		"/budgetExpendByPro/saveRengong",
		"/vProBusinessHall/getList",
		"/budgetExpendMonth/getHeadColumn",
		"/budgetUserSizeByCom/export",
		"/QuanShiElectricBill/getList",
		"/budgetUserSizeByCom/getList",
		"/TowerRentByPro/getList",
		"/vGridInstallationByUnit/getList",
		"/vGridInstallationByBuss/getList",
		"/vElecBillByProvince/export",
		"/BudgetPurchaseRate/getDynamic",
	}
	// urls := []string{
	// 	// //营业厅商圈  budgetBusinessHall
	// 	// "http://localhost:8080/gscw-zrys/budgetBusinessHall/getList",

	// 	// // 综维站 BudgetComprehensiveStation
	// 	// "http://localhost:8080/gscw-zrys/budgetComprehensiveStation/getList",

	// 	// // 每个城市的电费
	// 	// "http://localhost:8080/gscw-zrys/eachCityElectricCharge/getList",

	// 	// "http://localhost:8080/gscw-zrys/budgetExpendByBuss/getList",

	// 	// // 划小单元
	// 	// "http://localhost:8080/gscw-zrys/budgetByCom/getUnit",

	// 	// "http://localhost:8080/gscw-zrys/budgetExpendByPro/getList",

	// 	// // 网格装维 budgetGridInstallation
	// 	// "http://localhost:8080/gscw-zrys/budgetGridInstallation/getList",

	// 	// // 全口径人工成本-经营单位 budgetLaborCostByBuss
	// 	// "http://localhost:8080/gscw-zrys/budgetLaborCostByBuss/getList",
		
	// 	// // 全口径人工成本-划小单元 budgetLaborCostByCom
	// 	// "http://localhost:8080/gscw-zrys/budgetLaborCostByCom/getList",

	// 	// // 固网终端 1.4 LandlineTerminal
	// 	// "http://localhost:8080/gscw-zrys/LandlineTerminal/getList",

	// 	// // 固网终端 
	// 	// "http://localhost:8080/gscw-zrys/budgetFixedNetworkTerminal/getList",

	// 	// // 固网终端 budgetFixedNetworkTerminalMonth
	// 	// "http://localhost:8080/gscw-zrys/budgetFixedNetworkTerminalMonth/getList",

	// 	// // 采净比 BudgetPurchaseRate
	// 	// "http://localhost:8080/gscw-zrys/BudgetPurchaseRate/getAllList",

	// 	// // 全市电费 QuanShiElectricBill
	// 	// "http://localhost:8080/gscw-zrys/QuanShiElectricBill/getList",

	// 	// // 铁塔 towerElectric
	// 	// "http://localhost:8080/gscw-zrys/towerElectric/getList",

	// 	// // 用户规模-经营单位 budgetUserSizeByBuss	
	// 	// "http://localhost:8080/gscw-zrys/budgetUserSizeByBuss/getList",

	// 	// // 用户规模-划小单元 budgetUserSizeByCom
	// 	// "http://localhost:8080/gscw-zrys/budgetUserSizeByCom/getList",

	// 	// // 分公司塔租情况 TowerRentByPro
	// 	// "http://localhost:8080/gscw-zrys/TowerRentByPro/getList",

	// 	// // 全口径人工成本劳产率  BudgetLaborCostByPro
	// 	// "http://localhost:8080/gscw-zrys/budgetLaborCostByPro/getList",

	// 	// // 全口径人工成本工作量 BudgetJobCostByPro
	// 	// "http://localhost:8080/gscw-zrys/budgetJobCostByPro/getList",

	// 	// 用户规模-经营单位导出接口测试 budgetUserSizeByBuss
	// 	// "http://localhost:8080/gscw-zrys/budgetUserSizeByBuss/export",

	// 	// // 用户规模-划小单元导出测试导出接口测试 budgetUserSizeByCom
	// 	// "http://localhost:8080/gscw-zrys/budgetUserSizeByCom/export",

	// 	// // 全市电费-全市每月电费导出接口测试  QuanShiElectricBill
	// 	// "http://localhost:8080/gscw-zrys/QuanShiElectricBill/export",
		
	// 	// // 网格装维 budgetGridInstallation
	// 	// "http://localhost:8080/gscw-zrys/budgetGridInstallation/export",

	// 	// // 固网终端-天翼网关导出接口测试 budgetFixedNetworkTerminal
	// 	// "http://localhost:8080/gscw-zrys/budgetFixedNetworkTerminal/export",

	// 	// // 铁塔 导出接口测试 towerElectric
	// 	// "http://localhost:8080/gscw-zrys/towerElectric/export",

	// 	// // 塔租 导出接口测试 TowerRentByPro
	// 	// "http://localhost:8080/gscw-zrys/TowerRentByPro/export",

	// 	// // 全口径人工成本-工作量 导出接口测试 BudgetJobCostByPro
	// 	// "http://localhost:8080/gscw-zrys/BudgetJobCostByPro/export",

	// 	// 
	// 	"http://localhost:8080/gscw-zrys/vElecBillByCityByUnit/getList",

	// 	"http://localhost:8080/gscw-zrys/vElecBillByCity/getList",

	// 	"http://localhost:8080/gscw-zrys/vElecBillByProvince/getList",

	// 	// vCompreStaByCity
	// 	"http://localhost:8080/gscw-zrys/vCompreStaByCity/getList",

	// 	// vCompreStaByProvince
	// 	"http://localhost:8080/gscw-zrys/vCompreStaByProvince/getList",

	// 	// vGridInstallationByBuss
	// 	"http://localhost:8080/gscw-zrys/vGridInstallationByBuss/getList",

	// 	// vGridInstallationByProvince
	// 	"http://localhost:8080/gscw-zrys/vGridInstallationByProvince/getList",

	// 	// vGridInstallationByUnit
	// 	"http://localhost:8080/gscw-zrys/vGridInstallationByUnit/getList",
		
	// 	// vCityBusinessHall pyf
	// 	"http://localhost:8080/gscw-zrys/vCityBusinessHall/getList",

	// 	// vCityLaborCostByBuss pyf
	// 	"http://localhost:8080/gscw-zrys/vCityLaborCostByBuss/getList",

	// 	// vCityLaborCostByCom
	// 	"http://localhost:8080/gscw-zrys/vCityLaborCostByCom/getList",

	// 	// vCityUserSizeByBuss pyf
	// 	"http://localhost:8080/gscw-zrys/vCityUserSizeByBuss/getList",

	// 	// vCityUserSizeByCom pyf
	// 	"http://localhost:8080/gscw-zrys/vCityUserSizeByCom/getList",

	// 	// vProBusinessHall pyf
	// 	"http://localhost:8080/gscw-zrys/vProBusinessHall/getList",

	// 	// vProLaborCost pyf
	// 	"http://localhost:8080/gscw-zrys/vProLaborCost/getList",

	// 	// vUserSizeByPro pyf
	// 	"http://localhost:8080/gscw-zrys/vUserSizeByPro/getList",
	// }
	log.Println("共测试", len(urlsAll), "请求")
	for _, url := range urlsAll {
		res, err := Get(url, nil, nil)
		if err != nil {
			log.Panic("request error!!!")
		}
		if res.StatusCode == 200 {
			log.Println("request success")
		} else {
			log.Warnln(res.StatusCode, "请求失败")
		}
	}
}

func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail")
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	log.Printf("Go %s URL : %s \n",http.MethodGet, req.URL.String())
	return client.Do(req)
}


package common_tools

import (
	"fmt"
	"testing"
)

type SystemLogisticsTracking struct {
	LogisticsMode                                  string `resID:"24741" json:"logistics_mode"   `                                       //物流方式
	LogisticsSign                                  string `resID:"25046" json:"logistics_sign"   `                                       //来源标志
	LogisticsStatus                                string `resID:"24743" json:"logistics_status"   `                                     //物流状态
	AutomaticallyQueryTheCausesOfLogisticsFailures string `resID:"24752" json:"automatically_query_the_causes_of_logistics_failures"   ` //自动化查询物流失败原因
	DeliveryTime                                   string `resID:"24753" json:"delivery_time"   `                                        //发货时间
	SigningTime                                    string `resID:"24759" json:"signing_time"   `                                         //签收时间
	UpdatedName                                    string `resID:"24709_person" json:"updated_name"    `                                 //更新人姓名
	Updated                                        string `resID:"24709" json:"updated"   `                                              //更新人
	CauseOfAutomaticProxyCreationFailure           string `resID:"24715" json:"cause_of_automatic_proxy_creation_failure"   `            //自动创建代理失败原因
	AutomatedQueryOfLogisticsFailures              int    `resID:"24761" json:"automated_query_of_logistics_failures"   `                //自动化查询物流失败次数
	CurrentLogistics                               string `resID:"24742" json:"current_logistics"   `                                    //当前物流
	ServiceTime                                    string `resID:"24760" json:"service_time"   `                                         //送达时间
	FingerprintStartupTime                         string `resID:"24706" json:"fingerprint_startup_time"   `                             //指纹启动时间
	CreationTime                                   string `resID:"24708" json:"creation_time"   `                                        //创建时间
	LastAutomaticQueryOfLogisticsStatus            string `resID:"24755" json:"last_automatic_query_of_logistics_status"   `             //上次自动化查询物流状态
}
type DD struct {
	LogisticsSource                           string `resID:"24723" json:"logistics_source"   `                                      //物流来源
	LogisticsNumber                           string `resID:"24739" json:"logistics_number"   `                                      //物流单号
	LastAutomaticQueryOfLogisticsSuccessTime  string `resID:"24756,person" json:"last_automatic_query_of_logistics_success_time"   ` //上次自动化查询物流成功时间
	CreatedName                               string `resID:"24707_person" json:"created_name"    `                                  //创建人姓名
	Created                                   string `resID:"24707" json:"created"   `                                               //创建人
	CauseOfLastAutomaticInitializationFailure string `resID:"24713" json:"cause_of_last_automatic_initialization_failure"   `        //上次自动化初始化失败原因
}
type ZZ struct {
	LastUpdateTimeOfLogistics                  string `resID:"24778" json:"last_update_time_of_logistics"   `                   //物流最后更新时间
	CauseOfAutomaticFingerprintCreationFailure string `resID:"24716" json:"cause_of_automatic_fingerprint_creation_failure"   ` //自动创建指纹失败原因
	OrderNo                                    string `resID:"24740" json:"order_no."   `                                       //订单号
	LastAutomaticQueryOfLogisticsTime          string `resID:"24754" json:"last_automatic_query_of_logistics_time"   `          //上次自动化查询物流时间
	LogisticsRecords                           string `resID:"24744" json:"logistics_records"   `                               //物流记录
	UID                                        int    `resID:"uid" json:"uid"   `                                               //UID
	UpdateTime                                 string `resID:"24710" json:"update_time"   `                                     //更新时间
	AutomaticallyCreatedProxy                  string `resID:"24714" json:"automatically_created_proxy"   `                     //自动创建的代理
	Remarks                                    string `resID:"24733" json:"remarks"   `                                         //备注
	CurrentStatusRetentionDays                 int    `resID:"24750" json:"current_status_retention_days"   `                   //当前状态停留天数
	LogisticsCompany                           string `resID:"25604" json:"logistics_company"   `                               //物流公司
	DdD                                        DD     `resID:"DDC" json:"DDX"`                                                  //测试1
	DdD2                                       *DD    `resID:"DdD2" json:"DdD2"`                                                //测试1
	// DdD  DD     `resID:"DDC" json:"DDX"` //测试1
	DD  //测试2
	DDs []DD
}

func TestResolveTag(t *testing.T) {
	var z ZZ
	mp := make(map[string]string)
	ResolveTag(&z, mp, "json", "resID")
	for k, v := range mp {
		fmt.Println(k, v)
	}
}

func TestResolveVal(t *testing.T) {
	var z ZZ
	z.Remarks = "remark"
	mp := make(map[string]interface{})
	ResolveVal(&z, mp, "resID", "uid", false)
	for k, v := range mp {
		fmt.Println(k, v)
	}
}

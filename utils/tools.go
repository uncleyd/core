/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 00:24:25
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 10:12:06
***********************************************/

package utils

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
	"regexp"
	"strconv"
	"strings"
)

// 获取版本号
func TransVer2Str(ver int) string {
	var v1 = ver % 100
	var v2 = (ver / 100) % 100
	var v3 = ver / 10000
	return strings.Join([]string{strconv.Itoa(v3), strconv.Itoa(v2), strconv.Itoa(v1)}, ".")
}

// 获取版本号
func TransStr2Ver(verStr string) int {
	var p = strings.Split(verStr, ".")
	var plen = len(p)
	var iidx = 1
	var fver = 0
	for pidx := plen - 1; pidx >= 0; pidx-- {
		var pint, _ = strconv.Atoi(p[pidx])
		fver += pint * iidx
		iidx *= 100
	}
	return fver
}

// 校验手机号
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// 发送短信
func SendSms(bid, phone string, smsCode string, min string) bool {
	credential := common.NewCredential(
		"AKIDlJdz7ht744rQp0BVw1qOhsoU3VI0wAtZ",
		"h8KgneGAqKYWgcMrrc7Md5KK4fZGhsnU",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "", cpf)

	request := sms.NewSendSmsRequest()
	var name = "天气预报"
	if bid == "com.beemans.calendar.app" || bid == "com.calendar.jishi.app" || "com.live.calendar" == bid {
		name = "万年历"
	}
	params := "{\"PhoneNumberSet\":[\"+86" + phone + "\"],\"TemplateID\":\"541420\",\"Sign\":\"" + name + "\",\"TemplateParamSet\":[\"" + smsCode + "\",\"" + min + "\"],\"SmsSdkAppid\":\"1400320636\"}"

	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	_, err = client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return false
	}
	if err != nil {
		panic(err)
		return false
	}
	return true
}

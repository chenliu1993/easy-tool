package client

import (
	"fmt"

	"github.com/chenliu1993/easy-tool/pkg/types"
	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func GenerateCAMClient(id, key string) *cam.Client {
	credential := common.NewCredential(
		id,
		key,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cam.tencentcloudapi.com"

	client, _ := cam.NewClient(credential, "", cpf)
	return client
}

func UpdateUsers(client *cam.Client, accounts types.Accounts) error {
	for _, account := range accounts.Data {
		request := cam.NewUpdateUserRequest()
		// Currently we brutaly do a all update......
		request.Name = common.StringPtr(account.Name)
		request.Remark = common.StringPtr(account.Remark)
		request.ConsoleLogin = common.Uint64Ptr(uint64(account.ConsoleLogin))
		//request.Password = common.StringPtr(account.Password)
		//request.NeedResetPassword = common.Uint64Ptr(uint64(account.NeedResetPassword))
		request.PhoneNum = common.StringPtr(account.PhoneNum)
		request.CountryCode = common.StringPtr(account.CountryCode)
		request.Email = common.StringPtr(account.Email)

		// send the request
		response, err := client.UpdateUser(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			return fmt.Errorf("An API error has returned: %s", err)
		}
		if err != nil {
			return err
		}
		// Is this needed?
		fmt.Printf("%s", response.ToJsonString())
	}
	return nil
}

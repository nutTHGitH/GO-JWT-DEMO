package main

type JWTPayload struct {
    MerchantId string `json:"merchantId"`
    BranchId   string `json:"branchId"`
    DeviceId   string `json:"deviceId"`
    BranchSk   string `json:"branchSk"`
    LoginRefId string `json:"loginRefId"`
}

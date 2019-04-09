// This package contains the information for
// the pubsub messages to generate a licence

package pubsub_lgbkgenlic_v1

var lgbkgenlicV1Topic = "generate_logbook_licence_v1"

type GenerateLogbookLicenceMsgV1 struct {
	UserId         string `json:"userId"`
	IssuerId       string `json:"issuerId"`
	SubscriptionId string `json:"subscriptionId"`
	AddonAssigned  bool   `json:"addonAssigned"`
}
// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// CloneReceiptRuleSet
//
// The following example creates a receipt rule set by cloning an existing one:
func ExampleSES_CloneReceiptRuleSet_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.CloneReceiptRuleSetInput{
		OriginalRuleSetName: aws.String("RuleSetToClone"),
		RuleSetName:         aws.String("RuleSetToCreate"),
	}

	result, err := svc.CloneReceiptRuleSet(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeRuleSetDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleSetDoesNotExistException, aerr.Error())
			case ses.ErrCodeAlreadyExistsException:
				fmt.Println(ses.ErrCodeAlreadyExistsException, aerr.Error())
			case ses.ErrCodeLimitExceededException:
				fmt.Println(ses.ErrCodeLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// CreateReceiptFilter
//
// The following example creates a new IP address filter:
func ExampleSES_CreateReceiptFilter_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.CreateReceiptFilterInput{
		Filter: &ses.ReceiptFilter{
			IpFilter: &ses.ReceiptIpFilter{
				Cidr:   aws.String("1.2.3.4/24"),
				Policy: aws.String("Allow"),
			},
			Name: aws.String("MyFilter"),
		},
	}

	result, err := svc.CreateReceiptFilter(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeLimitExceededException:
				fmt.Println(ses.ErrCodeLimitExceededException, aerr.Error())
			case ses.ErrCodeAlreadyExistsException:
				fmt.Println(ses.ErrCodeAlreadyExistsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// CreateReceiptRule
//
// The following example creates a new receipt rule:
func ExampleSES_CreateReceiptRule_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.CreateReceiptRuleInput{
		After: aws.String(""),
		Rule: &ses.ReceiptRule{
			Actions: []*ses.ReceiptAction{
				{
					S3Action: &ses.S3Action{
						BucketName:      aws.String("MyBucket"),
						ObjectKeyPrefix: aws.String("email"),
					},
				},
			},
			Enabled:     aws.Bool(true),
			Name:        aws.String("MyRule"),
			ScanEnabled: aws.Bool(true),
			TlsPolicy:   aws.String("Optional"),
		},
		RuleSetName: aws.String("MyRuleSet"),
	}

	result, err := svc.CreateReceiptRule(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeInvalidSnsTopicException:
				fmt.Println(ses.ErrCodeInvalidSnsTopicException, aerr.Error())
			case ses.ErrCodeInvalidS3ConfigurationException:
				fmt.Println(ses.ErrCodeInvalidS3ConfigurationException, aerr.Error())
			case ses.ErrCodeInvalidLambdaFunctionException:
				fmt.Println(ses.ErrCodeInvalidLambdaFunctionException, aerr.Error())
			case ses.ErrCodeAlreadyExistsException:
				fmt.Println(ses.ErrCodeAlreadyExistsException, aerr.Error())
			case ses.ErrCodeRuleDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleDoesNotExistException, aerr.Error())
			case ses.ErrCodeRuleSetDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleSetDoesNotExistException, aerr.Error())
			case ses.ErrCodeLimitExceededException:
				fmt.Println(ses.ErrCodeLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// CreateReceiptRuleSet
//
// The following example creates an empty receipt rule set:
func ExampleSES_CreateReceiptRuleSet_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.CreateReceiptRuleSetInput{
		RuleSetName: aws.String("MyRuleSet"),
	}

	result, err := svc.CreateReceiptRuleSet(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeAlreadyExistsException:
				fmt.Println(ses.ErrCodeAlreadyExistsException, aerr.Error())
			case ses.ErrCodeLimitExceededException:
				fmt.Println(ses.ErrCodeLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DeleteIdentity
//
// The following example deletes an identity from the list of identities that have been
// submitted for verification with Amazon SES:
func ExampleSES_DeleteIdentity_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.DeleteIdentityInput{
		Identity: aws.String("user@example.com"),
	}

	result, err := svc.DeleteIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DeleteIdentityPolicy
//
// The following example deletes a sending authorization policy for an identity:
func ExampleSES_DeleteIdentityPolicy_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.DeleteIdentityPolicyInput{
		Identity:   aws.String("user@example.com"),
		PolicyName: aws.String("MyPolicy"),
	}

	result, err := svc.DeleteIdentityPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DeleteReceiptFilter
//
// The following example deletes an IP address filter:
func ExampleSES_DeleteReceiptFilter_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.DeleteReceiptFilterInput{
		FilterName: aws.String("MyFilter"),
	}

	result, err := svc.DeleteReceiptFilter(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DeleteReceiptRule
//
// The following example deletes a receipt rule:
func ExampleSES_DeleteReceiptRule_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.DeleteReceiptRuleInput{
		RuleName:    aws.String("MyRule"),
		RuleSetName: aws.String("MyRuleSet"),
	}

	result, err := svc.DeleteReceiptRule(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeRuleSetDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DeleteReceiptRuleSet
//
// The following example deletes a receipt rule set:
func ExampleSES_DeleteReceiptRuleSet_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.DeleteReceiptRuleSetInput{
		RuleSetName: aws.String("MyRuleSet"),
	}

	result, err := svc.DeleteReceiptRuleSet(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeCannotDeleteException:
				fmt.Println(ses.ErrCodeCannotDeleteException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DeleteVerifiedEmailAddress
//
// The following example deletes an email address from the list of identities that have
// been submitted for verification with Amazon SES:
func ExampleSES_DeleteVerifiedEmailAddress_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.DeleteVerifiedEmailAddressInput{
		EmailAddress: aws.String("user@example.com"),
	}

	result, err := svc.DeleteVerifiedEmailAddress(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DescribeActiveReceiptRuleSet
//
// The following example returns the metadata and receipt rules for the receipt rule
// set that is currently active:
func ExampleSES_DescribeActiveReceiptRuleSet_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.DescribeActiveReceiptRuleSetInput{}

	result, err := svc.DescribeActiveReceiptRuleSet(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DescribeReceiptRule
//
// The following example returns the details of a receipt rule:
func ExampleSES_DescribeReceiptRule_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.DescribeReceiptRuleInput{
		RuleName:    aws.String("MyRule"),
		RuleSetName: aws.String("MyRuleSet"),
	}

	result, err := svc.DescribeReceiptRule(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeRuleDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleDoesNotExistException, aerr.Error())
			case ses.ErrCodeRuleSetDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// DescribeReceiptRuleSet
//
// The following example returns the metadata and receipt rules of a receipt rule set:
func ExampleSES_DescribeReceiptRuleSet_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.DescribeReceiptRuleSetInput{
		RuleSetName: aws.String("MyRuleSet"),
	}

	result, err := svc.DescribeReceiptRuleSet(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeRuleSetDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetAccountSendingEnabled
//
// The following example returns if sending status for an account is enabled. (true
// / false):
func ExampleSES_GetAccountSendingEnabled_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.GetAccountSendingEnabledInput{}

	result, err := svc.GetAccountSendingEnabled(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetIdentityDkimAttributes
//
// The following example retrieves the Amazon SES Easy DKIM attributes for a list of
// identities:
func ExampleSES_GetIdentityDkimAttributes_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.GetIdentityDkimAttributesInput{
		Identities: []*string{
			aws.String("example.com"),
			aws.String("user@example.com"),
		},
	}

	result, err := svc.GetIdentityDkimAttributes(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetIdentityMailFromDomainAttributes
//
// The following example returns the custom MAIL FROM attributes for an identity:
func ExampleSES_GetIdentityMailFromDomainAttributes_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.GetIdentityMailFromDomainAttributesInput{
		Identities: []*string{
			aws.String("example.com"),
		},
	}

	result, err := svc.GetIdentityMailFromDomainAttributes(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetIdentityNotificationAttributes
//
// The following example returns the notification attributes for an identity:
func ExampleSES_GetIdentityNotificationAttributes_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.GetIdentityNotificationAttributesInput{
		Identities: []*string{
			aws.String("example.com"),
		},
	}

	result, err := svc.GetIdentityNotificationAttributes(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetIdentityPolicies
//
// The following example returns a sending authorization policy for an identity:
func ExampleSES_GetIdentityPolicies_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.GetIdentityPoliciesInput{
		Identity: aws.String("example.com"),
		PolicyNames: []*string{
			aws.String("MyPolicy"),
		},
	}

	result, err := svc.GetIdentityPolicies(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetIdentityVerificationAttributes
//
// The following example returns the verification status and the verification token
// for a domain identity:
func ExampleSES_GetIdentityVerificationAttributes_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.GetIdentityVerificationAttributesInput{
		Identities: []*string{
			aws.String("example.com"),
		},
	}

	result, err := svc.GetIdentityVerificationAttributes(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetSendQuota
//
// The following example returns the Amazon SES sending limits for an AWS account:
func ExampleSES_GetSendQuota_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.GetSendQuotaInput{}

	result, err := svc.GetSendQuota(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// GetSendStatistics
//
// The following example returns Amazon SES sending statistics:
func ExampleSES_GetSendStatistics_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.GetSendStatisticsInput{}

	result, err := svc.GetSendStatistics(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// ListIdentities
//
// The following example lists the email address identities that have been submitted
// for verification with Amazon SES:
func ExampleSES_ListIdentities_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.ListIdentitiesInput{
		IdentityType: aws.String("EmailAddress"),
		MaxItems:     aws.Int64(123),
		NextToken:    aws.String(""),
	}

	result, err := svc.ListIdentities(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// ListIdentityPolicies
//
// The following example returns a list of sending authorization policies that are attached
// to an identity:
func ExampleSES_ListIdentityPolicies_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.ListIdentityPoliciesInput{
		Identity: aws.String("example.com"),
	}

	result, err := svc.ListIdentityPolicies(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// ListReceiptFilters
//
// The following example lists the IP address filters that are associated with an AWS
// account:
func ExampleSES_ListReceiptFilters_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.ListReceiptFiltersInput{}

	result, err := svc.ListReceiptFilters(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// ListReceiptRuleSets
//
// The following example lists the receipt rule sets that exist under an AWS account:
func ExampleSES_ListReceiptRuleSets_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.ListReceiptRuleSetsInput{
		NextToken: aws.String(""),
	}

	result, err := svc.ListReceiptRuleSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// ListVerifiedEmailAddresses
//
// The following example lists all email addresses that have been submitted for verification
// with Amazon SES:
func ExampleSES_ListVerifiedEmailAddresses_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.ListVerifiedEmailAddressesInput{}

	result, err := svc.ListVerifiedEmailAddresses(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// PutIdentityPolicy
//
// The following example adds a sending authorization policy to an identity:
func ExampleSES_PutIdentityPolicy_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.PutIdentityPolicyInput{
		Identity:   aws.String("example.com"),
		Policy:     aws.String("{\"Version\":\"2008-10-17\",\"Statement\":[{\"Sid\":\"stmt1469123904194\",\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"arn:aws:iam::123456789012:root\"},\"Action\":[\"ses:SendEmail\",\"ses:SendRawEmail\"],\"Resource\":\"arn:aws:ses:us-east-1:EXAMPLE65304:identity/example.com\"}]}"),
		PolicyName: aws.String("MyPolicy"),
	}

	result, err := svc.PutIdentityPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeInvalidPolicyException:
				fmt.Println(ses.ErrCodeInvalidPolicyException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// ReorderReceiptRuleSet
//
// The following example reorders the receipt rules within a receipt rule set:
func ExampleSES_ReorderReceiptRuleSet_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.ReorderReceiptRuleSetInput{
		RuleNames: []*string{
			aws.String("MyRule"),
			aws.String("MyOtherRule"),
		},
		RuleSetName: aws.String("MyRuleSet"),
	}

	result, err := svc.ReorderReceiptRuleSet(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeRuleSetDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleSetDoesNotExistException, aerr.Error())
			case ses.ErrCodeRuleDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// SendEmail
//
// The following example sends a formatted email:
func ExampleSES_SendEmail_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{
				aws.String("recipient3@example.com"),
			},
			ToAddresses: []*string{
				aws.String("recipient1@example.com"),
				aws.String("recipient2@example.com"),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("This message body contains HTML formatting. It can, for example, contain links like this one: <a class=\"ulink\" href=\"http://docs.aws.amazon.com/ses/latest/DeveloperGuide\" target=\"_blank\">Amazon SES Developer Guide</a>."),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("This is the message body in text format."),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("Test email"),
			},
		},
		ReturnPath:    aws.String(""),
		ReturnPathArn: aws.String(""),
		Source:        aws.String("sender@example.com"),
		SourceArn:     aws.String(""),
	}

	result, err := svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			case ses.ErrCodeConfigurationSetSendingPausedException:
				fmt.Println(ses.ErrCodeConfigurationSetSendingPausedException, aerr.Error())
			case ses.ErrCodeAccountSendingPausedException:
				fmt.Println(ses.ErrCodeAccountSendingPausedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// SendRawEmail
//
// The following example sends an email with an attachment:
func ExampleSES_SendRawEmail_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.SendRawEmailInput{
		FromArn: aws.String(""),
		RawMessage: &ses.RawMessage{
			Data: []byte("From: sender@example.com\\nTo: recipient@example.com\\nSubject: Test email (contains an attachment)\\nMIME-Version: 1.0\\nContent-type: Multipart/Mixed; boundary=\"NextPart\"\\n\\n--NextPart\\nContent-Type: text/plain\\n\\nThis is the message body.\\n\\n--NextPart\\nContent-Type: text/plain;\\nContent-Disposition: attachment; filename=\"attachment.txt\"\\n\\nThis is the text in the attachment.\\n\\n--NextPart--"),
		},
		ReturnPathArn: aws.String(""),
		Source:        aws.String(""),
		SourceArn:     aws.String(""),
	}

	result, err := svc.SendRawEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			case ses.ErrCodeConfigurationSetSendingPausedException:
				fmt.Println(ses.ErrCodeConfigurationSetSendingPausedException, aerr.Error())
			case ses.ErrCodeAccountSendingPausedException:
				fmt.Println(ses.ErrCodeAccountSendingPausedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// SetActiveReceiptRuleSet
//
// The following example sets the active receipt rule set:
func ExampleSES_SetActiveReceiptRuleSet_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.SetActiveReceiptRuleSetInput{
		RuleSetName: aws.String("RuleSetToActivate"),
	}

	result, err := svc.SetActiveReceiptRuleSet(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeRuleSetDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// SetIdentityDkimEnabled
//
// The following example configures Amazon SES to Easy DKIM-sign the email sent from
// an identity:
func ExampleSES_SetIdentityDkimEnabled_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.SetIdentityDkimEnabledInput{
		DkimEnabled: aws.Bool(true),
		Identity:    aws.String("user@example.com"),
	}

	result, err := svc.SetIdentityDkimEnabled(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// SetIdentityFeedbackForwardingEnabled
//
// The following example configures Amazon SES to forward an identity's bounces and
// complaints via email:
func ExampleSES_SetIdentityFeedbackForwardingEnabled_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.SetIdentityFeedbackForwardingEnabledInput{
		ForwardingEnabled: aws.Bool(true),
		Identity:          aws.String("user@example.com"),
	}

	result, err := svc.SetIdentityFeedbackForwardingEnabled(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// SetIdentityHeadersInNotificationsEnabled
//
// The following example configures Amazon SES to include the original email headers
// in the Amazon SNS bounce notifications for an identity:
func ExampleSES_SetIdentityHeadersInNotificationsEnabled_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.SetIdentityHeadersInNotificationsEnabledInput{
		Enabled:          aws.Bool(true),
		Identity:         aws.String("user@example.com"),
		NotificationType: aws.String("Bounce"),
	}

	result, err := svc.SetIdentityHeadersInNotificationsEnabled(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// SetIdentityMailFromDomain
//
// The following example configures Amazon SES to use a custom MAIL FROM domain for
// an identity:
func ExampleSES_SetIdentityMailFromDomain_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.SetIdentityMailFromDomainInput{
		BehaviorOnMXFailure: aws.String("UseDefaultValue"),
		Identity:            aws.String("user@example.com"),
		MailFromDomain:      aws.String("bounces.example.com"),
	}

	result, err := svc.SetIdentityMailFromDomain(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// SetIdentityNotificationTopic
//
// The following example sets the Amazon SNS topic to which Amazon SES will publish
// bounce, complaint, and/or delivery notifications for emails sent with the specified
// identity as the Source:
func ExampleSES_SetIdentityNotificationTopic_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.SetIdentityNotificationTopicInput{
		Identity:         aws.String("user@example.com"),
		NotificationType: aws.String("Bounce"),
		SnsTopic:         aws.String("arn:aws:sns:us-west-2:111122223333:MyTopic"),
	}

	result, err := svc.SetIdentityNotificationTopic(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// SetReceiptRulePosition
//
// The following example sets the position of a receipt rule in a receipt rule set:
func ExampleSES_SetReceiptRulePosition_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.SetReceiptRulePositionInput{
		After:       aws.String("PutRuleAfterThisRule"),
		RuleName:    aws.String("RuleToReposition"),
		RuleSetName: aws.String("MyRuleSet"),
	}

	result, err := svc.SetReceiptRulePosition(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeRuleSetDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleSetDoesNotExistException, aerr.Error())
			case ses.ErrCodeRuleDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// UpdateAccountSendingEnabled
//
// The following example updated the sending status for this account.
func ExampleSES_UpdateAccountSendingEnabled_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.UpdateAccountSendingEnabledInput{
		Enabled: aws.Bool(true),
	}

	result, err := svc.UpdateAccountSendingEnabled(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// UpdateConfigurationSetReputationMetricsEnabled
//
// Set the reputationMetricsEnabled flag for a specific configuration set.
func ExampleSES_UpdateConfigurationSetReputationMetricsEnabled_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.UpdateConfigurationSetReputationMetricsEnabledInput{
		ConfigurationSetName: aws.String("foo"),
		Enabled:              aws.Bool(true),
	}

	result, err := svc.UpdateConfigurationSetReputationMetricsEnabled(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// UpdateConfigurationSetReputationMetricsEnabled
//
// Set the sending enabled flag for a specific configuration set.
func ExampleSES_UpdateConfigurationSetSendingEnabled_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.UpdateConfigurationSetSendingEnabledInput{
		ConfigurationSetName: aws.String("foo"),
		Enabled:              aws.Bool(true),
	}

	result, err := svc.UpdateConfigurationSetSendingEnabled(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// UpdateReceiptRule
//
// The following example updates a receipt rule to use an Amazon S3 action:
func ExampleSES_UpdateReceiptRule_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.UpdateReceiptRuleInput{
		Rule: &ses.ReceiptRule{
			Actions: []*ses.ReceiptAction{
				{
					S3Action: &ses.S3Action{
						BucketName:      aws.String("MyBucket"),
						ObjectKeyPrefix: aws.String("email"),
					},
				},
			},
			Enabled:     aws.Bool(true),
			Name:        aws.String("MyRule"),
			ScanEnabled: aws.Bool(true),
			TlsPolicy:   aws.String("Optional"),
		},
		RuleSetName: aws.String("MyRuleSet"),
	}

	result, err := svc.UpdateReceiptRule(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeInvalidSnsTopicException:
				fmt.Println(ses.ErrCodeInvalidSnsTopicException, aerr.Error())
			case ses.ErrCodeInvalidS3ConfigurationException:
				fmt.Println(ses.ErrCodeInvalidS3ConfigurationException, aerr.Error())
			case ses.ErrCodeInvalidLambdaFunctionException:
				fmt.Println(ses.ErrCodeInvalidLambdaFunctionException, aerr.Error())
			case ses.ErrCodeRuleSetDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleSetDoesNotExistException, aerr.Error())
			case ses.ErrCodeRuleDoesNotExistException:
				fmt.Println(ses.ErrCodeRuleDoesNotExistException, aerr.Error())
			case ses.ErrCodeLimitExceededException:
				fmt.Println(ses.ErrCodeLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// VerifyDomainDkim
//
// The following example generates DKIM tokens for a domain that has been verified with
// Amazon SES:
func ExampleSES_VerifyDomainDkim_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.VerifyDomainDkimInput{
		Domain: aws.String("example.com"),
	}

	result, err := svc.VerifyDomainDkim(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// VerifyDomainIdentity
//
// The following example starts the domain verification process with Amazon SES:
func ExampleSES_VerifyDomainIdentity_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.VerifyDomainIdentityInput{
		Domain: aws.String("example.com"),
	}

	result, err := svc.VerifyDomainIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// VerifyEmailAddress
//
// The following example starts the email address verification process with Amazon SES:
func ExampleSES_VerifyEmailAddress_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.VerifyEmailAddressInput{
		EmailAddress: aws.String("user@example.com"),
	}

	result, err := svc.VerifyEmailAddress(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// VerifyEmailIdentity
//
// The following example starts the email address verification process with Amazon SES:
func ExampleSES_VerifyEmailIdentity_shared00() {
	svc := ses.New(session.Must(session.NewSession()))
	input := &ses.VerifyEmailIdentityInput{
		EmailAddress: aws.String("user@example.com"),
	}

	result, err := svc.VerifyEmailIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

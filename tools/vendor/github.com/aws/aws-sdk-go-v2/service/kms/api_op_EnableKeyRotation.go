// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables automatic rotation of the key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) for the
// specified symmetric KMS key. You cannot enable automatic rotation of asymmetric
// KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-concepts.html#asymmetric-cmks),
// KMS keys with imported key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html), or
// KMS keys in a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
// To enable or disable automatic rotation of a set of related multi-Region keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html#mrk-replica-key),
// set the property on the primary key. The KMS key that you use for this operation
// must be in a compatible key state. For details, see Key state: Effect on your
// KMS key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a KMS key in a different Amazon Web Services account.
// Required permissions: kms:EnableKeyRotation
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations:
//
// * DisableKeyRotation
//
// * GetKeyRotationStatus
func (c *Client) EnableKeyRotation(ctx context.Context, params *EnableKeyRotationInput, optFns ...func(*Options)) (*EnableKeyRotationOutput, error) {
	if params == nil {
		params = &EnableKeyRotationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableKeyRotation", params, optFns, c.addOperationEnableKeyRotationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableKeyRotationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableKeyRotationInput struct {

	// Identifies a symmetric KMS key. You cannot enable automatic rotation of
	// asymmetric KMS keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-concepts.html#asymmetric-cmks),
	// KMS keys with imported key material
	// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html), or
	// KMS keys in a custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
	// To enable or disable automatic rotation of a set of related multi-Region keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html#mrk-replica-key),
	// set the property on the primary key. Specify the key ID or key ARN of the KMS
	// key. For example:
	//
	// * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	noSmithyDocumentSerde
}

type EnableKeyRotationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableKeyRotationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableKeyRotation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableKeyRotation{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpEnableKeyRotationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableKeyRotation(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opEnableKeyRotation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "EnableKeyRotation",
	}
}

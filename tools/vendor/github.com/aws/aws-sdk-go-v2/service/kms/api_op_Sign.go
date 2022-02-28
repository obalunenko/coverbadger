// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a digital signature (https://en.wikipedia.org/wiki/Digital_signature)
// for a message or message digest by using the private key in an asymmetric KMS
// key. To verify the signature, use the Verify operation, or use the public key in
// the same asymmetric KMS key outside of KMS. For information about symmetric and
// asymmetric KMS keys, see Using Symmetric and Asymmetric KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
// in the Key Management Service Developer Guide. Digital signatures are generated
// and verified by using asymmetric key pair, such as an RSA or ECC pair that is
// represented by an asymmetric KMS key. The key owner (or an authorized user) uses
// their private key to sign a message. Anyone with the public key can verify that
// the message was signed with that particular private key and that the message
// hasn't changed since it was signed. To use the Sign operation, provide the
// following information:
//
// * Use the KeyId parameter to identify an asymmetric KMS
// key with a KeyUsage value of SIGN_VERIFY. To get the KeyUsage value of a KMS
// key, use the DescribeKey operation. The caller must have kms:Sign permission on
// the KMS key.
//
// * Use the Message parameter to specify the message or message
// digest to sign. You can submit messages of up to 4096 bytes. To sign a larger
// message, generate a hash digest of the message, and then provide the hash digest
// in the Message parameter. To indicate whether the message is a full message or a
// digest, use the MessageType parameter.
//
// * Choose a signing algorithm that is
// compatible with the KMS key.
//
// When signing a message, be sure to record the KMS
// key and the signing algorithm. This information is required to verify the
// signature. To verify the signature that this operation generates, use the Verify
// operation. Or use the GetPublicKey operation to download the public key and then
// use the public key to verify the signature outside of KMS. The KMS key that you
// use for this operation must be in a compatible key state. For details, see Key
// state: Effect on your KMS key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// Key Management Service Developer Guide. Cross-account use: Yes. To perform this
// operation with a KMS key in a different Amazon Web Services account, specify the
// key ARN or alias ARN in the value of the KeyId parameter. Required permissions:
// kms:Sign
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations: Verify
func (c *Client) Sign(ctx context.Context, params *SignInput, optFns ...func(*Options)) (*SignOutput, error) {
	if params == nil {
		params = &SignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Sign", params, optFns, c.addOperationSignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SignInput struct {

	// Identifies an asymmetric KMS key. KMS uses the private key in the asymmetric KMS
	// key to sign the message. The KeyUsage type of the KMS key must be SIGN_VERIFY.
	// To find the KeyUsage of a KMS key, use the DescribeKey operation. To specify a
	// KMS key, use its key ID, key ARN, alias name, or alias ARN. When using an alias
	// name, prefix it with "alias/". To specify a KMS key in a different Amazon Web
	// Services account, you must use the key ARN or alias ARN. For example:
	//
	// * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// *
	// Alias name: alias/ExampleAlias
	//
	// * Alias ARN:
	// arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key
	// ARN for a KMS key, use ListKeys or DescribeKey. To get the alias name and alias
	// ARN, use ListAliases.
	//
	// This member is required.
	KeyId *string

	// Specifies the message or message digest to sign. Messages can be 0-4096 bytes.
	// To sign a larger message, provide the message digest. If you provide a message,
	// KMS generates a hash digest of the message and then signs it.
	//
	// This member is required.
	Message []byte

	// Specifies the signing algorithm to use when signing the message. Choose an
	// algorithm that is compatible with the type and size of the specified asymmetric
	// KMS key.
	//
	// This member is required.
	SigningAlgorithm types.SigningAlgorithmSpec

	// A list of grant tokens. Use a grant token when your permission to call this
	// operation comes from a new grant that has not yet achieved eventual consistency.
	// For more information, see Grant token
	// (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)
	// and Using a grant token
	// (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)
	// in the Key Management Service Developer Guide.
	GrantTokens []string

	// Tells KMS whether the value of the Message parameter is a message or message
	// digest. The default value, RAW, indicates a message. To indicate a message
	// digest, enter DIGEST.
	MessageType types.MessageType

	noSmithyDocumentSerde
}

type SignOutput struct {

	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the asymmetric KMS key that was used to sign the message.
	KeyId *string

	// The cryptographic signature that was generated for the message.
	//
	// * When used
	// with the supported RSA signing algorithms, the encoding of this value is defined
	// by PKCS #1 in RFC 8017 (https://tools.ietf.org/html/rfc8017).
	//
	// * When used with
	// the ECDSA_SHA_256, ECDSA_SHA_384, or ECDSA_SHA_512 signing algorithms, this
	// value is a DER-encoded object as defined by ANS X9.62–2005 and RFC 3279 Section
	// 2.2.3 (https://tools.ietf.org/html/rfc3279#section-2.2.3). This is the most
	// commonly used signature format and is appropriate for most uses.
	//
	// When you use
	// the HTTP API or the Amazon Web Services CLI, the value is Base64-encoded.
	// Otherwise, it is not Base64-encoded.
	Signature []byte

	// The signing algorithm that was used to sign the message.
	SigningAlgorithm types.SigningAlgorithmSpec

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSign{}, middleware.After)
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
	if err = addOpSignValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSign(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSign(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "Sign",
	}
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNotificationChannel(ctx *pulumi.Context, args *GetNotificationChannelArgs, opts ...pulumi.InvokeOption) (*GetNotificationChannelResult, error) {
	var rv GetNotificationChannelResult
	err := ctx.Invoke("sysdig:Secure/getNotificationChannel:GetNotificationChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetNotificationChannel.
type GetNotificationChannelArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by GetNotificationChannel.
type GetNotificationChannelResult struct {
	Account string `pulumi:"account"`
	ApiKey  string `pulumi:"apiKey"`
	Channel string `pulumi:"channel"`
	Enabled bool   `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string `pulumi:"id"`
	Name                 string `pulumi:"name"`
	NotifyWhenOk         bool   `pulumi:"notifyWhenOk"`
	NotifyWhenResolved   bool   `pulumi:"notifyWhenResolved"`
	Recipients           string `pulumi:"recipients"`
	RoutingKey           string `pulumi:"routingKey"`
	SendTestNotification bool   `pulumi:"sendTestNotification"`
	ServiceKey           string `pulumi:"serviceKey"`
	ServiceName          string `pulumi:"serviceName"`
	Topics               string `pulumi:"topics"`
	Type                 string `pulumi:"type"`
	Url                  string `pulumi:"url"`
	Version              int    `pulumi:"version"`
}

func GetNotificationChannelOutput(ctx *pulumi.Context, args GetNotificationChannelOutputArgs, opts ...pulumi.InvokeOption) GetNotificationChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNotificationChannelResult, error) {
			args := v.(GetNotificationChannelArgs)
			r, err := GetNotificationChannel(ctx, &args, opts...)
			return *r, err
		}).(GetNotificationChannelResultOutput)
}

// A collection of arguments for invoking GetNotificationChannel.
type GetNotificationChannelOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetNotificationChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationChannelArgs)(nil)).Elem()
}

// A collection of values returned by GetNotificationChannel.
type GetNotificationChannelResultOutput struct{ *pulumi.OutputState }

func (GetNotificationChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNotificationChannelResult)(nil)).Elem()
}

func (o GetNotificationChannelResultOutput) ToGetNotificationChannelResultOutput() GetNotificationChannelResultOutput {
	return o
}

func (o GetNotificationChannelResultOutput) ToGetNotificationChannelResultOutputWithContext(ctx context.Context) GetNotificationChannelResultOutput {
	return o
}

func (o GetNotificationChannelResultOutput) Account() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.Account }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.ApiKey }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) Channel() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.Channel }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNotificationChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) NotifyWhenOk() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) bool { return v.NotifyWhenOk }).(pulumi.BoolOutput)
}

func (o GetNotificationChannelResultOutput) NotifyWhenResolved() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) bool { return v.NotifyWhenResolved }).(pulumi.BoolOutput)
}

func (o GetNotificationChannelResultOutput) Recipients() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.Recipients }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) RoutingKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.RoutingKey }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) SendTestNotification() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) bool { return v.SendTestNotification }).(pulumi.BoolOutput)
}

func (o GetNotificationChannelResultOutput) ServiceKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.ServiceKey }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) Topics() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.Topics }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) string { return v.Url }).(pulumi.StringOutput)
}

func (o GetNotificationChannelResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v GetNotificationChannelResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNotificationChannelResultOutput{})
}
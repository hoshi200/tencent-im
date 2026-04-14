/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/9/3 18:47
 * @Desc: 离线推送
 */

package entity

import (
	"github.com/hoshi200/tencent-im/internal/conv"
	"github.com/hoshi200/tencent-im/internal/types"
)

type offlinePush struct {
	pushFlag    int                // 推送标识。0表示推送，1表示不离线推送。
	title       string             // 离线推送标题。该字段为 iOS 和 Android 共用。
	desc        string             // 离线推送内容。
	ext         string             // 离线推送透传内容。
	androidInfo *types.AndroidInfo // Android离线推送消息
	apnsInfo    *types.ApnsInfo    // IOS离线推送消息
}

func newOfflinePush() *offlinePush {
	return &offlinePush{}
}

// SetPushFlag 设置推送消息
func (o *offlinePush) SetPushFlag(pushFlag types.PushFlag) {
	o.pushFlag = int(pushFlag)
}

// SetTitle 设置离线推送标题
func (o *offlinePush) SetTitle(title string) {
	o.title = title
}

// SetDesc 设置离线推送内容
func (o *offlinePush) SetDesc(desc string) {
	o.desc = desc
}

// SetExt 设置离线推送透传内容
func (o *offlinePush) SetExt(ext interface{}) {
	o.ext = conv.String(ext)
}

// SetAndroidSound 设置Android离线推送声音文件路径
func (o *offlinePush) SetAndroidSound(sound string) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.Sound = sound
}

// SetAndroidHuaWeiChannelId 设置华为手机 EMUI 10.0 及以上的通知渠道字段
func (o *offlinePush) SetAndroidHuaWeiChannelId(channelId string) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.HuaWeiChannelID = channelId
}

// SetAndroidXiaoMiChannelId 设置小米手机 MIUI 10 及以上的通知类别（Channel）适配字段
func (o *offlinePush) SetAndroidXiaoMiChannelId(channelId string) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.XiaoMiChannelID = channelId
}

// SetAndroidOppoChannelId 设置OPPO手机 Android 8.0 及以上的 NotificationChannel 通知适配字段
func (o *offlinePush) SetAndroidOppoChannelId(channelId string) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.OPPOChannelID = channelId
}

// SetAndroidGoogleChannelId 设置Google 手机 Android 8.0 及以上的通知渠道字段
func (o *offlinePush) SetAndroidGoogleChannelId(channelId string) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.GoogleChannelID = channelId
}

// SetAndroidVivoClassification 设置VIVO 手机推送消息分类，“0”代表运营消息，“1”代表系统消息，不填默认为1
func (o *offlinePush) SetAndroidVivoClassification(classification types.VivoClassification) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.VIVOClassification = int(classification)
}

// SetAndroidHuaWeiImportance 设置华为推送通知消息分类
func (o *offlinePush) SetAndroidHuaWeiImportance(importance types.HuaWeiImportance) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.HuaWeiImportance = string(importance)
}

// SetAndroidHonorImportance 设置荣耀推送消息分类
func (o *offlinePush) SetAndroidHonorImportance(importance types.HonorImportance) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.HonorImportance = string(importance)
}

// SetAndroidExtAsHuaweiIntentParam 设置在控制台配置华为推送为“打开应用内指定页面”的前提下，传“1”表示将透传内容 Ext 作为 Intent 的参数，“0”表示将透传内容 Ext 作为 Action 参数。不填默认为0。
func (o *offlinePush) SetAndroidExtAsHuaweiIntentParam(param types.HuaweiIntentParam) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.ExtAsHuaweiIntentParam = int(param)
}

// SetAndroidPushStyle 设置Android通知栏样式，0代表默认样式，1代表长文本样式
func (o *offlinePush) SetAndroidPushStyle(style types.AndroidPushStyle) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.PushStyle = int(style)
}

// SetAndroidHuaWeiCategory 设置华为推送消息分类
func (o *offlinePush) SetAndroidHuaWeiCategory(category string) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.HuaWeiCategory = category
}

// SetAndroidOppoCategory 设置OPPO推送消息分类
func (o *offlinePush) SetAndroidOppoCategory(category string) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.OPPOCategory = category
}

// SetAndroidOppoPrivateMsgTemplateId 设置OPPO推送私信模板ID
func (o *offlinePush) SetAndroidOppoPrivateMsgTemplateId(templateId string) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.OPPOPrivateMsgTemplateId = templateId
}

// SetAndroidOppoPrivateTitleParameters 设置OPPO推送标题模板填充参数
func (o *offlinePush) SetAndroidOppoPrivateTitleParameters(params map[string]string) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.OPPOPrivateTitleParameters = params
}

// SetAndroidOppoPrivateContentParameters 设置OPPO推送内容模板填充参数
func (o *offlinePush) SetAndroidOppoPrivateContentParameters(params map[string]string) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.OPPOPrivateContentParameters = params
}

// SetAndroidOppoNotifyLevel 设置OPPO推送通知栏消息提醒等级
func (o *offlinePush) SetAndroidOppoNotifyLevel(level types.OPPONotifyLevel) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.OPPONotifyLevel = int(level)
}

// SetAndroidVivoCategory 设置vivo推送消息分类
func (o *offlinePush) SetAndroidVivoCategory(category string) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.VIVOCategory = category
}

// SetAndroidVivoNotifyType 设置vivo通知类型
func (o *offlinePush) SetAndroidVivoNotifyType(notifyType types.VIVONotifyType) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.VIVONotifyType = int(notifyType)
}

// SetAndroidGooglePriority 设置Google推送通知栏消息优先级
func (o *offlinePush) SetAndroidGooglePriority(priority types.GooglePriority) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.GooglePriority = string(priority)
}

// SetAndroidMeiZuNoticeMsgType 设置魅族推送消息分类
func (o *offlinePush) SetAndroidMeiZuNoticeMsgType(msgType types.MeiZuNoticeMsgType) {
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.androidInfo.MeiZuNoticeMsgType = int(msgType)
}

// SetApnsBadgeMode 设置IOS徽章计数模式
func (o *offlinePush) SetApnsBadgeMode(badgeMode types.BadgeMode) {
	if o.apnsInfo == nil {
		o.apnsInfo = &types.ApnsInfo{}
	}
	o.apnsInfo.BadgeMode = int(badgeMode)
}

// SetApnsTitle 设置APNs推送的标题
func (o *offlinePush) SetApnsTitle(title string) {
	if o.apnsInfo == nil {
		o.apnsInfo = &types.ApnsInfo{}
	}
	o.apnsInfo.Title = title
}

// SetApnsSubTitle 设置APNs推送的子标题
func (o *offlinePush) SetApnsSubTitle(subTitle string) {
	if o.apnsInfo == nil {
		o.apnsInfo = &types.ApnsInfo{}
	}
	o.apnsInfo.SubTitle = subTitle
}

// SetApnsImage 设置APNs携带的图片地址
func (o *offlinePush) SetApnsImage(image string) {
	if o.apnsInfo == nil {
		o.apnsInfo = &types.ApnsInfo{}
	}
	o.apnsInfo.Image = image
}

// SetApnsMutableContent 设置iOS10的推送扩展开关
func (o *offlinePush) SetApnsMutableContent(mutable types.MutableContent) {
	if o.apnsInfo == nil {
		o.apnsInfo = &types.ApnsInfo{}
	}
	o.apnsInfo.MutableContent = int(mutable)
}

// SetImage 设置右侧小图标图片地址
func (o *offlinePush) SetImage(imageUrl string) {
	if o.apnsInfo == nil {
		o.apnsInfo = &types.ApnsInfo{}
	}
	if o.androidInfo == nil {
		o.androidInfo = &types.AndroidInfo{}
	}
	o.apnsInfo.Image = imageUrl
	o.androidInfo.HuaWeiImage = imageUrl
	o.androidInfo.GoogleImage = imageUrl
	o.androidInfo.HonorImage = imageUrl
}

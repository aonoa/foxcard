package card

import "github.com/go-kratos/kratos/v2/errors"

// 错误
const (
	ErrSignInvalidCode int = 10010 // 无效的签名，请检查签名算法
	ErrSignExpireCode  int = 10011 // 签名已过期，签名只在1分钟内有效且是一次性的

	ErrTimestampCode int = 10013 // 时间戳与服务器时间相差过大，请检查timestamp参数
	ErrNonceCode     int = 10014 // 重复的nonce，需保证nonce对于当前软件在60s内不能重复出现

	ErrCardExpireCode int = 10014 // 卡密已过期
	// 10212	卡密已被冻结
	// 10213	卡密超过多开上限

	ErrCartLoginErrCode    int = 10014 // 卡密登陆失败已过期
	ErrCartLoginExpireCode int = 10015 // 卡密登录状态已失效 可能原因：超时未发心跳包、被挤下线、后台冻结、卡用户修改密码。
	ErrDeviceCode          int = 10016 // 未绑定此设备，只能在首次登录绑定的设备上使用
	// 10216	软件未开启解绑设备功能
	// 10217	超出可绑定设备上限，请在已绑定设备上使用

	ErrCardNotFoundCode int = 10019 //卡密不可用（卡密不存在）
	ErrAppNotFoundCode  int = 10020 //软件不存在，请检查app_key

	// 10240	卡密不存在或已被使用
	// 10241	卡密已被使用

	// 10242	账号已存在，用户注册时返回
	// 10243	用户登录密码错误
	// 10244	不能对未使用的卡密进行充值
	// 10250	用户已到期
	// 10252	用户已被冻结
	// 10253	用户超过多开上限
	// 10254	解绑密码不正确
	// 10255	该卡密不可解绑设备
	// 10304	已经是最新版本，获取软件最新版本返回
	// 10401	软件未开启试用，开发者需在后台软件管理配置
	// 10404	请先登录，原因：未登录就调用心跳接口或者登录后超过60秒再调用心跳接口
	// 10405	试用已到期，软件配置为一次性试用，心跳接口返回
	// 10406	本周期试用已到期，软件配置为间隔试用，心跳接口返回
)

var (
	ErrSignInvalid = errors.New(ErrSignInvalidCode, "SignInvalid", "无效的签名，请检查签名算法")

	ErrCartLoginExpire = errors.New(ErrCartLoginExpireCode, "CartLoginExpire", "卡密登录状态已失效")
	ErrDevice          = errors.New(ErrDeviceCode, "ErrDevice", "未绑定此设备")

	ErrCardExpire = errors.New(ErrCardExpireCode, "CardExpire", "卡密已过期")
)

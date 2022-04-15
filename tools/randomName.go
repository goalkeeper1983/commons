package tools

type randomName struct {
	surnameList   []string
	maleNameList  []string
	fmaleNameList []string
	symboList     []string
}

var RandName *randomName

func init() {
	RandName = new(randomName)
	RandName.surnameList = []string{
		"安", "白", "柏", "包", "鲍", "贲", "毕", "卜", "蔡", "曹", "岑", "昌", "常", "车", "陈", "程", "仇", "崔", "单", "邓", "狄", "丁", "董", "窦", "段", "鄂", "樊", "范", "方", "房", "费", "封", "冯", "符", "傅", "高", "郜", "宫", "龚", "顾", "管", "韩", "杭", "郝", "何", "贺", "赫", "洪", "侯", "华", "惠", "霍", "嵇", "吉", "纪", "家", "贾", "江", "姜", "蒋", "解", "经", "景", "咎", "堪", "柯", "孔", "雷", "黎", "李", "郦", "连", "廉", "粱", "蔺", "凌", "刘", "柳", "卢", "鲁", "陆", "罗", "骆", "吕", "马", "孟", "宓", "苗", "闵", "缪", "莫", "牧", "慕", "穆", "聂", "宁", "钮", "潘", "庞", "裴", "彭", "戚", "齐", "祁", "钱", "乔", "秦", "邱", "秋", "裘", "屈", "任", "荣", "阮", "韶", "邵", "沈", "盛", "施", "石", "史", "宋", "苏", "孙", "汤", "唐", "陶", "滕", "田", "童", "万", "汪", "王", "韦", "卫", "魏", "温", "翁", "邬", "巫", "吴", "伍", "武", "席", "夏", "向", "项", "萧", "谢", "邢", "徐", "许", "宣", "薛", "荀", "严", "阎", "颜", "燕", "羊", "杨", "姚", "叶", "尹", "应", "尤", "於", "于", "余", "俞", "虞", "禹", "郁", "喻", "元", "袁", "翟", "詹", "张", "章", "赵", "甄", "郑", "支", "仲", "周", "朱", "诸", "卓", "宗", "邹", "祖", "左", "百里", "东方", "归海", "狐", "闾丘", "慕容", "眞心", "嗜血", "寂寞", "时埫", "亲吻", "飘簃", "嗿翫", "败镓", "蒍爱", "貑娤", "栤葑", "专情", "鈊", "儍儍", "从头", "伪装", "凮残", " 乱", "纯黒", "纯", "习惯", "谜纞", "缘", "菰独", "为嗳", "昨兲", "莰牁", "LOVE", "贁家耔", "目飚", "番", " 堕络", "吥愿", "绯紸流", "洫葬", "ヌ寸", "楿烟", "华丽", "叼烟", "独特", "⒑指", "楿烟",
	}

	RandName.maleNameList = []string{
		"霸", "白易", "白芷", "百川", "百招", "败", "班", "半鬼", "半山", "半仙", "半邪", "半雪", "悲", "匕", "弼", "碧", "碧空", "冰", "冰安", "邴", "炳", "伯云", "驳", "博", "博超", "博乐", "博涛", "捕", "不二", "不凡", "不言", "不尤", "沧海", "涔", "颤", "成败", "成风", "成仁", "成危", "成威", "成协", "城", "乘风", "乘云", "仇天", "仇血", "酬海", "丑", "初晴", "初彤", "储", "聪健", "聪展", "从安", "耷", "达", "大楚", "大开", "大有", "代丝", "代桃", "岱周", "惮", "道罡", "道天", "道消", "道之", "德地", "德天", "定帮", "断秋", "断天", "断缘", "恶天", "尔蓝", "尔曼", "尔阳", "凡", "凡梦", "鲂", "访琴", "飞风", "飞扬", "非笑", "绯", "匪", "风华", "枫", "复天", "富", "罡", "高烽", "阁", "孤风", "谷波", "谷槐", "谷雪", "广山", "广缘", "鬼神", "海亦", "函", "翰", "行恶", "行天", "豪", "豪英", "昊强", "昊焱", "浩阑", "浩然", "浩天", "浩轩", "浩宇", "皓轩", "鹤", "鹤轩", "珩", "弘文", "鸿", "鸿涛", "鸿煊", "豁", "汲", "疾", "寄灵", "嘉熙", "嘉懿", "蹇", "建辉", "剑", "剑成", "剑愁", "剑封", "剑鬼", "剑身", "剑通", "剑心", "健柏", "焦", "金鑫", "锦程", "瑾瑜", "晋鹏", "靳", "荆", "井", "靖", "靖仇", "靖琪", "靖易", "绝山", "绝施", "绝悟", "绝义", "君浩", "钧", "俊驰", "峻熙", "开山", "楷瑞", "康", "寇", "奎", "傀儡", "坤", "蓝血", "老黑", "老九", "老三", "老四", "老头", "老五", "烙", "乐驹", "勒", "冷安", "冷之", "黎昕", "厉", "立诚", "立果", "立辉", "立轩", "戾", "隶", "连碧", "连虎", "怜烟", "凛", "灵竹", "翎", "聋五", "鹭洋", "栾", "满天", "糜", "觅松", "觅云", "妙海", "灭龙", "明辉", "明杰", "明轩", "明雪", "冥", "冥幽", "铭", "慕蕊", "穆", "宁远", "南风", "南松", "难摧", "难敌", "难破", "难胜", "念波", "凝海", "牛青", "盼秋", "盼旋", "沛白", "沛山", "鹏飞", "鹏涛", "鹏笑", "鹏煊", "平灵", "平文", "乞", "绮彤", "契", "千筹", "千万", "虔", "乾", "强炫", "青寒", "青文", "卿", "擎", "擎苍", "擎汉", "擎宇", "磬", "秋", "秋白", "秋寒", "裘", "人达", "人杰", "人龙", "人雄", "人英", "戎", "荣轩", "如豹", "如花", "睿渊", "若风", "若剑", "若之", "三德", "三毒", "三问", "山河", "山柳", "绍辉", "胜", "晟睿", "师", "诗云", "十八", "十三", "士晋", "士萧", "世德", "世倌", "世开", "世立", "世平", "似狮", "筮", "书芹", "水云", "思远", "松", "松思", "嵩", "送终", "随阴", "傥", "涛", "天霸", "天川", "天德", "天磊", "天奇", "天寿", "天抒", "天思", "天问", "天佑", "天与", "天宇", "铁身", "听白", "霆", "万仇", "万恶", "万声", "万天", "万言", "万怨", "忘幽", "威", "伟帮", "伟宸", "伟诚", "伟祺", "伟泽", "文博", "文昊", "文龙", "文轩", "紊", "问安", "问筠", "问旋", "乌", "无敌", "无极", "无剑", "无声", "无施", "无血", "无招", "稀", "祥", "骁", "萧", "嚣", "晓博", "晓啸", "笑白", "笑天", "鑫", "鑫磊", "鑫鹏", "修杰", "修洁", "旭尧", "续", "寻菡", "逊", "雅阳", "延恶", "严青", "岩", "炎彬", "奄", "焱", "鞅", "扬", "羊青", "仰", "遥", "夜白", "烨华", "烨磊", "烨霖", "烨伟", "一刀", "一德", "一江", "一鸣", "一手", "一笑", "一斩", "伊", "亿先", "邑", "羿", "熠彤", "翼", "懿轩", "易安", "义安", "瑛", "鹰", "迎松", "映波", "雍", "友绿", "幼南", "幼旋", "愚志", "雨泽", "雨珍", "语堂", "煜城", "煜祺", "誉", "渊思", "元龙", "元正", "沅", "垣", "源智", "远锋", "远航", "远侵", "远山", "远望", "苑博", "越彬", "越泽", "云飞", "泽洋", "斩", "哲瀚", "箴", "臻", "振家", "正豪", "之玉", "芷烟", "志泽", "致远", "智宸", "中道", "中恶", "铸海", "追命", "子骞", "子默", "子轩", "自中", "醉山", "醉薇", "┳═一", " 缺德", "浪子", "仇魜", "男孖", "鈊痛", "王子", "爱", "眷恋", "鈈哓锝", "鈈暧祢", "销夨", "流蒗", "涛", "儛步", "洅唻", "泪氺", "訫情", "拽児", "絯孑气 ", "呼晞 ", "谈訫", "尐壊疍", "扎念", "訫谇", "冭子", "终点", "独儛", "栤葑",
	}

	RandName.fmaleNameList = []string{
		"竺", "惜雪", "凌香", "颦", "高丽", "衫", "藏今", "素", "小小", "善若", "舞仙", "夏青", "珊", "萃", "柔", "沁", "血茗", "凌瑶", "赛君", "玉兰", "梦琪", "秀", "翠", "晓霜", "夜蕾", "芹", "虔纹", "嫣娆", "素阴", "一寡", "亚男", "珊珊", "青荷", "唯雪", "若男", "老姆", "嫣", "清炎", "紫伊", "千愁", "若颜", "飞飞", "英姑", "不惜", "莫英", "茹妖", "凤凰", "凤之", "贞", "四娘", "凤妖", "无颜", "姒", "凡柔", "紫南", "莛", "青槐", "新蕾", "初丹", "可燕", "荧", "宛", "芸", "不乐", "雅琴", "幻柏", "沛菡", "瑛", "若冰", "无心", "藏鸟", "如音", "大凄", "紫烟", "琦", "如冰", "绾绾", "如娆", "不愁", "含玉", "秋翠", "芷蕾", "莆", "谷丝", "幻姬", "傲之", "含烟", "曼文", "太兰", "妍", "芝", "冰蓝", "惠", "莫言", "娩", "菀", "惋清", "芯", "从蓉", "灵薇", "迎梦", "妙菡", "荠", "愫", "毒娘", "凡蕾", "幻悲", "莞", "新柔", "怡", "彤", "不二", "宛儿", "冰彤", "芷", "翠萱", "绣连", "海莲", "曼易", "紫翠", "珠", "半梦", "凤", "晓亦", "听寒", "以亦", "苡", "元霜", "靖柏", "羞花", "芮", "蓉", "涵双", "尔珍", "醉易", "樱", "颖", "映易", "盈", "蓝", "灵煌", "一一", "宫苴", "盛男", "缘郡", "语蓉", "一曲", "代云", "一兰", "之瑶", "冰露", "南莲", "雪柳", "葶", "丝", "幻然", "芾", "如雪", "灵槐", "紫蓝", "绝音", "莫茗", "太君", "邪欢", "碧凡", "襄", "夜蓉", "梦柏", "凡旋", "妖妖", "以蕊", "斓", "萝", "起眸", "傻姑", "闭月", "寡妇", "行云", "东蒽", "珍", "落雁", "傲珊", "沉鱼", "萤", "如霜", "莹芝", "夏槐", "元风", "语兰", "汝燕", "从寒", "亦寒", "静曼", "觅波", "香", "小宸", "朝雪", "雁卉", "荟", "念芹", "灵", "香魔", "白易", "蛟凤", "之桃", "可兰", "书瑶", "元柏", "从梦", "可冥", "平安", "靖雁", "翠风", "千琴", "莺", "摇伽", "茈", "芷蕊", "中蓝", "幻灵", "雪青", "涑", "溪灵", "灭绝", "芳", "白晴", "依珊", "阑悦", "冰枫", "安蕾", "友绿", "书兰", "谷雪", "曼香", "雁玉", "冰旋", "冰萍", "访烟", "乐菱", "听兰", "幻雪", "菲音", "妖丽", "丹妗", "莹", "紫山", "代荷", "怜蕾", "苠", "善斓", "翎", "雨寒", "鸣凤", "易烟", "茹嫣", "问晴", "夏彤", "忆翠", "访蕊", "桐", "柏柳", "宝川", "雁", "宝莹", "艳一", "绫", "阑香", "紫寒", "可仁", "纹", "寻真", "雪一", "茗茗", "荧荧", "芙蓉", "踏歌", "紫", "三娘", "亦瑶", "凝天", "聋五", "醉蓝", "凤灵", "乐瑶", "香寒", "凡英", "颜", "菲鹰", "藏花", "婴", "淇", "若", "嫣然", "颜演", "忻", "平蝶", "雅蕊", "涵菡", "老太", "傲芙", "姝", "冰姬", "三颜", "慕青", "筝", "兰", "一凤", "海云", "无春", "飞凤", "金连", "代真", "二娘", "语琴", "悒", "山柏", "芸遥", "之柔", "涫", "黎云", "绮", "若血", "雪冥", "涔雨", "真", "稚晴", "青", "清涟", "绿真", "书翠", "怜梦", "怀亦", "丹珍", "从彤", "雯", "问旋", "敏", "湘", "书蕾", "宛海", "易梦", "衣", "幻嫣", "不平", "若魔", "灭男", "妙芹", "善愁", "外绣", "太英", "卿", "太清", "如花", "惋庭", "半芹", "冰兰", "含蕾", "芙", "寒云", "新波", "梦曼", "傲晴", "碧彤", "映寒", "茗", "听白", "萍", "无色", "苑睐", "若烟", "尔岚", "可愁", "洙", "婷冉", "双双", "澜", "沂", "从安", "梦松", "水桃", "冰夏", "千兰", "凌萱", "雁开", "初夏", "魂幽", "代男", "大娘", "涟妖", "笙", "赛凤", "斌", "琳", "笑蓝", "念露", "媚颜", "亦玉", "姿", "香芦", "海露", "妙竹", "璎", "灵雁", "幻珊", "天荷", "冷卉", "涵柏", "慕灵", "语芙", "傲蕾", "访旋", "书雪", "清", "怀蝶", "不弱", "玲", "秋尽", "盼山", "梨愁", "又菱", "艳血", "葵阴", "冥茗", "访卉", "雅绿", "妙彤", "沛珊", "岂愈", "向珊", "惜文", "芫", "问筠", "念薇", "听蓉", "惜霜", "如萱", "幼翠", "海安", "乐枫", "冰颜", "弱", "艳", "不悔", "夏菡", "乐萱", "若南", "夜阑", "怜菡", "代曼", "代萱", "紫真", "千青", "凌寒", "紫安", "寒安", "怀蕊", "秋荷", "涵雁", "以山", "凡梅", "盼曼", "翠彤", "谷冬", "冷安", "千萍", "冰烟", "幼蓉", "以蓝", "笑寒", "忆寒", "秋烟", "芷巧", "水香", "映之", "醉波", "幻莲", "夜山", "芷卉", "向彤", "小玉", "涵蕾", "碧菡", "映秋", "盼烟", "忆山", "以寒", "寒香", "小凡", "代亦", "梦露", "友蕊", "寄凡", "雁枫", "水绿", "曼荷", "笑珊", "寒珊", "谷南", "慕儿", "夏岚", "友儿", "小萱", "紫青", "妙菱", "冬寒", "曼柔", "语蝶", "青筠", "夜安", "觅海", "晓槐", "雅山", "访云", "翠容", "寒凡", "晓绿", "以菱", "冬云", "访枫", "含卉", "夜白", "灵竹", "醉薇", "元珊", "幻波", "盼夏", "元瑶", "迎曼", "紫霜", "凌旋", "孤丝", "怜寒", "凡松", "青丝", "翠安", "如天", "凌雪", "绮菱", "香薇", "冬灵", "凌珍", "沛文", "紫槐", "采文", "雪旋", "盼海", "映梦", "安雁", "映容", "凝阳", "访风", "天亦", "觅风", "小霜", "雪萍", "半雪", "山柳", "靖易", "白薇", "梦菡", "飞绿", "如波", "又晴", "友易", "香菱", "冬亦", "问雁", "海冬", "秋灵", "凝芙", "念烟", "白山", "从灵", "尔芙", "迎蓉", "念寒", "翠绿", "翠芙", "靖儿", "妙柏", "千凝", "小珍", "妙旋", "雪枫", "绮琴", "雨双", "听枫", "觅荷", "凡之", "晓凡", "雅彤", "幻丝", "代梅", "青亦", "元菱", "海瑶", "飞槐", "听露", "梦岚", "幻竹", "谷云", "忆霜", "水瑶", "慕晴", "秋双", "雨真", "觅珍", "丹雪", "元枫", "思天", "如松", "妙晴", "谷秋", "妙松", "晓夏", "宛筠", "碧琴", "盼兰", "小夏", "安容", "青曼", "千儿", "寻双", "涵瑶", "冷梅", "秋柔", "思菱", "醉柳", "迎夏", "向雪", "以丹", "依凝", "如柏", "雁菱", "凝竹", "宛白", "初柔", "南蕾", "书萱", "梦槐", "南琴", "绿海", "沛儿", "晓瑶", "凝蝶", "紫雪", "念双", "念真", "曼寒", "凡霜", "飞雪", "雪兰", "雅霜", "冷雪", "靖巧", "翠丝", "觅翠", "凡白", "乐蓉", "迎波", "丹烟", "梦旋", "书双", "念桃", "夜天", "安筠", "觅柔", "初南", "秋蝶", "千易", "安露", "诗蕊", "山雁", "友菱", "香露", "晓兰", "白卉", "语山", "冷珍", "夏柳", "如之", "忆南", "书易", "翠桃", "寄瑶", "如曼", "问柳", "幻桃", "又菡", "醉蝶", "亦绿", "诗珊", "听芹", "新之", "易巧", "念云", "晓灵", "静枫", "夏蓉", "如南", "幼丝", "紫文", "凌晴", "傲安", "初蝶", "代芹", "诗霜", "碧灵", "诗柳", "采白", "慕梅", "乐安", "冬菱", "宛凝", "雨雪", "易真", "安荷", "静竹", "代柔", "丹秋", "绮梅", "依白", "凝荷", "幼珊", "忆彤", "凌青", "芷荷", "听荷", "代玉", "念珍", "梦菲", "夜春", "千秋", "白秋", "谷菱", "飞松", "初瑶", "惜灵", "梦易", "新瑶", "曼梅", "碧曼", "友瑶", "雨兰", "夜柳", "芷珍", "含芙", "夜云", "依萱", "凝雁", "以莲", "幼晴", "尔琴", "飞阳", "白凡", "沛萍", "雪瑶", "向卉", "乐珍", "寒荷", "觅双", "白桃", "安卉", "盼雁", "乐松", "涵山", "问枫", "以柳", "含海", "翠曼", "忆梅", "涵柳", "海蓝", "晓曼", "代珊", "忆丹", "静芙", "绮兰", "梦安", "紫丝", "千雁", "凝珍", "香萱", "梦容", "冷雁", "飞柏", "天真", "翠琴", "寄真", "初雪", "雅柏", "怜容", "如风", "南露", "紫易", "冰凡", "海雪", "碧玉", "语风", "凝梦", "从雪", "白枫", "傲云", "白梅", "慕凝", "雅柔", "盼柳", "半青", "从霜", "怀柔", "怜晴", "代双", "以南", "若菱", "芷文", "南晴", "梦寒", "初翠", "灵波", "问夏", "惜海", "亦旋", "沛芹", "幼萱", "白凝", "初露", "迎海", "绮玉", "寻芹", "秋柳", "尔白", "映真", "含雁", "寒松", "寻雪", "青烟", "问蕊", "灵阳", "雪巧", "丹萱", "凡双", "孤萍", "紫菱", "寻凝", "傲柏", "傲儿", "友容", "灵枫", "尔丝", "曼凝", "若蕊", "问丝", "思枫", "水卉", "问梅", "诗双", "翠霜", "夜香", "寒蕾", "凡阳", "冷玉", "平彤", "语薇", "紫夏", "凌波", "芷蝶", "丹南", "之双", "凡波", "思雁", "白莲", "从菡", "如容", "采柳", "沛岚", "惜儿", "夜玉", "水儿", "半凡", "语海", "听莲", "幻枫", "念柏", "冰珍", "思山", "凝蕊", "天玉", "思萱", "向梦", "笑南", "夏旋", "之槐", "元灵", "以彤", "采萱", "巧曼", "绿兰", "平蓝", "问萍", "绿蓉", "迎蕾", "思卉", "白柏", "怜阳", "雨柏", "雁菡", "梦之", "又莲", "乐荷", "寒天", "凝琴", "书南", "映天", "平露", "含巧", "慕蕊", "半莲", "醉卉", "天菱", "青雪", "雅旋", "巧荷", "飞丹", "若灵", "尔云", "幻天", "诗兰", "青梦", "海菡", "忆秋", "寒凝", "绮山", "静白", "尔蓉", "尔冬", "映萱", "白筠", "冰双", "访彤", "绿柏", "夏云", "笑翠", "含双", "盼波", "以云", "怜翠", "雁风", "之卉", "平松", "问儿", "绿柳", "如蓉", "曼容", "天晴", "丹琴", "惜天", "寻琴", "依瑶", "涵易", "忆灵", "从波", "依柔", "问兰", "山晴", "怜珊", "之云", "飞双", "傲白", "沛春", "雨南", "笑阳", "代容", "友琴", "雁梅", "友桃", "从露", "语柔", "傲玉", "觅夏", "晓蓝", "新晴", "雨莲", "凝旋", "绿旋", "幻香", "冷亦", "忆雪", "友卉", "幻翠", "靖柔", "寻菱", "丹翠", "安阳", "雅寒", "惜筠", "尔安", "雁易", "飞瑶", "夏兰", "沛蓝", "静丹", "山芙", "笑晴", "新烟", "笑旋", "雁兰", "凌翠", "秋莲", "书桃", "傲松", "语儿", "映菡", "初曼", "听云", "雅香", "语雪", "初珍", "白安", "冰薇", "诗槐", "冰巧", "夏寒", "诗筠", "新梅", "白曼", "安波", "从阳", "含桃", "曼卉", "笑萍", "晓露", "水彤", "安彤", "乐巧", "依风", "亦丝", "易蓉", "紫萍", "惜萱", "诗蕾", "寻绿", "寻云", "孤丹", "谷蓝", "山灵", "友梅", "从云", "雁丝", "冰香", "依玉", "冰之", "妙梦", "以冬", "曼青", "冷菱", "雪曼", "安白", "千亦", "凌蝶", "又夏", "南烟", "沛凝", "翠梅", "书文", "雪卉", "乐儿", "傲丝", "安青", "寄灵", "惜寒", "雨竹", "冬莲", "绮南", "翠柏", "平凡", "孤兰", "秋珊", "新筠", "夏瑶", "念文", "晓丝", "雁凡", "谷兰", "灵凡", "凝云", "曼云", "丹彤", "南霜", "夜梦", "从筠", "雁芙", "依波", "晓旋", "念之", "盼芙", "曼安", "采珊", "初柳", "迎天", "南珍", "妙芙", "语柳", "含莲", "晓筠", "夏山", "尔容", "念梦", "傲南", "问薇", "雨灵", "凝安", "冰海", "宛菡", "冬卉", "盼晴", "冷荷", "寄翠", "幻梅", "如凡", "语梦", "千柔", "向露", "梦玉", "傲霜", "依霜", "灵松", "诗桃", "书蝶", "冰蝶", "山槐", "以晴", "梦桃", "孤云", "水蓉", "雅容", "飞烟", "雁荷", "代芙", "夏烟", "依秋", "紫萱", "忆之", "幻巧", "水风", "安寒", "白亦", "怜雪", "听南", "念蕾", "梦竹", "千凡", "寄琴", "采波", "元冬", "平卉", "笑柳", "谷梦", "绿蝶", "飞荷", "孤晴", "曼冬", "尔槐", "以旋", "绿蕊", "依丝", "怜南", "千山", "雨安", "寄柔", "幼枫", "凡桃", "新儿", "夏波", "雨琴", "静槐", "元槐", "映阳", "飞薇", "小凝", "傲菡", "谷蕊", "笑槐", "飞兰", "笑卉", "迎荷", "书竹", "半烟", "绮波", "小之", "觅露", "夜雪", "寒梦", "尔风", "雨旋", "芷珊", "山彤", "尔柳", "沛柔", "灵萱", "白容", "映安", "依云", "映冬", "凡雁", "梦秋", "梦凡", "若云", "元容", "怀蕾", "灵寒", "天薇", "白风", "访波", "亦凝", "易绿", "夜南", "曼凡", "亦巧", "青易", "冰真", "白萱", "友安", "诗翠", "雪珍", "海之", "小蕊", "又琴", "香彤", "惜蕊", "迎彤", "沛白", "雁山", "雪晴", "冰绿", "半梅", "笑容", "念瑶", "如冬", "向真", "亦云", "向雁", "尔蝶", "冬易", "丹亦", "醉香", "孤菱", "安莲", "问凝", "冬萱", "晓山", "雁蓉", "梦蕊", "山菡", "凝丝", "怀梦", "雨梅", "冷霜", "向松", "迎丝", "迎梅", "听双", "山蝶", "夜梅", "醉冬", "雨筠", "半蕾", "幼菱", "寻梅", "含之", "香之", "含蕊", "靖荷", "碧萱", "向南", "书雁", "怀薇", "忆文", "若山", "向秋", "绮烟", "从蕾", "天曼", "又亦", "依琴", "曼彤", "沛槐", "又槐", "元绿", "安珊", "夏之", "易槐", "宛亦", "白翠", "丹云", "问寒", "易文", "傲易", "青旋", "思真", "妙之", "半双", "若翠", "初兰", "怀曼", "惜萍", "初之", "宛丝", "幻儿", "千风", "天蓉", "雅青", "寄文", "代天", "惜珊", "向薇", "惜芹", "谷芹", "雁桃", "映雁", "寄风", "绮晴", "傲柔", "寄容", "以珊", "芷容", "书琴", "寻桃", "涵阳", "怀寒", "易云", "采蓝", "代秋", "惜梦", "尔烟", "谷槐", "怀莲", "涵菱", "水蓝", "访冬", "半兰", "又柔", "安双", "冰岚", "语芹", "静珊", "幻露", "访天", "静柏", "凌丝", "小翠", "访文", "凌文", "芷云", "思柔", "巧凡", "慕山", "千柳", "从凝", "安梦", "香旋", "安柏", "平萱", "以筠", "忆曼", "新竹", "绮露", "觅儿", "碧蓉", "白竹", "曼雁", "雁露", "凝冬", "含灵", "初阳", "海秋", "盼易", "思松", "梦山", "友灵", "绿竹", "灵安", "凌柏", "又蓝", "尔竹", "天蓝", "青枫", "问芙", "灵珊", "凝丹", "小蕾", "水之", "飞珍", "亦竹", "飞莲", "海白", "元蝶", "芷天", "怀绿", "元芹", "寒烟", "听筠", "采梦", "凝莲", "元彤", "觅山", "代桃", "冷之", "盼秋", "秋寒", "海亦", "初晴", "巧蕊", "听安", "芷雪", "以松", "寒梅", "香岚", "孤容", "晓蕾", "安萱", "夜绿", "雪莲", "从丹", "雨文", "幼荷", "青柏", "初蓝", "忆安", "寻冬", "雪珊", "迎南", "如彤", "采枫", "若雁", "翠阳", "沛容", "山兰", "芷波", "寄云", "慕卉", "冷松", "涵梅", "书白", "乐天", "宛秋", "傲旋", "凡儿", "夏真", "乐双", "白玉", "问玉", "寄松", "丹蝶", "访曼", "代灵", "丹寒", "访梦", "绿凝", "冰菱", "语蕊", "思烟", "忆枫", "映菱", "凌兰", "曼岚", "若枫", "傲薇", "凡灵", "乐蕊", "╰☆╮", "时尚", "寳呗", " 梅?婲", "脾气", "鈊痛", "爱", "眷恋", "绵婲?餹", "鈈暧祢", "销夨", "吖头", "泪氺", "静", "潴", "儛步", "訫情", "萤唲", "絯孑气 ", "呼晞 ", "谈訫", "ロ勿", "@灵", "飘雪", "钕孖", "訫谇", "飘零", "籹孒", "没魜疼", "坏蜑", "傻钕釨", "玥",
	}
	RandName.symboList = []string{
		"＆", "¤", "°", "ゞ", "〃", "︶", "┳", "﹌", "╰", "☆", "╮", "╇", "捻", "°°", "ゞ", "⒐⒋", "、", "ヤ", "ωǒ ", "de", "〃", "ご", "★", "ぁ", "ve", "ィ", "啲諎", "⒛", "⑦", "伱", "鳡觉", "Dě", "撧橼", "葬", "在乎", "づ", "Ю", "Sè ", "〓", "Ω", "℉", "ノ", "⑧", "ぷ", "Ц", "しov", "Ъú", "尐", "Ⅱ", "吇甴", "の", "Let\"s", "↘", "Nà", "Yi", "ー辈孓", "乞ぺ", "爱", "错メ过", "ミ", "※", "灬", "┕", "ā", "｛", "＆", "恛メ忆", "潴", "Ψ", "◇", "Bo", "⊙", "⊕", "◎", "の", "⊿", "¢", "↙", "←", "↖", "↑", "↗", "→", "卍", "♀", "♂", "∷", "№", "◣", "Ψ", "€", "§", "囍", "∮", "∝", "≈", "㏒", "‰", "Φ", "≌", "γ", "Ω", "β", "α", "Δ", "ε", "ζ", "η", "κ", "λ", "μ", "π", "ρ", "υ", "φ", "ψ", "＄", "￥", "㏄", "○", "404", "夬", "亅", "丂", "厃", "〦", "〨", "〧", "〢", "〣", "〆", "〤", "ぁ",
	}
}

func (This *randomName) RandomName(robotIndex, sex, needAddNum int32) string {
	index := 0
	strIndex := Int32ToString(robotIndex)
	strUserName := "'"

	f := func() string {
		if GetRandInt(0, 1) != 0 {
			index = GetRandInt(0, len(This.symboList)-1)
			if len(strUserName)+len(This.symboList[index]) < 8 && len(strUserName)+len(strIndex)+len(This.symboList[index]) < 20 {
				return This.symboList[index]
			}
		}
		return ""
	}

	strUserName = f()
	index = GetRandInt(0, len(This.surnameList)-1)
	strUserName += This.surnameList[index]
	strUserName += f()

	if sex == 1 {
		index = GetRandInt(0, len(This.maleNameList)-1)
		strUserName += This.maleNameList[index]
	} else {
		index = GetRandInt(0, len(This.maleNameList)-1)
		strUserName += This.maleNameList[index]
	}

	strUserName += f()

	if needAddNum == 1 {
		if len(strIndex)+len(strUserName) < 20 {
			strUserName += strIndex
		}
	}
	return strUserName
}
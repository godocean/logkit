package transforms

import (
	"github.com/qiniu/logkit/sender"
	"github.com/qiniu/logkit/utils"
)

const (
	KeyType = "type"
)

const (
	TransformTypeString = "string"
	TransformTypeLong   = "long"
	TransformTypeFloat  = "float"
)

const (
	StageBeforeParser = "before_parser"
	StageAfterParser  = "after_parser"
)

//Transformer plugin做数据变换的接口
// 注意： transform的规则是，出错要把数据原样返回
type Transformer interface {
	Description() string
	SampleConfig() string
	ConfigOptions() []utils.Option
	Type() string
	Transform([]sender.Data) ([]sender.Data, error)
	RawTransform([]string) ([]string, error)
	Stage() string
	Stats() utils.StatsInfo
}

type Creator func() Transformer

var Transformers = map[string]Creator{}

func Add(name string, creator Creator) {
	Transformers[name] = creator
}

var (
	KeyStage = utils.Option{
		KeyName:       "stage",
		ChooseOnly:    true,
		ChooseOptions: []interface{}{StageAfterParser, StageBeforeParser},
		Default:       StageAfterParser,
		DefaultNoUse:  false,
		Description:   "transform运行的阶段(parser前还是parser后)(stage)",
		Type:          TransformTypeString,
	}
	KeyStageAfterOnly = utils.Option{
		KeyName:       "stage",
		ChooseOnly:    true,
		ChooseOptions: []interface{}{StageAfterParser},
		Default:       StageAfterParser,
		DefaultNoUse:  false,
		Description:   "transform运行的阶段(stage)",
		Type:          TransformTypeString,
	}
	KeyFieldName = utils.Option{
		KeyName:      "key",
		ChooseOnly:   false,
		Default:      "my_field_keyname",
		DefaultNoUse: true,
		Description:  "要进行Transform变化的键(key)",
		Type:         TransformTypeString,
	}
	KeyTimezoneoffset = utils.Option{
		KeyName:    "offset",
		ChooseOnly: true,
		ChooseOptions: []interface{}{0, -1, -2, -3, -4,
			-5, -6, -7, -8, -9, -10, -11, -12,
			1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12},
		Default:      "0",
		DefaultNoUse: false,
		Description:  "时区偏移量(offset)",
		CheckRegex:   "*",
		Type:         TransformTypeLong,
	}
)

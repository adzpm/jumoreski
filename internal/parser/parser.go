package parser

import "github.com/SevereCloud/vksdk/v2/api"

type (
	Parser struct {
		cfg   *Config
		vkapi *api.VK
		count uint64
	}
)

func NewParser(cfg *Config) *Parser {
	return &Parser{
		cfg: cfg,
	}
}

func (p *Parser) Start() error {
	p.vkapi = api.NewVK(p.cfg.AccessToken)
	return nil
}

func (p *Parser) GetCount() (uint64, error) {
	return 0, nil
}

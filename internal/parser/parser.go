package parser

import (
	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/adzpm/jumoreski/models"
	"gorm.io/gorm"
	"log"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
)

const (
	countOfRequestedPosts = 100
)

type (
	Parser struct {
		cfg     *Config               // cfg is a config
		storage *gorm.DB              // storage is a storage
		vkapi   *api.VK               // vkapi is a vk api
		count   int                   // count is a count of posts
		Posts   []*models.StoragePost // Posts is a slice of posts
	}
)

func NewParser(cfg *Config, storage *gorm.DB) *Parser {
	return &Parser{
		cfg:     cfg,
		storage: storage,
	}
}

func (p *Parser) Start() error {
	var (
		err   error
		count int
	)

	p.vkapi = api.NewVK(p.cfg.AccessToken)

	if count, err = p.getCount(); err != nil {
		return err
	}

	p.count = count

	if err = p.parseWall(); err != nil {
		return err
	}

	return nil
}

func (p *Parser) parseWall() error {
	var (
		err      error
		allPosts = make([]*models.StoragePost, 0)
		resp     api.WallGetResponse
	)

	for offset := 0; offset < p.count; offset += countOfRequestedPosts {
		log.Println("offset: ", offset)

		prm := params.NewWallGetBuilder().
			Domain(p.cfg.GroupDomain).
			Count(countOfRequestedPosts).
			Offset(offset).Params

		if resp, err = p.vkapi.WallGet(prm); err != nil {
			continue
		}

		for _, post := range resp.Items {
			if post.MarkedAsAds {
				continue
			}

			images := models.Images{}

			for _, attachment := range post.Attachments {
				switch attachment.Type {
				case object.AttachmentTypeGraffiti:
					images = append(images, attachment.Graffiti.Photo586)
				case object.AttachmentTypePhoto:
					images = append(images, attachment.Photo.Sizes[len(attachment.Photo.Sizes)-1].URL)
				case object.AttachmentTypePostedPhoto:
					images = append(images, attachment.PostedPhoto.Photo604)
				case object.AttachmentTypeDoc:
					if attachment.Doc.Ext == "gif" {
						images = append(images, attachment.Doc.URL)
					}
				}
			}

			if err = p.storage.Create(&models.StoragePost{
				Text:    post.Text,
				Watches: post.Views.Count,
			}).Error; err != nil {
				log.Println("error of insert into database: ", err)
				continue
			}
		}

	}

	p.Posts = allPosts

	return nil
}

func (p *Parser) getCount() (int, error) {
	prm := params.NewWallGetBuilder().
		Domain(p.cfg.GroupDomain).
		Count(1).
		Offset(0).Params

	resp, err := p.vkapi.WallGet(prm)
	if err != nil {
		return 0, err
	}

	return resp.Count, nil
}

func (p *Parser) CountOfPosts() int { return p.count }

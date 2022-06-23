package cmd

import (
	"fmt"
	"github.com/blugelabs/bluge"
	segment "github.com/blugelabs/bluge_segment_api"
	"github.com/urfave/cli/v2"
	"log"
)

var MemoCMD = &cli.Command{
	Name: "memo",
	Subcommands: []*cli.Command{
		{
			Name: "get",
			Action: func(ctx *cli.Context) error {
				key := ctx.Args().First()
				reader, err := bluge.OpenReader(bluge.DefaultConfig("memo"))
				if err != nil {
					log.Fatalf("error opening reader: %v", err)
				}
				q := bluge.NewMatchQuery(key)
				q.SetField("content")
				dmi, _ := reader.Search(ctx.Context, bluge.NewAllMatches(q))
				next, err := dmi.Next()
				for err == nil && next != nil {
					err = next.VisitStoredFields(func(field string, value []byte) bool {
						fmt.Println(field, string(value))
						return true
					})
					fmt.Println()
					if err != nil {
						log.Fatalf("error accessing stored fields: %v", err)
					}
					next, err = dmi.Next()
				}
				return nil
			},
		},
	},
	Action: func(ctx *cli.Context) error {
		content := ctx.Args().First()
		fmt.Println(content)
		config := bluge.DefaultConfig("memo")
		writer, err := bluge.OpenWriter(config)
		if err != nil {
			log.Fatalf("error opening writer: %v", err)
		}
		defer writer.Close()
		doc := bluge.NewDocument("memo").
			AddField(bluge.NewTextField("content", content))
		doc.EachField(func(field segment.Field) {
			fmt.Println(string(field.Name()), string(field.Value()))
		})
		err = writer.Update(doc.ID(), doc)
		if err != nil {
			log.Fatalf("error updating document: %v", err)
		}
		return nil
	},
}

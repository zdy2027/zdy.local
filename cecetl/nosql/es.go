package nosql

import (
	"context"
	"zdy.local/utils/nlogs"
	"gopkg.in/olivere/elastic.v5"
	"github.com/astaxie/beego"
	"encoding/json"
	"zdy.local/cecetl/message"
)

type ES struct {
	//var Esip = beego.AppConfig.String("es::host")+":"+beego.AppConfig.String("es::port")
	EClient *elastic.Client
}

func (e *ES)Init(){
	var err error
	e.EClient, err = elastic.NewClient(elastic.SetURL("http://"+beego.AppConfig.String("es::host")+":"+beego.AppConfig.String("es::port")))
	if err != nil {
		panic(err)
		// Handle error
	}
	info, code, err := e.EClient.Ping("http://"+beego.AppConfig.String("es::host")+":"+beego.AppConfig.String("es::port")).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	nlogs.ConsoleLogs.Debug("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
	exists, err := e.EClient.IndexExists("pacs").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists{
		//e.createIndex()
	}
}

func (e *ES)createIndex()  {
	mapping := `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":2
	},
	"mappings":{
		"_default_": {
			"_all": {
				"enabled": true
			}
		},
		"pacs":{
			"properties":{
				"orgname":{
				},
				"duns":{
				},
				"sickinfo":{
					"IDcard":{
					},
					"name":{
					},
					"namec":{
					},
					"age":{
					},
					"sex":{
					},
					"patientid":{
					},
					"birthday":{
					}
				},
				"seriesinfo":{
					"studyInstanceUID":{
					},
					"seriesID":{
					},
					"imgs":{
						"imgname":{
						},
						"url":{
						},
						"size":{
						},
						"status":{
						},
						"dpath":{
						},
						"spath":{
						}
					},
					"count":{
					},
					"datetimex":{
					}
				},
				"bodypart":{
				}
			}
		}
		"dicom":{
			"properties":{
				"pacs":{
					"group":{
					},
					"element":{
					},
					"name":{
					},
					"vr":{
					}，
					"vl":{
					},
					"value":{
					}
				},
				"IDcard":{
				},
				"studyInstanceUID":{
				},
				"seriesID":{
				},
				"datetimex":{
				}
			}
		}
	}
}
`
	createIndex, err := e.EClient.CreateIndex("pacs").Body(mapping).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
	}
}

func (e *ES)IndexData(uid,typex,obj string){
	_, err := e.EClient.Index().
		Index("pacs").
		Type(typex).
		Id(uid).
		BodyString(obj).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		//nlogs.ConsoleLogs.Debug("index error")
		nlogs.FileLogs.Error(obj)
		//nlogs.ConsoleLogs.Debug(uid)
		//panic(err)
	}
}

func (e *ES)UpdateIndex(uid,typex string,obj interface{})  {
	_,err := e.EClient.Update().
		Index("pacs").
		Type(typex).
		Id(uid).
		Script(elastic.NewScriptInline("ctx._source.retweets += params.num").Lang("painless").Param("num", 1)).
		Upsert(obj).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
		//nlogs.FileLogs.Error(obj)
	}
}

func (e *ES)GetPatients(idcard string) message.PrivateHits{
	termQuery := elastic.NewTermQuery("sickinfo.IDcard", idcard)
	searchResult, err := e.EClient.Search().
		Index("pacs").
		Type("pacs").
		Query(termQuery).   // specify the query
		//Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(context.Background())             // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	var hits message.PrivateHits
	if searchResult.Hits.TotalHits > 0 {
		nlogs.ConsoleLogs.Info("Found a total of %d pacs\n",searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			var t message.Private
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}
			hits.Info = append(hits.Info,t)
			// Work with tweet
			//fmt.Printf(t.SickInfo.SickName)
		}
	} else {
		// No hits
		nlogs.ConsoleLogs.Debug("Found no pacs\n")
	}
	//result,err := json.Marshal(hits)
	if err != nil {
		// Deserialization failed
	}
	//fmt.Println(string(result))
	return hits
}
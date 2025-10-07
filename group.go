package lightnovel

type ParentGroupID uint8

const (
	ParentGroupAll        ParentGroupID = 0 // 全部
	ParentGroupNews       ParentGroupID = 1 // 资讯
	ParentGroupLightNovel ParentGroupID = 3 // 轻小说

	ParentGroupManga    ParentGroupID = 33 // 漫画
	ParentGroupAnime    ParentGroupID = 34 // 动画
	ParentGroupMaterial ParentGroupID = 36 // 素材
	ParentGroupPicture  ParentGroupID = 37 // 图坊

	ParentGroupOther ParentGroupID = 40 // 其它資源
	ParentGroupGuild ParentGroupID = 42 // 公會大廳
)

type GroupID uint8

const (
	GroupAll GroupID = 0 // 全部 ParentGroupID=1(资讯)、33(漫画)、34(动画)、36(素材)、37(图坊)、40(其它資源)

	GroupNewsLightNovel GroupID = 100 // 輕小說 ParentGroupID=1(资讯)
	GroupNewsManga      GroupID = 101 // 漫畫 ParentGroupID=1(资讯)
	GroupNewsAnime      GroupID = 102 // 動漫 ParentGroupID=1(资讯)
	GroupNewsGame       GroupID = 103 // 遊戲 ParentGroupID=1(资讯)
	GroupNewsModel      GroupID = 104 // 手辦模型 ParentGroupID=1(资讯)
	GroupNewsOther      GroupID = 105 // 其他 ParentGroupID=1(资讯)

	GroupLightNovelLatest   GroupID = 106 // 最新 ParentGroupID=3(轻小说)
	GroupLightNovelComplete GroupID = 107 // 整卷 ParentGroupID=3(轻小说)
	GroupLightNovelDownload GroupID = 108 // 下载 ParentGroupID=3(轻小说)
	GroupLightNovelEpub     GroupID = 110 // Epub ParentGroupID=3(轻小说)
	GroupLightNovelOriginal GroupID = 111 // 原創 ParentGroupID=3(轻小说)

	GroupMangaRelease  GroupID = 112 // 發佈 ParentGroupID=33(漫畫)
	GroupMangaRepost   GroupID = 113 // 轉載 ParentGroupID=33(漫畫)
	GroupMangaDownload GroupID = 114 // 下載 ParentGroupID=33(漫畫)

	GroupAnimeRelease GroupID = 115 // 發佈 ParentGroupID=34(動畫)
	GroupAnimeRepost  GroupID = 116 // 轉載 ParentGroupID=34(動畫)

	GroupMaterialPopularScience GroupID = 119 // 科普 ParentGroupID=36(素材)
	GroupMaterialTutorial       GroupID = 120 // 教程 ParentGroupID=36(素材)
	GroupMaterialConceptSetting GroupID = 121 // 概念設定 ParentGroupID=36(素材)

	GroupPictureRepost   GroupID = 122 // 轉載 ParentGroupID=37(图坊)
	GroupPictureOriginal GroupID = 123 // 原創 ParentGroupID=37(图坊)
	GroupPictureOther    GroupID = 124 // 其他 ParentGroupID=37(图坊)

	GroupOtherMusic GroupID = 125 // 音樂 ParentGroupID=40(其它資源)
	GroupOtherGame  GroupID = 126 // 遊戲 ParentGroupID=40(其它資源)
	GroupRadioDrama GroupID = 127 // 廣播劇 ParentGroupID=40(其它資源)
	GroupOtherOther GroupID = 128 // 其它 ParentGroupID=40(其它資源)
)

package api

//ArticleResponse ...
type ArticleResponse struct {
	TypeOf                 string            `json:"type_of"`
	ID                     int32             `json:"id"`
	Title                  string            `json:"title"`
	Description            string            `json:"description"`
	CoverImage             string            `json:"cover_image"`
	ReadablePublishDate    string            `json:"readable_publish_date"`
	SocialImage            string            `json:"social_image"`
	TagList                []string          `json:"tag_list"`
	Tags                   string            `json:"tags"`
	Slug                   string            `json:"slug"`
	Path                   string            `json:"path"`
	URL                    string            `json:"url"`
	CanonicalURL           string            `json:"canonical_url"`
	CommentsCount          int32             `json:"comments_count"`
	PositiveReactionsCount int32             `json:"positive_reactions_count"`
	PublicReactionsCount   int32             `json:"public_reactions_count"`
	CreatedAt              string            `json:"created_at"`
	EditedAt               string            `json:"edited_at"`
	CrosspostedAt          string            `json:"crossposted_at"`
	PublishedAt            string            `json:"published_at"`
	LastCommentAt          string            `json:"last_comment_at"`
	PublishedTimestamp     string            `json:"published_timestamp"`
	BodyHtml               string            `json:"body_html"`
	BodyMarkdown           string            `json:"body_markdown"`
	User                   *UserType         `json:"user"`
	ReadingTimeMinutes     int32             `json:"reading_time_minutes"`
	Organization           *OrganizationType `json:"organization"`
	FlareTag               *FlareTagType     `json:"flare_tag"`
}

//GetArticlesResponse
type GetArticlesResponse []*ArticleResponse

//ModifiedArticle include the response from an Update or Create article
type ModifiedArticle struct {
	TypeOf                 string            `json:"type_of"`
	ID                     int32             `json:"id"`
	Title                  string            `json:"title"`
	Description            string            `json:"description"`
	CoverImage             string            `json:"cover_image"`
	ReadablePublishDate    string            `json:"readable_publish_date"`
	SocialImage            string            `json:"social_image"`
	TagList                string            `json:"tag_list"`
	Tags                   []string          `json:"tags"`
	Slug                   string            `json:"slug"`
	Path                   string            `json:"path"`
	URL                    string            `json:"url"`
	CanonicalURL           string            `json:"canonical_url"`
	CommentsCount          int32             `json:"comments_count"`
	PositiveReactionsCount int32             `json:"positive_reactions_count"`
	PublicReactionsCount   int32             `json:"public_reactions_count"`
	CreatedAt              string            `json:"created_at"`
	EditedAt               string            `json:"edited_at"`
	CrosspostedAt          string            `json:"crossposted_at"`
	PublishedAt            string            `json:"published_at"`
	LastCommentAt          string            `json:"last_comment_at"`
	PublishedTimestamp     string            `json:"published_timestamp"`
	BodyHtml               string            `json:"body_html"`
	BodyMarkdown           string            `json:"body_markdown"`
	User                   *UserType         `json:"user"`
	ReadingTimeMinutes     int32             `json:"reading_time_minutes"`
	Organization           *OrganizationType `json:"organization"`
	FlareTag               *FlareTagType     `json:"flare_tag"`
}

//UpdateArticleResponse ...
type UpdateArticleResponse ModifiedArticle

//ArticleCreatedResponse ...
type ArticleCreatedResponse ModifiedArticle

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int32  `json:"status"`
}

//ArticleEdit ...
type ArticleEdit struct {
	Article *ArticleEditType `json:"article"`
}

//ArticleEditType ...
type ArticleEditType struct {
	Title          string `json:"title"`
	BodyMarkdown   string `json:"body_markdown"`
	Published      bool   `json:"published"`
	Series         string `json:"series"`
	MainImage      string `json:"main_image"`
	CanonicalURL   string `json:"canonical_url"`
	Description    string `json:"description"`
	Tags           string `json:"tags"`
	OrganizationID int32  `json:"organization_id"`
}

//ArticleCreate ...
type ArticleCreate struct {
	Article *ArticleCreateType `json:"article"`
}

//ArticleCreateType ...
type ArticleCreateType struct {
	Title          string   `json:"title"`
	BodyMarkdown   string   `json:"body_markdown"`
	Published      bool     `json:"published"`
	Series         string   `json:"series"`
	MainImage      string   `json:"main_image"`
	CanonicalURL   string   `json:"canonical_url"`
	Description    string   `json:"description"`
	Tags           []string `json:"tags"`
	OrganizationID int32    `json:"organization_id"`
}

//User ...
type UserType struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	TwitterUsername string `json:"twitter_username"`
	GithubUsername  string `json:"github_username"`
	WebsiteURL      string `json:"website_url"`

	//Image 640x640
	ProfileImage string `json:"profile_image"`
	//Image 90x90
	ProfileImage90 string `json:"profile_image_90"`
}

//Organization ...
type OrganizationType struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Slug     string `json:"slug"`
	//Image 640x640
	ProfileImage string `json:"profile_image"`
	//Image 90x90
	ProfileImage90 string `json:"profile_image_90"`
}

type FlareTagType struct {
	Name         string `json:"name"`
	BgColorHex   string `json:"bg_color_hex"`
	TextColorHex string `json:"text_color_hex"`
}

//GetArticleQuery store the queries provided
//by the user on a Get articles
type GetArticleQuery struct {
	Page         string `json:"page"`
	PerPage      string `json:"per_page"`
	Tag          string `json:"tag"`
	Tags         string `json:"tags"`
	TagsExclude  string `json:"tags_exclude"`
	Username     string `json:"username"`
	State        string `json:"state"`
	Top          string `json:"top"`
	CollectionID string `json:"collection_id"`
}

package api

// RESPONSES FROM THE API

//ArticleResponse ...
type ArticleResponse struct {
	TypeOf                 string            `json:"type_of"`
	ID                     string            `json:"id"`
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
type GetArticlesResponse struct {
	Articles []*ArticleResponse `json:"articles"`
}

//UpdateArticleResponse ...
type UpdateArticleResponse struct {
	TypeOf                 string            `json:"type_of"`
	ID                     string            `json:"id"`
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

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int32  `json:"status"`
}

// REQUEST BODY TO API

//ArticleEdit ...
type ArticleEdit struct {
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
	BgColorText  string `json:"bg_color_text"`
	TextColorHex string `json:"text_color_hex"`
}

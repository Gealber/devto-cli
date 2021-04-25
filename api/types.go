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
	Page         int32  `json:"page"`
	PerPage      int32  `json:"per_page"`
	Tag          string `json:"tag"`
	Tags         string `json:"tags"`
	TagsExclude  string `json:"tags_exclude"`
	Username     string `json:"username"`
	State        string `json:"state"`
	Top          int32  `json:"top"`
	CollectionID int32  `json:"collection_id"`
}

//GetLatestArticleQuery store the queries provided
//by the user on a Get latest articles
type GetLatestArticleQuery struct {
	Page    int32 `json:"page"`
	PerPage int32 `json:"per_page"`
}

//ArticleVideoResponse ...
type ArticleVideoResponse struct {
	TypeOf                 string    `json:"type_of"`
	ID                     int32     `json:"id"`
	Path                   string    `json:"path"`
	CloudinaryVideoURL     string    `json:"cloudinary_video_url"`
	Title                  string    `json:"title"`
	VideoDurationInMinutes string    `json:"video_duration_in_minutes"`
	VideoSourceURL         string    `json:"video_source_url"`
	User                   *UserType `json:"user"`
}

//ArticlesVideoResponse ...
type ArticlesVideoResponse []*ArticleVideoResponse

//CommentResponse ...
type CommentType struct {
	TypeOf    string         `json:"type_of"`
	IDCode    string         `json:"id_code"`
	CreatedAt string         `json:"created_at"`
	BodyHtml  string         `json:"body_html"`
	User      *UserType      `json:"user"`
	Children  []*CommentType `json:"children"`
}

//CommentsResponse ...
type CommentsResponse []*CommentType

//CommentQuery ...
type CommentQuery struct {
	AID int32 `json:"a_id"`
	PID int32 `json:"p_id"`
}

//FollowTagsResponse ...
type FollowTagsResponse []*FollowsTagType

//FollowTagType ...
type FollowsTagType struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Points float64 `json:"points"`
}

//TagsResponse ...
type TagsResponse []*TagType

//TagType ...
type TagType struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	BgColorHex   string `json:"bg_color_hex"`
	TextColorHex string `json:"text_color_hex"`
}

//TagsQuery ...
type TagsQuery struct {
	Page    int32 `json:"page"`
	PerPage int32 `json:"per_page"`
}

//FollowerType ...
type FollowerType struct {
	TypeOf    string `json:"type_of"`
	CreatedAt string `json:"created_at"`
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Username  string `json:"username"`
	//image 60x60
	ProfileImage string `json:"profile_image"`
}

//FollowersResponse ...
type FollowersResponse []*FollowerType

//FollowersQuery ...
type FollowersQuery struct {
	Page    int32  `json:"page"`
	PerPage int32  `json:"per_page"`
	Sort    string `json:"sort"`
}

//ListingType ...
type ListingType struct {
	TypeOf        string            `json:"type_of"`
	ID            int64             `json:"id"`
	Title         string            `json:"title"`
	Slug          string            `json:"slug"`
	BodyMarkdown  string            `json:"body_markdown"`
	TagList       string            `json:"tag_list"`
	Tags          []string          `json:"tags"`
	Category      string            `json:"category"`
	ProcessedHtml string            `json:"precessed_html"`
	Published     bool              `json:"published"`
	User          *UserType         `json:"user"`
	Organization  *OrganizationType `json:"organization"`
}

//ListingCreateType ...
type ListingCreateType struct {
	Title               string   `json:"title"`
	BodyMarkdown        string   `json:"body_markdown"`
	Category            string   `json:"category"`
	TagList             string   `json:"tag_list"`
	Tags                []string `json:"tags"`
	ExpiresAt           string   `json:"expires_at"`
	Contact_via_connect bool     `json:"contact_via_connect"`
	Location            string   `json:"location"`
	OrganizationID      int64    `json:"organization_id"`
	Action              string   `json:"action"`
}

//ListingUpdateType ...
type ListingUpdateType struct {
	Title               string   `json:"title"`
	BodyMarkdown        string   `json:"body_markdown"`
	Category            string   `json:"category"`
	TagList             []string `json:"tag_list"`
	Tags                string   `json:"tags"`
	ExpiresAt           string   `json:"expires_at"`
	Contact_via_connect bool     `json:"contact_via_connect"`
	Location            string   `json:"location"`
	Action              string   `json:"action"`
}

//ListingUpdate ...
type ListingUpdate struct {
	Listing *ListingUpdateType `json:"listing"`
}

//ListingCreate ...
type ListingCreate struct {
	Listing *ListingCreateType `json:"listing"`
}

//ListingResponse ...
type ListingResponse []*ListingType

//ListingQuery ...
type ListingQuery struct {
	Page     int32  `json:"page"`
	PerPage  int32  `json:"per_page"`
	Category string `json:"category"`
}

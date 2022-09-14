package vrchat

import "time"

type UserData struct {
	ID                 string   `json:"id"`
	Username           string   `json:"username"`
	DisplayName        string   `json:"displayName"`
	UserIcon           string   `json:"userIcon"`
	Bio                string   `json:"bio"`
	BioLinks           []string `json:"bioLinks"`
	ProfilePicOverride string   `json:"profilePicOverride"`
	StatusDescription  string   `json:"statusDescription"`
	PastDisplayNames   []struct {
		DisplayName string    `json:"displayName"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"pastDisplayNames"`
	HasEmail                       bool          `json:"hasEmail"`
	HasPendingEmail                bool          `json:"hasPendingEmail"`
	ObfuscatedEmail                string        `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail         string        `json:"obfuscatedPendingEmail"`
	EmailVerified                  bool          `json:"emailVerified"`
	HasBirthday                    bool          `json:"hasBirthday"`
	Unsubscribe                    bool          `json:"unsubscribe"`
	StatusHistory                  []string      `json:"statusHistory"`
	StatusFirstTime                bool          `json:"statusFirstTime"`
	Friends                        []string      `json:"friends"`
	FriendGroupNames               []interface{} `json:"friendGroupNames"`
	CurrentAvatarImageURL          string        `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageURL string        `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatar                  string        `json:"currentAvatar"`
	CurrentAvatarAssetURL          string        `json:"currentAvatarAssetUrl"`
	FallbackAvatar                 string        `json:"fallbackAvatar"`
	AccountDeletionDate            interface{}   `json:"accountDeletionDate"`
	AcceptedTOSVersion             int           `json:"acceptedTOSVersion"`
	SteamID                        string        `json:"steamId"`
	SteamDetails                   struct {
	} `json:"steamDetails"`
	OculusID                 string    `json:"oculusId"`
	HasLoggedInFromClient    bool      `json:"hasLoggedInFromClient"`
	HomeLocation             string    `json:"homeLocation"`
	TwoFactorAuthEnabled     bool      `json:"twoFactorAuthEnabled"`
	TwoFactorAuthEnabledDate time.Time `json:"twoFactorAuthEnabledDate"`
	State                    string    `json:"state"`
	Tags                     []string  `json:"tags"`
	DeveloperType            string    `json:"developerType"`
	LastLogin                time.Time `json:"last_login"`
	LastPlatform             string    `json:"last_platform"`
	AllowAvatarCopying       bool      `json:"allowAvatarCopying"`
	Status                   string    `json:"status"`
	DateJoined               string    `json:"date_joined"`
	IsFriend                 bool      `json:"isFriend"`
	FriendKey                string    `json:"friendKey"`
	LastActivity             time.Time `json:"last_activity"`
	OnlineFriends            []string  `json:"onlineFriends"`
	ActiveFriends            []string  `json:"activeFriends"`
	OfflineFriends           []string  `json:"offlineFriends"`
}

type UserSearchData struct {
	AllowAvatarCopying             bool      `json:"allowAvatarCopying"`
	Bio                            string    `json:"bio"`
	BioLinks                       []string  `json:"bioLinks"`
	CurrentAvatarImageUrl          string    `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageUrl string    `json:"currentAvatarThumbnailImageUrl"`
	DateJoined                     time.Time `json:"date_joined"`
	DeveloperType                  string    `json:"developerType"`
	DisplayName                    string    `json:"displayName"`
	FriendKey                      string    `json:"friendKey"`
	FiendRequestStatus             string    `json:"friendRequestStatus"`
	ID                             string    `json:"id"`
	InstanceId                     string    `json:"instanceId"`
	IsFriend                       bool      `json:"isFriend"`
	LastActivity                   string    `json:"last_activity"`
	LastLogin                      string    `json:"last_login"`
	LastPlatform                   string    `json:"last_platform"`
	Location                       string    `json:"location"`
	Note                           string    `json:"note"`
	ProfilePicOverride             string    `json:"profilePicOverride"`
	State                          string    `json:"state"`
	Status                         string    `json:"status"`
	StatusDescription              string    `json:"statusDescription"`
	Tags                           []string  `json:"tags"`
	TravelingToInstance            string    `json:"travelingToInstance"`
	TravelingToLocation            string    `json:"travelingToLocation"`
	TravelingToWorld               string    `json:"travelingToWorld"`
	UserIcon                       string    `json:"userIcon"`
	Username                       string    `json:"username"`
	WorldId                        string    `json:"worldId"`
}

type WorldSearch struct {
	AssetUrl            string    `json:"assetUrl"`
	AuthorId            string    `json:"authorId"`
	AuthorName          string    `json:"authorName"`
	Capacity            int       `json:"capacity"`
	CreatedAt           time.Time `json:"created_at"`
	Description         string    `json:"description"`
	Favorites           int       `json:"favorites"`
	Featured            bool      `json:"featured"`
	Heat                int       `json:"heat"`
	ID                  string    `json:"id"`
	ImageUrl            string    `json:"imageUrl"`
	LabsPublicationDate string    `json:"labsPublicationDate"`
	Name                string    `json:"name"`
	Namespace           string    `json:"namespace"`
	Occupants           int       `json:"occupants"`
	Organization        string    `json:"organization"`
	Popularity          int       `json:"popularity"`
	PreviewYoutubeId    *string   `json:"previewYoutubeId"`
	PrivateOccupants    int       `json:"privateOccupants"`
	PublicOccupants     int       `json:"publicOccupants"`
	PublicationDate     string    `json:"publicationDate"`
	ReleaseStatus       string    `json:"releaseStatus"`
	Tags                []string  `json:"tags"`
	ThumbnailImageUrl   string    `json:"thumbnailImageUrl"`
	UpdatedAt           time.Time `json:"updated_at"`
	Version             int       `json:"version"`
	Visits              int       `json:"visits"`
	UnityPackages       []struct {
		AssetUrl        string    `json:"assetUrl"`
		AssetVersion    int       `json:"assetVersion"`
		CreatedAt       time.Time `json:"created_at"`
		ID              string    `json:"id"`
		Platform        string    `json:"platform"`
		PluginUrl       string    `json:"pluginUrl"`
		UnitySortNumber int       `json:"unitySortNumber"`
		UnityVersion    string    `json:"unityVersion"`
	}
}

type UserLoginRequireTwoFactorAuthResponseBody struct {
	RequiresTwoFactorAuth []string `json:"requiresTwoFactorAuth"`
}

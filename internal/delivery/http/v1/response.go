package v1

type response struct {
	Error string `json:"message,omitempty"`
}

type tokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type bannerIdResponse struct {
	BannerId int `json:"banner_id"`
}

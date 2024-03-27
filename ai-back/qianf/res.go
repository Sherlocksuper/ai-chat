package qianf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	PromptLimit         = 150
	NegativePromptLimit = 150
)

// 定义图片尺寸的枚举
type Size string

const (
	SizeAvatarSmall    Size = "768x768"
	SizeAvatarMedium   Size = "1024x1024"
	SizeAvatarLarge    Size = "1536x1536"
	SizeAvatarXLarge   Size = "2048x2048"
	SizeArticleMedium  Size = "1024x768"
	SizeArticleLarge   Size = "2048x1536"
	SizePosterMedium   Size = "768x1024"
	SizePosterLarge    Size = "1536x2048"
	SizeWallpaperSmall Size = "1024x576"
	SizeWallpaperLarge Size = "2048x1152"
	SizeFlyerSmall     Size = "576x1024"
	SizeFlyerLarge     Size = "1152x2048"
)

// 定义采样方式的枚举
type Sampler string

const (
	SamplerEuler          Sampler = "Euler"
	SamplerEulerA         Sampler = "Euler a"
	SamplerDPM2M          Sampler = "DPM++ 2M"
	SamplerDPM2MKarras    Sampler = "DPM++ 2M Karras"
	SamplerLMSKarras      Sampler = "LMS Karras"
	SamplerDPMSDE         Sampler = "DPM++ SDE"
	SamplerDPMSDEKarras   Sampler = "DPM++ SDE Karras"
	SamplerDPM2AKarras    Sampler = "DPM2 a Karras"
	SamplerHeun           Sampler = "Heun"
	SamplerDPM2MSDE       Sampler = "DPM++ 2M SDE"
	SamplerDPM2MSDEKarras Sampler = "DPM++ 2M SDE Karras"
	SamplerDPM2           Sampler = "DPM2"
	SamplerDPM2Karras     Sampler = "DPM2 Karras"
	SamplerDPM2A          Sampler = "DPM2 a"
	SamplerLMS            Sampler = "LMS"
)

// 定义生成风格的枚举
type Style string

const (
	StyleBase         Style = "Base"
	Style3DModel      Style = "3D Model"
	StyleAnalogFilm   Style = "Analog Film"
	StyleAnime        Style = "Anime"
	StyleCinematic    Style = "Cinematic"
	StyleComicBook    Style = "Comic Book"
	StyleCraftClay    Style = "Craft Clay"
	StyleDigitalArt   Style = "Digital Art"
	StyleEnhance      Style = "Enhance"
	StyleFantasyArt   Style = "Fantasy Art"
	StyleIsometric    Style = "Isometric"
	StyleLineArt      Style = "Line Art"
	StyleLowpoly      Style = "Lowpoly"
	StyleNeonpunk     Style = "Neonpunk"
	StyleOrigami      Style = "Origami"
	StylePhotographic Style = "Photographic"
	StylePixelArt     Style = "Pixel Art"
	StyleTexture      Style = "Texture"
)

// ImageRequestBody 定义了请求体的结构体
type ImageRequestBody struct {
	Prompt         string  `json:"prompt"`
	NegativePrompt string  `json:"negative_prompt,omitempty"`
	Size           Size    `json:"size,omitempty"`
	N              int     `json:"n,omitempty"`
	Steps          int     `json:"steps,omitempty"`
	SamplerIndex   Sampler `json:"sampler_index,omitempty"`
	Seed           int     `json:"seed,omitempty"`
	CfgScale       float64 `json:"cfg_scale,omitempty"`
	Style          Style   `json:"style,omitempty"`
	UserID         string  `json:"user_id,omitempty"`
}

// Validate 方法用于验证 ImageRequestBody 结构体的有效性
func (r *ImageRequestBody) Validate() error {
	// 验证 prompt 字段
	if r.Prompt == "" || len(r.Prompt) > PromptLimit {
		return errors.New("prompt is required and must be less than " + strconv.Itoa(PromptLimit) + " characters")
	}

	// 验证 negative_prompt 字段
	if len(r.NegativePrompt) > NegativePromptLimit {
		return errors.New("negative_prompt must be less than " + strconv.Itoa(NegativePromptLimit) + " characters")
	}

	// 验证 size 字段
	validSizes := map[Size]bool{
		SizeAvatarSmall:    true,
		SizeAvatarMedium:   true,
		SizeAvatarLarge:    true,
		SizeAvatarXLarge:   true,
		SizeArticleMedium:  true,
		SizeArticleLarge:   true,
		SizePosterMedium:   true,
		SizePosterLarge:    true,
		SizeWallpaperSmall: true,
		SizeWallpaperLarge: true,
		SizeFlyerSmall:     true,
		SizeFlyerLarge:     true,
	}
	if r.Size != "" && !validSizes[r.Size] {
		return errors.New("size is not valid")
	}

	// 验证 n 字段
	if r.N < 1 || r.N > 4 {
		return errors.New("n must be between 1 and 4")
	}

	// 验证 steps 字段
	if r.Steps < 10 || r.Steps > 50 {
		return errors.New("steps must be between 10 and 50")
	}

	// 验证 sampler_index 字段
	validSamplers := map[Sampler]bool{
		SamplerEuler:          true,
		SamplerEulerA:         true,
		SamplerDPM2M:          true,
		SamplerDPM2MKarras:    true,
		SamplerLMSKarras:      true,
		SamplerDPMSDE:         true,
		SamplerDPMSDEKarras:   true,
		SamplerDPM2AKarras:    true,
		SamplerHeun:           true,
		SamplerDPM2MSDE:       true,
		SamplerDPM2MSDEKarras: true,
		SamplerDPM2:           true,
		SamplerDPM2Karras:     true,
		SamplerDPM2A:          true,
		SamplerLMS:            true,
	}
	if r.SamplerIndex != "" && !validSamplers[r.SamplerIndex] {
		return errors.New("sampler_index is not valid")
	}

	// 验证 seed 字段
	if r.Seed < 0 || r.Seed > 4294967295 {
		return errors.New("seed must be between 0 and 4294967295")
	}

	// 验证 cfg_scale 字段
	if r.CfgScale < 0 || r.CfgScale > 30 {
		return errors.New("cfg_scale must be between 0 and 30")
	}

	// 验证 style 字段
	validStyles := map[Style]bool{
		StyleBase:         true,
		Style3DModel:      true,
		StyleAnalogFilm:   true,
		StyleAnime:        true,
		StyleCinematic:    true,
		StyleComicBook:    true,
		StyleCraftClay:    true,
		StyleDigitalArt:   true,
		StyleEnhance:      true,
		StyleFantasyArt:   true,
		StyleIsometric:    true,
		StyleLineArt:      true,
		StyleLowpoly:      true,
		StyleNeonpunk:     true,
		StyleOrigami:      true,
		StylePhotographic: true,
		StylePixelArt:     true,
		StyleTexture:      true,
	}
	if r.Style != "" && !validStyles[r.Style] {
		return errors.New("style is not valid")
	}

	// 如果所有验证都通过，则返回 nil
	return nil
}

func GenerateImage(accessToken string) {
	request := ImageRequestBody{
		Prompt: "cat",
		N:      1,
		// ... 其他字段
	}

	err := request.Validate()
	if err != nil {
	}

	js, _ := json.Marshal(request)

	resp, err := http.Post(DIFFUSSION_XL_BASEURL+"?access_token="+accessToken, "application/json", strings.NewReader(string(js)))
	//设置请求头
	resp.Header.Add("Content-Type", "application/json")
	if err != nil {
	}
	defer resp.Body.Close()

	//读取返回的数据
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

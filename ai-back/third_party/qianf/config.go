package qianf

import (
	"errors"
	"math/rand"
	"strconv"
)

type TokenService struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int    `json:"expires_in"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	Scope         string `json:"scope"`
	SessionSecret string `json:"session_secret"`
}

const (
	PromptLimit         = 150
	NegativePromptLimit = 150
)

// Size 定义图片尺寸的枚举
type Size string

const (
	Size1024x1024 Size = "1024x1024"
	Size2048x2048 Size = "2048x2048"
	Size1280x720  Size = "1280x720"
	Size720x1280  Size = "720x1280"
	Size2560x1440 Size = "2560x1440"
	Size1440x2560 Size = "1440x2560"
	Size1152x768  Size = "1152x768"
	Size2304x1536 Size = "2304x1536"
	Size1024x768  Size = "1024x768"
	Size2048x1536 Size = "2048x1536"
)

// Sampler 定义采样方式的枚举
type Sampler string

const (
	PNDM           Sampler = "pndm"
	EulerAncestral Sampler = "euler_ancestral"
	DPMSolver      Sampler = "dpm-solver"
	DDIM           Sampler = "ddim"
)

// Style 定义生成风格的枚举
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
	Size           Size    `json:"size,omitempty" `
	N              int     `json:"n,omitempty," `
	Steps          int     `json:"steps,omitempty" `
	SamplerIndex   Sampler `json:"sampler_index,omitempty" `
	Seed           int     `json:"seed,omitempty"`
	CfgScale       float64 `json:"cfg_scale,omitempty" `
	Style          Style   `json:"style,omitempty" `
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

	// 验证 size 字段 并设置默认值
	validSizes := map[Size]bool{
		Size1024x1024: true,
		Size2048x2048: true,
		Size1280x720:  true,
		Size720x1280:  true,
		Size2560x1440: true,
		Size1440x2560: true,
		Size1152x768:  true,
		Size2304x1536: true,
		Size1024x768:  true,
		Size2048x1536: true,
	}

	// 验证size
	if r.Size == "" {
		r.Size = Size1024x1024
	} else if r.Size != "" && !validSizes[r.Size] {
		return errors.New("size is not valid")
	}

	if r.N == 0 {
		r.N = 1
	} else if r.N < 1 || r.N > 2 {
		return errors.New("n must be between 1 and 2")

	}

	// 验证 steps 字段
	if r.Steps == 0 {
		r.Steps = 20
	} else if r.Steps > 50 {
		return errors.New("steps must be below and 50")
	}

	// 验证 sampler_index 字段
	validSamplers := map[Sampler]bool{
		PNDM:           true,
		EulerAncestral: true,
		DPMSolver:      true,
		DDIM:           true,
	}

	if r.SamplerIndex != "" && !validSamplers[r.SamplerIndex] {
		return errors.New("sampler_index is not valid")
	}

	// 验证 seed 字段
	if r.Seed == 0 {
		r.Seed = rand.Intn(4294967295-10) + 10
	} else if r.Seed < 10 || r.Seed > 4294967295 {
		return errors.New("seed must be between 10 and 4294967295")
	}

	// 验证 cfg_scale 字段
	if r.CfgScale == 0 {
		r.CfgScale = 8.0
	} else if r.CfgScale < 1.0 || r.CfgScale > 15.0 {
		return errors.New("cfg_scale must be between 1.0 and 15.0")
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
	if r.Style == "" {
		r.Style = StyleBase
	} else if r.Style != "" && !validStyles[r.Style] {
		return errors.New("style is not valid")
	}

	// 如果所有验证都通过，则返回 nil
	return nil
}

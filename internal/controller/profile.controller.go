package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/virgilIw/final-fase3/internal/dto"
	"github.com/virgilIw/final-fase3/internal/service"
	pkg "github.com/virgilIw/final-fase3/pkg/jwt"
)

type ProfileController struct {
	profilService *service.ProfileService
}

func NewProfileController(profileService *service.ProfileService) *ProfileController {
	return &ProfileController{
		profilService: profileService,
	}
}

// GetProfile godoc
//
//	@Summary		Get user profile
//	@Description	Get user account by ID
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	dto.Response
//	@Failure		400	{object}	dto.Response
//	@Failure		404	{object}	dto.Response
//	@Failure		500	{object}	dto.Response
//	@Router			/profile/{id} [get]
func (g *ProfileController) GetProfile(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Invalid ID",
			Success: false,
			Data:    []any{},
			Error:   err.Error(),
		})
		return
	}

	var req dto.GetProfileRequest
	req.ID = id

	profile, err := g.profilService.GetProfile(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Response{
			Msg:     "Profile not found",
			Success: false,
			Data:    []any{},
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Success To Get Profile",
		Success: true,
		Data: []any{
			gin.H{
				"profile": profile,
			},
		},
	})
}

// InputProfile godoc
//
//	@Summary		Input / Update Profile
//	@Description	Update user profile with name, bio, and profile image
//	@Tags			Profile
//	@Accept			mpfd
//	@Produce		json
//	@Security		BearerAuth
//	@Param			user_name	formData	string	true	"User name"
//	@Param			user_bio	formData	string	false	"User biography"
//	@Param			user_image	formData	file	false	"Profile image (jpg, jpeg, png)"
//	@Success		200	{object}	dto.Response
//	@Failure		400	{object}	dto.Response
//	@Failure		500	{object}	dto.Response
//	@Router			/profile/input [post]
func (p *ProfileController) InputProfile(c *gin.Context) {

	// ======================
	// 1. Binding form-data
	// ======================
	var req dto.InputProfileRequest
	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		log.Println("binding:", err.Error())

		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   err.Error(),
			Data:    []any{},
		})
		return
	}

	// ======================
	// 2. Ambil account id dari token
	// ======================
	token, _ := c.Get("token")
	claims, _ := token.(pkg.JwtClaims)

	req.AccountID = claims.UserID

	var imagePath string

	// ======================
	// 3. Kalau ada image
	// ======================
	if req.UserImage != nil {

		// validasi extension
		ext := strings.ToLower(filepath.Ext(req.UserImage.Filename))
		re := regexp.MustCompile(`^[.](jpg|jpeg|png)$`)

		if !re.MatchString(ext) {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     "invalid image extension",
				Success: false,
				Error:   "only jpg, jpeg, png allowed",
				Data:    []any{},
			})
			return
		}

		// buat nama file unik
		filename := fmt.Sprintf(
			"%d_profile_%d%s",
			time.Now().UnixNano(),
			req.AccountID,
			ext,
		)

		// lokasi simpan file
		savePath := filepath.Join("public", "profile", filename)

		// pastikan folder ada
		if err := os.MkdirAll(filepath.Join("public", "profile"), os.ModePerm); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, dto.Response{
				Msg:     "Internal Server Error",
				Success: false,
				Error:   "failed create folder",
				Data:    []any{},
			})
			return
		}

		// save file
		if err := c.SaveUploadedFile(req.UserImage, savePath); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, dto.Response{
				Msg:     "Internal Server Error",
				Success: false,
				Error:   "failed save image",
				Data:    []any{},
			})
			return
		}

		// path yang disimpan di DB
		imagePath = fmt.Sprintf("/profile/%s", filename)
	}

	// ======================
	// 4. Call service
	// ======================
	err := p.profilService.InputProfile(
		c.Request.Context(),
		req,
		imagePath,
	)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
			Data:    []any{},
		})
		return
	}

	// ======================
	// 5. Success response
	// ======================
	c.JSON(http.StatusOK, dto.Response{
		Msg:     "profile updated",
		Success: true,
		Data:    []any{},
	})
}

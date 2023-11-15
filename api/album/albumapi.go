package album

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type API struct {
	AlbumService Service
}

func ProvideAlbumAPI(as Service) API {
	return API{AlbumService: as}
}

func (api *API) FindAll(c *gin.Context) {
	albums, err := api.AlbumService.FindAll()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"albums": albums})
}

// FindByID godoc
// @Summary      Find an album
// @Description  get an album by ID
// @Tags         album
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "Album ID"
// @Success      200  {object}  Album
// @Failure      404
// @Failure      500
// @Router       /albums/{id} [get]
func (api *API) FindByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	album, err := api.AlbumService.FindByID(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if (AlbumDTO{}) == album {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"albums": album})
}

func (api *API) Add(c *gin.Context) {
	var newAlbum NewAlbumDTO

	if err := c.BindJSON(&newAlbum); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	album, err := api.AlbumService.Add(newAlbum)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, album)
}

func (api *API) Remove(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	removed, err := api.AlbumService.Remove(id)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if removed {
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusNotFound)
}

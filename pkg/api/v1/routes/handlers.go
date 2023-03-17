package routes

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/metal-toolbox/hollow-bomservice/internal/parse"
	sservice "go.hollow.sh/serverservice/pkg/api/v1"
)

func (r *Routes) uploadXlsxFile(c *gin.Context) (int, *sservice.ServerResponse) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return http.StatusBadRequest, &sservice.ServerResponse{Error: err.Error()}
	}
	boms, err := parse.ParseXlsxFile(data)
	// for _, bom := range boms {
	// 	fmt.Printf("boms = %v\n", bom)
	// }
	resp, err := r.repository.BatchBomsUpload(c.Request.Context(), boms)
	// fmt.Printf("upload %v, %v", resp, err)
	if err != nil {
		return http.StatusBadRequest, resp
	}
	return http.StatusOK, resp
}

func (r *Routes) getByAocMacAddress(c *gin.Context) (int, *sservice.ServerResponse) {
	_, resp, err := r.repository.GetAOCMacAddr(c.Request.Context(), c.Param("aoc_mac_address"))
	// fmt.Printf("get %v, %v", resp, err)
	if err != nil {
		return http.StatusBadRequest, resp
	}
	return http.StatusOK, resp
}

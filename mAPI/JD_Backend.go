package mAPI

import (
	"github.com/gin-gonic/gin"
)

// PolishJob
// @Tags 工作管理
// @Summary 擦亮工作
// @Description 对某一工作进行高亮标注
// @Param I'd query string true "希望高亮的JobId"
// @Router /api/PolishJob [GET]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
func PolishJob(ctx *gin.Context) {

}

// BatchPolishJobs
// @Tags 工作管理
// @Summary 批量擦亮工作
// @Description 对多个工作进行高亮标注
// @Param Ids query []string true "希望高亮的JobId组"
// @Router /api/BatchPolishJobs [GET]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
func BatchPolishJobs(ctx *gin.Context) {
}

// CollectJob
// @Tags 工作管理
// @Summary 收藏工作
// @Description 对某一工作进入收藏夹
// @Param JobId query string true "将要收藏的工作对应的JobId"
// @Param UserId query string true "UserId对应用户的收藏夹"
// @Router /api/CollectJob [GET]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
func CollectJob(ctx *gin.Context) {

}

// BatchPullOff
// @Tags 工作管理
// @Summary 批量下架工作
// @Description 将选中的工作进行下架
// @Param JobIds query []string true "将要下架的JobId组"
// @Router /api/BatchPullOff [GET]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
func BatchPullOff(ctx *gin.Context) {

}

// ListUploadedJobs
// @Tags 工作管理
// @Summary 列出用户已上传的工作
// @Description 列出用户已上传的工作
// @Param UserId query string true "将要列出工作的用户UserId"
// @Router /api/ListUploadedJobs [GET]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
func ListUploadedJobs(ctx *gin.Context) {

}

// ListCollectedJobs
// @Tags 工作管理
// @Summary 列出用户收藏的工作
// @Description 列出用户收藏的工作
// @Param UserId query string true "将要列出收藏工作的用户UserId"
// @Router /api/ListCollectedJobs [GET]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
func ListCollectedJobs(ctx *gin.Context) {

}

// ListViewedJobs
// @Tags 工作管理
// @Summary 列出最近浏览的工作
// @Description 列出最近浏览的工作
// @Param UserId query string true "将要列出最近浏览的用户UserId"
// @Router /api/ListViewedJobs [GET]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
func ListViewedJobs(ctx *gin.Context) {

}

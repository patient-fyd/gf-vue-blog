// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// PostTags is the golang structure for table post_tags.
type PostTags struct {
	PostId uint `json:"postId" orm:"post_id" description:"文章ID"` // 文章ID
	TagId  uint `json:"tagId"  orm:"tag_id"  description:"标签ID"` // 标签ID
}

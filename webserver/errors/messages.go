package errors

import "github.com/gin-gonic/gin"

var (
	InternalError      = gin.H{"error": "internal_error"}
	InvalidRequest     = gin.H{"error": "invalid_request"}
	CategoryNotFound   = gin.H{"error": "category_not_found"}
	CollectionNotFound = gin.H{"error": "collection_not_found"}
	ConceptNotFound    = gin.H{"error": "concept_not_found"}
	CriteriaNotFound   = gin.H{"error": "criteria_not_found"}
)

package errors

import "github.com/gin-gonic/gin"

var (
	NotFound                         = gin.H{"error": "not_found"}
	InternalError                    = gin.H{"error": "internal_error"}
	InvalidRequest                   = gin.H{"error": "invalid_request"}
	CategoryNotFound                 = gin.H{"error": "category_not_found"}
	CollectionNotFound               = gin.H{"error": "collection_not_found"}
	ConceptNotFound                  = gin.H{"error": "concept_not_found"}
	CriteriaNotFound                 = gin.H{"error": "criteria_not_found"}
	CommentNotFound                  = gin.H{"error": "comment_not_found"}
	MailAlreadyInUse                 = gin.H{"error": "mail_already_in_use"}
	InvalidLoginData                 = gin.H{"error": "invalid_login_data"}
	InvalidMailAddressOrUsername     = gin.H{"error": "invalid_mail_address_or_username"}
	UserAlreadyExists                = gin.H{"error": "user_already_exists"}
	MissingAuthenticationInformation = gin.H{"error": "missing_authentication"}
	NoPermission                     = gin.H{"error": "no_permission"}
)

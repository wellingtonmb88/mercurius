package auth

import "github.com/novatrixtech/mercurius/examples/simple/model"

// AccessTokenCache stores AccessToken generated and when it was generated
var AccessTokenCache map[string]model.AccessTokenData

func init() {
	AccessTokenCache = make(map[string]model.AccessTokenData, 0)
}

//RemoveUnusedAC remove from cache an Access Token generated before
func RemoveUnusedAC(contatoID int) {
	for key, ac := range AccessTokenCache {
		if ac.ContatoID == contatoID {
			delete(AccessTokenCache, key)
			break
		}
	}
}

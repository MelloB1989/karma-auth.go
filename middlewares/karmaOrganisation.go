package middlewares

import (
	orgs "karma_auth/helpers/organisations"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func IsKarmaOrganisation(c *fiber.Ctx) error {
    authHeaders := c.GetReqHeaders()
    for key, value := range authHeaders {
        log.Printf("Header: %s Value: %v\n", key, value)
    }
	log.Printf("Headers: %v\n", authHeaders)
    // Access headers with the correct case
    authID, okID := authHeaders["X-Karma-Org-Authid"]
    authToken, okToken := authHeaders["X-Karma-Org-Authtoken"]

    log.Println(authID, okID, authToken, okToken)

    if !okID || !okToken || len(authID) == 0 || len(authToken) == 0 || authID[0] == "" || authToken[0] == "" {
        return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized Token")
    }

    oid := strings.Split(authID[0], "---")[0]
    api := orgs.GetOrganisationApiToken(oid)
    if api.OID != oid || authToken[0] != api.Token {
        return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
    }

    c.Locals("oid", oid)
    c.Locals("tid", api.TID)
    c.Locals("token", authToken[0])
    return c.Next()
}
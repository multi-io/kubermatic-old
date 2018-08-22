package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kubermatic/kubermatic/cli/api"
)

var token = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImE1NGE5Y2YyNzU3M2QxNTUxZjg1NWMzYTg0ZWJhYTg3MzczOWYxY2QifQ.eyJpc3MiOiJodHRwczovL2Nsb3VkLmt1YmVybWF0aWMuaW8vZGV4Iiwic3ViIjoiQ2djek9EWTRPVE01RWdabmFYUm9kV0kiLCJhdWQiOiJrdWJlcm1hdGljIiwiZXhwIjoxNTM0OTI5NjY4LCJpYXQiOjE1MzQ4NDMyNjgsIm5vbmNlIjoicmFuZG9tIiwiYXRfaGFzaCI6ImNQN0dMX1hwNEswRkVHSTczd0hFbHciLCJlbWFpbCI6Imlnb3Iua29tbGV3QGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJuYW1lIjoiSWdvciBLb21sZXcifQ.sflsWtersAarZgLJrOzplpew21Jeq5RNCrIIM-Oc-3CQ9RDmD136YdRAA-QEKoFueCtfcGuHlkMk4OBQ9MwpowoPRCjQz0hgDGIOG8M_D2WVqWgYLEokN5O8FsS3W4OEWAZ8uy3Xi0VU4K31fbJUIi9NZCaXn3Rt455MMeSEJYmaBhdaiRWvMgsaRppbp4s0qGEQyUnM3lgxllT5nZliyc-tgt9Llz72KffAI9Bt3pBTE88Zm0SM_kpiZx8d_UUvBh6fCodUzOKiRaXR-7pr6CS2nVUpp7F-YxdPIg1X8OSKm1K2DwRHW5diUhBIh2CI9yvGrqQQf1sMHwYW_QDOEA"

func main() {
	cfg := api.NewConfiguration()
	cli := api.NewAPIClient(cfg)
	c := context.Background()
	ctx := context.WithValue(c, api.ContextAccessToken, token)

	user, code, err := cli.UsersApi.GetCurrentUser(ctx)
	if err != nil {
		log.Panic(err)
	}
	if code.StatusCode != 200 {
		log.Panic("Status code is not 200:", code.StatusCode)
	}
	printUserInfo(user)
}

func printUserInfo(user api.User) {
	fmt.Printf("Name: %s\nEmail: %s\n", user.Name, user.Email)
	fmt.Println("Projects:")
	for _, project := range user.Projects {
		fmt.Printf("\tGroup: %s, ID: %s\n", project.Group, project.Id)
	}
}

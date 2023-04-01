package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"valotips/model"

	"github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {
	post := new(model.Post)
	stand_img, err := c.FormFile("stand_img")
	if err != nil {
		return err
	}
	aim_img, err := c.FormFile("aim_img")
	if err != nil {
		return err
	}

	image, err := stand_img.Open()
	if err != nil {
		return err
	}
	stand_img_url, err := UploadImgur(image)
	if err != nil {
		return err
	}
	defer image.Close()

	image, err = aim_img.Open()
	if err != nil {
		return err
	}
	aim_img_url, err := UploadImgur(image)
	if err != nil {
		return err
	}
	defer image.Close()

	post.Title = c.FormValue("title")
	post.StandImg = stand_img_url
	post.AimImg = aim_img_url
	post.MapUUID = c.FormValue("map_uuid")
	post.AgentUUID = c.FormValue("agent_uuid")
	post.AbilitySlot = c.FormValue("ability_slot")

	model.DB.Create(&post)
	return c.JSON(http.StatusCreated, post)
}

func GetPost(c echo.Context) error {
	post := new(model.Post)
	model.DB.First(&post, c.Param("id"))
	return c.JSON(http.StatusOK, post)
}

func GetPosts(c echo.Context) error {
	posts := new([]model.Post)
	model.DB.Find(&posts)
	return c.JSON(http.StatusOK, posts)
}

func UploadImgur(image io.Reader) (string, error) {
	client_id := os.Getenv("IMGUR_CLIENT_ID")
	req, err := http.NewRequest("POST", "https://api.imgur.com/3/image", image)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Client-ID "+client_id)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// parse response to json
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var imgurResponse model.ImgurResponse
	json.Unmarshal(body, &imgurResponse)
	if !imgurResponse.Success {
		fmt.Println("Imgur upload failed")
		return "", err
	}
	return imgurResponse.Data.Link, nil
}

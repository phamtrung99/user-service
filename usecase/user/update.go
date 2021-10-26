package user

import (
	"context"
	"log"
	"mime/multipart"
	"strconv"
	"strings"

	authmyerror "github.com/phamtrung99/auth-service/util/myerror"
	"github.com/phamtrung99/gopkg/middleware"
	checkform "github.com/phamtrung99/movie-service/package/checkForm"
	imgvalid "github.com/phamtrung99/movie-service/package/fileValid"
	moviemyerror "github.com/phamtrung99/movie-service/util/myerror"
	"github.com/phamtrung99/user-service/model"
	"github.com/phamtrung99/user-service/util/myerror"
)

type UpdateInfoRequest struct {
	FormData *multipart.Form
}

func (u *Usecase) Update(ctx context.Context, req UpdateInfoRequest) (*model.User, error) {
	//Get current userId from Token.
	claim := middleware.GetClaim(ctx)
	userID := claim.UserID

	//Get current user info from userID
	user, err := u.userRepo.GetByID(ctx, userID)

	if err != nil {
		return nil, myerror.ErrUserGet(err)
	}

	//Check email
	if len(req.FormData.Value["email"]) != 0 {
		formEmail := req.FormData.Value["email"][0]
		if user.Email != formEmail {
			isMail, email := checkform.CheckFormatValue("email", formEmail)
			if !isMail {
				return nil, authmyerror.ErrEmailFormat(nil)
			}

			if u.userRepo.CheckEmailExist(ctx, email) {
				return nil, authmyerror.ErrExistedEmail(nil)
			}

			user.Email = email
		}
	}

	//Check full name
	if len(req.FormData.Value["full_name"]) != 0 {
		isName, fullName := checkform.CheckFormatValue("full_name", req.FormData.Value["full_name"][0])

		if !isName {
			return &model.User{}, authmyerror.ErrFullNameFormat(nil)
		}

		user.FullName = fullName
	}

	if len(req.FormData.Value["age"]) != 0 {
		isAge, ageStr := checkform.CheckFormatValue("age", req.FormData.Value["age"][0])
		if !isAge {
			return &model.User{}, moviemyerror.ErrAgeFormat(nil)
		}

		age, _ := strconv.Atoi(ageStr)
		user.Age = age
	}

	if len(req.FormData.File["avatar"]) != 0 {

		file := req.FormData.File["avatar"][0]
		pathFile := "/public/avatar/"

		filetype, err := imgvalid.CheckImage(file)

		if err != nil {
			return &model.User{}, err
		}

		imgFileName := strconv.FormatInt(user.ID, 10) + "." + strings.Split(filetype, "/")[1]

		err = imgvalid.CopyFile(file, imgFileName, "."+pathFile)

		if err != nil {
			log.Fatal(err)
		}

		user.Avatar = pathFile + imgFileName
	}

	result, err := u.userRepo.Update(ctx, user)

	if err != nil {
		return &model.User{}, myerror.ErrUserUpdate(err)
	}

	return result, nil
}

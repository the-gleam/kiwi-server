package timetables

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	timetablesModel "github.com/team-gleam/kiwi-basket/domain/model/timetables"
	"github.com/team-gleam/kiwi-basket/domain/model/user/token"
	timetablesRepository "github.com/team-gleam/kiwi-basket/domain/repository/timetables"
	credentialRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	errorResponse "github.com/team-gleam/kiwi-basket/interfaces/controllers/error"
	loginController "github.com/team-gleam/kiwi-basket/interfaces/controllers/user/login"
	timetablesUsecase "github.com/team-gleam/kiwi-basket/usecase/timetables"
	credentialUsecase "github.com/team-gleam/kiwi-basket/usecase/user/credential"
)

type TimetablesController struct {
	timetablesUsecase timetablesUsecase.TimetablesUsecase
}

func NewTimetablesController(
	c credentialRepository.ICredentialRepository,
	l loginRepository.ILoginRepository,
	t timetablesRepository.ITimetablesRepository,
) *TimetablesController {
	return &TimetablesController{
		timetablesUsecase.NewTimetablesUsecase(
			credentialUsecase.NewCredentialUsecase(c, l),
			t,
		),
	}
}

type TimetablesResponse struct {
	Timetables Timetables `json:"timetable"`
}

type Timetables struct {
	Mon Timetable `json:"mon"`
	Tue Timetable `json:"tue"`
	Wed Timetable `json:"wed"`
	Thu Timetable `json:"thu"`
	Fri Timetable `json:"fri"`
}

type Timetable struct {
	One   *Class `json:"1"`
	Two   *Class `json:"2"`
	Three *Class `json:"3"`
	Four  *Class `json:"4"`
	Five  *Class `json:"5"`
}

type Class struct {
	Subject string `json:"subject"`
	Room    string `json:"room"`
}

func (t TimetablesResponse) toTimetables() timetablesModel.Timetables {
	return timetablesModel.NewTimetables(
		t.Timetables.Mon.toTimetable(),
		t.Timetables.Tue.toTimetable(),
		t.Timetables.Wed.toTimetable(),
		t.Timetables.Thu.toTimetable(),
		t.Timetables.Fri.toTimetable(),
	)
}

func (t Timetable) toTimetable() timetablesModel.Timetable {
	return timetablesModel.NewTimetable(
		t.One.toClass(),
		t.Two.toClass(),
		t.Three.toClass(),
		t.Four.toClass(),
		t.Five.toClass(),
	)
}

func (t *Class) toClass() timetablesModel.Class {
	if t == nil {
		return timetablesModel.NoClass()
	}
	return timetablesModel.NewClass(t.Subject, t.Room)
}

func (c TimetablesController) Register(ctx echo.Context) error {
	t := ctx.Request().Header.Get("Token")
	if t == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(loginController.InvalidUsernameOrPassword)),
		)
	}

	res := new(TimetablesResponse)
	err := ctx.Bind(res)
	if err != nil || res.Timetables.Mon.One == new(Class) {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(loginController.InvalidJsonFormat)),
		)
	}

	timetables := res.toTimetables()

	err = c.timetablesUsecase.Add(token.NewToken(t), timetables)
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(loginController.InternalServerError)),
		)
	}

	return ctx.NoContent(http.StatusOK)
}
